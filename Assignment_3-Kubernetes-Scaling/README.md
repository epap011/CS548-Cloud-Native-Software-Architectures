# CS548 | Assignment 3 - Scalling apps on Kubernetes | Efthymios Papageorgiou - csdp1344

## Exercise 1

### a. Instead of "Hello from Python Flask!", the flask-hello container will return the value of the MESSAGE environment variable when someone uses the service (use Python's os.getenv). Provide the new Dockerfile and hello.py. Build and upload the new container to Docker Hub.

Just added **ENV MESSAGE="..." on the dockerfile

To build and upload the new image to Docker Hub

> docker login -u epap011

> sudo docker build -t "flask-env-var-image" .

> docker tag flask-env-var-image epap011/flask-env-var:latest 

> docker push epap011/flask-env-var:latest

### b. Provide two YAMLs to deploy the above container with all necessary resources (Deployment, Service, Ingress), so that "This is the first service!" is returned when someone visits the /first endpoint, and "This is the second service!" when someone visits the /second.

first.yaml: Deployment & Service of the irst service  
second.yaml: Deployment & Service of the second service   
ingress.yaml: Ingress


### c. Provide the commands needed to test the above two services with minikube (from running minikube, to curl or wget commands to use the services). Assume that the first deployment is in first.yaml and the second in second.yaml.

Apply Deployment, Service and Ingress resources  
> kubectl apply -f first.yaml -f second.yaml -f ingress.yaml  

Make sure that ingress addon is enabled
> minikube addons enable ingress  

Get the minikube ip  
> minikube ip  
Output: 192.168.49.2

Visit the /first endpoint  
> curl http://192.168.49.2/first  
Output: This is the first service  

Visit the /second endpoint  
> curl http://192.168.49.2/second  
Output: This is the second service  

## Exercise 2: 
Following on from the previous exercise, extend the YAML that implements the /first endpoint  

### a. To limit each Pod to a maximum of 20% CPU and 256MB RAM.  
check the lines 20,21 of the first.yaml  

### b. With a HorizontalPodAutoscaler manifest, which will increase the number of Pods in the Deployment when the average CPU usage exceeds 80%. Set a minimum of 1 Pod and a maximum of 8 for the Deployment. 

Make sure that ingress, dashboard, and metrics-server addons are enabled
> minikube addons enable ingress  

> minikube addons enable dashboard  

> minikube addons enable metrics-server  

Apply locust (to generate workload)
> kubectl apply -f locust.yaml  

Monitoring
> kubectl get hpa --watch
```
NAME              REFERENCE                           TARGETS   MINPODS   MAXPODS   REPLICAS   AGE
flask-first-hpa   Deployment/flask-first-deployment   0%/80%    1         8         1          71m
flask-first-hpa   Deployment/flask-first-deployment   39%/80%   1         8         1          71m
flask-first-hpa   Deployment/flask-first-deployment   100%/80%   1         8         1          72m
flask-first-hpa   Deployment/flask-first-deployment   100%/80%   1         8         2          72m
flask-first-hpa   Deployment/flask-first-deployment   59%/80%    1         8         2          73m
flask-first-hpa   Deployment/flask-first-deployment   43%/80%    1         8         2          74m
flask-first-hpa   Deployment/flask-first-deployment   44%/80%    1         8         2          75m
flask-first-hpa   Deployment/flask-first-deployment   44%/80%    1         8         2          76m
flask-first-hpa   Deployment/flask-first-deployment   45%/80%    1         8         2          77m
flask-first-hpa   Deployment/flask-first-deployment   44%/80%    1         8         2          78m
flask-first-hpa   Deployment/flask-first-deployment   43%/80%    1         8         2          79m
flask-first-hpa   Deployment/flask-first-deployment   44%/80%    1         8         2          80m
flask-first-hpa   Deployment/flask-first-deployment   44%/80%    1         8         2          83m
flask-first-hpa   Deployment/flask-first-deployment   43%/80%    1         8         2          84m
flask-first-hpa   Deployment/flask-first-deployment   43%/80%    1         8         2          85m
flask-first-hpa   Deployment/flask-first-deployment   42%/80%    1         8         2          86m
flask-first-hpa   Deployment/flask-first-deployment   43%/80%    1         8         2          87m
flask-first-hpa   Deployment/flask-first-deployment   43%/80%    1         8         2          88m
flask-first-hpa   Deployment/flask-first-deployment   43%/80%    1         8         2          89m
flask-first-hpa   Deployment/flask-first-deployment   43%/80%    1         8         2          91m
flask-first-hpa   Deployment/flask-first-deployment   29%/80%    1         8         2          93m
flask-first-hpa   Deployment/flask-first-deployment   0%/80%     1         8         2          94m
flask-first-hpa   Deployment/flask-first-deployment   0%/80%     1         8         2          97m
flask-first-hpa   Deployment/flask-first-deployment   0%/80%     1         8         1          98m
```
At 72m into the process, when the workload peaked at 100% and remained consistently at 80% of its usual level, Kubernetes added another pod. From then on, the CPU usage stayed around 45%, so there was no need for more pods since we were below the 80% threshold. At 93m the workload decreased and CPU usage dropped, Kubernetes removed the extra pod.

The scaling of the first service stops at 2 containers.

## Exercise 3  
If you have enabled the ingress addon in minikube, remove it. Issue the commands to install the Ingress controller implemented with Nginx using Helm. You will find the chart at https://artifacthub.io/packages/helm/ingress-nginx/ingress-nginx. Try again the services of the above exercises. What changes are needed in the YAML files to make them work (if any)?

> helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx

Output: "ingress-nginx" already exists with the same configuration, skipping  

> helm repo update  

Hang tight while we grab the latest from your chart repositories...  
...Successfully got an update from the "ingress-nginx" chart repository  
Update Complete. ⎈Happy Helming!⎈  

> helm install ingress-nginx ingress-nginx/ingress-nginx  

```
NAME: ingress-nginx
LAST DEPLOYED: Mon Apr 29 20:53:58 2024
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
The ingress-nginx controller has been installed.
It may take a few minutes for the load balancer IP to be available.
You can watch the status by running 'kubectl get service --namespace default ingress-nginx-controller --output wide --watch'

An example Ingress that makes use of the controller:
  apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    name: example
    namespace: foo
  spec:
    ingressClassName: nginx
    rules:
      - host: www.example.com
        http:
          paths:
            - pathType: Prefix
              backend:
                service:
                  name: exampleService
                  port:
                    number: 80
              path: /
    # This section is only required if TLS is to be enabled for the Ingress
    tls:
      - hosts:
        - www.example.com
        secretName: example-tls

If TLS is enabled for the Ingress, a Secret containing the certificate and key must also be provided:

  apiVersion: v1
  kind: Secret
  metadata:
    name: example-tls
    namespace: foo
  data:
    tls.crt: <base64 encoded cert>
    tls.key: <base64 encoded key>
  type: kubernetes.io/tls
```  

> helm list  

```
NAME            NAMESPACE       REVISION        UPDATED                                         STATUS          CHART                   APP VERSION
ingress-nginx   default         1               2024-04-29 20:53:58.533951515 +0300 EEST        deployed        ingress-nginx-4.10.1    1.10.1   
```  

changes to ingress.yaml: check line 8

## Exercise 4  
Create a Helm chart for the service that implements the /first endpoint of exercise 2. The chart should define variables for:

### a. The string to return.
### b. The endpoint to use for the service.
### c. The CPU and memory limits of each Pod (optional, default is no limits).
### d. The maximum number of Deployment Pods for the HorizontalPodAutoscaler (optional, default is 10).  

> helm install third-service ./third-service --set appName=third,msgString="This is a third service",ingress.path=/third,resources.limits.cpu=0.25,autoscaling.maxReplicas=20 

Output:  
```
NAME: third-service
LAST DEPLOYED: Mon Apr 29 22:02:27 2024
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:
1. Get the application URL by running these commands:
  export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=third-service,app.kubernetes.io/instance=third-service" -o jsonpath="{.items[0].metadata.name}")
  export CONTAINER_PORT=$(kubectl get pod --namespace default $POD_NAME -o jsonpath="{.spec.containers[0].ports[0].containerPort}")
  echo "Visit http://127.0.0.1:8080 to use your application"
  kubectl --namespace default port-forward $POD_NAME 8080:$CONTAINER_PORT

```  

> kubectl get all  

```
NAME                                            READY   STATUS    RESTARTS       AGE
pod/flask-first-deployment-54d46d5bd8-fd782     1/1     Running   0              42m
pod/flask-second-deployment-9fc5464cb-b7pkm     1/1     Running   0              42m
pod/ingress-nginx-controller-6dfcb8658d-cjgd8   1/1     Running   0              69m
pod/locust-5b98756679-l7s6r                     1/1     Running   1 (157m ago)   3h44m
pod/third-service-77956cc965-cf58t              1/1     Running   0              108s

NAME                                         TYPE           CLUSTER-IP       EXTERNAL-IP      PORT(S)                      AGE
service/flask-first-service                  ClusterIP      10.103.107.169   <none>           8080/TCP                     42m
service/flask-second-service                 ClusterIP      10.109.127.56    <none>           8080/TCP                     42m
service/ingress-nginx-controller             LoadBalancer   10.100.148.223   10.100.148.223   80:31535/TCP,443:30708/TCP   69m
service/ingress-nginx-controller-admission   ClusterIP      10.99.95.115     <none>           443/TCP                      69m
service/kubernetes                           ClusterIP      10.96.0.1        <none>           443/TCP                      3h50m
service/locust-service                       ClusterIP      10.99.59.70      <none>           8089/TCP                     42m
service/third-service                        ClusterIP      10.107.165.122   <none>           80/TCP                       108s

NAME                                       READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/flask-first-deployment     1/1     1            1           42m
deployment.apps/flask-second-deployment    1/1     1            1           42m
deployment.apps/ingress-nginx-controller   1/1     1            1           69m
deployment.apps/locust                     1/1     1            1           3h44m
deployment.apps/third-service              1/1     1            1           108s

NAME                                                  DESIRED   CURRENT   READY   AGE
replicaset.apps/flask-first-deployment-54d46d5bd8     1         1         1       42m
replicaset.apps/flask-second-deployment-9fc5464cb     1         1         1       42m
replicaset.apps/ingress-nginx-controller-6dfcb8658d   1         1         1       69m
replicaset.apps/locust-5b98756679                     1         1         1       3h44m
replicaset.apps/third-service-77956cc965              1         1         1       108s

NAME                                                  REFERENCE                           TARGETS   MINPODS   MAXPODS   REPLICAS   AGE
horizontalpodautoscaler.autoscaling/flask-first-hpa   Deployment/flask-first-deployment   0%/80%    1         8         1          3h44m
```  

> curl http://10.100.148.223/third  

Output: This is a third service!
