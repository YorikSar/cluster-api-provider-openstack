// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// Code generated by main. DO NOT EDIT.

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NetworkParam) DeepCopyInto(out *NetworkParam) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NetworkParam.
func (in *NetworkParam) DeepCopy() *NetworkParam {
	if in == nil {
		return nil
	}
	out := new(NetworkParam)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackClusterProviderStatus) DeepCopyInto(out *OpenstackClusterProviderStatus) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.CACertificate != nil {
		in, out := &in.CACertificate, &out.CACertificate
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.CAPrivateKey != nil {
		in, out := &in.CAPrivateKey, &out.CAPrivateKey
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackClusterProviderStatus.
func (in *OpenstackClusterProviderStatus) DeepCopy() *OpenstackClusterProviderStatus {
	if in == nil {
		return nil
	}
	out := new(OpenstackClusterProviderStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenstackClusterProviderStatus) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenstackProviderSpec) DeepCopyInto(out *OpenstackProviderSpec) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Networks != nil {
		in, out := &in.Networks, &out.Networks
		*out = make([]NetworkParam, len(*in))
		copy(*out, *in)
	}
	if in.SecurityGroups != nil {
		in, out := &in.SecurityGroups, &out.SecurityGroups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.RootVolume = in.RootVolume
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenstackProviderSpec.
func (in *OpenstackProviderSpec) DeepCopy() *OpenstackProviderSpec {
	if in == nil {
		return nil
	}
	out := new(OpenstackProviderSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenstackProviderSpec) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RootVolume) DeepCopyInto(out *RootVolume) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RootVolume.
func (in *RootVolume) DeepCopy() *RootVolume {
	if in == nil {
		return nil
	}
	out := new(RootVolume)
	in.DeepCopyInto(out)
	return out
}
