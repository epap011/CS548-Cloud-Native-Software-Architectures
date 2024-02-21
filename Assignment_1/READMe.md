# CS548 | Assignment 1 | Efthymios Papageorgiou - csdp1344

## Exercise 1

### a. Download the images tagged 1.23.3 and 1.23.3-alpine locally.

> docker pull nginx:1.23.3  

> docker pull nginx:1.23.3-alpine

### b. Compare the sizes of the two images.

> docker images  

nginx:1.23.3 is 142MB  
nginx:1.23.3-alpine is 40.7MB

### c. Start one of the two images in the background, with the appropriate network settings to forward port 80 locally and use a browser (or curl or wget) to see that calls are answered. What is the answer?

> docker run -p 80:80 -d nginx:1.23.3-alpine  

> curl http://localhost/  

Answer 
<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to nginx!</h1>
<p>If you see this page, the nginx web server is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using nginx.</em></p>
</body>
</html>  

### d. Confirm that the container is running in Docker

> docker ps  

CONTAINER ID   IMAGE                 COMMAND                  CREATED          STATUS          PORTS                               NAMES
fa5cebf3fe9c   nginx:1.23.3-alpine   "/docker-entrypoint.…"   12 seconds ago   Up 11 seconds   0.0.0.0:80->80/tcp, :::80->80/tcp   quizzical_mcnulty 

### e. Get the logs of the running container.
> docker logs fa5cebf3fe9c  

epap011@fedora:~$ sudo docker logs fa5cebf3fe9c
/docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
/docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
/docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
10-listen-on-ipv6-by-default.sh: info: Getting the checksum of /etc/nginx/conf.d/default.conf
10-listen-on-ipv6-by-default.sh: info: Enabled listen on IPv6 in /etc/nginx/conf.d/default.conf
/docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
/docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
/docker-entrypoint.sh: Configuration complete; ready for start up
2024/02/17 14:55:27 [notice] 1#1: using the "epoll" event method
2024/02/17 14:55:27 [notice] 1#1: nginx/1.23.3
2024/02/17 14:55:27 [notice] 1#1: built by gcc 12.2.1 20220924 (Alpine 12.2.1_git20220924-r4) 
2024/02/17 14:55:27 [notice] 1#1: OS: Linux 6.7.4-200.fc39.x86_64
2024/02/17 14:55:27 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1073741816:1073741816
2024/02/17 14:55:27 [notice] 1#1: start worker processes
2024/02/17 14:55:27 [notice] 1#1: start worker process 30
2024/02/17 14:55:27 [notice] 1#1: start worker process 31
2024/02/17 14:55:27 [notice] 1#1: start worker process 32
2024/02/17 14:55:27 [notice] 1#1: start worker process 33
2024/02/17 14:55:27 [notice] 1#1: start worker process 34
2024/02/17 14:55:27 [notice] 1#1: start worker process 35
2024/02/17 14:55:27 [notice] 1#1: start worker process 36
2024/02/17 14:55:27 [notice] 1#1: start worker process 37
172.17.0.1 - - [17/Feb/2024:14:55:35 +0000] "GET / HTTP/1.1" 200 615 "-" "curl/8.2.1" "-"  

### f. Stop the running container.
> docker stop fa5cebf3fe9c  

Output: fa5cebf3fe9c

### g. Start the stopped container.
> docker start fa5cebf3fe9c  

Output: fa5cebf3fe9c  

### h. Stop the container and remove it from Docker.
> docker rm -f fa5cebf3fe9c  

Output: fa5cebf3fe9c

## Exercise 2  

## a. Open a shell session inside the running container and change the first sentence of the default page to "Welcome to MY nginx!". Close the session.

> docker run -p 80:80 -d 2bc7edbc3cf2 

> docker exec -it d93d0072397f sh  

> vi /usr/share/nginx/html/index.html  

> curl http://localhost  

<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
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

## b. From your computer's terminal (outside the container) download the default page locally and upload another one in its place.

> docker cp d93d0072397f:/usr/share/nginx/html/index.html .  

Successfully copied 2.56kB to /home/epap011/Projects/CSD/CS548/Assignment1/.

> docker cp index.html d93d0072397f:/usr/share/nginx/html/index.html    

Successfully copied 2.56kB to d93d0072397f:/usr/share/nginx/html/index.html

> curl localhost  

<!DOCTYPE html>
<html>
<head>
<title>Welcome to nginx!</title>
<style>
html { color-scheme: light dark; }
body { width: 35em; margin: 0 auto;
font-family: Tahoma, Verdana, Arial, sans-serif; }
</style>
</head>
<body>
<h1>Welcome to the Candy Shop!</h1>
<p>If you see this page, the Candy Shop is successfully installed and
working. Further configuration is required.</p>

<p>For online documentation and support please refer to
<a href="http://nginx.org/">nginx.org</a>.<br/>
Commercial support is available at
<a href="http://nginx.com/">nginx.com</a>.</p>

<p><em>Thank you for using the Candy Shop.</em></p>
</body>
</html>

### c. Close the container, delete it and start another instance. Do you see the changes? Why;

Answer: I did not observe any changes. This is because when you close a container, any changes made to its filesystem (or configuration) are lost. 

## Exercise 3

### The code that produces the course's website is available on GitHub (https://github.com/chazapis/hy548). Write down the commands needed to download the repository (and submodules) and hugo (the tool that builds the website), build the website locally, and start an Nginx container to serve the CS-548 website instead of the default page.


> git clone https://github.com/chazapis/hy548  

> sudo dnf install hugo  

> git submodule update --init --recursive  

> make all  

> docker cp html/public/ sad_leavitt:/usr/share/nginx/html/  

> docker exec -it sad_leavitt sh  

updated location -> public  

> nginx -s reload  

Done  

# Exercise 4

### Following the previous exercise, create your own container image, based on Nginx, that will contain the CS-548 website instead of the default page. Downloading the CS-548 repository (and submodules), hugo and building the site should be done in the Dockerfile. Create a Docker Hub account and upload the image.

### a. The Dockerfile (inside the folder Assignment_1)

### b. commands needed to upload the image to Docker Hub
> docker login -u epap011

> docker build  -t epap011/hy548-site .  

> docker tag epap011/hy548-site epap011/hy548-site:latest  

> docker push epap011/hy548-site:latest  

### c. How much bigger is your own image than the image you were based on. Why;
Answer: 

### d. What have you done in the Dockerfile to keep the image as small as possible?
Answer: 

## Exercise 5

### Upload the Dockerfile from the previous exercise to your GitHub repository. Create a GitHub Action that will automatically build and push the image to your Docker Hub account (the workflow should be initiated by the user). Provide the YAML of the workflow you made.

The yaml file(main.yml) is located at CS548/.github/workflows/