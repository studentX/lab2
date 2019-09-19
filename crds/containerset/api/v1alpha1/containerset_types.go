/*
Copyright 2019 K8sland Training.

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
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ReplicaCount tracks the number of replicas.
!! YOUR_CODE !! Setup the min max replicas ie 1-5 included for valid replica counts
type ReplicaCount int32

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ContainerSetSpec defines the desired state of ContainerSet
type ContainerSetSpec struct {
	// Replicas specifies the desired number of pods
	Replicas ReplicaCount `json:"replicas,omitempty"`
	// Image the name and tag of the desired docker image
	Image string `json:"image,omitempty"`
}

// ContainerSetStatus defines the observed state of ContainerSet
type ContainerSetStatus struct {
	// HealthyReplicas the number of active replicas
	HealthyReplicas int32 `json:"healthyReplicas,omitEmpty"`
}

// +kubebuilder:object:root=true

// ContainerSet is the Schema for the containersets API
// +kubebuilder:printcolumn:name="Replicas",type=integer,JSONPath=`.spec.replicas`
// +kubebuilder:printcolumn:name="Image",type=string,JSONPath=`.spec.image`
!! YOUR_CODE !! Setup kubebuilder annotion for the css ContainerSet short name.
type ContainerSet struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContainerSetSpec   `json:"spec,omitempty"`
	Status ContainerSetStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContainerSetList contains a list of ContainerSet
type ContainerSetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContainerSet `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ContainerSet{}, &ContainerSetList{})
}
