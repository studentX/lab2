/*
Copyright 2020 K8sland Labs.

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

// PainterSpec defines the desired state of Painter
type PainterSpec struct {
	!!YOUR_CODE!! Add an annotation to ensure the color can only be Red, Blue, Green
	!!YOUR_CODE!! Add color field
}

// PainterStatus defines the observed state of Painter
type PainterStatus struct {
	!!YOUR_CODE!! Track painted pods count
}

// +kubebuilder:object:root=true
!!YOUR_CODE!! Add an annotation to specify namespaced and shortname
// +kubebuilder:subresource:status

// Painter is the Schema for the painters API
type Painter struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PainterSpec   `json:"spec,omitempty"`
	Status PainterStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PainterList contains a list of Painter
type PainterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Painter `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Painter{}, &PainterList{})
}
