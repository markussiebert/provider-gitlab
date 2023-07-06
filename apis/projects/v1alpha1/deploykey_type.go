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
	xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// DeployKeyParameters define desired state of Gitlab Deploy Key.
// https://docs.gitlab.com/ee/api/deploy_keys.html
// At least 1 of [ProjectID, ProjectIDRef, ProjectIDSelector] required.
type DeployKeyParameters struct {
	// The ID or URL-encoded path of the project owned by the authenticated user.
	// +optional
	// +immutable
	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-gitlab/apis/projects/v1alpha1.Project
	// +crossplane:generate:reference:refFieldName=ProjectIDRef
	// +crossplane:generate:reference:selectorFieldName=ProjectIDSelector
	ProjectID *string `json:"projectId,omitempty"`

	// ProjectIDRef is a reference to a project to retrieve its ProjectID.
	// +optional
	// +immutable
	ProjectIDRef *xpv1.Reference `json:"projectIdRef,omitempty"`

	// ProjectIDSelector selects reference to a project to retrieve its ProjectID.
	// +optional
	// +immutable
	ProjectIDSelector *xpv1.Selector `json:"projectIdSelector,omitempty"`

	// New Deploy Key’s title.
	// This property is required.
	Title string `json:"title"`

	// Can Deploy Key push to the project’s repository.
	// +optional
	CanPush *bool `json:"canPush,omitempty"`

	// Expiration date for the Deploy Key. Does not expire if no value is provided.
	// Expected in ISO 8601 format (2019-03-15T08:00:00Z).
	// +optional
	ExpiresAt *metav1.Time `json:"expiresAt,omitempty"`

	// KeySecretRef field representing reference to the key.
	// This property is required.
	KeySecretRef xpv1.SecretKeySelector `json:"keySecretRef"`
}

// DeployKeyObservation represents observed stated of Deploy Key.
// https://docs.gitlab.com/ee/api/deploy_keys.html
type DeployKeyObservation struct {
	ID        *int         `json:"id,omitempty"`
	CreatedAt *metav1.Time `json:"createdAt,omitempty"`
}

// DeployKeySpec defines desired state of Gitlab Deploy Key.
type DeployKeySpec struct {
	xpv1.ResourceSpec `json:","`
	ForProvider       DeployKeyParameters `json:"forProvider"`
}

// DeployKeyStatus represents observed state of Gitlab Deploy Key.
type DeployKeyStatus struct {
	xpv1.ResourceStatus `json:",inline"`
	AtProvider          DeployKeyObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A DeployKey is a managed resource that represents a Gitlab Deploy Key.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gitlab}
type DeployKey struct {
	metav1.TypeMeta   `json:","`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   DeployKeySpec   `json:"spec"`
	Status DeployKeyStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeployKeyList contains a list of Deploy Key items.
type DeployKeyList struct {
	metav1.TypeMeta `json:","`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeployKey `json:"items"`
}
