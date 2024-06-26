# CS548 | Assignment 2 - Kubernetes | Efthymios Papageorgiou - csdp1344

## Exercise 1
Provide the YAML that runs a Pod with Nginx 1.23.3-alpine as well as the kubectl commands needed to

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

## Exercise 2
The code that produces the course's website is available on GitHub (https://github.com/chazapis/hy548). Provide the YAML that creates a Job using Ubuntu 20.04, which when started will run a script (defined in a ConfigMap) that will download the repository (and submodules), hugo (the tool that builds the website), and build the website. Which command can you use to confirm that the Job completed successfully?

> kubectl apply -f configmaps/cs548-site-script.yaml  

> kubectl apply -f jobs/cs548-site-builder.yaml 

To confirm that the job completed successfully, we can do  

> kubectl get job cs548-site-builder  

and check the COMPLETIONS field. If the job is succussfully completed, then the field my show 1/1.  

#### A nice bug.

*When i submited the job, i was was waiting a couple of seconds for its completion. Sometime, i checked the pods that are created by this job. It were 5 pods with the name cs540-site-builder-*. When i saw them, i was like, Ohh the job failed, restart policy mechanism is being triggered 5 times (5 fails). So i started to check for the error. First i got all the pods related to the job by executing this command:*

> kubectl get pods --selector=job-name=cs548-site-builder  

All the related pods were listed. After that, i picked one of those and i did  

> kubectl logs cs548-site-builder-dsvq5  

The message was clear enough: **/bin/sh: 0: Can't open /scripts/build-script.sh**  

I missmached the script name between the job and configmap :)


## Exercise 3 
Following the previous two exercises, provide a single YAML that will run the Pod with Nginx, the above Job with the script, and a CronJob that will refresh the content every night at 2:15 (only if changes have been made to git). The Nginx Pod should serve the web pages produced by the Jobs instead of the default page. Briefly describe how data is communicated between containers  

Create the persistent volume for the cs548 site  
> kubectl apply -f persistent_volume_claims/cs548-site-vpc.yaml  

Create the persistent volume for the cs548 site repository  
> kubectl apply -f persistent_volume_claims/cs548-repo-vpc.yaml

Apply the, updated(compared to Exercise 2), ConfigMap  
> kubectl apply -f configmaps/cs548-site-script.yaml  

Deploy nginx  
> kubectl apply -f deployments/nginx.yaml

Apply the Job  

> kubectl apply -f jobs/cs548-site-builder.yaml  

Apply the CronJob  

> kubectl apply -f cronjobs/git-refresh.yaml  

Story about communication..  

For the data communication two persistent volumes were used. The first one, called cs548-repo-pvc, stores the course's repo. At first the cs548 site builder job downloads the repo and builds the site. At the same time the cronjob acceses this volume to detect any missmatches between the latest local commit and the latest remote commit.  

The other volume persistem claim, which is called cs548-site-vpc stores the public/* files of the html folder. At first the cs548 site builder job appends those file to this repo. If the CronJob detects any changes to the other volume, described above, then it deletes all the old files that are stored in the cs548-site-vpc and stores the new ones.  

This is all about the shared volumes and their usage.

## Exercise 4  

The exercise is answered by the yaml stored in deployments/nginx-deployment.yaml. The idea here is that the init container, named init container, spins into the folder cs548-site and when it will detect the index.html file it stops the spinning. At this point we know the job that builds the site has finished.