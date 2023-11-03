// Copyright (c) 2022-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

// Common ConditionTypes used by Image Registry Operator API objects.
const (
	// ReadyCondition defines the Ready condition type that summarizes the operational state of an Image Registry Operator API object.
	ReadyCondition ConditionType = "Ready"
)

// Common Condition.Reason used by Image Registry Operator API objects.
const (
	// DeletingReason (Severity=Info) documents a condition not in Status=True because the object is currently being deleted.
	DeletingReason = "Deleting"

	// DeletedReason (Severity=Error) documents a condition not in Status=True because the underlying object was deleted.
	DeletedReason = "Deleted"
)

// Condition.Reasons related to ClusterContentLibraryItem or ContentLibraryItem API objects.
const (
	ClusterContentLibraryRefValidationFailedReason = "ClusterContentLibraryRefValidationFailed"
	ContentLibraryRefValidationFailedReason        = "ContentLibraryRefValidationFailed"
	ContentLibraryItemFileUnavailableReason        = "ContentLibraryItemFileUnavailable"
)

// ConditionTypes used by ContentLibraryItemImportRequest API objects.
const (
	ContentLibraryItemImportRequestSourceValid               ConditionType = "SourceValid"
	ContentLibraryItemImportRequestTargetValid               ConditionType = "TargetValid"
	ContentLibraryItemImportRequestContentLibraryItemCreated ConditionType = "ContentLibraryItemCreated"
	ContentLibraryItemImportRequestTemplateUploaded          ConditionType = "TemplateUploaded"
	ContentLibraryItemImportRequestContentLibraryItemReady   ConditionType = "ContentLibraryItemReady"
	ContentLibraryItemImportRequestComplete                  ConditionType = "Complete"
)

// Condition.Reasons related to ContentLibraryItemImportRequest API objects.
const (
	// SourceURLInvalidReason documents that the source URL specified in the ContentLibraryItemImportRequest is invalid.
	SourceURLInvalidReason = "SourceURLInvalid"

	// SourceSchemeInvalidReason documents that the scheme in the source URL specified in the
	// ContentLibraryItemImportRequest is invalid.
	SourceSchemeInvalidReason = "SourceSchemeInvalid"

	// SourceHostInvalidReason documents that the host in the source URL specified in the
	// ContentLibraryItemImportRequest is invalid.
	SourceHostInvalidReason = "SourceHostInvalid"

	// TargetLibraryInvalidReason documents that the target ContentLibrary specified in the
	// ContentLibraryItemImportRequest is invalid.
	TargetLibraryInvalidReason = "TargetLibraryInvalid"

	// TargetLibraryItemInvalidReason documents that the specified target content library item in the
	// ContentLibraryItemImportRequest is invalid.
	TargetLibraryItemInvalidReason = "TargetLibraryItemInvalid"

	// TargetLibraryItemCreationFailureReason documents that the creation of the target ContentLibraryItem has failed.
	TargetLibraryItemCreationFailureReason = "TargetLibraryItemCreationFailure"

	// TargetLibraryItemUnavailableReason documents that the target ContentLibraryItem resource is not available in the
	// namespace yet.
	TargetLibraryItemUnavailableReason = "TargetLibraryItemUnavailable"

	// UploadInProgressReason documents that the files are in progress being uploaded to the target ContentLibraryItem.
	UploadInProgressReason = "UploadInProgress"

	// UploadFailureReason documents that uploading files to the target content library item has failed.
	UploadFailureReason = "UploadFailure"

	// UploadSessionExpiredReason documents that the uploading session for this ContentLibraryItemImportRequest has expired.
	UploadSessionExpiredReason = "UploadSessionExpired"

	// TargetLibraryItemNotReadyReason documents that the target ContentLibraryItem resource is not in ready status yet.
	TargetLibraryItemNotReadyReason = "TargetLibraryItemNotReady"
)
