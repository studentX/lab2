# <img src="../assets/lab.png" width="32" height="auto"/> Painter CRD Lab

> Leveraging Kubebuilder, define a painter CRD to color all pods in a given namespace.

1. *Coloring a pod* is defined as setting a color label on a pod equaling the CRD specific color
1. Define a CR to include a color property. Leverage the built-in enum validator
   to only include: Red, Blue and Green. Make sure to leverage annotations!
1. Implement a painter controller to monitor your painter CRD and pods as follows:
  1. When a painter CRD is created or updated, all pods in that namespace
     must be painted the CRD specified color.
  1. When a painter CRD is deleted, all pods color labels in that namespace must be removed!
  1. When a new pod is added or updated, its color label must be updated accordingly
1. Write tests to ensure your controller is operating nominally
1. In your local cluster, deploy your painter controller and ensure it exhibits the
   correct behaviors

## Commands

### Install Goland v 1.11

Ensure you have a recent [GO installation](https://vmw.gopherland.io/lessons/install)
on your local machine.

### Install kubebuilder

Please refer to the [CRD](https://vmw2.k8sland.io/lessons/crd_resource) lesson materials to install kubebuilder!

### Initialize kubebuilder

> NOTE: As of this writing Kubebuilder does NOT use go modules and required dep
> to manage GO package dependencies.

```shell
kubebuilder init --domain k8sland.io --license apache2 --owner "K8sland Training"
```

### Create your custom resource

```shell
kubebuilder create api --group workload --version v1alpha1 --kind Painter
```

### Install your Painter CRD

```shell
make install
```

### Run your Painter controller

```shell
make run
```

### Dockerize your controller

```shell
make docker-build IMG=CHANGE_ME_DOCKER_REGISTRY_USER_NAME/painter-manager:0.0.1
```

### Deploy your Docker image to a registry (DockerHub)

```shell
docker login
make docker-push IMG=CHANGE_ME_DOCKER_REGISTRY_USER_NAME/painter-manager:0.0.1
```

### Deploy your controller

```shell
make deploy
```

### Deploy your crd

```shell
kubectl apply -f config/samples/blue.yml
```

### Deploy nginx in your namespace

Using the lab template k8s directory...

```shell
kubectl apply -f k8s/nginx.yml
```

### Validate your pods are painted

```shell
kubectl get po --show-labels
```

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)