package main

import (
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"fmt"

	podapi "github.com/alejandroEsc/k8s-controller-example/api"
	"github.com/alejandroEsc/k8s-controller-example/pkg/util"
	"github.com/juju/loggo"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

// Code for pod app (grpc server)

// Init initializes the environment variables to be used by the app
func Init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("pod_server")
	viper.BindEnv("server_id")
	viper.BindEnv("port")
	viper.BindEnv("address")

}

var (
	logger = util.GetModuleLogger("cmd.podapp", loggo.INFO)
)

func main() {
	logger.Infof("Starting Example Pod Operator...")

	Init()
	port := viper.GetInt("port")
	address := viper.GetString("address")

	ctx := context.Background()
	_, cancel := context.WithCancel(ctx)
	defer cancel()

	//  Get notified that server is being asked to stop
	// Handle SIGINT and SIGTERM.
	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGINT, syscall.SIGTERM)

	if err := start(gracefulStop, port, address); err != nil {
		logger.Criticalf("failed to start server: %s", err)
		os.Exit(1)
	}

	logger.Infof("... Example Pod API stopped")

}

func start(gracefulStop chan os.Signal, port int, address string) error {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", address, port))
	if err != nil {
		logger.Criticalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	podapi.RegisterPodMessagingServiceServer(grpcServer, newPodServer())

	logger.Infof("attempting to start server in port %d", port)

	// Chance here to gracefully handle being stopped.
	go func() {
		sig := <-gracefulStop
		logger.Infof("caught sig: %+v", sig)
		logger.Infof("Wait for 2 second to finish processing")
		time.Sleep(2 * time.Second)
		grpcServer.Stop()
		logger.Infof("service terminated")
		os.Exit(0)
	}()

	if err := grpcServer.Serve(listener); err != nil {
		logger.Errorf("could not start service: %s", err)
		return err
	}

	return nil
}
