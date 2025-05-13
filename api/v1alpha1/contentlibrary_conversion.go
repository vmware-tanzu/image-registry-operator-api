// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	"k8s.io/apimachinery/pkg/types"
	ctrlconversion "sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/vmware-tanzu/image-registry-operator-api/api/utilconversion"
	imgregopv1 "github.com/vmware-tanzu/image-registry-operator-api/api/v1alpha2"
)

func Convert_v1alpha1_BaseContentLibrarySpec_To_v1alpha2_BaseContentLibrarySpec(in *BaseContentLibrarySpec, out *imgregopv1.BaseContentLibrarySpec, _ apiconversion.Scope) error {
	out.ID = string(in.UUID)
	out.Type = imgregopv1.LibraryTypeContentLibrary
	switch in.ResourceNamingStrategy {
	case ResourceNamingStrategyFromItemID:
		out.ResourceNamingStrategy = imgregopv1.ResourceNamingStrategyFromItemID
	case ResourceNamingStrategyPreferItemSourceID:
		out.ResourceNamingStrategy = imgregopv1.ResourceNamingStrategyPreferItemSourceID
	default:
		out.ResourceNamingStrategy = imgregopv1.ResourceNamingStrategy(in.ResourceNamingStrategy)
	}
	return nil
}

func Convert_v1alpha2_BaseContentLibrarySpec_To_v1alpha1_BaseContentLibrarySpec(in *imgregopv1.BaseContentLibrarySpec, out *BaseContentLibrarySpec, _ apiconversion.Scope) error {
	out.UUID = types.UID(in.ID)
	switch in.ResourceNamingStrategy {
	case imgregopv1.ResourceNamingStrategyFromItemID:
		out.ResourceNamingStrategy = ResourceNamingStrategyFromItemID
	case imgregopv1.ResourceNamingStrategyPreferItemSourceID:
		out.ResourceNamingStrategy = ResourceNamingStrategyPreferItemSourceID
	default:
		out.ResourceNamingStrategy = ResourceNamingStrategy(in.ResourceNamingStrategy)
	}
	return nil
}

func Convert_v1alpha1_ContentLibrarySpec_To_v1alpha2_ContentLibrarySpec(in *ContentLibrarySpec, out *imgregopv1.ContentLibrarySpec, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha1_ContentLibrarySpec_To_v1alpha2_ContentLibrarySpec(in, out, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_BaseContentLibrarySpec_To_v1alpha2_BaseContentLibrarySpec(&in.BaseContentLibrarySpec, &out.BaseContentLibrarySpec, nil); err != nil {
		return err
	}
	out.AllowDelete = in.Writable
	out.AllowPublish = in.Writable
	return nil
}

func Convert_v1alpha2_ContentLibrarySpec_To_v1alpha1_ContentLibrarySpec(in *imgregopv1.ContentLibrarySpec, out *ContentLibrarySpec, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha2_ContentLibrarySpec_To_v1alpha1_ContentLibrarySpec(in, out, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_BaseContentLibrarySpec_To_v1alpha1_BaseContentLibrarySpec(&in.BaseContentLibrarySpec, &out.BaseContentLibrarySpec, nil); err != nil {
		return err
	}
	out.Writable = in.AllowPublish || in.AllowDelete
	return nil
}

func Convert_v1alpha1_ClusterContentLibrarySpec_To_v1alpha2_ClusterContentLibrarySpec(in *ClusterContentLibrarySpec, out *imgregopv1.ClusterContentLibrarySpec, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha1_ClusterContentLibrarySpec_To_v1alpha2_ClusterContentLibrarySpec(in, out, s); err != nil {
		return err
	}
	return Convert_v1alpha1_BaseContentLibrarySpec_To_v1alpha2_BaseContentLibrarySpec(&in.BaseContentLibrarySpec, &out.BaseContentLibrarySpec, nil)
}

func Convert_v1alpha2_ClusterContentLibrarySpec_To_v1alpha1_ClusterContentLibrarySpec(in *imgregopv1.ClusterContentLibrarySpec, out *ClusterContentLibrarySpec, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha2_ClusterContentLibrarySpec_To_v1alpha1_ClusterContentLibrarySpec(in, out, s); err != nil {
		return err
	}
	return Convert_v1alpha2_BaseContentLibrarySpec_To_v1alpha1_BaseContentLibrarySpec(&in.BaseContentLibrarySpec, &out.BaseContentLibrarySpec, nil)
}

func Convert_v1alpha1_ContentLibraryStatus_To_v1alpha2_ContentLibraryStatus(in *ContentLibraryStatus, out *imgregopv1.ContentLibraryStatus, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha1_ContentLibraryStatus_To_v1alpha2_ContentLibraryStatus(in, out, s); err != nil {
		return err
	}
	if in.StorageBacking != nil {
		var sb imgregopv1.StorageBacking
		if err := Convert_v1alpha1_StorageBacking_To_v1alpha2_StorageBacking(in.StorageBacking, &sb, s); err != nil {
			return err
		}
		out.StorageBackings = []imgregopv1.StorageBacking{sb}
	}
	return nil
}

func Convert_v1alpha2_ContentLibraryStatus_To_v1alpha1_ContentLibraryStatus(in *imgregopv1.ContentLibraryStatus, out *ContentLibraryStatus, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha2_ContentLibraryStatus_To_v1alpha1_ContentLibraryStatus(in, out, s); err != nil {
		return err
	}
	if len(in.StorageBackings) > 0 {
		var sb StorageBacking
		if err := Convert_v1alpha2_StorageBacking_To_v1alpha1_StorageBacking(&in.StorageBackings[0], &sb, s); err != nil {
			return err
		}
		out.StorageBacking = &sb
	}
	return nil
}

// ConvertTo converts this ContentLibrary to the Hub version.
func (src *ContentLibrary) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*imgregopv1.ContentLibrary)
	if err := Convert_v1alpha1_ContentLibrary_To_v1alpha2_ContentLibrary(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &imgregopv1.ContentLibrary{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	// BEGIN RESTORE

	dst.Spec.StorageClass = restored.Spec.StorageClass
	dst.Spec.Type = restored.Spec.Type
	dst.Spec.AllowPublish = restored.Spec.AllowPublish
	dst.Spec.AllowDelete = restored.Spec.AllowDelete
	dst.Status = restored.Status

	// END RESTORE

	return nil
}

// ConvertFrom converts the hub version to this ContentLibrary.
func (dst *ContentLibrary) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*imgregopv1.ContentLibrary)
	if err := Convert_v1alpha2_ContentLibrary_To_v1alpha1_ContentLibrary(src, dst, nil); err != nil {
		return err
	}

	if src.Status.SubscriptionInfo == nil {
		dst.Status.Type = ContentLibraryTypeLocal
	} else {
		dst.Status.Type = ContentLibraryTypeSubscribed
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, dst)
}

// ConvertTo converts this ContentLibraryList to the Hub version.
func (src *ContentLibraryList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*imgregopv1.ContentLibraryList)
	return Convert_v1alpha1_ContentLibraryList_To_v1alpha2_ContentLibraryList(src, dst, nil)
}

// ConvertFrom converts the hub version to this ContentLibraryList.
func (dst *ContentLibraryList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*imgregopv1.ContentLibraryList)
	return Convert_v1alpha2_ContentLibraryList_To_v1alpha1_ContentLibraryList(src, dst, nil)
}

// ConvertTo converts this ClusterContentLibrary to the Hub version.
func (src *ClusterContentLibrary) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*imgregopv1.ClusterContentLibrary)
	if err := Convert_v1alpha1_ClusterContentLibrary_To_v1alpha2_ClusterContentLibrary(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &imgregopv1.ClusterContentLibrary{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	// BEGIN RESTORE

	dst.Spec.Type = restored.Spec.Type
	dst.Status = restored.Status

	// END RESTORE

	return nil
}

// ConvertFrom converts the hub version to this ClusterContentLibrary.
func (dst *ClusterContentLibrary) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*imgregopv1.ClusterContentLibrary)
	if err := Convert_v1alpha2_ClusterContentLibrary_To_v1alpha1_ClusterContentLibrary(src, dst, nil); err != nil {
		return err
	}

	if src.Status.SubscriptionInfo == nil {
		dst.Status.Type = ContentLibraryTypeLocal
	} else {
		dst.Status.Type = ContentLibraryTypeSubscribed
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, dst)
}

// ConvertTo converts this ClusterContentLibraryList to the Hub version.
func (src *ClusterContentLibraryList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*imgregopv1.ClusterContentLibraryList)
	return Convert_v1alpha1_ClusterContentLibraryList_To_v1alpha2_ClusterContentLibraryList(src, dst, nil)
}

// ConvertFrom converts the hub version to this ClusterContentLibraryList.
func (dst *ClusterContentLibraryList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*imgregopv1.ClusterContentLibraryList)
	return Convert_v1alpha2_ClusterContentLibraryList_To_v1alpha1_ClusterContentLibraryList(src, dst, nil)
}
