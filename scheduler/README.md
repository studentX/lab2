<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> Custom Scheduler Lab

> The party pooper scheduler! Write a party scheduler that will only schedule
> a pod if it has a costume.

NOTE: Skip to step 1-2 if no GO Chops!

1. Define a custom party scheduler aka partysched, that checks if a given pod
   has a costume label set to either `ghoul` or `goblin` and matches node costumes...
2. Using the provided makefile change to your Docker REGISTRY and build your scheduler image.
3. Use the following docker image: k8sland/go-partysched:0.0.2
4. Setup a partysched deployment with a service account and RBAC rules
5. Deploy your custom scheduler and pod
6. Monitor cluster event to see if your pod got scheduled
7. Change your pod costume and check if it can get scheduled on a node
8. Delete your application and scheduler!

## Commands

# Check RBAC Rules for serviceaccount

    ```shell
    kubectl auth can-i get pods -n default \
      --as system:serviceaccount:default:partysched
    ```

# Monitor cluster events

    ```shell
     kubectl get ev --field-selector reason=FailSchedule
    ```

<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
