# <img src="../assets/lab.png" width="32" height="auto"/> Kubebuider Lab

> Deploy a CRD using kubebuilder.

In this lab you are going to install and setup kubebuilder in your local development
environment. You are going to define and provision a sample CRD called containerset.
A containerset allows you to define a deployment by merely specifying an image and a
replica count.

> NOTE! The current version of kubebuilder cans the image and replicaCount to be
> nginx and 1 which is ok for this initial lab.

> NOTE! Sadly, as of this writing, kubebuilder does not leverage GO modules.
> So if you are running GO >= 1.11, we will need to turn modules off as
> kubebuilder currently leverages dep for package management

1. Install GO

The official instructions are [Here](https://golang.org/doc/install)

1. Install Kubebuilder

    ```shell
    export KB_ARCH=amd64
    export KB_REV=1.0.4

    # Download the release for your architecture
    cd Downloads
    curl -LO https://github.com/kubernetes-sigs/kubebuilder/releases/download/v${KB_REV}/kubebuilder_${KB_REV}_darwin_${KB_ARCH}.tar.gz
    # Extract the archive and setup directory
    tar -zxvf ~/Downloads/kubebuilder_${KB_REV}_darwin_${KB_ARCH}.tar.gz && sudo mv kubebuilder_${KB_REV}_darwin_${KB_ARCH} /usr/local/kubebuilder

    # update your PATH to include /usr/local/kubebuilder/bin
    # NOTE: We recommend setting this up in your terminal profile for your shell
    export PATH=$PATH:/usr/local/kubebuilder/bin

    # Verify!
    kubebuilder --help
    ```

1. Install Kustomize

  Kubebuilder requires [Kustomize](https://github.com/kubernetes-sigs/kustomize)
  to manipulate your deployment artifacts. Kustomize version must be > 1.0.4

    ```shell
    # For OSX install use homebrew
    brew install kustomize
    # For other platforms please see install instruction on the link above!
    ```

1. Setup a GO workspace

    ```shell
    md -p $HOME/k8sland/src
    export GOPATH=$HOME/k8sland
    export GOBIN=$GOPATH/bin
    export GO111MODULE=off
    export PATH=$PATH:$GOBIN
    ```

1. Install Dep for GO dependency management

  Please follow the installation recipes [Here](https://github.com/golang/dep)

1. Initialize KubeBuilder and install dependencies

    > NOTE! This will take a while...

    ```shell
    kubebuilder init --domain k8sland.io --license apache2 --owner "K8sland Training"
    ```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)