// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha2

// Hub marks ContentLibrary as a conversion hub.
func (*ContentLibrary) Hub() {}

// Hub marks ContentLibraryList as a conversion hub.
func (*ContentLibraryList) Hub() {}

// Hub marks ClusterContentLibrary as a conversion hub.
func (*ClusterContentLibrary) Hub() {}

// Hub marks ClusterContentLibraryList as a conversion hub.
func (*ClusterContentLibraryList) Hub() {}
