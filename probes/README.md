<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> Probes Lab

> Add Liveness/Readiness probes to an Iconoflix application

1. In the given Iconoflix deployment (k8s/iconoflix.yml) add a HTTP Liveness probe
   1. Use /check/health
2. Monitor your endpoints
3. What do you see?
4. Now redeploy your Iconoflix pod but with a Readiness delay to 20s
5. What do you notice in your Iconoflix endpoint?
6. Next define a HTTP Readiness probe on the Iconoflix container
   1. /check/alive
7. Watch the Iconoflix pod logs
8. What do you notice?
9. Delete your application

<br/>

---
## <img src="../assets/fox.png" width="32" height="auto"/> Commands

### Watching endpoints

    ```shell
    watch kubectl get ep
    ```

### Check probes are working

    ```shell
    # Watch events to make sure the probes are not failing
    kubectl get ev --field-selector reason=Unhealthy
    # Tail the logs to see the http requests
    kubectl logs -f `kubectl get po -l app=iconoflix -o go-template='{{(index .items 0).metadata.name}}'`
    ```

<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2019 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)