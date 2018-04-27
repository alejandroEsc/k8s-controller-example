package controller

import (
	"github.com/alejandroEsc/k8s-controller-example/pkg/util"
	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
)

// CreateControllerResource creates the resource object to be consumed by kubernetes
func CreateControllerResource() *apiextensionsv1beta1.CustomResourceDefinition {
	subResource := apiextensionsv1beta1.CustomResourceSubresources{
		Status: &apiextensionsv1beta1.CustomResourceSubresourceStatus{},
	}

	// No validation yet.
	var validation *apiextensionsv1beta1.CustomResourceValidation

	return util.NewCRD(
		"controller.alejandroesc.com",
		"SampleController",
		"podcontroller",
		"podcontrollers",
		[]string{"cc"},
		&subResource,
		validation)
}
