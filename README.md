## memhog

An app that wastefully allocates memory.

It is used in tandem with [memhog-operator](https://github.com/metral/memhog-operator) to demo a simple vertical auto-scaler based on memory consumption.

## Requirements
* Kubernetes v1.7.0+
* glide v0.11.1

### Building & Running

```
// Build
$ glide install -s -u -v
$ make
```

#### Run Locally

```
$ $GOPATH/bin/memhog -v2
```

#### Run on Kubernetes

```
// Create AppMonitor for Pod
$ kubectl create -f k8s/appmonitor/johnny-cache.yaml

$ kubectl create -f k8s/pod/memhog.yaml
```

Or in one command:

```
$ kubectl create -f k8s/ -R
```
