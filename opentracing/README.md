<img src="../assets/k8sland.png" align="right" width="auto" height="128"/>

<br/>
<br/>
<br/>

# OpenTracing G.O.T Lab

For this lab you are going to decorate a web server by leverage opentracing to
decorate a request. There are 2 services involved namely castle and knight. The
Knight will attempt to melt the castle, but if you're a G.O.T fan, you already
know that only the nightking can melt a castle using his undead dragon.

1. Instrument the castle code base to add tracing upon receiving a melt request
1. If the given night is not the nightking, then the trace should report a failure
1. If the given night is indeed the nightking then the trace should indicate the
  castle was melted.
1. Your castle trace needs to add the following tags to the trace:
  1. http.method
  2. http.url
  3. knight
1. In the event of a nightking the castle span should log a message indicating
   the castle was melted.
1. All other knights should produce a span error.
1. Span errors are indicated by:
   1. Setting a span tag error=true
   2. Adding a structured log on the span using:
    1. event=error
    2. message=only the nightkind can melt
1. You will need to modify the provided makefile to use your own docker registry
1. Build and push your own docker images with your new code
1. Deploy jaeger, castle and knight service to your local cluster
   1. You will need to modify the K8s manifest to use your image names (castle, knight)
1. Validate that your traces are correctly showing the microservices flow using
   different knights

<br/>

## Commands


1. Setup Jaeger

  ```shell
  kubectl apply -f k8s/jaeger.yml
  # Jaeger UI
  open http://$(minikube ip):30600/
  ```

2. Build your code and publish your own docker images

    ```shell
    make push
    ```

1. Deploy your services

    ```shell
    kubectl deploy -f k8s/got
    ```

1. Test your endpoints and traces

   ```shell
   http $(minikube ip):30501/v1/melt knight=fred
   ```

<br/>

---
<img src="../assets/imhotep_logo.png" width="32" height="auto"/> © 2018 Imhotep Software LLC.
All materials licensed under [Apache v2.0](http://www.apache.org/licenses/LICENSE-2.0)
