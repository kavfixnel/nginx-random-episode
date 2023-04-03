# Random episode webserver

This is a very simple Nginx configuration that utilizes Lua with the help of
[this custom Nginx image](https://github.com/fabiocicerchia/nginx-lua) to randomly redirect
you to an Episode of a given show

## How it works

The scripts have a list of all espisodes of a given show, will randomly pick one for you and tell
Nginx to respond with a 302 redirect to that episode.

```bash
$ curl http://localhost:80/netflix/seinfeld -v
*   Trying 127.0.0.1:80...
* Connected to localhost (127.0.0.1) port 80 (#0)
> GET /netflix/seinfeld HTTP/1.1
> Host: localhost
> User-Agent: curl/7.83.1
> Accept: */*
> 
* Mark bundle as not supporting multiuse
< HTTP/1.1 302 Moved Temporarily
< Server: nginx/1.23.3
< Date: Sat, 01 Apr 2023 06:26:09 GMT
< Content-Type: text/html
< Content-Length: 145
< Connection: keep-alive
< Location: https://www.netflix.com/watch/80132933
< 
<html>
<head><title>302 Found</title></head>
<body>
<center><h1>302 Found</h1></center>
<hr><center>nginx/1.23.3</center>
</body>
</html>
* Connection #0 to host localhost left intact
```

## How to run locally

The docker compose configuration makes setup trivial

```bash
$ docker compose up
[+] Running 1/0
 ⠿ Container nginx-random-redirect-nginx-lua-1  Created                                                                                                        0.0s
Attaching to nginx-random-redirect-nginx-lua-1
nginx-random-redirect-nginx-lua-1  | 💗 Support the Project 💗
nginx-random-redirect-nginx-lua-1  | This project is only maintained by one person, Fabio Cicerchia <https://github.com/fabiocicerchia>.
nginx-random-redirect-nginx-lua-1  | It started as a simple docker image, now it updates automatically periodically and provides support for multiple disto 😎
nginx-random-redirect-nginx-lua-1  | Maintaining a project is a very time consuming activity, especially when done alone 💪
nginx-random-redirect-nginx-lua-1  | I really want to make this project better and become super cool 🚀
nginx-random-redirect-nginx-lua-1  | 
nginx-random-redirect-nginx-lua-1  | If you'd like to support this open-source project I'll appreciate any kind of contribution <https://github.com/sponsors/fabiocicerchia>.
nginx-random-redirect-nginx-lua-1  | 
nginx-random-redirect-nginx-lua-1  | ---
nginx-random-redirect-nginx-lua-1  | 
nginx-random-redirect-nginx-lua-1  |   % Total    % Received % Xferd  Average Speed   Time    Time     Time  Current
nginx-random-redirect-nginx-lua-1  |                                  Dload  Upload   Total   Spent    Left  Speed
  0     0    0     0    0     0      0      0 --:--:-- --:--:-- --:--:--     0
nginx-random-redirect-nginx-lua-1  | /docker-entrypoint.sh: /docker-entrypoint.d/ is not empty, will attempt to perform configuration
nginx-random-redirect-nginx-lua-1  | /docker-entrypoint.sh: Looking for shell scripts in /docker-entrypoint.d/
nginx-random-redirect-nginx-lua-1  | /docker-entrypoint.sh: Launching /docker-entrypoint.d/10-listen-on-ipv6-by-default.sh
nginx-random-redirect-nginx-lua-1  | 10-listen-on-ipv6-by-default.sh: info: IPv6 listen already enabled
nginx-random-redirect-nginx-lua-1  | /docker-entrypoint.sh: Launching /docker-entrypoint.d/20-envsubst-on-templates.sh
nginx-random-redirect-nginx-lua-1  | /docker-entrypoint.sh: Launching /docker-entrypoint.d/30-tune-worker-processes.sh
nginx-random-redirect-nginx-lua-1  | /docker-entrypoint.sh: Configuration complete; ready for start up
nginx-random-redirect-nginx-lua-1  | 2023/04/01 06:27:17 [notice] 1#1: using the "epoll" event method
nginx-random-redirect-nginx-lua-1  | 2023/04/01 06:27:17 [notice] 1#1: nginx/1.23.3
nginx-random-redirect-nginx-lua-1  | 2023/04/01 06:27:17 [notice] 1#1: built by gcc 12.2.1 20220924 (Alpine 12.2.1_git20220924-r4) 
nginx-random-redirect-nginx-lua-1  | 2023/04/01 06:27:17 [notice] 1#1: OS: Linux 5.15.49-linuxkit
nginx-random-redirect-nginx-lua-1  | 2023/04/01 06:27:17 [notice] 1#1: getrlimit(RLIMIT_NOFILE): 1048576:1048576
nginx-random-redirect-nginx-lua-1  | 2023/04/01 06:27:17 [notice] 1#1: start worker processes
nginx-random-redirect-nginx-lua-1  | 2023/04/01 06:27:17 [notice] 1#1: start worker process 25
nginx-random-redirect-nginx-lua-1  | 2023/04/01 06:27:17 [notice] 1#1: start worker process 26
nginx-random-redirect-nginx-lua-1  | 2023/04/01 06:27:17 [notice] 1#1: start worker process 27
nginx-random-redirect-nginx-lua-1  | 2023/04/01 06:27:17 [notice] 1#1: start worker process 28
nginx-random-redirect-nginx-lua-1  | 2023/04/01 06:27:17 [notice] 1#1: start worker process 29
nginx-random-redirect-nginx-lua-1  | 2023/04/01 06:27:17 [notice] 1#1: start worker process 30
```

## How to run locally with TLS

To test TLS with this setup, we need to create our own self-signed certificate for testing. We can do this by
following [this tutorial](https://www.digitalocean.com/community/tutorials/how-to-create-a-self-signed-ssl-certificate-for-nginx-in-ubuntu-18-04), or just running the following:

```bash
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout nginx-selfsigned.key -out nginx-selfsigned.crt
```

Aterwards, we can spin up with TLS version of this application with

```bash
$ docker compose up -f docker-compose.tls.yaml
```

## Deploying to a hosted instance

Included in this repo are the neccesary Terraform and Ansible configurations to deploy this application to
an environment in AWS. In order to spin up an environment, execute the following in the `terraform` subdirectory

```bash
terraform init
terraform apply
```

The apply command will spit out a variable named `public_dns`. Take the variable and put
it in a file named `hosts` in the ansible subdirectory:

```
[aws]
<public_dns variable>
```

followed by the following commands in said subdirectory

```bash
ansible-playbook -i hosts setup_machines.yaml
ansible-playbook -i hosts deploy.yaml
```

The app should now be running and reachable under the dns name from earlier