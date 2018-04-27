package main

import (
	podapi "github.com/alejandroEsc/k8s-controller-example/api"
	"golang.org/x/net/context"
	"github.com/spf13/viper"
)

type podServer struct {
	healthStatus bool
}

func (p *podServer) Health(context.Context, *podapi.HealthStatusRequest) (*podapi.HealthStatus, error) {

	return &podapi.HealthStatus{IsOK: p.healthStatus}, nil
}

func (p *podServer) PodInfo(context.Context, *podapi.PodInfoRequest) (*podapi.PodInfo, error) {
	id := viper.GetString("id")
	return &podapi.PodInfo{ID: id}, nil
}

func (p *podServer) updatedHealthStatus(isOK bool) {
	p.healthStatus = isOK
}

func newPodServer() *podServer {
	return &podServer{}
}
