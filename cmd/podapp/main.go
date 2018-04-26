package podapp

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/alejandroEsc/k8s-controller-example/internal/util"
	"github.com/juju/loggo"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
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

	logger.Infof("... Kraken-Cluster API stopped")

}

func start(gracefulStop chan os.Signal) error {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("kraken_cluster")

	return nil
}
