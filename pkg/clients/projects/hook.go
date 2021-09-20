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
	"strings"

	"github.com/google/go-cmp/cmp"
	"github.com/xanzy/go-gitlab"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"github.com/crossplane-contrib/provider-gitlab/apis/projects/v1alpha1"
	"github.com/crossplane-contrib/provider-gitlab/pkg/clients"
)

const (
	errHookNotFound = "404 Not found"
)

// HookClient defines Gitlab Hook service operations
type HookClient interface {
	GetProjectHook(pid interface{}, hook int, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectHook, *gitlab.Response, error)
	AddProjectHook(pid interface{}, opt *gitlab.AddProjectHookOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectHook, *gitlab.Response, error)
	EditProjectHook(pid interface{}, hook int, opt *gitlab.EditProjectHookOptions, options ...gitlab.RequestOptionFunc) (*gitlab.ProjectHook, *gitlab.Response, error)
	DeleteProjectHook(pid interface{}, hook int, options ...gitlab.RequestOptionFunc) (*gitlab.Response, error)
}

// NewHookClient returns a new Gitlab Project service
func NewHookClient(cfg clients.Config) HookClient {
	git := clients.NewClient(cfg)
	return git.Projects
}

// IsErrorHookNotFound helper function to test for errProjectNotFound error.
func IsErrorHookNotFound(err error) bool {
	if err == nil {
		return false
	}
	return strings.Contains(err.Error(), errHookNotFound)
}

// LateInitializeHook fills the empty fields in the hook spec with the
// values seen in gitlab.Hook.
func LateInitializeHook(in *v1alpha1.HookParameters, hook *gitlab.ProjectHook) { // nolint:gocyclo
	if hook == nil {
		return
	}

	if in.ConfidentialNoteEvents == nil {
		in.ConfidentialNoteEvents = &hook.ConfidentialNoteEvents
	}
	if in.PushEvents == nil {
		in.PushEvents = &hook.PushEvents
	}
	if in.IssuesEvents == nil {
		in.IssuesEvents = &hook.IssuesEvents
	}
	in.PushEventsBranchFilter = clients.LateInitializeStringPtr(in.PushEventsBranchFilter, hook.PushEventsBranchFilter)
	if in.ConfidentialIssuesEvents == nil {
		in.ConfidentialIssuesEvents = &hook.ConfidentialIssuesEvents
	}
	if in.MergeRequestsEvents == nil {
		in.MergeRequestsEvents = &hook.MergeRequestsEvents
	}
	if in.TagPushEvents == nil {
		in.TagPushEvents = &hook.TagPushEvents
	}
	if in.NoteEvents == nil {
		in.NoteEvents = &hook.NoteEvents
	}
	if in.JobEvents == nil {
		in.JobEvents = &hook.JobEvents
	}
	if in.PipelineEvents == nil {
		in.PipelineEvents = &hook.PipelineEvents
	}
	if in.WikiPageEvents == nil {
		in.WikiPageEvents = &hook.WikiPageEvents
	}
	if in.EnableSSLVerification == nil {
		in.EnableSSLVerification = &hook.EnableSSLVerification
	}
}

// GenerateHookObservation is used to produce v1alpha1.HookObservation from
// gitlab.Hook.
func GenerateHookObservation(hook *gitlab.ProjectHook) v1alpha1.HookObservation { // nolint:gocyclo
	if hook == nil {
		return v1alpha1.HookObservation{}
	}

	o := v1alpha1.HookObservation{
		ID: hook.ID,
	}

	if hook.CreatedAt != nil {
		o.CreatedAt = &metav1.Time{Time: *hook.CreatedAt}
	}
	return o
}

// GenerateCreateHookOptions generates project creation options
func GenerateCreateHookOptions(p *v1alpha1.HookParameters) *gitlab.AddProjectHookOptions {
	hook := &gitlab.AddProjectHookOptions{
		URL:                      p.URL,
		ConfidentialNoteEvents:   p.ConfidentialNoteEvents,
		PushEvents:               p.PushEvents,
		PushEventsBranchFilter:   p.PushEventsBranchFilter,
		IssuesEvents:             p.IssuesEvents,
		ConfidentialIssuesEvents: p.ConfidentialIssuesEvents,
		MergeRequestsEvents:      p.MergeRequestsEvents,
		TagPushEvents:            p.TagPushEvents,
		NoteEvents:               p.NoteEvents,
		JobEvents:                p.JobEvents,
		PipelineEvents:           p.PipelineEvents,
		WikiPageEvents:           p.WikiPageEvents,
		EnableSSLVerification:    p.EnableSSLVerification,
		Token:                    p.Token,
	}

	return hook
}

// GenerateEditHookOptions generates project edit options
func GenerateEditHookOptions(p *v1alpha1.HookParameters) *gitlab.EditProjectHookOptions {
	o := &gitlab.EditProjectHookOptions{
		URL:                      p.URL,
		ConfidentialNoteEvents:   p.ConfidentialNoteEvents,
		PushEvents:               p.PushEvents,
		PushEventsBranchFilter:   p.PushEventsBranchFilter,
		IssuesEvents:             p.IssuesEvents,
		ConfidentialIssuesEvents: p.ConfidentialIssuesEvents,
		MergeRequestsEvents:      p.MergeRequestsEvents,
		TagPushEvents:            p.TagPushEvents,
		NoteEvents:               p.NoteEvents,
		JobEvents:                p.JobEvents,
		PipelineEvents:           p.PipelineEvents,
		WikiPageEvents:           p.WikiPageEvents,
		EnableSSLVerification:    p.EnableSSLVerification,
		Token:                    p.Token,
	}

	return o
}

// IsHookUpToDate checks whether there is a change in any of the modifiable fields.
func IsHookUpToDate(p *v1alpha1.HookParameters, g *gitlab.ProjectHook) bool { // nolint:gocyclo
	if !cmp.Equal(p.URL, clients.StringToPtr(g.URL)) {
		return false
	}
	if !clients.IsBoolEqualToBoolPtr(p.ConfidentialNoteEvents, g.ConfidentialNoteEvents) {
		return false
	}
	if !clients.IsBoolEqualToBoolPtr(p.PushEvents, g.PushEvents) {
		return false
	}
	if !cmp.Equal(p.PushEventsBranchFilter, clients.StringToPtr(g.PushEventsBranchFilter)) {
		return false
	}
	if !clients.IsBoolEqualToBoolPtr(p.IssuesEvents, g.IssuesEvents) {
		return false
	}
	if !clients.IsBoolEqualToBoolPtr(p.ConfidentialIssuesEvents, g.ConfidentialIssuesEvents) {
		return false
	}
	if !clients.IsBoolEqualToBoolPtr(p.MergeRequestsEvents, g.MergeRequestsEvents) {
		return false
	}
	if !clients.IsBoolEqualToBoolPtr(p.TagPushEvents, g.TagPushEvents) {
		return false
	}
	if !clients.IsBoolEqualToBoolPtr(p.NoteEvents, g.NoteEvents) {
		return false
	}
	if !clients.IsBoolEqualToBoolPtr(p.JobEvents, g.JobEvents) {
		return false
	}
	if !clients.IsBoolEqualToBoolPtr(p.PipelineEvents, g.PipelineEvents) {
		return false
	}
	if !clients.IsBoolEqualToBoolPtr(p.WikiPageEvents, g.WikiPageEvents) {
		return false
	}
	if !clients.IsBoolEqualToBoolPtr(p.EnableSSLVerification, g.EnableSSLVerification) {
		return false
	}

	return true
}
