Golang
======

Golang Interactive Playground
-----------------------------

`Bazel` is designed to work at scale and supports incremental hermetic builds across a distributed infrastructure, which is necessary for large codebase. With **Bazel Go ruleset**, you are able to manage the Go toolchain and external libraries without depending on locally installed ones. 

`Gazelle` is used to generate Go and Protocol Buffers rules. With `Gazelle`, you are able to generate `Bazel` rules for most Go packages in our Go monorepo with minimal human input. `Gazelle` can also import the versions of Go modules into **Bazel rules** so we can conveniently and efficiently build external libraries. 

```
𝜆 git clone https://github.com/terrencemiao/golang src

𝜆 go mod init github.com/terrencemiao/golang
```

A file `go.mod` is created:

```
𝜆 cat go.mod
module github.com/terrencemiao/golang

go 1.16
```

Now run the command:

```
𝜆 bazel run //:gazelle 
```

which tells `bazel` to run the `gazelle` target specified in the `BUILD` file. This will autogenerate the `BUILD.bazel` files for all of the packages.

```
𝜆 tree -C
.
├── BUILD
├── LICENSE
├── README.md
├── WORKSPACE
├── bazel
│   ├── docker
│   │   ├── BUILD
│   │   ├── def.bzl
│   │   └── repos.bzl
│   └── go
│       ├── BUILD
│       ├── WORKSPACE
│       ├── def.bzl
│       └── repos.bzl
├── bazel-bin -> /private/var/tmp/_bazel_terrence/3...6/execroot/__main__/bazel-out/darwin-fastbuild/bin
├── bazel-out -> /private/var/tmp/_bazel_terrence/3...6/execroot/__main__/bazel-out
├── bazel-src -> /private/var/tmp/_bazel_terrence/3...6/execroot/__main__
├── bazel-testlogs -> /private/var/tmp/_bazel_terrence/3...6/execroot/__main__/bazel-out/darwin-fastbuild/testlogs
├── go.mod
├── go_third_party.bzl
├── link_go.sh
├── protos
│   ├── common
│   │   ├── BUILD.bazel
│   │   └── common.proto
│   └── hello
│       ├── BUILD.bazel
│       ├── hello.proto
│       └── hello_service.proto
└── services
    └── hello
        ├── BUILD.bazel
        ├── main.go
        └── server
            ├── BUILD.bazel
            ├── server.go
            └── server_test.go

13 directories, 24 files
```

In addition, `*.pb.go` artefact files also generated:

```
𝜆 find bazel-out/ -name "*.pb.go"
bazel-out//darwin-fastbuild/bin/protos/common/common_go_proto_/github.com/terrencemiao/golang/protos/common/common.pb.go
bazel-out//darwin-fastbuild/bin/protos/hello/hello_go_proto_/github.com/terrencemiao/golang/protos/hello/hello_service.pb.go
bazel-out//darwin-fastbuild/bin/protos/hello/hello_go_proto_/github.com/terrencemiao/golang/protos/hello/hello.pb.go
```

Now, inform `bazel` about the dependencies mentioned in `go.mod` file. Either:

```
𝜆 go get github.com/bazelbuild/bazel-gazelle/cmd/gazelle
𝜆 gazelle -go_prefix github.com/terrencemiao/golang
𝜆 gazelle update-repos --from_file=go.mod -to_macro=go_third_party.bzl%go_deps
```

or, with `bazel`:

```
𝜆 bazel run //:gazelle -- update-repos -from_file=go.mod -to_macro=go_third_party.bzl%go_deps
```

Compile hello service:

```
𝜆 bazel build //services/hello
INFO: Analyzed target //services/hello:hello (117 packages loaded, 1553 targets configured).
INFO: Found 1 target...
Target //services/hello:hello up-to-date:
  bazel-bin/services/hello/hello_/hello
INFO: Elapsed time: 2.331s, Critical Path: 0.06s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
```

Run hello service, with default proxy port **24689**:

```
𝜆 bazel run //services/hello
INFO: Analyzed target //services/hello:hello (0 packages loaded, 0 targets configured).
INFO: Found 1 target...
Target //services/hello:hello up-to-date:
  bazel-bin/services/hello/hello_/hello
INFO: Elapsed time: 0.430s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
INFO: Build completed successfully, 1 total action
2021/07/17 20:04:25 Setting proxy server port 24689
```

Run hello service, with proxy port **8082**:

```
𝜆 bazel run //services/hello -- --proxy-port 8082
INFO: Analyzed target //services/hello:hello (0 packages loaded, 0 targets configured).
INFO: Found 1 target...
Target //services/hello:hello up-to-date:
  bazel-bin/services/hello/hello_/hello
INFO: Elapsed time: 0.538s, Critical Path: 0.01s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
INFO: Build completed successfully, 1 total action
2021/07/17 20:06:03 Setting proxy server port 8082
```

Testing
-------

Error thrown when run:

```
𝜆 bazel test //...
...
ERROR: golang/src/services/hello/server/BUILD.bazel:14:8: no such package '@com_github_stretchr_testify//require': The repository '@com_github_stretchr_testify' could not be resolved and referenced by '//services/hello/server:server_test'
...
```

Solution fix this issue:

```
𝜆 bazel run //:gazelle -- update-repos github.com/stretchr/testify
```

Building Docker images with bazel
---------------------------------

Add `k8s` dependency at first.

```
𝜆 bazel run //:gazelle -- update-repos github.com/vdemeester/k8s-pkg-credentialprovider@v1.21.0-1
INFO: Analyzed target //:gazelle (0 packages loaded, 0 targets configured).
INFO: Found 1 target...
Target //:gazelle up-to-date:
  bazel-bin/gazelle-runner.bash
  bazel-bin/gazelle
INFO: Elapsed time: 0.528s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
INFO: Build completed successfully, 1 total action
```

Need to add `rules_docker` to build Docker image. `rules_docker` can push the container image to ECR or GCR.

```
𝜆 bazel build //services/hello:image
INFO: Analyzed target //services/hello:image (1 packages loaded, 4 targets configured).
INFO: Found 1 target...
Target //services/hello:image up-to-date:
  bazel-bin/services/hello/image-layer.tar
INFO: Elapsed time: 0.744s, Critical Path: 0.01s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
```

Also need add docker/container image in `services/hello/BUILD.bazel` file.

```
𝜆 bazel run //services/hello:image
INFO: Analyzed target //services/hello:image (0 packages loaded, 0 targets configured).
INFO: Found 1 target...
Target //services/hello:image up-to-date:
  bazel-bin/services/hello/image-layer.tar
INFO: Elapsed time: 0.344s, Critical Path: 0.00s
INFO: 1 process: 1 internal.
INFO: Build completed successfully, 1 total action
INFO: Build completed successfully, 1 total action
8651333b21e7: Loading layer [==================================================>]  3.031MB/3.031MB
ba16d454860a: Loading layer [==================================================>]  15.44MB/15.44MB
1c255866fed0: Loading layer [==================================================>]  11.69MB/11.69MB
84ff92691f90: Loading layer [==================================================>]  10.24kB/10.24kB
Loaded image ID: sha256:2889e51d148a67b1df2b1d7394ce5ac93bc022c6a2513c02259c3167dc6790b8
Tagging 2889e51d148a67b1df2b1d7394ce5ac93bc022c6a2513c02259c3167dc6790b8 as bazel/services/hello:image
```

A docker image and it will show on your local docker:

```
𝜆 docker images
REPOSITORY                  TAG       IMAGE ID       CREATED         SIZE
bazel/services/hello        image     2889e51d148a   2 minutes ago   28.5MB
```

Run Golang application in Docker:

```
𝜆 docker run -it --rm bazel/services/hello:image
2021/07/28 02:33:20 Setting proxy server port 24689

...

^C2021/07/28 02:33:23 Shutting down
```

A few issues raised during building Docker image:

- `standard_init_linux.go:228: exec user process caused: exec format error` thrown when run Docker image, missing `goarch="amd64"`
-  `fail("goos must be set if goarch is set")` thrown during Docker image building, missing `goos = "linux"`

Publishing a module
-------------------

Removes any dependencies the module might have accumulated that are no longer necessary.

```
𝜆 go mod tidy
go mod tidy
warning: ignoring symlink /Users/terrence/Projects/golang/src/bazel-bin
warning: ignoring symlink /Users/terrence/Projects/golang/src/bazel-out
warning: ignoring symlink /Users/terrence/Projects/golang/src/bazel-src
warning: ignoring symlink /Users/terrence/Projects/golang/src/bazel-testlogs
```

Tag the project with a version number.

```
𝜆 git tag -a v1.2.6 -m "Publish module version v1.2.6"
```

```
𝜆 git push origin v1.2.6
Enumerating objects: 1, done.
Counting objects: 100% (1/1), done.
Writing objects: 100% (1/1), 829 bytes | 829.00 KiB/s, done.
Total 1 (delta 0), reused 0 (delta 0), pack-reused 0
To https://github.com/terrencemiao/golang.git
 * [new tag]         v1.2.6 -> v1.2.6
```

Publish Golang module.

Golang packages are given lower case, single-word names; there should be no need for underscores or mixedCaps.

```
𝜆 env GOPROXY=proxy.golang.org go list -m github.com/terrencemiao/golang@v1.2.6
github.com/terrencemiao/golang v1.2.6
```

Can find the published Golang module at _https://pkg.go.dev/github.com/terrencemiao/golang_

Reference
---------

- How to Golang Monorepo, _https://medium.com/goc0de/how-to-golang-monorepo-4f62320a01fd_
- Go rules for Bazel, _https://github.com/bazelbuild/rules_go_
