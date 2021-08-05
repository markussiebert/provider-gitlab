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

// DeployTokenParameters define the desired state of a Gitlab deploy token
// https://docs.gitlab.com/ee/api/deploy_tokens.html
type DeployTokenParameters struct {
	// ProjectID is the ID of the project to create the deploy token in.
	// +optional
	// +immutable
	ProjectID *int `json:"projectId,omitempty"`

	// ProjectIDRef is a reference to a project to retrieve its projectId
	// +optional
	// +immutable
	ProjectIDRef *xpv1.Reference `json:"projectIdRef,omitempty"`

	// ProjectIDSelector selects reference to a project to retrieve its projectId.
	// +optional
	ProjectIDSelector *xpv1.Selector `json:"projectIdSelector,omitempty"`

	// Expiration date for the deploy token. Does not expire if no value is provided.
	// Expected in ISO 8601 format (2019-03-15T08:00:00Z)
	// +optional
	// +immutable
	ExpiresAt *metav1.Time `json:"expiresAt,omitempty"`

	// Username for deploy token. Default is gitlab+deploy-token-{n}
	// +optional
	// +immutable
	Username *string `json:"username,omitempty"`

	// Scopes indicates the deploy token scopes.
	// Must be at least one of read_repository, read_registry, write_registry,
	// read_package_registry, or write_package_registry.
	// +immutable
	Scopes []string `json:"scopes"`
}

// DeployTokenObservation represents a deploy token.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/deploy_tokens.html
type DeployTokenObservation struct{}

// A DeployTokenSpec defines the desired state of a Gitlab Project.
type DeployTokenSpec struct {
	xpv1.ResourceSpec `json:",inline"`
	ForProvider       DeployTokenParameters `json:"forProvider"`
}

// A DeployTokenStatus represents the observed state of a Gitlab Project.
type DeployTokenStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DeployTokenObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A DeployToken is a managed resource that represents a Gitlab deploy token
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gitlab}
type DeployToken struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeployTokenSpec   `json:"spec"`
	Status DeployTokenStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeployTokenList contains a list of Project items
type DeployTokenList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeployToken `json:"items"`
}
