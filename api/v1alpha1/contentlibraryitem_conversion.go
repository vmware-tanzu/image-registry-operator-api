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

func Convert_v1alpha1_ContentLibraryItemSpec_To_v1alpha2_ContentLibraryItemSpec(in *ContentLibraryItemSpec, out *imgregopv1.ContentLibraryItemSpec, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha1_ContentLibraryItemSpec_To_v1alpha2_ContentLibraryItemSpec(in, out, s); err != nil {
		return err
	}
	out.ID = string(in.UUID)
	return nil
}

func Convert_v1alpha2_ContentLibraryItemSpec_To_v1alpha1_ContentLibraryItemSpec(in *imgregopv1.ContentLibraryItemSpec, out *ContentLibraryItemSpec, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha2_ContentLibraryItemSpec_To_v1alpha1_ContentLibraryItemSpec(in, out, s); err != nil {
		return err
	}
	out.UUID = types.UID(in.ID)
	return nil
}

func Convert_v1alpha1_ContentLibraryItemStatus_To_v1alpha2_ContentLibraryItemStatus(in *ContentLibraryItemStatus, out *imgregopv1.ContentLibraryItemStatus, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha1_ContentLibraryItemStatus_To_v1alpha2_ContentLibraryItemStatus(in, out, s); err != nil {
		return err
	}
	out.Version = in.MetadataVersion
	return nil
}

func Convert_v1alpha2_ContentLibraryItemStatus_To_v1alpha1_ContentLibraryItemStatus(in *imgregopv1.ContentLibraryItemStatus, out *ContentLibraryItemStatus, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha2_ContentLibraryItemStatus_To_v1alpha1_ContentLibraryItemStatus(in, out, s); err != nil {
		return err
	}
	out.MetadataVersion = in.Version
	return nil
}

func Convert_v1alpha1_FileInfo_To_v1alpha2_FileInfo(in *FileInfo, out *imgregopv1.FileInfo, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha1_FileInfo_To_v1alpha2_FileInfo(in, out, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha2_FileInfo_To_v1alpha1_FileInfo(in *imgregopv1.FileInfo, out *FileInfo, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha2_FileInfo_To_v1alpha1_FileInfo(in, out, s); err != nil {
		return err
	}
	return nil
}

func Convert_v1alpha1_CertificateVerificationInfo_To_v1alpha2_CertificateVerificationInfo(in *CertificateVerificationInfo, out *imgregopv1.CertificateVerificationInfo, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha1_CertificateVerificationInfo_To_v1alpha2_CertificateVerificationInfo(in, out, s); err != nil {
		return err
	}
	switch in.Status {
	case CertVerificationStatusInternal:
		out.Status = imgregopv1.CertVerificationStatusInternal
	case CertVerificationStatusNotAvailable:
		out.Status = imgregopv1.CertVerificationStatusNotAvailable
	case CertVerificationStatusUntrusted:
		out.Status = imgregopv1.CertVerificationStatusUntrusted
	case CertVerificationStatusVerificationFailure:
		out.Status = imgregopv1.CertVerificationStatusVerificationFailure
	case CertVerificationStatusVerificationInProgress:
		out.Status = imgregopv1.CertVerificationStatusVerificationInProgress
	case CertVerificationStatusVerified:
		out.Status = imgregopv1.CertVerificationStatusVerified
	default:
		out.Status = imgregopv1.CertVerificationStatus(in.Status)
	}
	return nil
}

func Convert_v1alpha2_CertificateVerificationInfo_To_v1alpha1_CertificateVerificationInfo(in *imgregopv1.CertificateVerificationInfo, out *CertificateVerificationInfo, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha2_CertificateVerificationInfo_To_v1alpha1_CertificateVerificationInfo(in, out, s); err != nil {
		return err
	}
	switch in.Status {
	case imgregopv1.CertVerificationStatusInternal:
		out.Status = CertVerificationStatusInternal
	case imgregopv1.CertVerificationStatusNotAvailable:
		out.Status = CertVerificationStatusNotAvailable
	case imgregopv1.CertVerificationStatusUntrusted:
		out.Status = CertVerificationStatusUntrusted
	case imgregopv1.CertVerificationStatusVerificationFailure:
		out.Status = CertVerificationStatusVerificationFailure
	case imgregopv1.CertVerificationStatusVerificationInProgress:
		out.Status = CertVerificationStatusVerificationInProgress
	case imgregopv1.CertVerificationStatusVerified:
		out.Status = CertVerificationStatusVerified
	default:
		out.Status = CertVerificationStatus(in.Status)
	}
	return nil
}

// ConvertTo converts this ContentLibraryItem to the Hub version.
func (src *ContentLibraryItem) ConvertTo(dstRaw ctrlconversion.Hub) error {

	dst := dstRaw.(*imgregopv1.ContentLibraryItem)
	if err := Convert_v1alpha1_ContentLibraryItem_To_v1alpha2_ContentLibraryItem(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &imgregopv1.ContentLibraryItem{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	// BEGIN RESTORE

	dst.Spec.LibraryName = restored.Spec.LibraryName
	dst.Status = restored.Status

	// END RESTORE

	if dst.Spec.LibraryName == "" && src.Status.ContentLibraryRef != nil {
		dst.Spec.LibraryName = src.Status.ContentLibraryRef.Name
	}

	return nil
}

// ConvertFrom converts the hub version to this ContentLibraryItem.
func (dst *ContentLibraryItem) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*imgregopv1.ContentLibraryItem)
	if err := Convert_v1alpha2_ContentLibraryItem_To_v1alpha1_ContentLibraryItem(src, dst, nil); err != nil {
		return err
	}

	dst.Status.ContentLibraryRef = &NameAndKindRef{
		Kind: "ContentLibrary",
		Name: src.Spec.LibraryName,
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, dst)
}

// ConvertTo converts this ContentLibraryItemList to the Hub version.
func (src *ContentLibraryItemList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*imgregopv1.ContentLibraryItemList)
	return Convert_v1alpha1_ContentLibraryItemList_To_v1alpha2_ContentLibraryItemList(src, dst, nil)
}

// ConvertFrom converts the hub version to this ContentLibraryItemList.
func (dst *ContentLibraryItemList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*imgregopv1.ContentLibraryItemList)
	return Convert_v1alpha2_ContentLibraryItemList_To_v1alpha1_ContentLibraryItemList(src, dst, nil)
}

// ConvertTo converts this ClusterContentLibraryItem to the Hub version.
func (src *ClusterContentLibraryItem) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*imgregopv1.ClusterContentLibraryItem)
	if err := Convert_v1alpha1_ClusterContentLibraryItem_To_v1alpha2_ClusterContentLibraryItem(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &imgregopv1.ClusterContentLibraryItem{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	// BEGIN RESTORE

	dst.Spec.LibraryName = restored.Spec.LibraryName
	dst.Status = restored.Status

	// END RESTORE

	if dst.Spec.LibraryName == "" && src.Status.ContentLibraryRef != nil {
		dst.Spec.LibraryName = src.Status.ContentLibraryRef.Name
	}

	return nil
}

// ConvertFrom converts the hub version to this ClusterContentLibraryItem.
func (dst *ClusterContentLibraryItem) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*imgregopv1.ClusterContentLibraryItem)
	if err := Convert_v1alpha2_ClusterContentLibraryItem_To_v1alpha1_ClusterContentLibraryItem(src, dst, nil); err != nil {
		return err
	}

	dst.Status.ContentLibraryRef = &NameAndKindRef{
		Kind: "ClusterContentLibrary",
		Name: src.Spec.LibraryName,
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, dst)
}

// ConvertTo converts this ClusterContentLibraryItemList to the Hub version.
func (src *ClusterContentLibraryItemList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*imgregopv1.ClusterContentLibraryItemList)
	return Convert_v1alpha1_ClusterContentLibraryItemList_To_v1alpha2_ClusterContentLibraryItemList(src, dst, nil)
}

// ConvertFrom converts the hub version to this ClusterContentLibraryItemList.
func (dst *ClusterContentLibraryItemList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*imgregopv1.ClusterContentLibraryItemList)
	return Convert_v1alpha2_ClusterContentLibraryItemList_To_v1alpha1_ClusterContentLibraryItemList(src, dst, nil)
}
