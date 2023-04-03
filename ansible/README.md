# Deploying application with Ansible

These playbooks will deploy nginx-random-redirect to an Ubuntu host with docker-compose

They assume that there is a hosts file with a tag called `[aws]` and hosts underneith that liks so:

```
[aws]
thisis.and.example.host.com
```

## Playbooks

There are two palybooks in this folder that will do the following:

### setup_machines.yaml

Will install packages and the docker runtime as well as the compose plugin for docker

### deploy.yaml

Will go out and deploy a new version of the application (as is commited in Git at the
the moment of execution)
