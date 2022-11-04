# image-registry-operator-api

## Overview

The image-registry-operator-api provides the object model and generated
client libraries for the VM Image Registry project, which is part of
vSphere's [Kubernetes](https://kubernetes.io) support.

Image Registry Operator allows users to manage the lifecycle of vSphere
Content Library Items using a declarative Kubernetes consumption
model and is an integral part of Project Pacific in vSphere with Kubernetes.

The state of ContentLibrary and ContentLibraryItem CRDs in a Kubernetes
cluster is monitored by the built-in controllers which reconcile the
CRD specification into Content Libraries and Content Library Items in
vSphere. The state of the Content Libraries and Content Library Items
in vSphere is reflected back in the ContentLibrary and ContentLibraryItem
CRD status.

## Try it out

This guide will walk you through the process of integrating
image-registry-operator-api into your project.

### Prerequisites

The image-registry-operator-api project currently requires a ESXi
cluster in vSphere 8 with Kubernetes.
What this means in functional terms is that you can manage workloads
in a given Workload Namespace using a Kubernetes client connected
directly to an embedded Kubernetes API Server running in the
vSphere cluster. The image registry operator APIs currently
allow you to monitor ClusterContentLibrary and ClusterContentLibraryItem
objects that exist in the target cluster and ContentLibrary and
ContentLibraryItem objects that exist in the target namespace.
The support to be able to create and manage them will be added soon.

### Testing sample code

Package envtest provides libraries for integration testing by starting
a local control plane. Control plane binaries (etcd and kube-apiserver)
are loaded by default from /usr/local/kubebuilder/bin. This can be
overridden by setting the KUBEBUILDER_ASSETS environment variable, or
by directly creating a control plane for the Environment to use.

By building the `list-ctrl` target, the tools target gets built too
and the kubebuilder tool binaries are downloaded to the `tools/bin`
folder and the KUBEBUILDER_ASSETS points to the location where the binaries
are downloaded.

### Build and Test the sample code

make list-ctrl
./hack/samples/bin/list-ctrl

## Contributing

The image-registry-operator-api project team welcomes contributions
from the community. Before you start working with image-registry-operator-api,
please read our [Developer Certificate of Origin](https://cla.vmware.com/dco).
All contributions to this repository must be signed as described on that page.
Your signature certifies that you wrote the patch or have the right to pass
it on as an open-source patch. For more detailed information,
refer to [CONTRIBUTING.md](CONTRIBUTING_CLA.md).

## License

Image Registry Operator API is licensed under the [Apache License, version 2.0](LICENSE)
