package podServer

import (
	podapi "github.com/alejandroEsc/k8s-controller-example/api"
	"golang.org/x/net/context"
)

// PodServer represents the api interface and
type PodServer struct {
	HealthStatus bool
	PodID        string
}

// NewPodServer creates a new PodServer
func NewPodServer(healthStatus bool, podID string) *PodServer {
	return &PodServer{HealthStatus: healthStatus, PodID: podID}
}

// Health is the impl of the interface for the podserver
func (p *PodServer) Health(context.Context, *podapi.HealthStatusRequest) (*podapi.HealthStatus, error) {

	return &podapi.HealthStatus{IsOK: p.HealthStatus}, nil
}

// PodInfo is the impl of the interface for the podserver
func (p *PodServer) PodInfo(context.Context, *podapi.PodInfoRequest) (*podapi.PodInfo, error) {
	return &podapi.PodInfo{ID: p.PodID}, nil
}

func (p *PodServer) updatedHealthStatus(isOK bool) {
	p.HealthStatus = isOK
}
