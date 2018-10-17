# <img src="../assets/lab.png" width="32" height="auto"/> Init Container Lab

> Leverage Init Containers to provision dictionaries for a Dictionary service.

1. In this lab, you will leverage your local minikube cluster
1. Define a pod using the following Docker image: k8sland/dictionary-svc-go:0.0.2
1. The dictionary service is launched using the following command:
   ```shell
   /app/dictionary -a asset_dir -d dictionary
   ```
1. This service runs on port 4000 and exposes /words endpoint to list out the words
   contained in the dictionary loaded via *-d* option above.
1. Define an init container to prepopulate a volume to be used by the dictionary
   service
1. Your init container will need to clone this repo [Dictionaries](https://github.com/k8sland/dictionaries.git) in order to provision the volume


---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)