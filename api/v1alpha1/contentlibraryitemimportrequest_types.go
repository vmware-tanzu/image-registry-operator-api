// Copyright (c) 2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ContentLibraryItemImportRequestSource contains the specification of the source for the import request.
type ContentLibraryItemImportRequestSource struct {
	// URL is the endpoint that points to an OVF/OVA template that
	// has to be imported as a new Content Library Item in the
	// vSphere Content Library.
	// +required
	URL string `json:"url"`
}

// ContentLibraryItemImportRequestTargetItem contains the specification of the target
// content library item for the import request.
type ContentLibraryItemImportRequestTargetItem struct {
	// Name is the name of the new content library item that will be created in vSphere.
	// If omitted, the content library item will be created with the same name as the name
	// of the image specified in the spec.source.url if an item with the same name does not exist
	// in the specified Content Library in vSphere.
	// +optional
	Name string `json:"name,omitempty"`

	// Description is a description for the content library item that will be created in vSphere.
	// +optional
	Description string `json:"description,omitempty"`
}

// ContentLibraryItemImportRequestTarget is the target specification of an import request.
type ContentLibraryItemImportRequestTarget struct {
	// Item contains information about the content library item to which
	// the template will be imported in vSphere.
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
	// +required
	Source ContentLibraryItemImportRequestSource `json:"source"`

	// Target is the target of the import request which includes the content library item
	// information and a ContentLibrary resource.
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
}

// ContentLibraryItemImportRequestStatus defines the observed state of a
// ContentLibraryItemImportRequest.
type ContentLibraryItemImportRequestStatus struct {
	// ItemRef is the reference to the target ContentLibraryItem resource of the import request.
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

	// Ready is set to true only when the content library item has been
	// created and the template is imported successfully in vSphere
	// and the new ContentLibraryItem resource is ready.
	// Readiness is determined by waiting until there is status condition
	// Type=Complete and ensuring it and all other status conditions present
	// have a Status=True. The conditions present will be:
	//   * SourceValid
	//   * TargetValid
	//   * ContentLibraryItemCreated
	//   * TemplateUploaded
	//   * ContentLibraryItemReady
	//   * Complete
	//
	// +optional
	Ready bool `json:"ready,omitempty"`

	// FileUpload indicates the upload status of files belonging to the template.
	// +optional
	FileUpload *ContentLibraryItemFileUploadStatus `json:"fileUpload,omitempty"`

	// Conditions describes the current condition information of the ContentLibraryItemImportRequest.
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
// +kubebuilder:printcolumn:name="Ready",type="boolean",JSONPath=".status.ready"
// +kubebuilder:printcolumn:name="ContentLibraryRef",type="string",JSONPath=".spec.target.library.name"
// +kubebuilder:printcolumn:name="ContentLibraryItemRef",type="string",JSONPath=".status.itemRef.name"

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
