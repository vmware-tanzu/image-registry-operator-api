// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha2

// Hub marks ContentLibraryItem as a conversion hub.
func (*ContentLibraryItem) Hub() {}

// Hub marks ContentLibraryItemList as a conversion hub.
func (*ContentLibraryItemList) Hub() {}

// Hub marks ClusterContentLibraryItem as a conversion hub.
func (*ClusterContentLibraryItem) Hub() {}

// Hub marks ClusterContentLibraryItemList as a conversion hub.
func (*ClusterContentLibraryItemList) Hub() {}
