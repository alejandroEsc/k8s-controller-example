package main

import (
	"github.com/spf13/viper"
)

// Init initializes the environment variables to be used by the app
func Init() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix("cluster_controller")
	viper.BindEnv("kubeconfig")
	viper.BindEnv("master_url")
	viper.BindEnv("worker_thread_count")
}
