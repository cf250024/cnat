/*
Copyright 2021 cf250024.

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

const (
	PhasePending = "PENDING"
	PhaseRunning = "RUNNING"
	PhaseDone    = "DONE"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AtSpec defines the desired state of At
type AtSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Schedule is the desired time the command is supposed to be executed.
	// Note: the format used here is UTC time https://www.utctime.net
	Schedule string `json:"schedule,omitempty"`
	// Command is the desired command (executed in a Bash shell) to be executed.
	Command string `json:"command,omitempty"`
}

// AtStatus defines the observed state of At
type AtStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// Phase represents the state of the schedule: until the command is executed
	// it is PENDING, afterwards it is DONE.
	Phase string `json:"phase,omitempty"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// At is the Schema for the ats API
// +k8s:openapi-gen=true

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status

// At is the Schema for the ats API
type At struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AtSpec   `json:"spec,omitempty"`
	Status AtStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// AtList contains a list of At
type AtList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []At `json:"items"`
}

func init() {
	SchemeBuilder.Register(&At{}, &AtList{})
}
