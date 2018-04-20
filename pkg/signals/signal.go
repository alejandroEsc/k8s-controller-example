package signals

import (
	"os"
	"os/signal"
)

var onlyOneSignalHandler = make(chan struct{})

// from https://github.com/kubernetes/sample-controller/tree/master/pkg/signals
func CreateGracefulStopCh() (stopCh <-chan struct{}) {
	close(onlyOneSignalHandler)

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