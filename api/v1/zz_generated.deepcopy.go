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

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AffinitySpec) DeepCopyInto(out *AffinitySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AffinitySpec.
func (in *AffinitySpec) DeepCopy() *AffinitySpec {
	if in == nil {
		return nil
	}
	out := new(AffinitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSetStatus) DeepCopyInto(out *InstanceSetStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSetStatus.
func (in *InstanceSetStatus) DeepCopy() *InstanceSetStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaParamsSpec) DeepCopyInto(out *MetaParamsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaParamsSpec.
func (in *MetaParamsSpec) DeepCopy() *MetaParamsSpec {
	if in == nil {
		return nil
	}
	out := new(MetaParamsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MetaSpec) DeepCopyInto(out *MetaSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.DataVolumeClaimSpec.DeepCopyInto(&out.DataVolumeClaimSpec)
	in.Resources.DeepCopyInto(&out.Resources)
	out.Parameters = in.Parameters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MetaSpec.
func (in *MetaSpec) DeepCopy() *MetaSpec {
	if in == nil {
		return nil
	}
	out := new(MetaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Metadata) DeepCopyInto(out *Metadata) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Annotations != nil {
		in, out := &in.Annotations, &out.Annotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Metadata.
func (in *Metadata) DeepCopy() *Metadata {
	if in == nil {
		return nil
	}
	out := new(Metadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringSpec) DeepCopyInto(out *MonitoringSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringSpec.
func (in *MonitoringSpec) DeepCopy() *MonitoringSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenGeminiCluster) DeepCopyInto(out *OpenGeminiCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenGeminiCluster.
func (in *OpenGeminiCluster) DeepCopy() *OpenGeminiCluster {
	if in == nil {
		return nil
	}
	out := new(OpenGeminiCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenGeminiCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenGeminiClusterList) DeepCopyInto(out *OpenGeminiClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OpenGeminiCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenGeminiClusterList.
func (in *OpenGeminiClusterList) DeepCopy() *OpenGeminiClusterList {
	if in == nil {
		return nil
	}
	out := new(OpenGeminiClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OpenGeminiClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenGeminiClusterSpec) DeepCopyInto(out *OpenGeminiClusterSpec) {
	*out = *in
	if in.Paused != nil {
		in, out := &in.Paused, &out.Paused
		*out = new(bool)
		**out = **in
	}
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(Metadata)
		(*in).DeepCopyInto(*out)
	}
	in.SQL.DeepCopyInto(&out.SQL)
	in.Meta.DeepCopyInto(&out.Meta)
	in.Store.DeepCopyInto(&out.Store)
	out.Monitoring = in.Monitoring
	out.Affinity = in.Affinity
	if in.EnableHttpAuth != nil {
		in, out := &in.EnableHttpAuth, &out.EnableHttpAuth
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenGeminiClusterSpec.
func (in *OpenGeminiClusterSpec) DeepCopy() *OpenGeminiClusterSpec {
	if in == nil {
		return nil
	}
	out := new(OpenGeminiClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenGeminiClusterStatus) DeepCopyInto(out *OpenGeminiClusterStatus) {
	*out = *in
	if in.InstanceSets != nil {
		in, out := &in.InstanceSets, &out.InstanceSets
		*out = make([]InstanceSetStatus, len(*in))
		copy(*out, *in)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenGeminiClusterStatus.
func (in *OpenGeminiClusterStatus) DeepCopy() *OpenGeminiClusterStatus {
	if in == nil {
		return nil
	}
	out := new(OpenGeminiClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLParamsSpec) DeepCopyInto(out *SQLParamsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLParamsSpec.
func (in *SQLParamsSpec) DeepCopy() *SQLParamsSpec {
	if in == nil {
		return nil
	}
	out := new(SQLParamsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLSpec) DeepCopyInto(out *SQLSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.Resources.DeepCopyInto(&out.Resources)
	out.Parameters = in.Parameters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLSpec.
func (in *SQLSpec) DeepCopy() *SQLSpec {
	if in == nil {
		return nil
	}
	out := new(SQLSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreParamsSpec) DeepCopyInto(out *StoreParamsSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreParamsSpec.
func (in *StoreParamsSpec) DeepCopy() *StoreParamsSpec {
	if in == nil {
		return nil
	}
	out := new(StoreParamsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StoreSpec) DeepCopyInto(out *StoreSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.DataVolumeClaimSpec.DeepCopyInto(&out.DataVolumeClaimSpec)
	in.Resources.DeepCopyInto(&out.Resources)
	out.Parameters = in.Parameters
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StoreSpec.
func (in *StoreSpec) DeepCopy() *StoreSpec {
	if in == nil {
		return nil
	}
	out := new(StoreSpec)
	in.DeepCopyInto(out)
	return out
}
