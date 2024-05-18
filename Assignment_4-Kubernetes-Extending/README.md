# CS548 | Assignment 4 - Kubernetes-Extending | Efthymios Papageorgiou - csdp1344

## Exercise 1
Provide the YAML that allows you to manage a custom resource of type "Fruit" with Kubernetes. An example instance is the following:

```
apiVersion: hy548.csd.uoc.gr/v1
kind: Fruit
metadata:
  name: apple
spec:
  origin: Krousonas
  count: 3
  grams: 980 
```

### a. Install the custom resource.
> kubectl apply -f fruit-crd.yaml

### b. Create the above instance.
> kubectl apply -f apple.yaml

### c. Return the new instance in YAML format.
> kubectl get fruit apple -o yaml

Output:
```
apiVersion: hy548.csd.uoc.gr/v1
kind: Fruit
metadata:
  annotations:
    kubectl.kubernetes.io/last-applied-configuration: |
      {"apiVersion":"hy548.csd.uoc.gr/v1","kind":"Fruit","metadata":{"annotations":{},"name":"apple","namespace":"default"},"spec":{"count":3,"grams":980,"origin":"Krousonas"}}
  creationTimestamp: "2024-05-16T09:01:25Z"
  generation: 1
  name: apple
  namespace: default
  resourceVersion: "3492"
  uid: 778e6648-51a7-42ed-82ab-6537d361df37
spec:
  count: 3
  grams: 980
  origin: Krousonas
```

### d. Return a list of all available instances.
> kubectl get fruits

Output:
```
NAME    AGE
apple   30m
```

## Exercise 2
Extend the example so that:  

### a.  The example code for CRDs is available on GitHub (https://github.com/chazapis/hy548). The executable controller.py runs inside a container. Provide the Dockerfile. Build and upload the new container to Docker Hub.  

Docker Image can be found at: https://hub.docker.com/repository/docker/epap011/greetings-controller/general  

### b. Provide greeting-controller.yaml, which will create a Deployment with the container you made. Make sure the necessary permissions are set so the controller can read the "Greeting" CRDs and create the corresponding Deployments with the hello-kubernetes container in all namespaces.  

check files:  
- ClusterRole & ClusterRoleBinding : greeting/greeting-controller.cr.yaml
- Deployent: greeting/greeting-controller-dm.yaml  


Provide the commands you used to verify that the deployment works correctly.

> kubectl get pods -l app=greeting-controller  

Output: 
```
NAME                                   READY   STATUS    RESTARTS   AGE
greeting-controller-69586d8767-ktft7   1/1     Running   0          33m
```

## Exercise 3 
The example code for the webhooks is available on GitHub (https://github.com/chazapis/hy548). Extend the example so that:  

### a. The executable controller.py runs inside a container. Provide the Dockerfile. Build and upload the new container to Docker Hub.

Docker image can be found: https://hub.docker.com/repository/docker/epap011/webhook-controller/general  

### b. The webhook.yaml should use the new container instead of the proxy with Nginx. Provide the new webhook.yaml.  

Provide the commands you used to verify the webhook works correctly.

> kubectl describe deployment controller -n custom-label-injector  

Output:
```
Name:                   controller
Namespace:              custom-label-injector
CreationTimestamp:      Sat, 18 May 2024 09:30:01 +0300
Labels:                 app=custom-label-injector
Annotations:            deployment.kubernetes.io/revision: 3
Selector:               app=custom-label-injector
Replicas:               1 desired | 1 updated | 1 total | 1 available | 0 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  25% max unavailable, 25% max surge
Pod Template:
  Labels:  app=custom-label-injector
  Containers:
   controller:
    Image:        epap011/webhook-controller:latest
    Port:         8000/TCP
    Host Port:    0/TCP
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>
Conditions:
  Type           Status  Reason
  ----           ------  ------
  Available      True    MinimumReplicasAvailable
  Progressing    True    NewReplicaSetAvailable
OldReplicaSets:  controller-7888887588 (0/0 replicas created), controller-5ccc656cc5 (0/0 replicas created)
NewReplicaSet:   controller-5d9d6664c9 (1/1 replicas created)
Events:
  Type    Reason             Age    From                   Message
  ----    ------             ----   ----                   -------
  Normal  ScalingReplicaSet  5m37s  deployment-controller  Scaled up replica set controller-7888887588 to 1
  Normal  ScalingReplicaSet  4m31s  deployment-controller  Scaled up replica set controller-5ccc656cc5 to 1
  Normal  ScalingReplicaSet  27s    deployment-controller  Scaled down replica set controller-7888887588 to 0 from 1
  Normal  ScalingReplicaSet  27s    deployment-controller  Scaled up replica set controller-5d9d6664c9 to 1 from 0
  Normal  ScalingReplicaSet  19s    deployment-controller  Scaled down replica set controller-5ccc656cc5 to 0 from 1
  ```  

  > kubectl get pods -n custom-label-injector  

  Output:  
  ```
  NAME                          READY   STATUS    RESTARTS   AGE
controller-5d9d6664c9-j8lsq   1/1     Running   0          11m  
```