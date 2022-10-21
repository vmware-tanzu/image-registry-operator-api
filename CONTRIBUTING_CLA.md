# Contributing to image-registry-operator-api

We welcome contributions from the community and first
want to thank you for taking the time to contribute!

Please familiarize yourself with the
[Code of Conduct](https://github.com/vmware/.github/blob/main/CODE_OF_CONDUCT.md)
before contributing.

Before you start working with image-registry-operator-api, please read and
sign our Contributor License Agreement [CLA](https://cla.vmware.com/cla/1/preview).
If you wish to contribute code and you have not signed our contributor
license agreement (CLA), our bot will prompt you to do so when you open a Pull Request.
For any questions about the CLA process, please refer to our
[FAQ]([https://cla.vmware.com/faq](https://cla.vmware.com/faq)).

## Ways to contribute

We welcome many different types of contributions and not all of them
need a Pull request. Contributions may include:

* New features and proposals
* Documentation
* Bug fixes
* Issue Triage
* Answering questions and giving feedback
* Helping to onboard new contributors
* Other related activities

## Getting started

The image-registry-operator-api project currently requires a ESXi
cluster in vSphere 7 with Kubernetes.

What this means in functional terms is that you can manage workloads
in a given Workload Namespace using a Kubernetes client connected
directly to the an embedded Kubernetes API Server running in the
vSphere cluster. The image-registry-operator-api APIs currently allow
you to monitor ClusterContentLibrary and ClusterContentLibraryItem
objects that exist in the target cluster and ContentLibrary and
ContentLibraryItem objects that exist in the target namespace. The
support to be able to create and manager them will be added soon.

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

## Contribution Flow

This is a rough outline of what a contributor's workflow looks like:

* Make a fork of the repository within your GitHub account
* Create a topic branch in your fork from where you want to base your work
* Make commits of logical units
* Make sure your commit messages are with the proper format,
  quality and descriptiveness (see below)
* Push your changes to the topic branch in your fork
* Create a pull request containing that commit

We follow the GitHub workflow and you can find more details on the
[GitHub flow documentation](https://docs.github.com/en/get-started/quickstart/github-flow).

### Pull Request Checklist

Before submitting your pull request, we advise you to use the following:

* Check if your code changes will pass both code linting checks and unit tests.
* Ensure your commit messages are descriptive. We follow the conventions on
  [How to Write a Git Commit Message](http://chris.beams.io/posts/git-commit/).
  Be sure to include any related GitHub issue references in the commit message.
  See [GFM syntax](https://guides.github.com/features/mastering-markdown/#GitHub-flavored-markdown)
  for referencing issues and commits.
* Check the commits and commits messages and ensure they are free from typos.

## Reporting Bugs and Creating Issues

For specifics on what to include in your report, please follow the guidelines
in the issue and pull request templates when available.

Given that the image-registry-operator-api is currently the API surface to a
larger project, please be aware that we must consider the wider impact of
any API changes.

## Ask for Help

The best way to reach us with a question when contributing is to ask on:

* The original GitHub issue

