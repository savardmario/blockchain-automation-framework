# General setup


## developer setup



## Clone forked repo

1. If you have not already done so, fork [blockchain-automation-framework](https://github.com/hyperledger-labs/blockchain-automation-framework) and clone the forked repo to your machine.

   ```bash
   cd ~/project
   git clone git@github.com:<githubuser>/blockchain-automation-framework.git
   
   ```

1. Add a “local” branch to your machine
   ```bash
   cd ~/project/blockchain-automation-framework
   git fetch --all
   git checkout tags/v0.8.1.0 -b local 
   git push --set-upstream origin local 
   ```

> Delete a “local” branch 
   ```bash
   git branch -d local
   git push origin --delete local
   ```

### create ssh key
```shell
# Generate you ssh key and name it RTI_BLOCKCHAIN_GITOPS_SVC_ACC
# No password
ssh-keygen -t ed25519 -C "rti.blockchain@ca.thalesgroup.com"
eval "$(ssh-agent -s)"

```
> follow this guide to add your ssh-key to [gitlab-profile](https://sc01-trt.thales-systems.ca/gitlab/help/ssh/README#see-if-you-have-an-existing-ssh-key-pair)

### create gitlab access-token
> follow this guide to add rti-gitops access-token to [gitlab-access-token](https://sc01-trt.thales-systems.ca/gitlab/-/profile/personal_access_tokens) 

## Minikube

Follow the documentation here [baf_minikube_setup](./docs/source/developer/baf_minikube_setup.md)

### Configure minikube to use 4GB memory and default kubernetes version
```shell
minikube start --vm-driver=virtualbox --kubernetes-version=v1.20.7
```

### Check status of minikube by running
```shell
The Kubernetes config file is generated at ~/.kube/config
minikube status

```

### To stop (do not delete) minikube execute the following
```shell
minikube stop
```


## Update kubeconfig file

1. Create a `build` folder inside your BAF repository:
   ```bash
   cd ~/project/blockchain-automation-framework
   mkdir build
   ```
1. Copy ca.crt, client.key, client.crt from `~/.minikube` to build:

   ```bash
   cp ~/.minikube/ca.crt build/
   cp ~/.minikube/profiles/minikube/client.key build/
   cp ~/.minikube/profiles/minikube/client.crt build/
   ```

1. Copy `~/.kube/config` file to build:

   ```bash
   cp ~/.kube/config build/
   ```

1. Copy RTI_BLOCKCHAIN_GITOPS_SVC_ACC file from ~/.ssh to build. (This is the private key file which you used to authenticate to your GitHub in pre-requisites)
   ```bash
   cp ~/.ssh/RTI_BLOCKCHAIN_GITOPS_SVC_ACC build/
   ```

docker build . -t baf-build-run
docker build Dockerfile.reset -t baf-build-reset

docker run -it -v $(pwd):/home/blockchain-automation-framework/ baf-build-run
docker run -it -v $(pwd):/home/blockchain-automation-framework/ baf-build-reset



docker run --rm -it -v $(pwd):/home/blockchain-automation-framework/ baf-build-run:latest