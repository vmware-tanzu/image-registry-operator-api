// Copyright (c) 2022-2023 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/vmware-tanzu/image-registry-operator-api/api/v1alpha1"

	core "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	k8sRuntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"

	ctrlClient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

var testEnv *envtest.Environment

// List ContentLibraries in a target namespace to stdout using a controller client
func main() {
	fmt.Printf("Starting test env...\n")
	testClient, err := startTestEnv()
	if err != nil {
		panic(err)
	}
	defer func() {
		fmt.Printf("Stopping test env...\n")
		testEnv.Stop()
	}()

	imgRegOpClient, err := getImgRegOpClient(testClient)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Populating test env...\n")
	fmt.Printf("Creating test namespace...\n")
	namespaceName := "test-ns"
	err = createTestNamespace(imgRegOpClient, namespaceName)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Creating content library resources...\n")
	err = populateTestEnv(imgRegOpClient, "test-cl-1", namespaceName)
	if err != nil {
		panic(err)
	}
	err = populateTestEnv(imgRegOpClient, "test-cl-2", namespaceName)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Listing Content Libraries:\n")
	clList := v1alpha1.ContentLibraryList{}
	err = imgRegOpClient.List(context.TODO(), &clList)
	if err != nil {
		panic(err)
	}
	for _, cl := range clList.Items {
		fmt.Printf("- %s\n", cl.GetName())
	}
}

// Get a image-registry-operator-api client from the generated clientset
func getImgRegOpClient(config *rest.Config) (ctrlClient.Client, error) {
	scheme := k8sRuntime.NewScheme()
	_ = v1alpha1.AddToScheme(scheme)
	_ = core.AddToScheme(scheme)
	client, err := ctrlClient.New(config, ctrlClient.Options{
		Scheme: scheme,
	})
	return client, err
}

func startTestEnv() (*rest.Config, error) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	crd_filepath := filepath.Join(dir, "..", "..", "..", "config", "crd", "bases")
	kubebuilder_tools_filepath := filepath.Join(dir, "..", "..", "tools", "bin")
	os.Setenv("KUBEBUILDER_ASSETS", kubebuilder_tools_filepath)
	testEnv = &envtest.Environment{
		CRDDirectoryPaths: []string{
			crd_filepath,
		},
	}
	return testEnv.Start()
}

func populateTestEnv(client ctrlClient.Client, name string, namespace string) error {
	newCL := v1alpha1.ContentLibrary{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
	return client.Create(context.TODO(), &newCL)
}

func createTestNamespace(client ctrlClient.Client, namespace string) error {
	ns := &core.Namespace{}
	*ns = core.Namespace{
		ObjectMeta: metav1.ObjectMeta{Name: namespace},
	}
	return client.Create(context.TODO(), ns)
}
