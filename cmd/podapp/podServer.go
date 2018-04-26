package main

import(
	"golang.org/x/net/context"
	podapi "github.com/alejandroEsc/k8s-controller-example/api"
)

type podServer struct {
	healthStatus bool
}

func (p *podServer) Health(context.Context, *podapi.HealthStatusRequest) (*podapi.HealthStatus, error) {

	return &podapi.HealthStatus{ IsOK: p.healthStatus}, nil
}

func (p *podServer) PodInfo(context.Context, *podapi.PodInfoRequest) (*podapi.PodInfo, error) {
	return &podapi.PodInfo{ID: "myID"}, nil
}

func (p *podServer) updatedHealthStatus(isOK bool){
	p.healthStatus = isOK
}

func newPodServer() *podServer {
	return &podServer{}
}