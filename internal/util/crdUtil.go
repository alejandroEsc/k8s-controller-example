package util

/**
taken from gitlab/mvenezia/redis-operator
*/

import (
	"fmt"
	"log"
	"time"

	api "github.com/alejandroEsc/k8s-controller-example/pkg/apis/controller/v1alpha1"

	apiextensionsv1beta1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1beta1"
	apiextensionsclient "k8s.io/apiextensions-apiserver/pkg/client/clientset/clientset"
	k8serrors "k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NewCRD creates crd apixextension object to be consumed by k8s api
func NewCRD(
	name string,
	kind string,
	singular string,
	plural string,
	shortNames []string,
	subResource *apiextensionsv1beta1.CustomResourceSubresources,
	validation *apiextensionsv1beta1.CustomResourceValidation) *apiextensionsv1beta1.CustomResourceDefinition {
	return &apiextensionsv1beta1.CustomResourceDefinition{
		ObjectMeta: metav1.ObjectMeta{
			Name: name,
		},
		Spec: apiextensionsv1beta1.CustomResourceDefinitionSpec{
			Group:   api.SchemeGroupVersion.Group,
			Version: api.SchemeGroupVersion.Version,
			Scope:   apiextensionsv1beta1.NamespaceScoped,
			Names: apiextensionsv1beta1.CustomResourceDefinitionNames{
				Plural:     plural,
				Kind:       kind,
				ShortNames: shortNames,
				Singular:   singular,
			},
			Subresources: subResource,
			Validation:   validation,
		},
	}
}

// DeployCRD creates the objects in kubernetes
func DeployCRD(clientset apiextensionsclient.Interface, crd *apiextensionsv1beta1.CustomResourceDefinition) error {
	_, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Create(crd)
	crdExists := k8serrors.IsAlreadyExists(err)
	if err != nil && !crdExists {
		return err
	} else if crdExists {
		log.Print("CRD already exists, skipping installation\n")
	}
	return nil
}

// WaitCRDDeployReady waits until proper condition is obtained.
func WaitCRDDeployReady(clientset apiextensionsclient.Interface, crdName string) error {
	err := Retry(5*time.Second, 20, func() (bool, error) {
		crd, err := clientset.ApiextensionsV1beta1().CustomResourceDefinitions().Get(crdName, metav1.GetOptions{})
		if err != nil {
			return false, err
		}
		for _, cond := range crd.Status.Conditions {
			switch cond.Type {
			case apiextensionsv1beta1.Established:
				if cond.Status == apiextensionsv1beta1.ConditionTrue {
					return true, nil
				}
			case apiextensionsv1beta1.NamesAccepted:
				if cond.Status == apiextensionsv1beta1.ConditionFalse {
					return false, fmt.Errorf("Name conflict: %v", cond.Reason)
				}
			}
		}
		return false, nil
	})
	if err != nil {
		return fmt.Errorf("wait CRD created failed: %v", err)
	}
	return nil
}
