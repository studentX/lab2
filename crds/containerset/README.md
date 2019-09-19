<img src="../../../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../../../assets/lab.png" width="32" height="auto"/> CRD Controller Lab

> Using KubeBuilder, create a ContainerSet CRD

A ContainerSet CRD defines an image and a number of replicas and ensures a
deployment is created in the current namespace with the image/replicas specified
by the CRD.

> NOTE!! We will use an already pre-generated ContainerSet for this lab but you will
> have to tweak a few things prior to deploying the CRD on your cluster
> the right thing in the next lab. For now the default generated code ensures
> an nginx deployment is created with a single replicas.

1. Before you begin, make sure to install GO and kubebuilder per the previous lesson's instructions!
   The ContainerSet has already been generated and implemented for you with a few caveats...

2. Clone the [K8sland Lab Repo](https://github.com/k8sland/lab2)

    ```shell
    git clone https://github.com/k8sland/lab2.git
    ```

3. Cd to crds/containerset

4. Fill in the ContainerSet CRD manifest in config/samples/zorg.yaml
   1. Specify the container image nginx:1.17
   2. Specify 1 replica

5. Verify your enviroment is cool by installing the CRD schema and running the controller locally

    ```shell
    make install
    # Verify!
    kubectl get crd # Ensure containersets.workload.k8sland.io exists!
    ```

6. Run the ContainerSet controller

    ```shell
    make run
    ```

7. Watch your local pods and deployments

    ```shell
    watch kubectl get po,deploy
    ```

8. Instantiate the Zorg CRD

   ```shell
   kubectl apply -f config/samples/zorg.yaml
   # In your watch window you should see a new zorg deployment and one zorg pod
   ```

9. Using kubebuilder annotations under api/v1alpha1/containerset_types.go
   1. Ensure the CRD replicas can only range between 1-5 replicas
   2. Ensure you can use css as the CRDs shortName

   See the class material references for validators and CRD settings

10. Reinstall your CRD and check the validation and shortName are working as expected.

11. Delete your containerset and crd

    ```shell
    kubectl delete -f config/samples -f config/crd/bases
    ```

12. Deploy your controller as a container
    1. If you don't have an account yet, create a [DockerHub account](https://hub.docker.com/)
    2. Build your CRD Docker image and push it on DockerHub

    ```shell
    docker login -u YOUR_USER -p YOUR_PASSWD # Authenticate with DockerHub so you can push your image
    make docker-build docker-push IMG=YOUR_USER/cs:0.0.1 # Build and push your Docker image
    ```

13. Deploy your new controller image as a container on your local cluster

    ```shell
    make deploy IMG=YOUR_USER/cs:0.0.1
    ```

14. Rinse and repeat by deploying your Zorg CRD

    ```shell
    kubectl apply -f config/samples/zorg.yaml
    ```

<br/>

---
<img src="../../../assets/imhotep_logo.png" width="32" height="auto"/> © 2019 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)