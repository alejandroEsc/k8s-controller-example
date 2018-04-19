package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	c "github.com/alejandroEsc/k8s-provisioner-juju-example/internal/controller"
	"github.com/alejandroEsc/k8s-provisioner-juju-example/internal/util"
	clientset "github.com/alejandroEsc/k8s-provisioner-juju-example/pkg/client/clientset/versioned"
	informers "github.com/alejandroEsc/k8s-provisioner-juju-example/pkg/client/informers/externalversions"
	"github.com/juju/loggo"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	kubeinformers "k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"github.com/alejandroEsc/k8s-provisioner-juju-example/pkg/signals"
)

var (
	logger               loggo.Logger
	masterURL            string
	kubeconfig           string
	workerThreads        int
	err                  error
)

// Init initializes the environment variables to be used by the app
func Init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("cluster_controller")
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
	logger.Infof("Starting Cluster-Controller Operator...")

	cfg, err := clientcmd.BuildConfigFromFlags(masterURL, kubeconfig)
	if err != nil {
		logger.Criticalf("Error building kubeconfig: %s", err.Error())
	}

	kubeClient, err := kubernetes.NewForConfig(cfg)
	if err != nil {
		logger.Criticalf("Error building kubernetes clientset: %s", err.Error())
	}

	client, err := clientset.NewForConfig(cfg)
	if err != nil {
		logger.Criticalf("Error building example clientset: %s", err.Error())
	}

	kubeInformerFactory := kubeinformers.NewSharedInformerFactory(kubeClient, time.Second*30)
	exampleInformerFactory := informers.NewSharedInformerFactory(client, time.Second*30)

	controller := c.NewController(kubeClient, client, kubeInformerFactory, exampleInformerFactory)

	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	//  Get notified that server is being asked to stop
	// Handle SIGINT and SIGTERM.
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	if err := controller.Run(workerThreads, signals.CreateGracefulStopCh()); err != nil {
		logger.Criticalf("failed to start server: %s", err)
		os.Exit(1)
	}

	logger.Infof("... Cluster-Controller Operator stopped")

}


