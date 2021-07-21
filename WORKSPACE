load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

# Download the Go rules
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "8e968b5fcea1d2d64071872b12737bbb5514524ee5f0a4f54f5920266c261acb",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.28.0/rules_go-v0.28.0.zip",
    ],
)

# Download Gazelle
http_archive(
    name = "bazel_gazelle",
    sha256 = "62ca106be173579c0a167deb23358fdfe71ffa1e4cfdddf5582af26520f1c66f",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.23.0/bazel-gazelle-v0.23.0.tar.gz",
    ],
)

# Load macros and repository rules
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

# Declare indirect dependencies and register toolchains
go_rules_dependencies()

go_register_toolchains(version = "1.16.5")

gazelle_dependencies()


load("//:go_third_party.bzl", "go_deps")
# Declare Go direct dependencies

# ERROR: src/services/hello/BUILD.bazel:3:11: no such package '@com_github_jessevdk_go_flags//':
#   The repository '@com_github_jessevdk_go_flags' could not be resolved and referenced by ...
# AS packge is not declared in your WORKSPACE, run: bazel run //:gazelle -- update-repos github.com/jessevdk/go-flags
go_repository(
    name = "com_github_jessevdk_go_flags",
    importpath = "github.com/jessevdk/go-flags",
    sum = "h1:1jKYvbxEjfUl0fmqTCOfonvskHHXMjBySTLW4y9LFvc=",
    version = "v1.5.0",
)

# ERROR: src/services/hello/BUILD.bazel:3:11: no such package '@org_golang_google_grpc//':
#   The repository '@org_golang_google_grpc' could not be resolved and referenced by ...
# AS packge is not declared in your WORKSPACE, run: bazel run //:gazelle -- update-repos google.golang.org/grpc
go_repository(
    name = "org_golang_google_grpc",
    importpath = "google.golang.org/grpc",
    sum = "h1:Klz8I9kdtkIN6EpHHUOMLCYhTn/2WAe5a0s1hcBkdTI=",
    version = "v1.39.0",
)

# gazelle:repository_macro go_third_party.bzl%go_deps
go_deps()


http_archive(
    name = "com_google_protobuf",
    sha256 = "d0f5f605d0d656007ce6c8b5a82df3037e1d8fe8b121ed42e536f569dec16113",
    strip_prefix = "protobuf-3.14.0",
    urls = [
        "https://mirror.bazel.build/github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
        "https://github.com/protocolbuffers/protobuf/archive/v3.14.0.tar.gz",
    ],
)

load("@com_google_protobuf//:protobuf_deps.bzl", "protobuf_deps")

protobuf_deps()
