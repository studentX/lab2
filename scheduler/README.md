<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> Custom Scheduler Lab

> The party pooper💩scheduler!
> Write a party scheduler that will only schedule a pod if it has a costume on nodes that matches this costume.

NOTE: [NO GO CHOPS!] Skip to step #5

1. Implement a custom party scheduler aka partysched, that checks if a given pod
   has a costume label set to either ghoul or goblin
2. Make sure your pod gets scheduled on a node with the same costume as to pod's costume label
3. For all other cases ie no label or incorrect attire, make sure your scheduler
   selected no nodes and spews out a log.
4. Using the provided Makefile change the REGISTRY to your own docker registy and publish your party scheduler as a Docker container
5. Setup a partysched deployment with a service account and RBAC rules
6. Deploy your party scheduler in your cluster
7. Rinse and repeat your checks in a deployed configuration
8. Delete your application and scheduler!

## Commands

### [GO CHOPS] Run your party scheduler manually

    ```shell
    go run main.go
    ```

### Build and deploy a Docker image

  > NOTE! Update the Makefile and setup for your DockerHub registry

    ```shell
    make img push
    ```

# Check RBAC Rules for serviceaccount

    ```shell
    kubectl auth can-i get nodes \
      --namespace default \
      --as system:serviceaccount:default:partysched
    ```

<br/>

---

<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
