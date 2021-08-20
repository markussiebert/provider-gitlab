/*
Copyright 2021 The Crossplane Authors.

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

	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

// VariableType indicates the type of the GitLab CI variable.
type VariableType string

// List of variable type values.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/project_level_variables.html
const (
	VariableTypeEnvVar VariableType = "env_var"
	VariableTypeFile   VariableType = "file"
)

// VariableParameters define the desired state of a Gitlab CI Variable
// https://docs.gitlab.com/ee/api/project_level_variables.html
type VariableParameters struct {
	// ProjectID is the ID of the project to create the variable on.
	// +optional
	// +immutable
	ProjectID *int `json:"projectId,omitempty"`

	// ProjectIDRef is a reference to a project to retrieve its projectId.
	// +optional
	// +immutable
	ProjectIDRef *xpv1.Reference `json:"projectIdRef,omitempty"`

	// ProjectIDSelector selects reference to a project to retrieve its projectId.
	// +optional
	ProjectIDSelector *xpv1.Selector `json:"projectIdSelector,omitempty"`

	// Key for the variable.
	// +kubebuilder:validation:Pattern:=^[a-zA-Z0-9\_]+$
	// +kubebuilder:validation:MaxLength:=255
	// +immutable
	Key string `json:"key"`

	// Value for the variable.
	Value string `json:"value"`

	// Masked enables or disables variable masking.
	// +optional
	Masked *bool `json:"masked,omitempty"`

	// Protected enables or disables variable protection.
	// +optional
	Protected *bool `json:"protected,omitempty"`

	// VariableType is the type of the variable.
	// +kubebuilder:validation:Enum:=env_var;file
	// +optional
	VariableType *VariableType `json:"variableType,omitempty"`

	// EnvironmentScope indicates the environment scope
	// that this variable is applied to.
	// +optional
	EnvironmentScope *string `json:"environmentScope,omitempty"`
}

// A VariableSpec defines the desired state of a Gitlab Project CI
// Variable.
type VariableSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       VariableParameters `json:"forProvider"`
}

// A VariableStatus represents the observed state of a Gitlab Project CI
// Variable.
type VariableStatus struct {
	xpv1.ResourceStatus `json:",inline"`
}

// +kubebuilder:object:root=true

// A Variable is a managed resource that represents a Gitlab CI variable.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gitlab}
type Variable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VariableSpec   `json:"spec"`
	Status VariableStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VariableList contains a list of Variable items.
type VariableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Variable `json:"items"`
}
