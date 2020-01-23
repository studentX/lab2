<img src="../../assets/k8sland.png" align="right" width="128" height="auto"/>

<br/>

# <img src="../../assets/lab.png" width="32" height="auto"/> Fred CRD Lab

> Using KubeBuilder, create a sample Fred CRD
> NOTE! We will use the generated code here without modifications. We will do
> the right thing in the next lab.

1. Define a Fred resource

    ```shell
    git clone https://github.com/k8sland/lab2/tree/master $HOME/k8sland_level2_labs
    cd $HOME/k8sland_level2_labs/crd_fred
    kubebuilder create api --group blee --version v1alpha1 --kind Fred
    ```

2. Install the CRD schema

    ```shell
    make install
    # Verify!
    kubectl get crd | grep fred
    ```

3. Run the sample controller

    ```shell
    make run
    ```

4. Watch your new containersets

    ```shell
    kubectl get freds
    ```

5. Install the sample CRD

   ```shell
   kubectl apply -f config/blee_v1alpha1_fred.yaml
   # In your watch window you should see a new fred-sample
   ```

6. Describe the resource to make sure the Foo field is correct

7. Update the fred sample manifest and change the Foo field to some other value.

8. Update the manifest and make sure the Foo field is changed.

9. Stop your controller and delete the CRD

   ```shell
   make uninstall
   ```

10. Setup a docker a docker image (OPTIONAL push to a docker registry if you have one...)

    ```shell
    make docker-build IMG=CHANGE_ME_IMAGE_NAME:CHANGE_ME_IMAGE_REV
    ```

11. Run your controller in cluster

    ```shell
    make deploy IMG=CHANGE_ME_IMAGE_NAME:CHANGE_ME_IMAGE_REV
    ```

12. Check your fred's controller logs to make sure all is cool.

13. Rinse repeat installing the crd and instance. Do it in cluster this time!!

---
<img src="../../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)