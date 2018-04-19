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
)

var (
	logger               loggo.Logger
	masterURL            string
	kubeconfig           string
	workerThreads        int
	err                  error
	onlyOneSignalHandler = make(chan struct{})
)

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

	if err := controller.Run(workerThreads, createGracefulStopCh()); err != nil {
		logger.Criticalf("failed to start server: %s", err)
		os.Exit(1)
	}

	logger.Infof("... Cluster-Controller Operator stopped")

}

// from https://github.com/kubernetes/sample-controller/tree/master/pkg/signals
func createGracefulStopCh() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler)
	shutdownSignals := []os.Signal{os.Interrupt, syscall.SIGTERM, syscall.SIGINT}

	stop := make(chan struct{})
	c := make(chan os.Signal, 2)
	signal.Notify(c, shutdownSignals...)
	go func() {
		<-c
		close(stop)
		<-c
		os.Exit(1) // second signal. Exit directly.
	}()

	return stop

}
