<img src="../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../assets/lab.png" width="32" height="auto"/> Probes Lab

> Add Liveness/Readiness probes to an Iconoflix application

1. In the given Iconoflix deployment add an HTTP Liveness probe
   1. Use /check/alive
2. Monitor your endpoints
   1. kubectl get ep --watch
3. What do you see?
4. Now redeploy your Iconoflix pod but with a Liveness delay to 20s
5. What do you notice in your Iconoflix endpoint?
6. Next define a HTTP Readiness probe on the Iconoflix container
   1. /check/health
7. Watch the Iconoflix pod logs
8. What do you notice?
9. Delete your application

<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2020 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)