/*
Copyright 2020 K8sland Training.

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

// ClusterDepotSpec defines the desired state of ClusterDepot
type ClusterDepotSpec struct {
	// kubebuilder:validation:Enum=Red,Blue,Green
	Color string `json:"color,omitempty"`
}

// ClusterDepotStatus defines the observed state of ClusterDepot
type ClusterDepotStatus struct {
	PaintedPods int32 `json:"paintedPods,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=cd
// +kubebuilder:subresource:status

// ClusterDepot is the Schema for the clusterdepots API
type ClusterDepot struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ClusterDepotSpec   `json:"spec,omitempty"`
	Status ClusterDepotStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ClusterDepotList contains a list of ClusterDepot
type ClusterDepotList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ClusterDepot `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ClusterDepot{}, &ClusterDepotList{})
}
