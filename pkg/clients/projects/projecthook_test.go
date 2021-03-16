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

package projects

import (
	"testing"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/google/go-cmp/cmp"
	"github.com/xanzy/go-gitlab"

	"github.com/crossplane-contrib/provider-gitlab/apis/projects/v1alpha1"
)

var (
	url                      = "https://my-project.example.com"
	confidentialNoteEvents   = true
	pushEvents               = true
	pushEventsBranchFilter   = "foo"
	issuesEvents             = true
	confidentialIssuesEvents = true
	mergeRequestsEvents      = true
	tagPushEvents            = true
	noteEvents               = true
	jobEvents                = true
	pipelineEvents           = true
	wikiPageEvents           = true
	enableSSLVerification    = true
	token                    = "84B9C651-9025-47D2-9124-DD951BD268E8"
)

func TestGenerateProjectHookObservation(t *testing.T) {
	id := 0
	createdAt := time.Now()

	type args struct {
		ph *gitlab.ProjectHook
	}

	cases := map[string]struct {
		args args
		want v1alpha1.ProjectHookObservation
	}{
		"Full": {
			args: args{
				ph: &gitlab.ProjectHook{
					ID:        id,
					CreatedAt: &createdAt,
				},
			},
			want: v1alpha1.ProjectHookObservation{
				ID:        id,
				CreatedAt: &metav1.Time{Time: createdAt},
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := GenerateProjectHookObservation(tc.args.ph)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
		})
	}
}
func TestLateInitializeProjectHook(t *testing.T) {
	cases := map[string]struct {
		parameters  *v1alpha1.ProjectHookParameters
		projecthook *gitlab.ProjectHook
		want        *v1alpha1.ProjectHookParameters
	}{
		"AllOptionalFields": {
			parameters: &v1alpha1.ProjectHookParameters{},
			projecthook: &gitlab.ProjectHook{
				ConfidentialNoteEvents:   confidentialNoteEvents,
				PushEvents:               pushEvents,
				PushEventsBranchFilter:   pushEventsBranchFilter,
				IssuesEvents:             issuesEvents,
				ConfidentialIssuesEvents: confidentialIssuesEvents,
				MergeRequestsEvents:      mergeRequestsEvents,
				TagPushEvents:            tagPushEvents,
				NoteEvents:               noteEvents,
				JobEvents:                jobEvents,
				PipelineEvents:           pipelineEvents,
				WikiPageEvents:           wikiPageEvents,
				EnableSSLVerification:    enableSSLVerification,
			},
			want: &v1alpha1.ProjectHookParameters{
				ConfidentialNoteEvents:   &confidentialNoteEvents,
				PushEvents:               &pushEvents,
				PushEventsBranchFilter:   &pushEventsBranchFilter,
				IssuesEvents:             &issuesEvents,
				ConfidentialIssuesEvents: &confidentialIssuesEvents,
				MergeRequestsEvents:      &mergeRequestsEvents,
				TagPushEvents:            &tagPushEvents,
				NoteEvents:               &noteEvents,
				JobEvents:                &jobEvents,
				PipelineEvents:           &pipelineEvents,
				WikiPageEvents:           &wikiPageEvents,
				EnableSSLVerification:    &enableSSLVerification,
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			LateInitializeProjectHook(tc.parameters, tc.projecthook)
			if diff := cmp.Diff(tc.want, tc.parameters); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
		})
	}
}
func TestGenerateCreateProjectHookOptions(t *testing.T) {
	type args struct {
		parameters *v1alpha1.ProjectHookParameters
	}
	cases := map[string]struct {
		args args
		want *gitlab.AddProjectHookOptions
	}{
		"AllFields": {
			args: args{
				parameters: &v1alpha1.ProjectHookParameters{
					URL:                      &url,
					ConfidentialNoteEvents:   &confidentialNoteEvents,
					PushEvents:               &pushEvents,
					PushEventsBranchFilter:   &pushEventsBranchFilter,
					IssuesEvents:             &issuesEvents,
					ConfidentialIssuesEvents: &confidentialIssuesEvents,
					MergeRequestsEvents:      &mergeRequestsEvents,
					TagPushEvents:            &tagPushEvents,
					NoteEvents:               &noteEvents,
					JobEvents:                &jobEvents,
					PipelineEvents:           &pipelineEvents,
					WikiPageEvents:           &wikiPageEvents,
					EnableSSLVerification:    &enableSSLVerification,
					Token:                    &token,
				},
			},
			want: &gitlab.AddProjectHookOptions{
				URL:                      &url,
				ConfidentialNoteEvents:   &confidentialNoteEvents,
				PushEvents:               &pushEvents,
				PushEventsBranchFilter:   &pushEventsBranchFilter,
				IssuesEvents:             &issuesEvents,
				ConfidentialIssuesEvents: &confidentialIssuesEvents,
				MergeRequestsEvents:      &mergeRequestsEvents,
				TagPushEvents:            &tagPushEvents,
				NoteEvents:               &noteEvents,
				JobEvents:                &jobEvents,
				PipelineEvents:           &pipelineEvents,
				WikiPageEvents:           &wikiPageEvents,
				EnableSSLVerification:    &enableSSLVerification,
				Token:                    &token,
			},
		},
		"SomeFields": {
			args: args{
				parameters: &v1alpha1.ProjectHookParameters{
					PushEvents:             &pushEvents,
					PushEventsBranchFilter: &pushEventsBranchFilter,
					IssuesEvents:           &issuesEvents,
				},
			},
			want: &gitlab.AddProjectHookOptions{
				PushEvents:             &pushEvents,
				PushEventsBranchFilter: &pushEventsBranchFilter,
				IssuesEvents:           &issuesEvents,
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := GenerateCreateProjectHookOptions(tc.args.parameters)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
		})
	}
}
func TestGenerateEditProjectHookOptions(t *testing.T) {
	type args struct {
		parameters *v1alpha1.ProjectHookParameters
	}
	cases := map[string]struct {
		args args
		want *gitlab.EditProjectHookOptions
	}{
		"AllFields": {
			args: args{
				parameters: &v1alpha1.ProjectHookParameters{
					URL:                      &url,
					ConfidentialNoteEvents:   &confidentialNoteEvents,
					PushEvents:               &pushEvents,
					PushEventsBranchFilter:   &pushEventsBranchFilter,
					IssuesEvents:             &issuesEvents,
					ConfidentialIssuesEvents: &confidentialIssuesEvents,
					MergeRequestsEvents:      &mergeRequestsEvents,
					TagPushEvents:            &tagPushEvents,
					NoteEvents:               &noteEvents,
					JobEvents:                &jobEvents,
					PipelineEvents:           &pipelineEvents,
					WikiPageEvents:           &wikiPageEvents,
					EnableSSLVerification:    &enableSSLVerification,
					Token:                    &token,
				},
			},
			want: &gitlab.EditProjectHookOptions{
				URL:                      &url,
				ConfidentialNoteEvents:   &confidentialNoteEvents,
				PushEvents:               &pushEvents,
				PushEventsBranchFilter:   &pushEventsBranchFilter,
				IssuesEvents:             &issuesEvents,
				ConfidentialIssuesEvents: &confidentialIssuesEvents,
				MergeRequestsEvents:      &mergeRequestsEvents,
				TagPushEvents:            &tagPushEvents,
				NoteEvents:               &noteEvents,
				JobEvents:                &jobEvents,
				PipelineEvents:           &pipelineEvents,
				WikiPageEvents:           &wikiPageEvents,
				EnableSSLVerification:    &enableSSLVerification,
				Token:                    &token,
			},
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := GenerateEditProjectHookOptions(tc.args.parameters)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
		})
	}
}
func TestIsProjectHookUpToDate(t *testing.T) {
	type args struct {
		projecthook *gitlab.ProjectHook
		p           *v1alpha1.ProjectHookParameters
	}

	cases := map[string]struct {
		args args
		want bool
	}{
		"SameFields": {
			args: args{
				p: &v1alpha1.ProjectHookParameters{
					URL:                      &url,
					ConfidentialNoteEvents:   &confidentialNoteEvents,
					PushEvents:               &pushEvents,
					PushEventsBranchFilter:   &pushEventsBranchFilter,
					IssuesEvents:             &issuesEvents,
					ConfidentialIssuesEvents: &confidentialIssuesEvents,
					MergeRequestsEvents:      &mergeRequestsEvents,
					TagPushEvents:            &tagPushEvents,
					NoteEvents:               &noteEvents,
					JobEvents:                &jobEvents,
					PipelineEvents:           &pipelineEvents,
					WikiPageEvents:           &wikiPageEvents,
					EnableSSLVerification:    &enableSSLVerification,
					Token:                    &token,
				},
				projecthook: &gitlab.ProjectHook{
					URL:                      url,
					ConfidentialNoteEvents:   confidentialNoteEvents,
					PushEvents:               pushEvents,
					PushEventsBranchFilter:   pushEventsBranchFilter,
					IssuesEvents:             issuesEvents,
					ConfidentialIssuesEvents: confidentialIssuesEvents,
					MergeRequestsEvents:      mergeRequestsEvents,
					TagPushEvents:            tagPushEvents,
					NoteEvents:               noteEvents,
					JobEvents:                jobEvents,
					PipelineEvents:           pipelineEvents,
					WikiPageEvents:           wikiPageEvents,
					EnableSSLVerification:    enableSSLVerification,
				},
			},
			want: true,
		},
		"DifferentFields": {
			args: args{
				p: &v1alpha1.ProjectHookParameters{
					URL:                      &url,
					ConfidentialNoteEvents:   &confidentialNoteEvents,
					PushEvents:               &pushEvents,
					PushEventsBranchFilter:   &pushEventsBranchFilter,
					IssuesEvents:             &issuesEvents,
					ConfidentialIssuesEvents: &confidentialIssuesEvents,
					MergeRequestsEvents:      &mergeRequestsEvents,
					TagPushEvents:            &tagPushEvents,
					NoteEvents:               &noteEvents,
					JobEvents:                &jobEvents,
					PipelineEvents:           &pipelineEvents,
					WikiPageEvents:           &wikiPageEvents,
					EnableSSLVerification:    &enableSSLVerification,
					Token:                    &token,
				},
				projecthook: &gitlab.ProjectHook{
					URL:                      "http://some.other.url",
					ConfidentialNoteEvents:   false,
					PushEvents:               false,
					PushEventsBranchFilter:   "bar",
					IssuesEvents:             false,
					ConfidentialIssuesEvents: false,
					MergeRequestsEvents:      false,
					TagPushEvents:            false,
					NoteEvents:               false,
					JobEvents:                false,
					PipelineEvents:           false,
					WikiPageEvents:           false,
					EnableSSLVerification:    false,
				},
			},
			want: false,
		},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			got := IsProjectHookUpToDate(tc.args.p, tc.args.projecthook)
			if diff := cmp.Diff(tc.want, got); diff != "" {
				t.Errorf("r: -want, +got:\n%s", diff)
			}
		})
	}

}
