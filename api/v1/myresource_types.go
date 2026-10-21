package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +kubebuilder:object:root=true
// +kubebuilder:ac:generate=true
type MyResource struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec MyResourceSpec `json:"spec,omitempty"`
}

type MyResourceSpec struct {
	Selector ObjectSelector `json:"selector,omitempty"`
}

type ObjectSelector struct {
	metav1.LabelSelector `json:",inline"` // source of the panic

	// APIVersion of the target resources
	APIVersion string `yaml:"apiVersion,omitempty" json:"apiVersion,omitempty"`

	// Kind of the target resources
	Kind string `yaml:"kind,omitempty" json:"kind,omitempty"`

	// Name of the target resource
	// +optional
	Name *string `yaml:"name,omitempty" json:"name,omitempty"`
}
