<img src="../../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../../assets/lab.png" width="32" height="auto"/> ContainerSet CRD Lab

> Using KubeBuilder, create a ContainerSet CRD

A ContainerSet CRD defines an image and a number of replicas and ensures a
deployment is created in the current namespace with the image/replicas specified
by the CRD.

> NOTE!! We will use an already generated/implemented ContainerSet for this lab but you will
> have to tweak a few things prior to deploying the CRD on your cluster ;)

1. Before you begin, make sure to install GO and kubebuilder per the previous lesson's instructions!

2. Clone the [K8sland Lab Repo](https://github.com/k8sland/lab2)

    ```shell
    git clone https://github.com/k8sland/lab2.git
    ```

3. Cd to crds/containerset

4. Fill in the ContainerSet CRD manifest in config/samples/zorg.yaml
   1. Specify 1 replica
   2. Specify the container image as nginx:1.17

5. Verify your enviroment is cool by installing the CRD schema and running the controller locally

    ```shell
    make install
    # Verify!
    kubectl get crd # Ensure containersets.workload.k8sland.io exists and the OpenAPI schema is correct!
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
   # Verify the crd instance exist and is correct
   kubectl get containerset
   kubectl get zorg -oyaml
   # In your watch window you should see a new zorg deployment and one zorg pod
   ```

9. Using kubebuilder annotations under api/v1alpha1/containerset_types.go
   1. Ensure the CRD replicas can only range between 1-5 replicas
   2. Ensure you can use `css` as the CRDs shortName

   See the class material References section for validators and CRD settings

10. Reinstall your CRD and check the validation and shortName are working as expected.

    ```shell
    kubectl apply -f config/samples/zorg.yaml # Fails if replicas is 0 or 6
    kubectl get css # Displays the zorg ContainerSet instance
    ```

11. Delete your containerset and crd

    ```shell
    kubectl delete -f config/samples -f config/crd/bases
    ```

12. Deploy your controller as a container
    1. NOTE! If you don't want to create a DockerHub account just skip the docker-push make target below!
    2. If you don't have an account yet, create a [DockerHub account](https://hub.docker.com/)
    3. Build your CRD Docker image and push it to DockerHub

    ```shell
    docker login -u YOUR_DOCKERHUB_USER -p YOUR_DOCKERHUB_PASSWD # Authenticate with DockerHub so you can push your image
    make docker-build docker-push IMG=YOUR_USER/cs:0.0.1 # Build and push your Docker image
    ```

13. Deploy your new controller image as a container on your local cluster

    ```shell
    make deploy IMG=YOUR_DOCKERHUB_USER/cs:0.0.1
    ```

14. Rinse and repeat by deploying and verifying your Zorg CRD is working as expected.

    ```shell
    kubectl apply -f config/samples/zorg.yaml
    ```

15. Well Done!!

<br/>

---
<img src="../../assets/imhotep_logo.png" width="32" height="auto"/> © 2019 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)