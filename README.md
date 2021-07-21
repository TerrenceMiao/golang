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

Publishing a module
-------------------

Removes any dependencies the module might have accumulated that are no longer necessary.

```
𝜆 go mod tidy
warning: ignoring symlink /Users/terrence/Projects/golang/src/bazel-bin
warning: ignoring symlink /Users/terrence/Projects/golang/src/bazel-out
warning: ignoring symlink /Users/terrence/Projects/golang/src/bazel-src
warning: ignoring symlink /Users/terrence/Projects/golang/src/bazel-testlogs
go: finding module for package github.com/terrencemiao/golang/protos/hello
go: finding module for package github.com/jessevdk/go-flags
go: finding module for package github.com/terrencemiao/golang/protos/common
go: finding module for package google.golang.org/grpc
go: finding module for package github.com/stretchr/testify/require
go: found github.com/jessevdk/go-flags in github.com/jessevdk/go-flags v1.5.0
go: found google.golang.org/grpc in google.golang.org/grpc v1.39.0
go: found github.com/stretchr/testify/require in github.com/stretchr/testify v1.7.0
go: downloading golang.org/x/net v0.0.0-20200822124328-c89045814202
go: downloading github.com/golang/protobuf v1.4.3
go: downloading gopkg.in/check.v1 v0.0.0-20161208181325-20d25e280405
go: downloading google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013
go: downloading github.com/google/go-cmp v0.5.0
go: downloading google.golang.org/protobuf v1.25.0
```

Tag the project with a version number.

```
𝜆 git tag -a v1.1.0 -m "Publish module version v1.1.0"

𝜆 git push origin v1.1.0
Enumerating objects: 1, done.
Counting objects: 100% (1/1), done.
Writing objects: 100% (1/1), 829 bytes | 829.00 KiB/s, done.
Total 1 (delta 0), reused 0 (delta 0), pack-reused 0
To https://github.com/terrencemiao/golang.git
 * [new tag]         v1.1.0 -> v1.1.0
```

Publish Golang module.

Golang packages are given lower case, single-word names; there should be no need for underscores or mixedCaps.

```
𝜆 env GOPROXY=proxy.golang.org go list -m github.com/terrencemiao/golang@v1.1.0
github.com/terrencemiao/golang v1.1.0
```

Can find the published Golang module at _https://pkg.go.dev/github.com/terrencemiao/golang_

Reference
---------

- How to Golang Monorepo, _https://medium.com/goc0de/how-to-golang-monorepo-4f62320a01fd_
- Go rules for Bazel, _https://github.com/bazelbuild/rules_go_
