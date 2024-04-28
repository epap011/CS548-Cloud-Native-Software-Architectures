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

locust.yaml  
> kubectl apply -f locust.yaml  

> kubectl get hpa --watch
```
NAME                         REFERENCE                               TARGETS   MINPODS   MAXPODS   REPLICAS   AGE  
flask-env-var-1-deployment   Deployment/flask-env-var-1-deployment   0%/20%    1         8         1          8h  
flask-env-var-1-deployment   Deployment/flask-env-var-1-deployment   56%/20%   1         8         1          8h  
flask-env-var-1-deployment   Deployment/flask-env-var-1-deployment   56%/20%   1         8         3          9h  
flask-env-var-1-deployment   Deployment/flask-env-var-1-deployment   28%/20%   1         8         3          9h  
flask-env-var-1-deployment   Deployment/flask-env-var-1-deployment   21%/20%   1         8         3          9h  
flask-env-var-1-deployment   Deployment/flask-env-var-1-deployment   20%/20%   1         8         3          9h  
flask-env-var-1-deployment   Deployment/flask-env-var-1-deployment   21%/20%   1         8         3          9h  
flask-env-var-1-deployment   Deployment/flask-env-var-1-deployment   20%/20%   1         8         3          9h  
flask-env-var-1-deployment   Deployment/flask-env-var-1-deployment   21%/20%   1         8         3          9h  
flask-env-var-1-deployment   Deployment/flask-env-var-1-deployment   20%/20%   1         8         3          9h  
```