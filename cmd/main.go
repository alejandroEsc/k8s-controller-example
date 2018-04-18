package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/spf13/viper"
	"github.com/juju/loggo"
	"github.com/alejandroEsc/k8s-provisioner-juju-example/internal/util"
	"golang.org/x/net/context"
	"github.com/alejandroEsc/k8s-provisioner-juju-example/internal/juju"

)

// Code for the Samsung-Cluster operator

func main() {
	logger := util.GetModuleLogger("cmd", loggo.INFO)
	logger.Infof("Starting Samsung-Cluster Operator...")


	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	//  Get notified that server is being asked to stop
	// Handle SIGINT and SIGTERM.
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	if err := start(gracefulStop); err != nil {
		logger.Criticalf("failed to start server: %s", err)
		os.Exit(1)
	}

	logger.Infof("... Samsung-Cluster Operator stopped")

}

func start(gracefulStop chan os.Signal) error {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("samsung_cluster")


	return nil
}