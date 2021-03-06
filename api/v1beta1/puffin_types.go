/*


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

package v1beta1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// ColorPhase determines whether a puffin has been assigned a color or not!
type ColorPhase string

const (
	// ColorPhasePending states that the puffin is yet to be colored
	ColorPhasePending ColorPhase = "PENDING"

	// ColorPhaseColored states that the puffin has been colored
	ColorPhaseColored ColorPhase = "COLORED"
)

// PuffinSpec defines the desired state of Puffin
type PuffinSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Color is an example field of Puffin. Edit Puffin_types.go to remove/update
	Color string `json:"color,omitempty"`
}

// PuffinStatus defines the observed state of Puffin
type PuffinStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Specifies the puffin's color status.
	// Valid values are:
	// - "PENDING" (default): The puffin is yet to be assigned a color;
	// - "COLORED": The puffin has been assigned a color;
	Message ColorPhase `json:"message,omitempty"`
}

// +kubebuilder:object:root=true

// Puffin is the Schema for the puffins API
type Puffin struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PuffinSpec   `json:"spec,omitempty"`
	Status PuffinStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PuffinList contains a list of Puffin
type PuffinList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Puffin `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Puffin{}, &PuffinList{})
}
