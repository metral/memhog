## memhog

An app that wastefully allocates memory.

It is used in tandem with [memhog-operator](https://github.com/memhog-operator) to demo a simple vertical auto-scaler based on memory consumption.


### Building & Running

```
$ glide up -v
$ make
$ $GOPATH/bin/memhog -v2
```
