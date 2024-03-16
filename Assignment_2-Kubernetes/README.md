# CS548 | Assignment 2 - Kubernetes | Efthymios Papageorgiou - csdp1344

## Exercise 1 - Provide the YAML that runs a Pod with Nginx 1.23.3-alpine as well as the kubectl commands needed to

### a. Install the manifest on Kubernetes and start the Pod.

> kubectl apply -f nginx.yaml  

### b. Forward port 80 locally, so that it answers calls through a browser (or curl or wget).

> kubectl port-forward assignment2-csdp1344 80:80  

### c. See the logs of the running container.

> kubectl logs assignment2-csdp1344  

### d. Open a shell session inside the running container and change the first sentence of the default page to "Welcome to MY nginx!". Close the session.

> kubectl exec -it assignment2-csdp1344 -- /bin/sh  

i changed the welcome message to "Welcome to MY nginx!" at /usr/share/nginx/html/index.html

> wget http://localhost:8080/  

Output:  
<!DOCTYPE html>
<html>
<head>
<title>Welcome to MY nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to MY nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>  

### e. From your computer terminal (outside the container), download the default page locally and upload another one in its place.  

Download the default page
> kubectl cp default/assignment2-csdp1344:usr/share/nginx/html/index.html 

The file changed to:

<!DOCTYPE html>
<html>
<head>
<title>Welcome to Necronomikon !</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to Necronomikon !</h1>
<p>No new horror can be more terrible than the daily torture of the commonplace.</p>

<p><em>-H. P. Lovecraft</em></p>
</body>
</html>

So now, this index.html file will take the place of the default nginx page  

> kubectl cp index.html default/assignment2-csdp1344:usr/share/nginx/html/index.html

So now, if we run (outisde the container):

>  wget http://localhost:8080/  

We get as output the Nekronomikon page

### f. Stop the Pod and remove the manifest from Kubernetes.  

> kubectl delete pod assignment2-csdp1344  