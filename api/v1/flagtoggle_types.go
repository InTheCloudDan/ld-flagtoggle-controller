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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// FlagToggleSpec defines the desired state of FlagToggle
type FlagToggleSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	// Feature flag key to identify it
	FlagKey string `json:"flagkey"`
	// desired state of the feature flag
	Enabled bool `json:"enabled"`
	// project key to target
	ProjectKey string `json:"projectkey"`
	// environment to run against
	EnvironmentKey []string `json:"environmentkey"`
	// Comment to go along with status changes
	Comment string `json:"comment"`
}

// FlagToggleStatus defines the observed state of FlagToggle
type FlagToggleStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// FlagToggle is the Schema for the flagtoggles API
type FlagToggle struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   FlagToggleSpec   `json:"spec,omitempty"`
	Status FlagToggleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FlagToggleList contains a list of FlagToggle
type FlagToggleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FlagToggle `json:"items"`
}

func init() {
	SchemeBuilder.Register(&FlagToggle{}, &FlagToggleList{})
}
