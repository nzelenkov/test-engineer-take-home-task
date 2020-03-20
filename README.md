# Test Engineer - take home task
This is a solution for the following take home task 
https://github.com/mxinden/test-engineer-take-home-task/blob/master/README.md

# Features:
- Configuration with EVN vars suitable for CI/CD pipelines
- Wrap makefile for easy deployment
- Kubernetes deployment with helm chart
- Slim docker image - 7.2Mb

# System requirements
- Docker engine
- Kubernetes cluster
- Configured kubectrl and helm
- GNU Make
```
$ docker -v
Docker version 19.03.7-ce, build 7141c199a2
$ kubectl version --short
Client Version: v1.16.7
Server Version: v1.16.7
$ helm version --short
Client: v2.15.1+gcf1de4f
Server: v2.15.1+gcf1de4f
$ make -v
GNU Make 4.3
```

# Fast track
```
nano .env
source .env
make build
make deploy
make clean
```

# Step by step guide
### Configure build and application by modifing the .env file
```
nano .env
```
### Source the configuration .env file
```
source .env
```
### Build docker image and push to hub.docker.com 
```
$ make build
build docker image
Sending build context to Docker daemon  95.23kB
Step 1/16 : FROM golang:1.12.5-alpine3.9 AS builder
 ---> c7330979841b
Step 2/16 : LABEL Author="Nikola.Zelenkov <z@aapack.live>"
 ---> Running in 6833847bf60c
Removing intermediate container 6833847bf60c
 ---> 817219a60000
Step 3/16 : LABEL stage=builder
 ---> Running in 69cf119be19a
Removing intermediate container 69cf119be19a
 ---> 17f8afffd691
Step 4/16 : ENV APPNAME=go-sbst-http
 ---> Running in 2fddac831663
Removing intermediate container 2fddac831663
 ---> acd9023eed9e
Step 5/16 : ENV DOMAIN=3zed
 ---> Running in 94b42186e5aa
Removing intermediate container 94b42186e5aa
 ---> 0f28fe2f9e2b
Step 6/16 : ENV GOBIN=/go/bin
 ---> Running in e6a2bc1c801b
Removing intermediate container e6a2bc1c801b
 ---> 62ee52424228
Step 7/16 : ENV GOPATH=/go
 ---> Running in fd5b1123b2a5
Removing intermediate container fd5b1123b2a5
 ---> f4cc49cabd76
Step 8/16 : ENV CGO_ENABLED=0
 ---> Running in 6468946fe81a
Removing intermediate container 6468946fe81a
 ---> a51da5c42eda
Step 9/16 : ENV GOOS=linux
 ---> Running in 8821bd09185c
Removing intermediate container 8821bd09185c
 ---> 20f7bc15c946
Step 10/16 : COPY . go-sbst-http
 ---> 9a39e37af7dd
Step 11/16 : RUN apk update && apk add git && cd go-sbst-http go get -d -v && go build -v && go install -v && go test -v ./...
 ---> Running in f957188798de
fetch http://dl-cdn.alpinelinux.org/alpine/v3.9/main/x86_64/APKINDEX.tar.gz
fetch http://dl-cdn.alpinelinux.org/alpine/v3.9/community/x86_64/APKINDEX.tar.gz
v3.9.5-18-g276b549fe5 [http://dl-cdn.alpinelinux.org/alpine/v3.9/main]
v3.9.5-14-g437bc75312 [http://dl-cdn.alpinelinux.org/alpine/v3.9/community]
OK: 9778 distinct packages available
(1/6) Installing nghttp2-libs (1.35.1-r1)
(2/6) Installing libssh2 (1.9.0-r1)
(3/6) Installing libcurl (7.64.0-r3)
(4/6) Installing expat (2.2.8-r0)
(5/6) Installing pcre2 (10.32-r1)
(6/6) Installing git (2.20.2-r0)
Executing busybox-1.29.3-r10.trigger
OK: 20 MiB in 21 packages
net
internal/x/net/http/httpproxy
net/textproto
crypto/x509
internal/x/net/http/httpguts
crypto/tls
net/http/httptrace
net/http
_/go/go-sbst-http
?       _/go/go-sbst-http       [no test files]
Removing intermediate container f957188798de
 ---> f2be0bf78c4c
Step 12/16 : FROM scratch
 ---> 
Step 13/16 : ENV APPNAME=go-sbst-http
 ---> Using cache
 ---> 2c8846a39b16
Step 14/16 : WORKDIR /usr/local/bin
 ---> Using cache
 ---> c2bf1b5a9bca
Step 15/16 : COPY --from=builder /go/bin/go-sbst-http .
 ---> Using cache
 ---> 4797a8761523
Step 16/16 : CMD ["go-sbst-http"]
 ---> Using cache
 ---> e01ba26fb526
Successfully built e01ba26fb526
Successfully tagged 3zed/go-sbst-http:latest
The push refers to repository [docker.io/3zed/go-sbst-http]
866c81f274e5: Layer already exists 
383cb0f4fa2e: Layer already exists  
latest: digest: sha256:1528238b511851080bf86a40ed617ce23d962dc1b8d70376e4f186456bc5f547 size: 735
```

### Deployment of Substrate to Kubernetes, helm deploy and test
```
$ make deploy
deploy to k8s
NAME:   deploy
LAST DEPLOYED: Fri Mar 20 20:07:15 2020
NAMESPACE: default
STATUS: DEPLOYED

RESOURCES:
==> v1/Deployment
NAME    READY  UP-TO-DATE  AVAILABLE  AGE
deploy  0/1    0           0          0s

==> v1/Pod(related)
NAME                     READY  STATUS   RESTARTS  AGE
deploy-659c9557fb-xd5mf  0/1    Pending  0         0s

==> v1/Service
NAME    TYPE       CLUSTER-IP      EXTERNAL-IP  PORT(S)   AGE
deploy  ClusterIP  10.152.183.172  <none>       9933/TCP  0s


NOTES:
************************************
Substrate deployment to k8s launched
************************************
This deployment runs until the Subsrate block height reaches 10.
When tests is PASSED the deployment is deleted and resources released.
RUNNING: deploy-test-block-height
PASSED: deploy-test-block-height
release "deploy" deleted
pod "deploy-test-block-height" deleted
```
### Clean and release all resouces
```
$ make clean
clean up resources
Deleted Images:
deleted: sha256:f2be0bf78c4c20ca73505599e3a2e2515062a79a9f2c7a13fc656a46cc74c2a3
deleted: sha256:bc85d5c9eb8582284df0c82b5c5b8adf4523427874e81c211b6908ec776ca432
deleted: sha256:9a39e37af7dddeafe129aceb272701a19685c2dd1bfb067cb43e303c07a5a42a
deleted: sha256:851693f170b415a7db6491838b1b9343b4b2340072a0ecdefbd6432f7fa17f5e
deleted: sha256:20f7bc15c946e73a4a6f46e5adf8ab5b3980b21e1bca7876095b031174b94b8d
deleted: sha256:a51da5c42eda7915cd543c0d3e649364a3964aeceae0bb43ec5766040c20ee24
deleted: sha256:f4cc49cabd76a42cb5e5125fe8540b4a7141f92a3e70809df3e6822d8321bcab
deleted: sha256:62ee524242288de4fe0cd16435a05016c3b681618fd6e3259da05c3f094c9ac3
deleted: sha256:0f28fe2f9e2b65ba41163679b0704823635a907efbf4b8b7ab20144dcab53168
deleted: sha256:acd9023eed9eb7742ef3d874119ea3a3f1ae349c9fa4b8d605a94448589f93a8
deleted: sha256:17f8afffd6919f78b2a27760795a4b07cc35fbaacdd626fde0960be1fc396a1e
deleted: sha256:817219a60000d1b6894376f7247de8c3637a27a4749c9adfd9a897ddc453a000

Total reclaimed space: 31.49MB
```

# TODO
### NICE TO HAVE
- MVC, handlers and routing (mutex)
- RPC - use ws://host:9944 with non-blocking async/await
- Use SCALE encode/decode libs
- Open/expose other (than 9933) ports - 9944, 30333?
- Liveliness, Readiness probes and Metrics
