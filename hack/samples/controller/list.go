// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"context"
	"fmt"
	"path/filepath"

	"github.com/vmware-tanzu/image-registry-operator-api/api/v1alpha1"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/rest"

	ctrlClient "sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/envtest"
)

var testEnv *envtest.Environment

const namespace string = "default"

// List ContentLibraries in a target cluster to stdout using a controller client
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
	err = populateTestEnv(imgRegOpClient, "test-cl-1")
	if err != nil {
		panic(err)
	}
	err = populateTestEnv(imgRegOpClient, "test-cl-2")
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
	scheme := runtime.NewScheme()
	_ = v1alpha1.AddToScheme(scheme)
	client, err := ctrlClient.New(config, ctrlClient.Options{
		Scheme: scheme,
	})
	return client, err
}

func startTestEnv() (*rest.Config, error) {
	testEnv = &envtest.Environment{
		CRDDirectoryPaths: []string{
			filepath.Join("..", "..", "config", "crd", "bases"),
		},
	}

	return testEnv.Start()
}

func populateTestEnv(client ctrlClient.Client, name string) error {
	newCL := v1alpha1.ContentLibrary{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
	return client.Create(context.TODO(), &newCL)
}
