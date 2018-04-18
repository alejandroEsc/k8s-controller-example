package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterCreator is the spec for the ClusterCreator resource
type ClusterCreator struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec ClusterCreatorSpec `json:"spec"`
	Status ClusterCreatorStatus `json:"status"`
}

// ClusterCreatorSpec is the spec of the ClusterCreatorSpec
type ClusterCreatorSpec struct {
	ClusterName string `json:"clusterName"`
}

// ClusterCreatorStatus is the status of the ClusterCreator resource
type ClusterCreatorStatus struct {
	AvailableReplicas int32 `json:"availableReplicas"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterCreatorList is the list of ClusterCreator resources
type ClusterCreatorList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []ClusterCreator `json:"items"`
}