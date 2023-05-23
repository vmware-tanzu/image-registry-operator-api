// Copyright (c) 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContentLibraryItemImportRequestSource contains the specification of the source for the import request.
type ContentLibraryItemImportRequestSource struct {
	// URL is the endpoint that points to an OVF/OVA template that
	// is to be imported as a new Content Library Item in a
	// vSphere Content Library.
	// +required
	URL string `json:"url"`
}

// ContentLibraryItemImportRequestTargetItem contains the specification of the target
// content library item for the import request.
type ContentLibraryItemImportRequestTargetItem struct {
	// Name is the name of the new content library item that will be created in vSphere.
	// If omitted, the content library item will be created with the same name as the name
	// of the image specified in the spec.source.url in the specified vSphere Content Library.
	// If an item with the same name already exists in the specified vSphere Content Library,
	// the TargetValid condition will become false in the status.
	// +optional
	Name string `json:"name,omitempty"`

	// Description is a description for a vSphere Content Library Item.
	// +optional
	Description string `json:"description,omitempty"`
}

// ContentLibraryItemImportRequestTarget is the target specification of an import request.
type ContentLibraryItemImportRequestTarget struct {
	// Item contains information about the content library item to which
	// the template will be imported in vSphere.
	// If omitted, the content library item will be created with the same name as the name
	// of the image specified in the spec.source.url in the specified vSphere Content Library.
        // If an item with the same name already exists in the specified vSphere Content Library,
        // the TargetValid condition will become false in the status.
	// +optional
	Item ContentLibraryItemImportRequestTargetItem `json:"item,omitempty"`

	// Library contains information about the library in which the library item
	// will be created in vSphere.
	// +kubebuilder:default={apiVersion: imageregistry.vmware.com/v1alpha1, kind: ContentLibraryItem}
	// +required
	Library LocalObjectRef `json:"library"`
}

// ContentLibraryItemImportRequestSpec defines the desired state of a
// ContentLibraryItemImportRequest.
type ContentLibraryItemImportRequestSpec struct {
	// Source is the source of the import request which includes an external URL
	// pointing to a VM image template.
	// Source and Target will be immutable if the SourceValid and TargetValid conditions are true.
	// +required
	Source ContentLibraryItemImportRequestSource `json:"source"`

	// Target is the target of the import request which includes the content library item
	// information and a ContentLibrary resource.
	// Source and Target will be immutable if the SourceValid and TargetValid conditions are true.
	// +required
	Target ContentLibraryItemImportRequestTarget `json:"target"`

	// TTLSecondsAfterFinished is the time-to-live duration for how long this
	// resource will be allowed to exist once the import operation
	// completes. After the TTL expires, the resource will be automatically
	// deleted without the user having to take any direct action.
	// If this field is unset then the request resource will not be
	// automatically deleted. If this field is set to zero then the request
	// resource is eligible for deletion immediately after it finishes.
	// +optional
	// +kubebuilder:validation:Minimum=0
	TTLSecondsAfterFinished *int64 `json:"ttlSecondsAfterFinished,omitempty"`
}

// ContentLibraryItemFileUploadStatus indicates the upload status of files belonging to the template.
type ContentLibraryItemFileUploadStatus struct {
	// InProgress lists the names of files that are being imported into vSphere.
	// +optional
	InProgress []string `json:"inProgress,omitempty"`

	// Completed lists the names of files that have been successfully imported into vSphere.
	// +optional
	Completed []string `json:"completed,omitempty"`

	// Failed lists the names of files that failed to be imported into vSphere.
	// +optional
	Failed []string `json:"failed,omitempty"`
}

// ContentLibraryItemImportRequestStatus defines the observed state of a
// ContentLibraryItemImportRequest.
type ContentLibraryItemImportRequestStatus struct {
	// ItemRef is the reference to the target ContentLibraryItem resource of the import request.
	// If the ContentLibraryItemImportRequest is deleted when the import operation fails or before
	// the Complete condition is set to true, the import operation will be cancelled in vSphere
	// and the corresponding vSphere Content Library Item will be deleted.
	// +optional
	ItemRef *LocalObjectRef `json:"itemRef,omitempty"`

	// CompletionTime represents time when the request was completed.
	// The value of this field should be equal to the value of the
	// LastTransitionTime for the status condition Type=Complete.
	// +optional
	CompletionTime metav1.Time `json:"completionTime,omitempty"`

	// StartTime represents time when the request was acknowledged by the
	// controller.
	// +optional
	StartTime metav1.Time `json:"startTime,omitempty"`

	// FileUpload indicates the upload status of files belonging to the template.
	// +optional
	FileUploadStatus *ContentLibraryItemFileUploadStatus `json:"fileUploadStatus,omitempty"`

	// Conditions describes the current condition information of the ContentLibraryItemImportRequest.
	// The conditions present will be:
	//   * SourceValid
	//   * TargetValid
	//   * ContentLibraryItemCreated
	//   * TemplateUploaded
	//   * ContentLibraryItemReady
	//   * Complete
	// +optional
	Conditions []Condition `json:"conditions,omitempty"`
}

func (clItemImportRequest *ContentLibraryItemImportRequest) GetConditions() Conditions {
	return clItemImportRequest.Status.Conditions
}

func (clItemImportRequest *ContentLibraryItemImportRequest) SetConditions(conditions Conditions) {
	clItemImportRequest.Status.Conditions = conditions
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced,shortName=clitemimport
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:name="ContentLibraryRef",type="string",JSONPath=".spec.target.library.name"
// +kubebuilder:printcolumn:name="ContentLibraryItemRef",type="string",JSONPath=".status.itemRef.name"
// +kubebuilder:printcolumn:name="Ready",type="string",JSONPath=".status.conditions[?(.type=='Complete')].status"

// ContentLibraryItemImportRequest defines the information necessary to import a VM image
// template as a ContentLibraryItem to a Content Library in vSphere.
type ContentLibraryItemImportRequest struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ContentLibraryItemImportRequestSpec   `json:"spec,omitempty"`
	Status ContentLibraryItemImportRequestStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ContentLibraryItemImportRequestList contains a list of
// ContentLibraryItemImportRequest resources.
type ContentLibraryItemImportRequestList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ContentLibraryItemImportRequest `json:"items"`
}

func init() {
	RegisterTypeWithScheme(&ContentLibraryItemImportRequest{}, &ContentLibraryItemImportRequestList{})
}
