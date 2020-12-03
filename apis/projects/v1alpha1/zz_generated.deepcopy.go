// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomAttribute) DeepCopyInto(out *CustomAttribute) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomAttribute.
func (in *CustomAttribute) DeepCopy() *CustomAttribute {
	if in == nil {
		return nil
	}
	out := new(CustomAttribute)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ForkParent) DeepCopyInto(out *ForkParent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ForkParent.
func (in *ForkParent) DeepCopy() *ForkParent {
	if in == nil {
		return nil
	}
	out := new(ForkParent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GroupAccess) DeepCopyInto(out *GroupAccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GroupAccess.
func (in *GroupAccess) DeepCopy() *GroupAccess {
	if in == nil {
		return nil
	}
	out := new(GroupAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Links) DeepCopyInto(out *Links) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Links.
func (in *Links) DeepCopy() *Links {
	if in == nil {
		return nil
	}
	out := new(Links)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Permissions) DeepCopyInto(out *Permissions) {
	*out = *in
	if in.ProjectAccess != nil {
		in, out := &in.ProjectAccess, &out.ProjectAccess
		*out = new(ProjectAccess)
		**out = **in
	}
	if in.GroupAccess != nil {
		in, out := &in.GroupAccess, &out.GroupAccess
		*out = new(GroupAccess)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Permissions.
func (in *Permissions) DeepCopy() *Permissions {
	if in == nil {
		return nil
	}
	out := new(Permissions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Project) DeepCopyInto(out *Project) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Project.
func (in *Project) DeepCopy() *Project {
	if in == nil {
		return nil
	}
	out := new(Project)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Project) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectAccess) DeepCopyInto(out *ProjectAccess) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectAccess.
func (in *ProjectAccess) DeepCopy() *ProjectAccess {
	if in == nil {
		return nil
	}
	out := new(ProjectAccess)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectList) DeepCopyInto(out *ProjectList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Project, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectList.
func (in *ProjectList) DeepCopy() *ProjectList {
	if in == nil {
		return nil
	}
	out := new(ProjectList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ProjectList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectNamespace) DeepCopyInto(out *ProjectNamespace) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectNamespace.
func (in *ProjectNamespace) DeepCopy() *ProjectNamespace {
	if in == nil {
		return nil
	}
	out := new(ProjectNamespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectObservation) DeepCopyInto(out *ProjectObservation) {
	*out = *in
	if in.Owner != nil {
		in, out := &in.Owner, &out.Owner
		*out = new(User)
		(*in).DeepCopyInto(*out)
	}
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.LastActivityAt != nil {
		in, out := &in.LastActivityAt, &out.LastActivityAt
		*out = (*in).DeepCopy()
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(ProjectNamespace)
		**out = **in
	}
	if in.Permissions != nil {
		in, out := &in.Permissions, &out.Permissions
		*out = new(Permissions)
		(*in).DeepCopyInto(*out)
	}
	if in.MarkedForDeletionAt != nil {
		in, out := &in.MarkedForDeletionAt, &out.MarkedForDeletionAt
		*out = (*in).DeepCopy()
	}
	if in.ForkedFromProject != nil {
		in, out := &in.ForkedFromProject, &out.ForkedFromProject
		*out = new(ForkParent)
		**out = **in
	}
	if in.SharedWithGroups != nil {
		in, out := &in.SharedWithGroups, &out.SharedWithGroups
		*out = make([]SharedWithGroups, len(*in))
		copy(*out, *in)
	}
	if in.Statistics != nil {
		in, out := &in.Statistics, &out.Statistics
		*out = new(ProjectStatistics)
		**out = **in
	}
	if in.Links != nil {
		in, out := &in.Links, &out.Links
		*out = new(Links)
		**out = **in
	}
	if in.CustomAttributes != nil {
		in, out := &in.CustomAttributes, &out.CustomAttributes
		*out = make([]CustomAttribute, len(*in))
		copy(*out, *in)
	}
	if in.ComplianceFrameworks != nil {
		in, out := &in.ComplianceFrameworks, &out.ComplianceFrameworks
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectObservation.
func (in *ProjectObservation) DeepCopy() *ProjectObservation {
	if in == nil {
		return nil
	}
	out := new(ProjectObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectParameters) DeepCopyInto(out *ProjectParameters) {
	*out = *in
	if in.Path != nil {
		in, out := &in.Path, &out.Path
		*out = new(string)
		**out = **in
	}
	if in.NamespaceID != nil {
		in, out := &in.NamespaceID, &out.NamespaceID
		*out = new(int)
		**out = **in
	}
	if in.DefaultBranch != nil {
		in, out := &in.DefaultBranch, &out.DefaultBranch
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.IssuesAccessLevel != nil {
		in, out := &in.IssuesAccessLevel, &out.IssuesAccessLevel
		*out = new(AccessControlValue)
		**out = **in
	}
	if in.RepositoryAccessLevel != nil {
		in, out := &in.RepositoryAccessLevel, &out.RepositoryAccessLevel
		*out = new(AccessControlValue)
		**out = **in
	}
	if in.MergeRequestsAccessLevel != nil {
		in, out := &in.MergeRequestsAccessLevel, &out.MergeRequestsAccessLevel
		*out = new(AccessControlValue)
		**out = **in
	}
	if in.ForkingAccessLevel != nil {
		in, out := &in.ForkingAccessLevel, &out.ForkingAccessLevel
		*out = new(AccessControlValue)
		**out = **in
	}
	if in.BuildsAccessLevel != nil {
		in, out := &in.BuildsAccessLevel, &out.BuildsAccessLevel
		*out = new(AccessControlValue)
		**out = **in
	}
	if in.WikiAccessLevel != nil {
		in, out := &in.WikiAccessLevel, &out.WikiAccessLevel
		*out = new(AccessControlValue)
		**out = **in
	}
	if in.SnippetsAccessLevel != nil {
		in, out := &in.SnippetsAccessLevel, &out.SnippetsAccessLevel
		*out = new(AccessControlValue)
		**out = **in
	}
	if in.PagesAccessLevel != nil {
		in, out := &in.PagesAccessLevel, &out.PagesAccessLevel
		*out = new(AccessControlValue)
		**out = **in
	}
	if in.EmailsDisabled != nil {
		in, out := &in.EmailsDisabled, &out.EmailsDisabled
		*out = new(bool)
		**out = **in
	}
	if in.ResolveOutdatedDiffDiscussions != nil {
		in, out := &in.ResolveOutdatedDiffDiscussions, &out.ResolveOutdatedDiffDiscussions
		*out = new(bool)
		**out = **in
	}
	if in.ContainerRegistryEnabled != nil {
		in, out := &in.ContainerRegistryEnabled, &out.ContainerRegistryEnabled
		*out = new(bool)
		**out = **in
	}
	if in.SharedRunnersEnabled != nil {
		in, out := &in.SharedRunnersEnabled, &out.SharedRunnersEnabled
		*out = new(bool)
		**out = **in
	}
	if in.Visibility != nil {
		in, out := &in.Visibility, &out.Visibility
		*out = new(VisibilityValue)
		**out = **in
	}
	if in.ImportURL != nil {
		in, out := &in.ImportURL, &out.ImportURL
		*out = new(string)
		**out = **in
	}
	if in.PublicBuilds != nil {
		in, out := &in.PublicBuilds, &out.PublicBuilds
		*out = new(bool)
		**out = **in
	}
	if in.OnlyAllowMergeIfPipelineSucceeds != nil {
		in, out := &in.OnlyAllowMergeIfPipelineSucceeds, &out.OnlyAllowMergeIfPipelineSucceeds
		*out = new(bool)
		**out = **in
	}
	if in.OnlyAllowMergeIfAllDiscussionsAreResolved != nil {
		in, out := &in.OnlyAllowMergeIfAllDiscussionsAreResolved, &out.OnlyAllowMergeIfAllDiscussionsAreResolved
		*out = new(bool)
		**out = **in
	}
	if in.MergeMethod != nil {
		in, out := &in.MergeMethod, &out.MergeMethod
		*out = new(MergeMethodValue)
		**out = **in
	}
	if in.RemoveSourceBranchAfterMerge != nil {
		in, out := &in.RemoveSourceBranchAfterMerge, &out.RemoveSourceBranchAfterMerge
		*out = new(bool)
		**out = **in
	}
	if in.LFSEnabled != nil {
		in, out := &in.LFSEnabled, &out.LFSEnabled
		*out = new(bool)
		**out = **in
	}
	if in.RequestAccessEnabled != nil {
		in, out := &in.RequestAccessEnabled, &out.RequestAccessEnabled
		*out = new(bool)
		**out = **in
	}
	if in.TagList != nil {
		in, out := &in.TagList, &out.TagList
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PrintingMergeRequestLinkEnabled != nil {
		in, out := &in.PrintingMergeRequestLinkEnabled, &out.PrintingMergeRequestLinkEnabled
		*out = new(bool)
		**out = **in
	}
	if in.BuildGitStrategy != nil {
		in, out := &in.BuildGitStrategy, &out.BuildGitStrategy
		*out = new(string)
		**out = **in
	}
	if in.BuildTimeout != nil {
		in, out := &in.BuildTimeout, &out.BuildTimeout
		*out = new(int)
		**out = **in
	}
	if in.AutoCancelPendingPipelines != nil {
		in, out := &in.AutoCancelPendingPipelines, &out.AutoCancelPendingPipelines
		*out = new(string)
		**out = **in
	}
	if in.BuildCoverageRegex != nil {
		in, out := &in.BuildCoverageRegex, &out.BuildCoverageRegex
		*out = new(string)
		**out = **in
	}
	if in.CIConfigPath != nil {
		in, out := &in.CIConfigPath, &out.CIConfigPath
		*out = new(string)
		**out = **in
	}
	if in.AutoDevopsEnabled != nil {
		in, out := &in.AutoDevopsEnabled, &out.AutoDevopsEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AutoDevopsDeployStrategy != nil {
		in, out := &in.AutoDevopsDeployStrategy, &out.AutoDevopsDeployStrategy
		*out = new(string)
		**out = **in
	}
	if in.ApprovalsBeforeMerge != nil {
		in, out := &in.ApprovalsBeforeMerge, &out.ApprovalsBeforeMerge
		*out = new(int)
		**out = **in
	}
	if in.ExternalAuthorizationClassificationLabel != nil {
		in, out := &in.ExternalAuthorizationClassificationLabel, &out.ExternalAuthorizationClassificationLabel
		*out = new(string)
		**out = **in
	}
	if in.Mirror != nil {
		in, out := &in.Mirror, &out.Mirror
		*out = new(bool)
		**out = **in
	}
	if in.MirrorTriggerBuilds != nil {
		in, out := &in.MirrorTriggerBuilds, &out.MirrorTriggerBuilds
		*out = new(bool)
		**out = **in
	}
	if in.InitializeWithReadme != nil {
		in, out := &in.InitializeWithReadme, &out.InitializeWithReadme
		*out = new(bool)
		**out = **in
	}
	if in.TemplateName != nil {
		in, out := &in.TemplateName, &out.TemplateName
		*out = new(string)
		**out = **in
	}
	if in.TemplateProjectID != nil {
		in, out := &in.TemplateProjectID, &out.TemplateProjectID
		*out = new(int)
		**out = **in
	}
	if in.UseCustomTemplate != nil {
		in, out := &in.UseCustomTemplate, &out.UseCustomTemplate
		*out = new(bool)
		**out = **in
	}
	if in.GroupWithProjectTemplatesID != nil {
		in, out := &in.GroupWithProjectTemplatesID, &out.GroupWithProjectTemplatesID
		*out = new(int)
		**out = **in
	}
	if in.PackagesEnabled != nil {
		in, out := &in.PackagesEnabled, &out.PackagesEnabled
		*out = new(bool)
		**out = **in
	}
	if in.ServiceDeskEnabled != nil {
		in, out := &in.ServiceDeskEnabled, &out.ServiceDeskEnabled
		*out = new(bool)
		**out = **in
	}
	if in.AutocloseReferencedIssues != nil {
		in, out := &in.AutocloseReferencedIssues, &out.AutocloseReferencedIssues
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectParameters.
func (in *ProjectParameters) DeepCopy() *ProjectParameters {
	if in == nil {
		return nil
	}
	out := new(ProjectParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectSpec) DeepCopyInto(out *ProjectSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectSpec.
func (in *ProjectSpec) DeepCopy() *ProjectSpec {
	if in == nil {
		return nil
	}
	out := new(ProjectSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectStatistics) DeepCopyInto(out *ProjectStatistics) {
	*out = *in
	out.StorageStatistics = in.StorageStatistics
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectStatistics.
func (in *ProjectStatistics) DeepCopy() *ProjectStatistics {
	if in == nil {
		return nil
	}
	out := new(ProjectStatistics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProjectStatus) DeepCopyInto(out *ProjectStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProjectStatus.
func (in *ProjectStatus) DeepCopy() *ProjectStatus {
	if in == nil {
		return nil
	}
	out := new(ProjectStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SharedWithGroups) DeepCopyInto(out *SharedWithGroups) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SharedWithGroups.
func (in *SharedWithGroups) DeepCopy() *SharedWithGroups {
	if in == nil {
		return nil
	}
	out := new(SharedWithGroups)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageStatistics) DeepCopyInto(out *StorageStatistics) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageStatistics.
func (in *StorageStatistics) DeepCopy() *StorageStatistics {
	if in == nil {
		return nil
	}
	out := new(StorageStatistics)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *User) DeepCopyInto(out *User) {
	*out = *in
	if in.CreatedAt != nil {
		in, out := &in.CreatedAt, &out.CreatedAt
		*out = (*in).DeepCopy()
	}
	if in.LastActivityOn != nil {
		in, out := &in.LastActivityOn, &out.LastActivityOn
		*out = (*in).DeepCopy()
	}
	if in.CurrentSignInAt != nil {
		in, out := &in.CurrentSignInAt, &out.CurrentSignInAt
		*out = (*in).DeepCopy()
	}
	if in.LastSignInAt != nil {
		in, out := &in.LastSignInAt, &out.LastSignInAt
		*out = (*in).DeepCopy()
	}
	if in.ConfirmedAt != nil {
		in, out := &in.ConfirmedAt, &out.ConfirmedAt
		*out = (*in).DeepCopy()
	}
	if in.Identities != nil {
		in, out := &in.Identities, &out.Identities
		*out = make([]*UserIdentity, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(UserIdentity)
				**out = **in
			}
		}
	}
	if in.CustomAttributes != nil {
		in, out := &in.CustomAttributes, &out.CustomAttributes
		*out = make([]*CustomAttribute, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(CustomAttribute)
				**out = **in
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new User.
func (in *User) DeepCopy() *User {
	if in == nil {
		return nil
	}
	out := new(User)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserIdentity) DeepCopyInto(out *UserIdentity) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserIdentity.
func (in *UserIdentity) DeepCopy() *UserIdentity {
	if in == nil {
		return nil
	}
	out := new(UserIdentity)
	in.DeepCopyInto(out)
	return out
}
