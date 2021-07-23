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

package manualv1alpha1

import (
	"github.com/crossplane/crossplane-runtime/apis/common/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BlockDeviceMapping) DeepCopyInto(out *BlockDeviceMapping) {
	*out = *in
	if in.DeviceName != nil {
		in, out := &in.DeviceName, &out.DeviceName
		*out = new(string)
		**out = **in
	}
	if in.EBS != nil {
		in, out := &in.EBS, &out.EBS
		*out = new(EBSBlockDevice)
		(*in).DeepCopyInto(*out)
	}
	if in.NoDevice != nil {
		in, out := &in.NoDevice, &out.NoDevice
		*out = new(string)
		**out = **in
	}
	if in.VirtualName != nil {
		in, out := &in.VirtualName, &out.VirtualName
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BlockDeviceMapping.
func (in *BlockDeviceMapping) DeepCopy() *BlockDeviceMapping {
	if in == nil {
		return nil
	}
	out := new(BlockDeviceMapping)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CPUOptionsRequest) DeepCopyInto(out *CPUOptionsRequest) {
	*out = *in
	if in.CoreCount != nil {
		in, out := &in.CoreCount, &out.CoreCount
		*out = new(int64)
		**out = **in
	}
	if in.ThreadsPerCore != nil {
		in, out := &in.ThreadsPerCore, &out.ThreadsPerCore
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CPUOptionsRequest.
func (in *CPUOptionsRequest) DeepCopy() *CPUOptionsRequest {
	if in == nil {
		return nil
	}
	out := new(CPUOptionsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapacityReservationSpecification) DeepCopyInto(out *CapacityReservationSpecification) {
	*out = *in
	if in.CapacityReservationPreference != nil {
		in, out := &in.CapacityReservationPreference, &out.CapacityReservationPreference
		*out = new(string)
		**out = **in
	}
	if in.CapacityReservationTarget != nil {
		in, out := &in.CapacityReservationTarget, &out.CapacityReservationTarget
		*out = new(CapacityReservationTarget)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacityReservationSpecification.
func (in *CapacityReservationSpecification) DeepCopy() *CapacityReservationSpecification {
	if in == nil {
		return nil
	}
	out := new(CapacityReservationSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CapacityReservationTarget) DeepCopyInto(out *CapacityReservationTarget) {
	*out = *in
	if in.CapacityReservationID != nil {
		in, out := &in.CapacityReservationID, &out.CapacityReservationID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CapacityReservationTarget.
func (in *CapacityReservationTarget) DeepCopy() *CapacityReservationTarget {
	if in == nil {
		return nil
	}
	out := new(CapacityReservationTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CreditSpecificationRequest) DeepCopyInto(out *CreditSpecificationRequest) {
	*out = *in
	if in.CPUCredits != nil {
		in, out := &in.CPUCredits, &out.CPUCredits
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CreditSpecificationRequest.
func (in *CreditSpecificationRequest) DeepCopy() *CreditSpecificationRequest {
	if in == nil {
		return nil
	}
	out := new(CreditSpecificationRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EBSBlockDevice) DeepCopyInto(out *EBSBlockDevice) {
	*out = *in
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	if in.Encrypted != nil {
		in, out := &in.Encrypted, &out.Encrypted
		*out = new(bool)
		**out = **in
	}
	if in.IOps != nil {
		in, out := &in.IOps, &out.IOps
		*out = new(int64)
		**out = **in
	}
	if in.KmsKeyID != nil {
		in, out := &in.KmsKeyID, &out.KmsKeyID
		*out = new(string)
		**out = **in
	}
	if in.SnapshotID != nil {
		in, out := &in.SnapshotID, &out.SnapshotID
		*out = new(string)
		**out = **in
	}
	if in.VolumeSize != nil {
		in, out := &in.VolumeSize, &out.VolumeSize
		*out = new(int64)
		**out = **in
	}
	if in.VolumeType != nil {
		in, out := &in.VolumeType, &out.VolumeType
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EBSBlockDevice.
func (in *EBSBlockDevice) DeepCopy() *EBSBlockDevice {
	if in == nil {
		return nil
	}
	out := new(EBSBlockDevice)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticGpuSpecification) DeepCopyInto(out *ElasticGpuSpecification) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticGpuSpecification.
func (in *ElasticGpuSpecification) DeepCopy() *ElasticGpuSpecification {
	if in == nil {
		return nil
	}
	out := new(ElasticGpuSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ElasticInferenceAccelerator) DeepCopyInto(out *ElasticInferenceAccelerator) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int64)
		**out = **in
	}
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ElasticInferenceAccelerator.
func (in *ElasticInferenceAccelerator) DeepCopy() *ElasticInferenceAccelerator {
	if in == nil {
		return nil
	}
	out := new(ElasticInferenceAccelerator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HibernationOptionsRequest) DeepCopyInto(out *HibernationOptionsRequest) {
	*out = *in
	if in.Configured != nil {
		in, out := &in.Configured, &out.Configured
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HibernationOptionsRequest.
func (in *HibernationOptionsRequest) DeepCopy() *HibernationOptionsRequest {
	if in == nil {
		return nil
	}
	out := new(HibernationOptionsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *IamInstanceProfileSpecification) DeepCopyInto(out *IamInstanceProfileSpecification) {
	*out = *in
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new IamInstanceProfileSpecification.
func (in *IamInstanceProfileSpecification) DeepCopy() *IamInstanceProfileSpecification {
	if in == nil {
		return nil
	}
	out := new(IamInstanceProfileSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Instance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceIPV6Address) DeepCopyInto(out *InstanceIPV6Address) {
	*out = *in
	if in.IPV6Address != nil {
		in, out := &in.IPV6Address, &out.IPV6Address
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceIPV6Address.
func (in *InstanceIPV6Address) DeepCopy() *InstanceIPV6Address {
	if in == nil {
		return nil
	}
	out := new(InstanceIPV6Address)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceList) DeepCopyInto(out *InstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceList.
func (in *InstanceList) DeepCopy() *InstanceList {
	if in == nil {
		return nil
	}
	out := new(InstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceMarketOptionsRequest) DeepCopyInto(out *InstanceMarketOptionsRequest) {
	*out = *in
	if in.MarketType != nil {
		in, out := &in.MarketType, &out.MarketType
		*out = new(string)
		**out = **in
	}
	if in.SpotOptions != nil {
		in, out := &in.SpotOptions, &out.SpotOptions
		*out = new(SpotMarketOptions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceMarketOptionsRequest.
func (in *InstanceMarketOptionsRequest) DeepCopy() *InstanceMarketOptionsRequest {
	if in == nil {
		return nil
	}
	out := new(InstanceMarketOptionsRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceNetworkInterfaceSpecification) DeepCopyInto(out *InstanceNetworkInterfaceSpecification) {
	*out = *in
	if in.AssociatePublicIPAddress != nil {
		in, out := &in.AssociatePublicIPAddress, &out.AssociatePublicIPAddress
		*out = new(bool)
		**out = **in
	}
	if in.DeleteOnTermination != nil {
		in, out := &in.DeleteOnTermination, &out.DeleteOnTermination
		*out = new(bool)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.DeviceIndex != nil {
		in, out := &in.DeviceIndex, &out.DeviceIndex
		*out = new(int64)
		**out = **in
	}
	if in.Groups != nil {
		in, out := &in.Groups, &out.Groups
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InterfaceType != nil {
		in, out := &in.InterfaceType, &out.InterfaceType
		*out = new(string)
		**out = **in
	}
	if in.Ipv6AddressCount != nil {
		in, out := &in.Ipv6AddressCount, &out.Ipv6AddressCount
		*out = new(int64)
		**out = **in
	}
	if in.IPV6Addresses != nil {
		in, out := &in.IPV6Addresses, &out.IPV6Addresses
		*out = make([]InstanceIPV6Address, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkInterfaceID != nil {
		in, out := &in.NetworkInterfaceID, &out.NetworkInterfaceID
		*out = new(string)
		**out = **in
	}
	if in.PrivateIPAddress != nil {
		in, out := &in.PrivateIPAddress, &out.PrivateIPAddress
		*out = new(string)
		**out = **in
	}
	if in.PrivateIPAddresses != nil {
		in, out := &in.PrivateIPAddresses, &out.PrivateIPAddresses
		*out = make([]PrivateIPAddressSpecification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SecondaryPrivateIPAddressCount != nil {
		in, out := &in.SecondaryPrivateIPAddressCount, &out.SecondaryPrivateIPAddressCount
		*out = new(int64)
		**out = **in
	}
	if in.SubnetID != nil {
		in, out := &in.SubnetID, &out.SubnetID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceNetworkInterfaceSpecification.
func (in *InstanceNetworkInterfaceSpecification) DeepCopy() *InstanceNetworkInterfaceSpecification {
	if in == nil {
		return nil
	}
	out := new(InstanceNetworkInterfaceSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceObservation) DeepCopyInto(out *InstanceObservation) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceObservation.
func (in *InstanceObservation) DeepCopy() *InstanceObservation {
	if in == nil {
		return nil
	}
	out := new(InstanceObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceParameters) DeepCopyInto(out *InstanceParameters) {
	*out = *in
	if in.ClientToken != nil {
		in, out := &in.ClientToken, &out.ClientToken
		*out = new(string)
		**out = **in
	}
	if in.DryRun != nil {
		in, out := &in.DryRun, &out.DryRun
		*out = new(bool)
		**out = **in
	}
	if in.ImageID != nil {
		in, out := &in.ImageID, &out.ImageID
		*out = new(string)
		**out = **in
	}
	if in.MaxCount != nil {
		in, out := &in.MaxCount, &out.MaxCount
		*out = new(int64)
		**out = **in
	}
	if in.MinCount != nil {
		in, out := &in.MinCount, &out.MinCount
		*out = new(int64)
		**out = **in
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(RunInstancesMonitoringEnabled)
		(*in).DeepCopyInto(*out)
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.TagSpecifications != nil {
		in, out := &in.TagSpecifications, &out.TagSpecifications
		*out = make([]TagSpecification, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.UserData != nil {
		in, out := &in.UserData, &out.UserData
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceParameters.
func (in *InstanceParameters) DeepCopy() *InstanceParameters {
	if in == nil {
		return nil
	}
	out := new(InstanceParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	out.AtProvider = in.AtProvider
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LaunchTemplateSpecification) DeepCopyInto(out *LaunchTemplateSpecification) {
	*out = *in
	if in.LaunchTemplateID != nil {
		in, out := &in.LaunchTemplateID, &out.LaunchTemplateID
		*out = new(string)
		**out = **in
	}
	if in.LaunchTemplateName != nil {
		in, out := &in.LaunchTemplateName, &out.LaunchTemplateName
		*out = new(string)
		**out = **in
	}
	if in.Version != nil {
		in, out := &in.Version, &out.Version
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LaunchTemplateSpecification.
func (in *LaunchTemplateSpecification) DeepCopy() *LaunchTemplateSpecification {
	if in == nil {
		return nil
	}
	out := new(LaunchTemplateSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LicenseConfigurationRequest) DeepCopyInto(out *LicenseConfigurationRequest) {
	*out = *in
	if in.LicenseConfigurationARN != nil {
		in, out := &in.LicenseConfigurationARN, &out.LicenseConfigurationARN
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LicenseConfigurationRequest.
func (in *LicenseConfigurationRequest) DeepCopy() *LicenseConfigurationRequest {
	if in == nil {
		return nil
	}
	out := new(LicenseConfigurationRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Placement) DeepCopyInto(out *Placement) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(string)
		**out = **in
	}
	if in.AvailabilityZone != nil {
		in, out := &in.AvailabilityZone, &out.AvailabilityZone
		*out = new(string)
		**out = **in
	}
	if in.GroupName != nil {
		in, out := &in.GroupName, &out.GroupName
		*out = new(string)
		**out = **in
	}
	if in.HostID != nil {
		in, out := &in.HostID, &out.HostID
		*out = new(string)
		**out = **in
	}
	if in.HostResourceGroupARN != nil {
		in, out := &in.HostResourceGroupARN, &out.HostResourceGroupARN
		*out = new(string)
		**out = **in
	}
	if in.PartitionNumber != nil {
		in, out := &in.PartitionNumber, &out.PartitionNumber
		*out = new(int64)
		**out = **in
	}
	if in.SpreadDomain != nil {
		in, out := &in.SpreadDomain, &out.SpreadDomain
		*out = new(string)
		**out = **in
	}
	if in.Tenancy != nil {
		in, out := &in.Tenancy, &out.Tenancy
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Placement.
func (in *Placement) DeepCopy() *Placement {
	if in == nil {
		return nil
	}
	out := new(Placement)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrivateIPAddressSpecification) DeepCopyInto(out *PrivateIPAddressSpecification) {
	*out = *in
	if in.Primary != nil {
		in, out := &in.Primary, &out.Primary
		*out = new(bool)
		**out = **in
	}
	if in.PrivateIPAddress != nil {
		in, out := &in.PrivateIPAddress, &out.PrivateIPAddress
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrivateIPAddressSpecification.
func (in *PrivateIPAddressSpecification) DeepCopy() *PrivateIPAddressSpecification {
	if in == nil {
		return nil
	}
	out := new(PrivateIPAddressSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunInstancesMonitoringEnabled) DeepCopyInto(out *RunInstancesMonitoringEnabled) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunInstancesMonitoringEnabled.
func (in *RunInstancesMonitoringEnabled) DeepCopy() *RunInstancesMonitoringEnabled {
	if in == nil {
		return nil
	}
	out := new(RunInstancesMonitoringEnabled)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SpotMarketOptions) DeepCopyInto(out *SpotMarketOptions) {
	*out = *in
	if in.BlockDurationMinutes != nil {
		in, out := &in.BlockDurationMinutes, &out.BlockDurationMinutes
		*out = new(int64)
		**out = **in
	}
	if in.InstanceInterruptionBehavior != nil {
		in, out := &in.InstanceInterruptionBehavior, &out.InstanceInterruptionBehavior
		*out = new(string)
		**out = **in
	}
	if in.MaxPrice != nil {
		in, out := &in.MaxPrice, &out.MaxPrice
		*out = new(string)
		**out = **in
	}
	if in.SpotInstanceType != nil {
		in, out := &in.SpotInstanceType, &out.SpotInstanceType
		*out = new(string)
		**out = **in
	}
	if in.ValidUntil != nil {
		in, out := &in.ValidUntil, &out.ValidUntil
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SpotMarketOptions.
func (in *SpotMarketOptions) DeepCopy() *SpotMarketOptions {
	if in == nil {
		return nil
	}
	out := new(SpotMarketOptions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TagSpecification) DeepCopyInto(out *TagSpecification) {
	*out = *in
	if in.ResourceType != nil {
		in, out := &in.ResourceType, &out.ResourceType
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]Tag, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TagSpecification.
func (in *TagSpecification) DeepCopy() *TagSpecification {
	if in == nil {
		return nil
	}
	out := new(TagSpecification)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCCIDRBlock) DeepCopyInto(out *VPCCIDRBlock) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCCIDRBlock.
func (in *VPCCIDRBlock) DeepCopy() *VPCCIDRBlock {
	if in == nil {
		return nil
	}
	out := new(VPCCIDRBlock)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCCIDRBlock) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCCIDRBlockList) DeepCopyInto(out *VPCCIDRBlockList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VPCCIDRBlock, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCCIDRBlockList.
func (in *VPCCIDRBlockList) DeepCopy() *VPCCIDRBlockList {
	if in == nil {
		return nil
	}
	out := new(VPCCIDRBlockList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VPCCIDRBlockList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCCIDRBlockObservation) DeepCopyInto(out *VPCCIDRBlockObservation) {
	*out = *in
	if in.AssociationID != nil {
		in, out := &in.AssociationID, &out.AssociationID
		*out = new(string)
		**out = **in
	}
	if in.CIDRBlock != nil {
		in, out := &in.CIDRBlock, &out.CIDRBlock
		*out = new(string)
		**out = **in
	}
	if in.IPv6CIDRBlock != nil {
		in, out := &in.IPv6CIDRBlock, &out.IPv6CIDRBlock
		*out = new(string)
		**out = **in
	}
	if in.IPv6CIDRBlockState != nil {
		in, out := &in.IPv6CIDRBlockState, &out.IPv6CIDRBlockState
		*out = new(VPCCIDRBlockState)
		(*in).DeepCopyInto(*out)
	}
	if in.IPv6Pool != nil {
		in, out := &in.IPv6Pool, &out.IPv6Pool
		*out = new(string)
		**out = **in
	}
	if in.NetworkBorderGroup != nil {
		in, out := &in.NetworkBorderGroup, &out.NetworkBorderGroup
		*out = new(string)
		**out = **in
	}
	if in.CIDRBlockState != nil {
		in, out := &in.CIDRBlockState, &out.CIDRBlockState
		*out = new(VPCCIDRBlockState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCCIDRBlockObservation.
func (in *VPCCIDRBlockObservation) DeepCopy() *VPCCIDRBlockObservation {
	if in == nil {
		return nil
	}
	out := new(VPCCIDRBlockObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCCIDRBlockParameters) DeepCopyInto(out *VPCCIDRBlockParameters) {
	*out = *in
	if in.AmazonProvidedIPv6CIDRBlock != nil {
		in, out := &in.AmazonProvidedIPv6CIDRBlock, &out.AmazonProvidedIPv6CIDRBlock
		*out = new(bool)
		**out = **in
	}
	if in.CIDRBlock != nil {
		in, out := &in.CIDRBlock, &out.CIDRBlock
		*out = new(string)
		**out = **in
	}
	if in.IPv6CIDRBlock != nil {
		in, out := &in.IPv6CIDRBlock, &out.IPv6CIDRBlock
		*out = new(string)
		**out = **in
	}
	if in.IPv6CIDRBlockNetworkBorderGroup != nil {
		in, out := &in.IPv6CIDRBlockNetworkBorderGroup, &out.IPv6CIDRBlockNetworkBorderGroup
		*out = new(string)
		**out = **in
	}
	if in.IPv6Pool != nil {
		in, out := &in.IPv6Pool, &out.IPv6Pool
		*out = new(string)
		**out = **in
	}
	if in.VPCID != nil {
		in, out := &in.VPCID, &out.VPCID
		*out = new(string)
		**out = **in
	}
	if in.VPCIDRef != nil {
		in, out := &in.VPCIDRef, &out.VPCIDRef
		*out = new(v1.Reference)
		**out = **in
	}
	if in.VPCIDSelector != nil {
		in, out := &in.VPCIDSelector, &out.VPCIDSelector
		*out = new(v1.Selector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCCIDRBlockParameters.
func (in *VPCCIDRBlockParameters) DeepCopy() *VPCCIDRBlockParameters {
	if in == nil {
		return nil
	}
	out := new(VPCCIDRBlockParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCCIDRBlockSpec) DeepCopyInto(out *VPCCIDRBlockSpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCCIDRBlockSpec.
func (in *VPCCIDRBlockSpec) DeepCopy() *VPCCIDRBlockSpec {
	if in == nil {
		return nil
	}
	out := new(VPCCIDRBlockSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCCIDRBlockState) DeepCopyInto(out *VPCCIDRBlockState) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
	if in.StatusMessage != nil {
		in, out := &in.StatusMessage, &out.StatusMessage
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCCIDRBlockState.
func (in *VPCCIDRBlockState) DeepCopy() *VPCCIDRBlockState {
	if in == nil {
		return nil
	}
	out := new(VPCCIDRBlockState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VPCCIDRBlockStatus) DeepCopyInto(out *VPCCIDRBlockStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VPCCIDRBlockStatus.
func (in *VPCCIDRBlockStatus) DeepCopy() *VPCCIDRBlockStatus {
	if in == nil {
		return nil
	}
	out := new(VPCCIDRBlockStatus)
	in.DeepCopyInto(out)
	return out
}
