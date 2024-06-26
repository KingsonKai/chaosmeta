//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2023.

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
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonMeasure) DeepCopyInto(out *CommonMeasure) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonMeasure.
func (in *CommonMeasure) DeepCopy() *CommonMeasure {
	if in == nil {
		return nil
	}
	out := new(CommonMeasure)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CommonMeasure) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonMeasureList) DeepCopyInto(out *CommonMeasureList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CommonMeasure, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonMeasureList.
func (in *CommonMeasureList) DeepCopy() *CommonMeasureList {
	if in == nil {
		return nil
	}
	out := new(CommonMeasureList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CommonMeasureList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonMeasureSpec) DeepCopyInto(out *CommonMeasureSpec) {
	*out = *in
	out.Judgement = in.Judgement
	if in.Args != nil {
		in, out := &in.Args, &out.Args
		*out = make([]MeasureArgs, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonMeasureSpec.
func (in *CommonMeasureSpec) DeepCopy() *CommonMeasureSpec {
	if in == nil {
		return nil
	}
	out := new(CommonMeasureSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonMeasureStatus) DeepCopyInto(out *CommonMeasureStatus) {
	*out = *in
	if in.Measures != nil {
		in, out := &in.Measures, &out.Measures
		*out = make([]MeasureTask, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonMeasureStatus.
func (in *CommonMeasureStatus) DeepCopy() *CommonMeasureStatus {
	if in == nil {
		return nil
	}
	out := new(CommonMeasureStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Judgement) DeepCopyInto(out *Judgement) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Judgement.
func (in *Judgement) DeepCopy() *Judgement {
	if in == nil {
		return nil
	}
	out := new(Judgement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeasureArgs) DeepCopyInto(out *MeasureArgs) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeasureArgs.
func (in *MeasureArgs) DeepCopy() *MeasureArgs {
	if in == nil {
		return nil
	}
	out := new(MeasureArgs)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MeasureTask) DeepCopyInto(out *MeasureTask) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MeasureTask.
func (in *MeasureTask) DeepCopy() *MeasureTask {
	if in == nil {
		return nil
	}
	out := new(MeasureTask)
	in.DeepCopyInto(out)
	return out
}
