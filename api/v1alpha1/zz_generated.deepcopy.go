//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CertificateVerificationInfo) DeepCopyInto(out *CertificateVerificationInfo) {
	*out = *in
	if in.CertChain != nil {
		in, out := &in.CertChain, &out.CertChain
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CertificateVerificationInfo.
func (in *CertificateVerificationInfo) DeepCopy() *CertificateVerificationInfo {
	if in == nil {
		return nil
	}
	out := new(CertificateVerificationInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterContentLibrary) DeepCopyInto(out *ClusterContentLibrary) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterContentLibrary.
func (in *ClusterContentLibrary) DeepCopy() *ClusterContentLibrary {
	if in == nil {
		return nil
	}
	out := new(ClusterContentLibrary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterContentLibrary) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterContentLibraryItem) DeepCopyInto(out *ClusterContentLibraryItem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterContentLibraryItem.
func (in *ClusterContentLibraryItem) DeepCopy() *ClusterContentLibraryItem {
	if in == nil {
		return nil
	}
	out := new(ClusterContentLibraryItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterContentLibraryItem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterContentLibraryItemList) DeepCopyInto(out *ClusterContentLibraryItemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterContentLibraryItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterContentLibraryItemList.
func (in *ClusterContentLibraryItemList) DeepCopy() *ClusterContentLibraryItemList {
	if in == nil {
		return nil
	}
	out := new(ClusterContentLibraryItemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterContentLibraryItemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterContentLibraryList) DeepCopyInto(out *ClusterContentLibraryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterContentLibrary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterContentLibraryList.
func (in *ClusterContentLibraryList) DeepCopy() *ClusterContentLibraryList {
	if in == nil {
		return nil
	}
	out := new(ClusterContentLibraryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterContentLibraryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterContentLibrarySpec) DeepCopyInto(out *ClusterContentLibrarySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterContentLibrarySpec.
func (in *ClusterContentLibrarySpec) DeepCopy() *ClusterContentLibrarySpec {
	if in == nil {
		return nil
	}
	out := new(ClusterContentLibrarySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in Conditions) DeepCopyInto(out *Conditions) {
	{
		in := &in
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Conditions.
func (in Conditions) DeepCopy() Conditions {
	if in == nil {
		return nil
	}
	out := new(Conditions)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibrary) DeepCopyInto(out *ContentLibrary) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibrary.
func (in *ContentLibrary) DeepCopy() *ContentLibrary {
	if in == nil {
		return nil
	}
	out := new(ContentLibrary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContentLibrary) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryItem) DeepCopyInto(out *ContentLibraryItem) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryItem.
func (in *ContentLibraryItem) DeepCopy() *ContentLibraryItem {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContentLibraryItem) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryItemFileUploadStatus) DeepCopyInto(out *ContentLibraryItemFileUploadStatus) {
	*out = *in
	if in.InProgress != nil {
		in, out := &in.InProgress, &out.InProgress
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Completed != nil {
		in, out := &in.Completed, &out.Completed
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryItemFileUploadStatus.
func (in *ContentLibraryItemFileUploadStatus) DeepCopy() *ContentLibraryItemFileUploadStatus {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryItemFileUploadStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryItemImportRequest) DeepCopyInto(out *ContentLibraryItemImportRequest) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryItemImportRequest.
func (in *ContentLibraryItemImportRequest) DeepCopy() *ContentLibraryItemImportRequest {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryItemImportRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContentLibraryItemImportRequest) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryItemImportRequestList) DeepCopyInto(out *ContentLibraryItemImportRequestList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContentLibraryItemImportRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryItemImportRequestList.
func (in *ContentLibraryItemImportRequestList) DeepCopy() *ContentLibraryItemImportRequestList {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryItemImportRequestList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContentLibraryItemImportRequestList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryItemImportRequestSource) DeepCopyInto(out *ContentLibraryItemImportRequestSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryItemImportRequestSource.
func (in *ContentLibraryItemImportRequestSource) DeepCopy() *ContentLibraryItemImportRequestSource {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryItemImportRequestSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryItemImportRequestSpec) DeepCopyInto(out *ContentLibraryItemImportRequestSpec) {
	*out = *in
	out.Source = in.Source
	out.Target = in.Target
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryItemImportRequestSpec.
func (in *ContentLibraryItemImportRequestSpec) DeepCopy() *ContentLibraryItemImportRequestSpec {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryItemImportRequestSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryItemImportRequestStatus) DeepCopyInto(out *ContentLibraryItemImportRequestStatus) {
	*out = *in
	if in.TargetRef != nil {
		in, out := &in.TargetRef, &out.TargetRef
		*out = new(LocalObjectRef)
		**out = **in
	}
	in.CompletionTime.DeepCopyInto(&out.CompletionTime)
	in.StartTime.DeepCopyInto(&out.StartTime)
	if in.FileUpload != nil {
		in, out := &in.FileUpload, &out.FileUpload
		*out = new(ContentLibraryItemFileUploadStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryItemImportRequestStatus.
func (in *ContentLibraryItemImportRequestStatus) DeepCopy() *ContentLibraryItemImportRequestStatus {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryItemImportRequestStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryItemImportRequestTarget) DeepCopyInto(out *ContentLibraryItemImportRequestTarget) {
	*out = *in
	out.Item = in.Item
	out.Library = in.Library
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryItemImportRequestTarget.
func (in *ContentLibraryItemImportRequestTarget) DeepCopy() *ContentLibraryItemImportRequestTarget {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryItemImportRequestTarget)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryItemImportRequestTargetItem) DeepCopyInto(out *ContentLibraryItemImportRequestTargetItem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryItemImportRequestTargetItem.
func (in *ContentLibraryItemImportRequestTargetItem) DeepCopy() *ContentLibraryItemImportRequestTargetItem {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryItemImportRequestTargetItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryItemList) DeepCopyInto(out *ContentLibraryItemList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContentLibraryItem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryItemList.
func (in *ContentLibraryItemList) DeepCopy() *ContentLibraryItemList {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryItemList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContentLibraryItemList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryItemSpec) DeepCopyInto(out *ContentLibraryItemSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryItemSpec.
func (in *ContentLibraryItemSpec) DeepCopy() *ContentLibraryItemSpec {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryItemSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryItemStatus) DeepCopyInto(out *ContentLibraryItemStatus) {
	*out = *in
	if in.ContentLibraryRef != nil {
		in, out := &in.ContentLibraryRef, &out.ContentLibraryRef
		*out = new(NameAndKindRef)
		**out = **in
	}
	out.SizeInBytes = in.SizeInBytes.DeepCopy()
	if in.SecurityCompliance != nil {
		in, out := &in.SecurityCompliance, &out.SecurityCompliance
		*out = new(bool)
		**out = **in
	}
	if in.CertificateVerificationInfo != nil {
		in, out := &in.CertificateVerificationInfo, &out.CertificateVerificationInfo
		*out = new(CertificateVerificationInfo)
		(*in).DeepCopyInto(*out)
	}
	if in.FileInfo != nil {
		in, out := &in.FileInfo, &out.FileInfo
		*out = make([]FileInfo, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.CreationTime.DeepCopyInto(&out.CreationTime)
	in.LastModifiedTime.DeepCopyInto(&out.LastModifiedTime)
	in.LastSyncTime.DeepCopyInto(&out.LastSyncTime)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryItemStatus.
func (in *ContentLibraryItemStatus) DeepCopy() *ContentLibraryItemStatus {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryItemStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryList) DeepCopyInto(out *ContentLibraryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ContentLibrary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryList.
func (in *ContentLibraryList) DeepCopy() *ContentLibraryList {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ContentLibraryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibrarySpec) DeepCopyInto(out *ContentLibrarySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibrarySpec.
func (in *ContentLibrarySpec) DeepCopy() *ContentLibrarySpec {
	if in == nil {
		return nil
	}
	out := new(ContentLibrarySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContentLibraryStatus) DeepCopyInto(out *ContentLibraryStatus) {
	*out = *in
	if in.StorageBacking != nil {
		in, out := &in.StorageBacking, &out.StorageBacking
		*out = new(StorageBacking)
		**out = **in
	}
	if in.PublishInfo != nil {
		in, out := &in.PublishInfo, &out.PublishInfo
		*out = new(PublishInfo)
		**out = **in
	}
	if in.SubscriptionInfo != nil {
		in, out := &in.SubscriptionInfo, &out.SubscriptionInfo
		*out = new(SubscriptionInfo)
		**out = **in
	}
	in.CreationTime.DeepCopyInto(&out.CreationTime)
	in.LastModifiedTime.DeepCopyInto(&out.LastModifiedTime)
	in.LastSyncTime.DeepCopyInto(&out.LastSyncTime)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContentLibraryStatus.
func (in *ContentLibraryStatus) DeepCopy() *ContentLibraryStatus {
	if in == nil {
		return nil
	}
	out := new(ContentLibraryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FileInfo) DeepCopyInto(out *FileInfo) {
	*out = *in
	out.SizeInBytes = in.SizeInBytes.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FileInfo.
func (in *FileInfo) DeepCopy() *FileInfo {
	if in == nil {
		return nil
	}
	out := new(FileInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LocalObjectRef) DeepCopyInto(out *LocalObjectRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LocalObjectRef.
func (in *LocalObjectRef) DeepCopy() *LocalObjectRef {
	if in == nil {
		return nil
	}
	out := new(LocalObjectRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NameAndKindRef) DeepCopyInto(out *NameAndKindRef) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NameAndKindRef.
func (in *NameAndKindRef) DeepCopy() *NameAndKindRef {
	if in == nil {
		return nil
	}
	out := new(NameAndKindRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PublishInfo) DeepCopyInto(out *PublishInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PublishInfo.
func (in *PublishInfo) DeepCopy() *PublishInfo {
	if in == nil {
		return nil
	}
	out := new(PublishInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageBacking) DeepCopyInto(out *StorageBacking) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageBacking.
func (in *StorageBacking) DeepCopy() *StorageBacking {
	if in == nil {
		return nil
	}
	out := new(StorageBacking)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SubscriptionInfo) DeepCopyInto(out *SubscriptionInfo) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SubscriptionInfo.
func (in *SubscriptionInfo) DeepCopy() *SubscriptionInfo {
	if in == nil {
		return nil
	}
	out := new(SubscriptionInfo)
	in.DeepCopyInto(out)
	return out
}
