// © Broadcom. All Rights Reserved.
// The term "Broadcom" refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/google/go-cmp/cmp"
	apiequality "k8s.io/apimachinery/pkg/api/equality"
	ctrlconversion "sigs.k8s.io/controller-runtime/pkg/conversion"

	imgregopv1a1 "github.com/vmware-tanzu/image-registry-operator-api/api/v1alpha1"
	imgregopv1 "github.com/vmware-tanzu/image-registry-operator-api/api/v1alpha2"
)

func TestContentLibraryItemConversion(t *testing.T) {
	t.Run("spoke-hub", func(t *testing.T) {
		testCases := []struct {
			name  string
			hub   ctrlconversion.Hub
			after ctrlconversion.Hub
			spoke ctrlconversion.Convertible
		}{
			{
				name: "status.libraryRef.name",
				spoke: &imgregopv1a1.ContentLibraryItem{
					Status: imgregopv1a1.ContentLibraryItemStatus{
						ContentLibraryRef: &imgregopv1a1.NameAndKindRef{
							Name: "fake",
						},
					},
				},
				after: &imgregopv1.ContentLibraryItem{},
				hub: &imgregopv1.ContentLibraryItem{
					Spec: imgregopv1.ContentLibraryItemSpec{
						LibraryName: "fake",
					},
				},
			},
		}

		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				g := NewWithT(t)

				// Convert spoke to hub.
				g.Expect(tc.spoke.ConvertTo(tc.after)).To(Succeed())

				// Check that everything is equal.
				g.Expect(apiequality.Semantic.DeepEqual(tc.hub, tc.after)).To(BeTrue(), cmp.Diff(tc.hub, tc.after))
			})
		}
	})
}

func TestContentLibraryItemListConversion(t *testing.T) {
	t.Run("spoke-hub", func(t *testing.T) {
		testCases := []struct {
			name  string
			hub   ctrlconversion.Hub
			after ctrlconversion.Hub
			spoke ctrlconversion.Convertible
		}{
			{
				name: "status.libraryRef.name",
				spoke: &imgregopv1a1.ContentLibraryItemList{
					Items: []imgregopv1a1.ContentLibraryItem{
						{
							Status: imgregopv1a1.ContentLibraryItemStatus{
								ContentLibraryRef: &imgregopv1a1.NameAndKindRef{
									Name: "fake",
								},
							},
						},
					},
				},
				after: &imgregopv1.ContentLibraryItemList{},
				hub: &imgregopv1.ContentLibraryItemList{
					Items: []imgregopv1.ContentLibraryItem{
						{
							Spec: imgregopv1.ContentLibraryItemSpec{
								LibraryName: "fake",
							},
						},
					},
				},
			},
		}

		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				g := NewWithT(t)

				// Convert spoke to hub.
				g.Expect(tc.spoke.ConvertTo(tc.after)).To(Succeed())

				// Check that everything is equal.
				g.Expect(apiequality.Semantic.DeepEqual(tc.hub, tc.after)).To(BeTrue(), cmp.Diff(tc.hub, tc.after))
			})
		}
	})
}

func TestClusterContentLibraryItemConversion(t *testing.T) {
	t.Run("spoke-hub", func(t *testing.T) {
		testCases := []struct {
			name  string
			hub   ctrlconversion.Hub
			after ctrlconversion.Hub
			spoke ctrlconversion.Convertible
		}{
			{
				name: "status.libraryRef.name",
				spoke: &imgregopv1a1.ClusterContentLibraryItem{
					Status: imgregopv1a1.ContentLibraryItemStatus{
						ContentLibraryRef: &imgregopv1a1.NameAndKindRef{
							Name: "fake",
						},
					},
				},
				after: &imgregopv1.ClusterContentLibraryItem{},
				hub: &imgregopv1.ClusterContentLibraryItem{
					Spec: imgregopv1.ContentLibraryItemSpec{
						LibraryName: "fake",
					},
				},
			},
		}

		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				g := NewWithT(t)

				// Convert spoke to hub.
				g.Expect(tc.spoke.ConvertTo(tc.after)).To(Succeed())

				// Check that everything is equal.
				g.Expect(apiequality.Semantic.DeepEqual(tc.hub, tc.after)).To(BeTrue(), cmp.Diff(tc.hub, tc.after))
			})
		}
	})
}

func TestClusterContentLibraryItemListConversion(t *testing.T) {
	t.Run("spoke-hub", func(t *testing.T) {
		testCases := []struct {
			name  string
			hub   ctrlconversion.Hub
			after ctrlconversion.Hub
			spoke ctrlconversion.Convertible
		}{
			{
				name: "status.libraryRef.name",
				spoke: &imgregopv1a1.ClusterContentLibraryItemList{
					Items: []imgregopv1a1.ClusterContentLibraryItem{
						{
							Status: imgregopv1a1.ContentLibraryItemStatus{
								ContentLibraryRef: &imgregopv1a1.NameAndKindRef{
									Name: "fake",
								},
							},
						},
					},
				},
				after: &imgregopv1.ClusterContentLibraryItemList{},
				hub: &imgregopv1.ClusterContentLibraryItemList{
					Items: []imgregopv1.ClusterContentLibraryItem{
						{
							Spec: imgregopv1.ContentLibraryItemSpec{
								LibraryName: "fake",
							},
						},
					},
				},
			},
		}

		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				g := NewWithT(t)

				// Convert spoke to hub.
				g.Expect(tc.spoke.ConvertTo(tc.after)).To(Succeed())

				// Check that everything is equal.
				g.Expect(apiequality.Semantic.DeepEqual(tc.hub, tc.after)).To(BeTrue(), cmp.Diff(tc.hub, tc.after))
			})
		}
	})
}
