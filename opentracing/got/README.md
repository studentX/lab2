<img src="../../assets/k8sland.png" align="right" width="auto" height="128"/>

<br/>
<br/>
<br/>

# OpenTracing G.O.T Lab

<br/>

This repo contains sample GO code for a words dictionary HTTP service.


## Commands

1. Enable ingress on Minikube

  ```shell
  minikube addons enable ingress
  ```

2. Setup Jaeger Operator

  ```shell
  # Install CRD
  kubectl apply -f k8s/jaeger
  # Create CRD instance
  kubectl apply -f k8s/jaeger.yml
  # Jaeger UI
  open http://$(minikube ip)/
  ```

## Endpoints

The service exposes two endpoints on port 5000:

1. **/words** to fetch a collection of dictionary words
2. **/healthz** to assess the health of the service

<br/>

---
<img src="../../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
