<img src="../../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../../assets/lab.png" width="32" height="auto"/> ClusterDepot CRD Lab

> Build a ClusterDepot CRD -- `How Sudoers Get More Done!`

Leveraging Kubebuilder, define a painter CRD to color all pods in that CRD's namespace.

NOTE!: If you have GO chops, checkout the go branch and write some of the CRD code!!

1. *Coloring a pod* is defined as setting a color label on a pod equaling
   the CRD specific color
2. Create Painter CR schema to include a color property.
   1. Be sure to leverage the built-in **Enum** validator to only include colors Red, Blue and Green.
   2. Using annotations, ensure the CR is namespaced
   3. Using annotations, set a shortName for it ie `pt`
3. Implement a painter controller to monitor your painter CRD and pods:
   1. When a painter CRD is created or updated, all pods in that namespace
     must be painted the CRD specified color.
   2. When a painter CRD is deleted, all pods color labels in that namespace
     must be removed!
4. Install your CRD schema on your cluster
5. Launch your controller locally and make sure the pods are getting painted/unpainted!
6. Next, setup RBAC policies for your controller
7. Build your controller docker image
8. Deploy your painter controller on your cluster and ensure it exhibits the correct behaviors.

## Commands

1. Create the CRD

   ```shell
   kubebuilder create api --group clusterdepot --version v1alpha1 --kind Painter --resource --controller
   ```

1. Install the CRD

   ```shell
   make install
   ```

1. Run the CRD (local)

   ```shell
   make run
   ```

1. Create a Docker image

   ```shell
   make docker-img IMG=MY_REPO:MY_REV
   ```

1. Run the controller in cluster

   ```shell
   make deploy IMG=MY_REPO:MY_REV
   ```

1. Uninstall the CRD

   ```shell
   make uninstall
   ```

---
<img src="../../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)