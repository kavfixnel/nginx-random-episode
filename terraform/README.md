# Terraform for AWS

This subdirectory defines all neccessary configuration for spinning up your own environment in AWS.
The following will be deployed:

- A small t2.nano instance with an Ubuntu 20.04 installation
- A a key-pair to authenticate with the `ubuntu` user on the machine
- An ingress security rule that allows ssh and http traffic into the host

> In order to deploy things correctly, update the ssh key to reflect the correct public key
