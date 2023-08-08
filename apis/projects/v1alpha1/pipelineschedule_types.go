/*
Copyright 2023 The Crossplane Authors.

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

// PipelineScheduleParameters represents a pipeline schedule.
//
// GitLab API docs:
// https://docs.gitlab.com/ee/api/pipeline_schedules.html
// At least 1 of [ProjectID, ProjectIDRef, ProjectIDSelector] required.
type PipelineScheduleParameters struct {
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

	// Description is a description of the pipeline schedule.
	// +required
	Description string `json:"description"`

	// Ref is the branch or tag name that is triggered.
	// +required
	Ref string `json:"ref"`

	// Cron is the cron schedule, for example: 0 1 * * *.
	// +required
	Cron string `json:"cron"`

	// CronTimezone is the time zone supported by ActiveSupport::TimeZone,
	// for example: Pacific Time (US & Canada) (default: UTC).
	// +optional
	CronTimezone *string `json:"cronTimezone,omitempty"`

	// Active is the activation of pipeline schedule.
	// If false is set, the pipeline schedule is initially deactivated (default: true).
	// +optional
	Active *bool `json:"active,omitempty"`

	// PipelineVariables is a type of environment variable.
	Variables []PipelineVariable `json:"variables,omitempty"`
}

// PipelineScheduleObservation represents observed stated of Gitlab Pipeline Schedule.
// https://docs.gitlab.com/ee/api/pipeline_schedules.htm
type PipelineScheduleObservation struct {
	ID           *int          `json:"id,omitempty"`
	NextRunAt    *metav1.Time  `json:"nextRunAt,omitempty"`
	CreatedAt    *metav1.Time  `json:"createdAt,omitempty"`
	UpdatedAt    *metav1.Time  `json:"updatedAt,omitempty"`
	Owner        *User         `json:"owner,omitempty"`
	LastPipeline *LastPipeline `json:"lastPipeline,omitempty"`
}

// LastPipeline represents the last pipeline ran by schedule
// this will be returned only for individual schedule get operation
type LastPipeline struct {
	ID     int    `json:"id"`
	SHA    string `json:"sha"`
	Ref    string `json:"ref"`
	Status string `json:"status"`
}

// PipelineVariable represents a pipeline variable.
//
// GitLab API docs: https://docs.gitlab.com/ee/api/pipelines.html
type PipelineVariable struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	// +optional
	VariableType *string `json:"variableType,omitempty"`
}

// PipelineScheduleSpec defines desired state of Gitlab Pipeline Schedule.
type PipelineScheduleSpec struct {
	xpv1.ResourceSpec `json:","`
	ForProvider       PipelineScheduleParameters `json:"forProvider"`
}

// PipelineScheduleStatus represents observed state of Gitlab Pipeline Schedule.
type PipelineScheduleStatus struct {
	xpv1.ResourceStatus `json:","`
	AtProvider          PipelineScheduleObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// A PipelineSchedule is a managed resource that represents a Gitlab Pipeline Schedule.
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,gitlab}
type PipelineSchedule struct {
	metav1.TypeMeta   `json:","`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PipelineScheduleSpec   `json:"spec"`
	Status PipelineScheduleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PipelineScheduleList contains a list of Pipeline Schedule items.
type PipelineScheduleList struct {
	metav1.TypeMeta `json:","`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []PipelineSchedule `json:"items"`
}
