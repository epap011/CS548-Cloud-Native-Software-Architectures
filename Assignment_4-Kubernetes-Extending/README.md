# CS548 | Assignment 3 - Scalling apps on Kubernetes | Efthymios Papageorgiou - csdp1344

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

> 