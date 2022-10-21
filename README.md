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
cluster in vSphere 7 with Kubernetes.
What this means in functional terms is that you can manage workloads
in a given Workload Namespace using a Kubernetes client connected
directly to the an embedded Kubernetes API Server running in the
vSphere cluster. The image-registry-operator-api APIs currently
allow you to monitor ClusterContentLibrary and ClusterContentLibraryItem
objects that exist in the target cluster and ContentLibrary and
ContentLibraryItem objects that exist in the target namespace.
The support to be able to create and manager them will be added soon.

### Install kubebuilder tools

Install the latest stable version of
[kubebuilder tools](https://storage.googleapis.com/kubebuilder-tools/)

### Check your client access by viewing ContentLibraries in the target cluster

```bash
kubectl get contentlibraries
```

### Build and Test the sample code

cd hack/samples
make all
bin/list-ctrl

## Contributing

The image-registry-operator-api project team welcomes contributions
from the community. Before you start working with image-registry-operator-api,
please read our [Developer Certificate of Origin](https://cla.vmware.com/dco).
All contributions to this repository must be signed as described on that page.
Your signature certifies that you wrote the patch or have the right to pass
it on as an open-source patch. For more detailed information,
refer to [CONTRIBUTING.md](CONTRIBUTING.md).

## License

Image Registry Operator API is licensed under the [Apache License, version 2.0](LICENSE)
