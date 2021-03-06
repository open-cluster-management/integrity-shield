// +build !ignore_autogenerated

//
// Copyright 2020 IBM Corporation
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	common "github.com/IBM/integrity-enforcer/shield/pkg/common"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileStatusDetail) DeepCopyInto(out *ProfileStatusDetail) {
	*out = *in
	if in.Request != nil {
		in, out := &in.Request, &out.Request
		*out = new(common.Request)
		**out = **in
	}
	if in.Result != nil {
		in, out := &in.Result, &out.Result
		*out = (*in).DeepCopy()
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileStatusDetail.
func (in *ProfileStatusDetail) DeepCopy() *ProfileStatusDetail {
	if in == nil {
		return nil
	}
	out := new(ProfileStatusDetail)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProfileStatusSummary) DeepCopyInto(out *ProfileStatusSummary) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProfileStatusSummary.
func (in *ProfileStatusSummary) DeepCopy() *ProfileStatusSummary {
	if in == nil {
		return nil
	}
	out := new(ProfileStatusSummary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSigningProfile) DeepCopyInto(out *ResourceSigningProfile) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSigningProfile.
func (in *ResourceSigningProfile) DeepCopy() *ResourceSigningProfile {
	if in == nil {
		return nil
	}
	out := new(ResourceSigningProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceSigningProfile) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSigningProfileList) DeepCopyInto(out *ResourceSigningProfileList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ResourceSigningProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSigningProfileList.
func (in *ResourceSigningProfileList) DeepCopy() *ResourceSigningProfileList {
	if in == nil {
		return nil
	}
	out := new(ResourceSigningProfileList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ResourceSigningProfileList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSigningProfileSpec) DeepCopyInto(out *ResourceSigningProfileSpec) {
	*out = *in
	if in.TargetNamespaceSelector != nil {
		in, out := &in.TargetNamespaceSelector, &out.TargetNamespaceSelector
		*out = (*in).DeepCopy()
	}
	if in.ProtectRules != nil {
		in, out := &in.ProtectRules, &out.ProtectRules
		*out = make([]*common.Rule, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.IgnoreRules != nil {
		in, out := &in.IgnoreRules, &out.IgnoreRules
		*out = make([]*common.Rule, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.ForceCheckRules != nil {
		in, out := &in.ForceCheckRules, &out.ForceCheckRules
		*out = make([]*common.Rule, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.KustomizePatterns != nil {
		in, out := &in.KustomizePatterns, &out.KustomizePatterns
		*out = make([]*common.KustomizePattern, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.ProtectAttrs != nil {
		in, out := &in.ProtectAttrs, &out.ProtectAttrs
		*out = make([]*common.AttrsPattern, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.UnprotectAttrs != nil {
		in, out := &in.UnprotectAttrs, &out.UnprotectAttrs
		*out = make([]*common.AttrsPattern, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	if in.IgnoreAttrs != nil {
		in, out := &in.IgnoreAttrs, &out.IgnoreAttrs
		*out = make([]*common.AttrsPattern, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = (*in).DeepCopy()
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSigningProfileSpec.
func (in *ResourceSigningProfileSpec) DeepCopy() *ResourceSigningProfileSpec {
	if in == nil {
		return nil
	}
	out := new(ResourceSigningProfileSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceSigningProfileStatus) DeepCopyInto(out *ResourceSigningProfileStatus) {
	*out = *in
	if in.Summary != nil {
		in, out := &in.Summary, &out.Summary
		*out = make([]*ProfileStatusSummary, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ProfileStatusSummary)
				**out = **in
			}
		}
	}
	if in.Latest != nil {
		in, out := &in.Latest, &out.Latest
		*out = make([]*ProfileStatusDetail, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(ProfileStatusDetail)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceSigningProfileStatus.
func (in *ResourceSigningProfileStatus) DeepCopy() *ResourceSigningProfileStatus {
	if in == nil {
		return nil
	}
	out := new(ResourceSigningProfileStatus)
	in.DeepCopyInto(out)
	return out
}
