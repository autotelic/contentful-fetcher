load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "manifest.go",
        "query.go",
    ],
    importpath = "github.com/autotelic/contentful-fetcher/lib/contentful",
    visibility = ["//visibility:public"],
    deps = ["//lib/filesystem:go_default_library"],
)
