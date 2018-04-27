package main

import (
	"os"
	"time"

	c "github.com/alejandroEsc/k8s-controller-example/internal/controller"
	clientset "github.com/alejandroEsc/k8s-controller-example/pkg/client/clientset/versioned"
	informers "github.com/alejandroEsc/k8s-controller-example/pkg/client/informers/externalversions"
	"github.com/alejandroEsc/k8s-controller-example/pkg/signals"
	"github.com/alejandroEsc/k8s-controller-example/pkg/util"
	"github.com/juju/loggo"
	"github.com/spf13/viper"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	logger        loggo.Logger
	masterURL     string
	kubeconfig    string
	workerThreads int
	err           error
)

// Init initializes the environment variables to be used by the app
func Init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("sample_controller")
	viper.BindEnv("kubeconfig")
	viper.BindEnv("master_url")
	viper.BindEnv("worker_thread_count")
}

func main() {
	Init()

	kubeconfig = viper.GetString("kubeconfig")
	masterURL = viper.GetString("master_url")
	workerThreads = viper.GetInt("worker_thread_count")

	logger = util.GetModuleLogger("cmd.controller", loggo.INFO)
	logger.Infof("Starting Sample-Controller Operator...")

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		logger.Criticalf("Error building kubeconfig: %s", err)
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		logger.Criticalf("Error building kubernetes clientset: %s", err)
	}

	client, err := clientset.NewForConfig(cfg)
	if err != nil {
		logger.Criticalf("Error building example clientset: %s", err)
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	exampleInformerFactory := informers.NewSharedInformerFactory(client, time.Second*30)

	controller := c.NewController(kubeClient, client, kubeInformerFactory, exampleInformerFactory)

	stopCh := signals.CreateGracefulStopCh()

	go kubeInformerFactory.Start(stopCh)
	go exampleInformerFactory.Start(stopCh)

	kubeExtClient := apiextensionsclient.NewForConfigOrDie(cfg)

	err = controller.InitCRD(kubeExtClient, c.CreateClusterCreatorRD())
	if err != nil {
		logger.Criticalf("Error create CRD: %s", err)
	}

	// start controller
	if err := controller.Run(workerThreads, stopCh); err != nil {
		logCriticalErrorAndExit("failed to start server: %s", err)
	}

	logger.Infof("... Sample-Controller Operator stopped")

}

func logCriticalErrorAndExit(msg string, args ...interface{}) {
	logger.Criticalf(msg, args)
	os.Exit(1)
}
