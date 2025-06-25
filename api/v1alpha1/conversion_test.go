// © Broadcom. All Rights Reserved.
// The term “Broadcom” refers to Broadcom Inc. and/or its subsidiaries.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1_test

import (
	"testing"

	. "github.com/onsi/gomega"
	"k8s.io/apimachinery/pkg/api/apitesting/fuzzer"
	"k8s.io/apimachinery/pkg/runtime"
	runtimeserializer "k8s.io/apimachinery/pkg/runtime/serializer"
	"sigs.k8s.io/controller-runtime/pkg/conversion"
	"sigs.k8s.io/randfill"

	"github.com/vmware-tanzu/image-registry-operator-api/api/utilconversion"
	imgregopv1a1 "github.com/vmware-tanzu/image-registry-operator-api/api/v1alpha1"
	imgregopv1 "github.com/vmware-tanzu/image-registry-operator-api/api/v1alpha2"
)

func TestFuzzyConversion(t *testing.T) {

	testCases := []struct {
		name        string
		hub         conversion.Hub
		spoke       conversion.Convertible
		fuzzerFuncs func(runtimeserializer.CodecFactory) []interface{}
	}{
		{
			name: "ContentLibrary",
			hub:  &imgregopv1.ContentLibrary{},
			spoke: &imgregopv1a1.ContentLibrary{
				Status: imgregopv1a1.ContentLibraryStatus{
					Type: imgregopv1a1.ContentLibraryTypeLocal,
				},
			},
			fuzzerFuncs: overrideContentLibraryFieldFuncs,
		},
		{
			name: "ClusterContentLibrary",
			hub:  &imgregopv1.ClusterContentLibrary{},
			spoke: &imgregopv1a1.ClusterContentLibrary{
				Status: imgregopv1a1.ContentLibraryStatus{
					Type: imgregopv1a1.ContentLibraryTypeLocal,
				},
			},
			fuzzerFuncs: overrideContentLibraryFieldFuncs,
		},
		{
			name:        "ContentLibraryItem",
			hub:         &imgregopv1.ContentLibraryItem{},
			spoke:       &imgregopv1a1.ContentLibraryItem{},
			fuzzerFuncs: overrideContentLibraryItemFieldFuncs,
		},

		{
			name:        "ClusterContentLibraryItem",
			hub:         &imgregopv1.ClusterContentLibraryItem{},
			spoke:       &imgregopv1a1.ClusterContentLibraryItem{},
			fuzzerFuncs: overrideClusterContentLibraryItemFieldFuncs,
		},
		{
			name:        "ContentLibraryItemImportRequest",
			hub:         &imgregopv1.ContentLibraryItemImportRequest{},
			spoke:       &imgregopv1a1.ContentLibraryItemImportRequest{},
			fuzzerFuncs: overrideContentLibraryItemImportRequestFieldFuncs,
		},
	}

	for i := range testCases {
		tc := testCases[i]

		g := NewWithT(t)
		scheme := runtime.NewScheme()
		g.Expect(imgregopv1a1.AddToScheme(scheme)).To(Succeed())
		g.Expect(imgregopv1.AddToScheme(scheme)).To(Succeed())

		input := utilconversion.FuzzTestFuncInput{
			Scheme: scheme,
			Hub:    tc.hub.DeepCopyObject().(conversion.Hub),
			Spoke:  tc.spoke.DeepCopyObject().(conversion.Convertible),
			FuzzerFuncs: []fuzzer.FuzzerFuncs{
				tc.fuzzerFuncs,
			},
		}

		t.Run(tc.name, func(t *testing.T) {
			t.Run("Spoke-Hub-Spoke", func(t *testing.T) {
				t.Parallel()
				utilconversion.SpokeHubSpoke(NewWithT(t), input)
			})
			t.Run("Hub-Spoke-Hub", func(t *testing.T) {
				t.Parallel()
				utilconversion.HubSpokeHub(NewWithT(t), input)
			})
		})
	}
}

func overrideContentLibraryFieldFuncs(_ runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{
		func(status *imgregopv1a1.ContentLibraryStatus, c randfill.Continue) {
			c.Fill(status)
			overrideConditionsSeverity(status.Conditions)
		},
	}
}

func overrideContentLibraryItemFieldFuncs(_ runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{
		func(status *imgregopv1a1.ContentLibraryItemStatus, c randfill.Continue) {
			c.Fill(status)
			overrideConditionsSeverity(status.Conditions)
			if status.ContentLibraryRef == nil {
				status.ContentLibraryRef = &imgregopv1a1.NameAndKindRef{
					Kind: "ContentLibrary",
				}
			}
		},
	}
}

func overrideClusterContentLibraryItemFieldFuncs(_ runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{
		func(status *imgregopv1a1.ContentLibraryItemStatus, c randfill.Continue) {
			c.Fill(status)
			overrideConditionsSeverity(status.Conditions)
			if status.ContentLibraryRef == nil {
				status.ContentLibraryRef = &imgregopv1a1.NameAndKindRef{
					Kind: "ClusterContentLibrary",
				}
			}
		},
	}
}

func overrideContentLibraryItemImportRequestFieldFuncs(_ runtimeserializer.CodecFactory) []interface{} {
	return []interface{}{
		func(status *imgregopv1a1.ContentLibraryItemImportRequestStatus, c randfill.Continue) {
			c.Fill(status)
			overrideConditionsSeverity(status.Conditions)
		},
		func(spec *imgregopv1a1.ContentLibraryItemImportRequestSpec, c randfill.Continue) {
			c.Fill(spec)
			spec.Target.Library.APIVersion = imgregopv1a1.GroupVersion.String()
			spec.Target.Library.Kind = "ContentLibrary"
		},
	}
}

func overrideConditionsSeverity(conditions []imgregopv1a1.Condition) {
	// metav1.Conditions do not have this field, so on down conversions it will
	// always be empty.
	for i := range conditions {
		conditions[i].Severity = ""
	}
}
