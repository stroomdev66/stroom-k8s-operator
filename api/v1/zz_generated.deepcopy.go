// +build !ignore_autogenerated

/*
Copyright 2021.

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
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackupSettings) DeepCopyInto(out *BackupSettings) {
	*out = *in
	if in.DatabaseNames != nil {
		in, out := &in.DatabaseNames, &out.DatabaseNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.TargetVolume.DeepCopyInto(&out.TargetVolume)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackupSettings.
func (in *BackupSettings) DeepCopy() *BackupSettings {
	if in == nil {
		return nil
	}
	out := new(BackupSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapRef) DeepCopyInto(out *ConfigMapRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapRef.
func (in *ConfigMapRef) DeepCopy() *ConfigMapRef {
	if in == nil {
		return nil
	}
	out := new(ConfigMapRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseServer) DeepCopyInto(out *DatabaseServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
	out.StroomClusterRef = in.StroomClusterRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseServer.
func (in *DatabaseServer) DeepCopy() *DatabaseServer {
	if in == nil {
		return nil
	}
	out := new(DatabaseServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseServerList) DeepCopyInto(out *DatabaseServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]DatabaseServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseServerList.
func (in *DatabaseServerList) DeepCopy() *DatabaseServerList {
	if in == nil {
		return nil
	}
	out := new(DatabaseServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DatabaseServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseServerRef) DeepCopyInto(out *DatabaseServerRef) {
	*out = *in
	out.ServerRef = in.ServerRef
	out.ServerAddress = in.ServerAddress
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseServerRef.
func (in *DatabaseServerRef) DeepCopy() *DatabaseServerRef {
	if in == nil {
		return nil
	}
	out := new(DatabaseServerRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseServerSpec) DeepCopyInto(out *DatabaseServerSpec) {
	*out = *in
	out.Image = in.Image
	if in.DatabaseNames != nil {
		in, out := &in.DatabaseNames, &out.DatabaseNames
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.AdditionalConfig != nil {
		in, out := &in.AdditionalConfig, &out.AdditionalConfig
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Resources.DeepCopyInto(&out.Resources)
	in.VolumeClaim.DeepCopyInto(&out.VolumeClaim)
	in.Backup.DeepCopyInto(&out.Backup)
	out.ReadinessProbeTimings = in.ReadinessProbeTimings
	out.LivenessProbeTimings = in.LivenessProbeTimings
	if in.PodAnnotations != nil {
		in, out := &in.PodAnnotations, &out.PodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.PodSecurityContext.DeepCopyInto(&out.PodSecurityContext)
	in.SecurityContext.DeepCopyInto(&out.SecurityContext)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Affinity.DeepCopyInto(&out.Affinity)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseServerSpec.
func (in *DatabaseServerSpec) DeepCopy() *DatabaseServerSpec {
	if in == nil {
		return nil
	}
	out := new(DatabaseServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DatabaseServerStatus) DeepCopyInto(out *DatabaseServerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DatabaseServerStatus.
func (in *DatabaseServerStatus) DeepCopy() *DatabaseServerStatus {
	if in == nil {
		return nil
	}
	out := new(DatabaseServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Image) DeepCopyInto(out *Image) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Image.
func (in *Image) DeepCopy() *Image {
	if in == nil {
		return nil
	}
	out := new(Image)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IngressSettings) DeepCopyInto(out *IngressSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IngressSettings.
func (in *IngressSettings) DeepCopy() *IngressSettings {
	if in == nil {
		return nil
	}
	out := new(IngressSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LogSenderSettings) DeepCopyInto(out *LogSenderSettings) {
	*out = *in
	out.Image = in.Image
	in.SecurityContext.DeepCopyInto(&out.SecurityContext)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LogSenderSettings.
func (in *LogSenderSettings) DeepCopy() *LogSenderSettings {
	if in == nil {
		return nil
	}
	out := new(LogSenderSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeSet) DeepCopyInto(out *NodeSet) {
	*out = *in
	in.LocalDataVolumeClaim.DeepCopyInto(&out.LocalDataVolumeClaim)
	in.SharedDataVolume.DeepCopyInto(&out.SharedDataVolume)
	in.Resources.DeepCopyInto(&out.Resources)
	out.ReadinessProbeTimings = in.ReadinessProbeTimings
	out.LivenessProbeTimings = in.LivenessProbeTimings
	if in.PodAnnotations != nil {
		in, out := &in.PodAnnotations, &out.PodAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	in.PodSecurityContext.DeepCopyInto(&out.PodSecurityContext)
	in.SecurityContext.DeepCopyInto(&out.SecurityContext)
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Affinity.DeepCopyInto(&out.Affinity)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeSet.
func (in *NodeSet) DeepCopy() *NodeSet {
	if in == nil {
		return nil
	}
	out := new(NodeSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProbeTimings) DeepCopyInto(out *ProbeTimings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProbeTimings.
func (in *ProbeTimings) DeepCopy() *ProbeTimings {
	if in == nil {
		return nil
	}
	out := new(ProbeTimings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResourceRef) DeepCopyInto(out *ResourceRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResourceRef.
func (in *ResourceRef) DeepCopy() *ResourceRef {
	if in == nil {
		return nil
	}
	out := new(ResourceRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServerAddress) DeepCopyInto(out *ServerAddress) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServerAddress.
func (in *ServerAddress) DeepCopy() *ServerAddress {
	if in == nil {
		return nil
	}
	out := new(ServerAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StroomCluster) DeepCopyInto(out *StroomCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StroomCluster.
func (in *StroomCluster) DeepCopy() *StroomCluster {
	if in == nil {
		return nil
	}
	out := new(StroomCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StroomCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StroomClusterList) DeepCopyInto(out *StroomClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StroomCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StroomClusterList.
func (in *StroomClusterList) DeepCopy() *StroomClusterList {
	if in == nil {
		return nil
	}
	out := new(StroomClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StroomClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StroomClusterSpec) DeepCopyInto(out *StroomClusterSpec) {
	*out = *in
	out.Image = in.Image
	out.DatabaseServerRef = in.DatabaseServerRef
	out.ConfigMapRef = in.ConfigMapRef
	out.Ingress = in.Ingress
	if in.NodeSets != nil {
		in, out := &in.NodeSets, &out.NodeSets
		*out = make([]NodeSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraEnv != nil {
		in, out := &in.ExtraEnv, &out.ExtraEnv
		*out = make([]corev1.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraVolumes != nil {
		in, out := &in.ExtraVolumes, &out.ExtraVolumes
		*out = make([]corev1.Volume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.ExtraVolumeMounts != nil {
		in, out := &in.ExtraVolumeMounts, &out.ExtraVolumeMounts
		*out = make([]corev1.VolumeMount, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.LogSender.DeepCopyInto(&out.LogSender)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StroomClusterSpec.
func (in *StroomClusterSpec) DeepCopy() *StroomClusterSpec {
	if in == nil {
		return nil
	}
	out := new(StroomClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StroomClusterStatus) DeepCopyInto(out *StroomClusterStatus) {
	*out = *in
	if in.Nodes != nil {
		in, out := &in.Nodes, &out.Nodes
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StroomClusterStatus.
func (in *StroomClusterStatus) DeepCopy() *StroomClusterStatus {
	if in == nil {
		return nil
	}
	out := new(StroomClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StroomTaskAutoscaler) DeepCopyInto(out *StroomTaskAutoscaler) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StroomTaskAutoscaler.
func (in *StroomTaskAutoscaler) DeepCopy() *StroomTaskAutoscaler {
	if in == nil {
		return nil
	}
	out := new(StroomTaskAutoscaler)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StroomTaskAutoscaler) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StroomTaskAutoscalerList) DeepCopyInto(out *StroomTaskAutoscalerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StroomTaskAutoscaler, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StroomTaskAutoscalerList.
func (in *StroomTaskAutoscalerList) DeepCopy() *StroomTaskAutoscalerList {
	if in == nil {
		return nil
	}
	out := new(StroomTaskAutoscalerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StroomTaskAutoscalerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StroomTaskAutoscalerSpec) DeepCopyInto(out *StroomTaskAutoscalerSpec) {
	*out = *in
	out.StroomClusterRef = in.StroomClusterRef
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StroomTaskAutoscalerSpec.
func (in *StroomTaskAutoscalerSpec) DeepCopy() *StroomTaskAutoscalerSpec {
	if in == nil {
		return nil
	}
	out := new(StroomTaskAutoscalerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StroomTaskAutoscalerStatus) DeepCopyInto(out *StroomTaskAutoscalerStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StroomTaskAutoscalerStatus.
func (in *StroomTaskAutoscalerStatus) DeepCopy() *StroomTaskAutoscalerStatus {
	if in == nil {
		return nil
	}
	out := new(StroomTaskAutoscalerStatus)
	in.DeepCopyInto(out)
	return out
}
