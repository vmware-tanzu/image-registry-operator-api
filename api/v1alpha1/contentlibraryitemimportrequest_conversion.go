// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	apiconversion "k8s.io/apimachinery/pkg/conversion"
	ctrlconversion "sigs.k8s.io/controller-runtime/pkg/conversion"

	"github.com/vmware-tanzu/image-registry-operator-api/api/utilconversion"
	imgregopv1 "github.com/vmware-tanzu/image-registry-operator-api/api/v1alpha2"
)

func Convert_v1alpha1_ContentLibraryItemImportRequestStatus_To_v1alpha2_ContentLibraryItemImportRequestStatus(in *ContentLibraryItemImportRequestStatus, out *imgregopv1.ContentLibraryItemImportRequestStatus, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha1_ContentLibraryItemImportRequestStatus_To_v1alpha2_ContentLibraryItemImportRequestStatus(in, out, s); err != nil {
		return err
	}
	if in.ItemRef != nil {
		out.ItemName = in.ItemRef.Name
	}
	return nil
}

func Convert_v1alpha2_ContentLibraryItemImportRequestStatus_To_v1alpha1_ContentLibraryItemImportRequestStatus(in *imgregopv1.ContentLibraryItemImportRequestStatus, out *ContentLibraryItemImportRequestStatus, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha2_ContentLibraryItemImportRequestStatus_To_v1alpha1_ContentLibraryItemImportRequestStatus(in, out, s); err != nil {
		return err
	}
	if in.ItemName != "" {
		out.ItemRef = &LocalObjectRef{
			APIVersion: GroupVersion.String(),
			Kind:       "ContentLibraryItem",
			Name:       in.ItemName,
		}
	}
	return nil
}

func Convert_v1alpha1_ContentLibraryItemImportRequestTarget_To_v1alpha2_ContentLibraryItemImportRequestTarget(in *ContentLibraryItemImportRequestTarget, out *imgregopv1.ContentLibraryItemImportRequestTarget, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha1_ContentLibraryItemImportRequestTarget_To_v1alpha2_ContentLibraryItemImportRequestTarget(in, out, s); err != nil {
		return err
	}
	out.LibraryName = in.Library.Name
	return nil
}

func Convert_v1alpha2_ContentLibraryItemImportRequestTarget_To_v1alpha1_ContentLibraryItemImportRequestTarget(in *imgregopv1.ContentLibraryItemImportRequestTarget, out *ContentLibraryItemImportRequestTarget, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha2_ContentLibraryItemImportRequestTarget_To_v1alpha1_ContentLibraryItemImportRequestTarget(in, out, s); err != nil {
		return err
	}
	out.Library.APIVersion = GroupVersion.String()
	out.Library.Kind = "ContentLibrary"
	out.Library.Name = in.LibraryName
	return nil
}

func Convert_v1alpha1_FileTransferStatus_To_v1alpha2_FileTransferStatus(in *FileTransferStatus, out *imgregopv1.FileTransferStatus, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha1_FileTransferStatus_To_v1alpha2_FileTransferStatus(in, out, s); err != nil {
		return err
	}
	switch in.Status {
	case TransferStatusError:
		out.Status = imgregopv1.TransferStatusError
	case TransferStatusReady:
		out.Status = imgregopv1.TransferStatusReady
	case TransferStatusTransferring:
		out.Status = imgregopv1.TransferStatusTransferring
	case TransferStatusValidating:
		out.Status = imgregopv1.TransferStatusValidating
	case TransferStatusWaiting:
		out.Status = imgregopv1.TransferStatusWaiting
	default:
		out.Status = imgregopv1.TransferStatus(in.Status)
	}
	return nil
}

func Convert_v1alpha2_FileTransferStatus_To_v1alpha1_FileTransferStatus(in *imgregopv1.FileTransferStatus, out *FileTransferStatus, s apiconversion.Scope) error {
	if err := autoConvert_v1alpha2_FileTransferStatus_To_v1alpha1_FileTransferStatus(in, out, s); err != nil {
		return err
	}
	switch in.Status {
	case imgregopv1.TransferStatusError:
		out.Status = TransferStatusError
	case imgregopv1.TransferStatusReady:
		out.Status = TransferStatusReady
	case imgregopv1.TransferStatusTransferring:
		out.Status = TransferStatusTransferring
	case imgregopv1.TransferStatusValidating:
		out.Status = TransferStatusValidating
	case imgregopv1.TransferStatusWaiting:
		out.Status = TransferStatusWaiting
	default:
		out.Status = ""
	}
	return nil
}

// ConvertTo converts this ContentLibraryItemImportRequest to the Hub version.
func (src *ContentLibraryItemImportRequest) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*imgregopv1.ContentLibraryItemImportRequest)
	if err := Convert_v1alpha1_ContentLibraryItemImportRequest_To_v1alpha2_ContentLibraryItemImportRequest(src, dst, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &imgregopv1.ContentLibraryItemImportRequest{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	// BEGIN RESTORE

	dst.Status = restored.Status

	// END RESTORE

	return nil
}

// ConvertFrom converts the hub version to this ContentLibraryItemImportRequest.
func (dst *ContentLibraryItemImportRequest) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*imgregopv1.ContentLibraryItemImportRequest)
	if err := Convert_v1alpha2_ContentLibraryItemImportRequest_To_v1alpha1_ContentLibraryItemImportRequest(src, dst, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, dst)
}

// ConvertTo converts this ContentLibraryItemImportRequestList to the Hub version.
func (src *ContentLibraryItemImportRequestList) ConvertTo(dstRaw ctrlconversion.Hub) error {
	dst := dstRaw.(*imgregopv1.ContentLibraryItemImportRequestList)
	return Convert_v1alpha1_ContentLibraryItemImportRequestList_To_v1alpha2_ContentLibraryItemImportRequestList(src, dst, nil)
}

// ConvertFrom converts the hub version to this ContentLibraryItemImportRequestList.
func (dst *ContentLibraryItemImportRequestList) ConvertFrom(srcRaw ctrlconversion.Hub) error {
	src := srcRaw.(*imgregopv1.ContentLibraryItemImportRequestList)
	return Convert_v1alpha2_ContentLibraryItemImportRequestList_To_v1alpha1_ContentLibraryItemImportRequestList(src, dst, nil)
}
