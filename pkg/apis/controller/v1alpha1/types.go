package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SampleResource is the spec for the controller
type SampleResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SampleResourceSpec   `json:"spec"`
	Status SampleResourceStatus `json:"status"`
}

// SampleResourceSpec is the spec of the SampleResourceSpec
type SampleResourceSpec struct {
	SampleResourceName string `json:"clusterName"`
}

// SampleResourceStatus is the status of the SampleResource resource
type SampleResourceStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SampleResourceList is the list of SampleResources
type SampleResourceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []SampleResource `json:"items"`
}
