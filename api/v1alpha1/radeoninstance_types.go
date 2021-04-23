/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// RadeonInstanceSpec defines the desired state of RadeonInstance
type RadeonInstanceSpec struct {
	// +kubebuilder:validation:Pattern=[a-zA-Z0-9\.\-\/]+
	Namespace string `json:"namespace,omitempty"`

	// +kubebuilder:validation:Pattern=[a-zA-Z0-9\-]+
	Image string `json:"image,omitempty"`

	// Image pull policy
	// +kubebuilder:validation:Optional
	ImagePullPolicy string `json:"imagePullPolicy,omitempty"`
}

// RadeonInstanceStatus defines the observed state of RadeonInstance
type RadeonInstanceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

//+kubebuilder:object:root=true
//+kubebuilder:subresource:status

// RadeonInstance is the Schema for the radeoninstances API
type RadeonInstance struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RadeonInstanceSpec   `json:"spec,omitempty"`
	Status RadeonInstanceStatus `json:"status,omitempty"`
}

//+kubebuilder:object:root=true

// RadeonInstanceList contains a list of RadeonInstance
type RadeonInstanceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RadeonInstance `json:"items"`
}

func init() {
	SchemeBuilder.Register(&RadeonInstance{}, &RadeonInstanceList{})
}

// ImagePath returns a compiled full valid image string
func (r *RadeonInstanceSpec) ImagePath() string {
	return r.Image
}

// ImagePolicy returns a valid corev1.PullPolicy from the string in the CR
func (r *RadeonInstanceSpec) ImagePolicy(pullPolicy string) corev1.PullPolicy {
	switch corev1.PullPolicy(pullPolicy) {
	case corev1.PullAlways:
		return corev1.PullAlways
	case corev1.PullNever:
		return corev1.PullNever
	}
	return corev1.PullIfNotPresent
}
