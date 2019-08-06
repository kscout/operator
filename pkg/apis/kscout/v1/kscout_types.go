package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// NOTE: json tags are required.  Any new fields you add must have json tags for
// the fields to be serialized.

// KScoutSpec defines the desired state of KScout
// +k8s:openapi-gen=true
type KScoutSpec struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after
	// modifying this file
	// Add custom validation using kubebuilder tags:
	// https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// Environment is a name used to signify the stability and audience of the
	// KScout resources. This value must be unique among KScout resources.
	//
	// The special value "prod" is used for stable production code presented
	// to users.
	Environment string `json:"environment"`

	// MinReplicas is the lowest number of pods each service will be allowed
	// to run.
	MinReplicas int32 `json:"minReplicas"`

	// CatalogAPI is the desired catalogue API state
	CatalogAPI CatalogAPISpec `json:"catalogAPI"`
}

// CatalogAPISpec defines the desired state of the catalog API
type CatalogAPISpec struct {
	// ImageVersion is the container image tag
	ImageVersion string `json:"imageVersion"`
}

// KScoutStatus defines the observed state of KScout
// +k8s:openapi-gen=true
type KScoutStatus struct {
	// Important: Run "operator-sdk generate k8s" to regenerate code after
	// modifying this file
	// Add custom validation using kubebuilder tags:
	// https://book.kubebuilder.io/beyond_basics/generating_crd.html

	// CatalogAPI is the catalog API's status
	CatalogAPI CatalogAPIStatus `json:"imageVersion"`
}

// CatalogAPIStatus defines the status of the catalog API
type CatalogAPIStatus struct {
	// ImageVersion is the container image tag
	ImageVersion string `json:"imageVersion"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KScout is the Schema for the kscouts API
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status
type KScout struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KScoutSpec   `json:"spec,omitempty"`
	Status KScoutStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KScoutList contains a list of KScout
type KScoutList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KScout `json:"items"`
}

func init() {
	SchemeBuilder.Register(&KScout{}, &KScoutList{})
}
