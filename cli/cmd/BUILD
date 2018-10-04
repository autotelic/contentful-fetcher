load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["root.go"],
    importpath = "github.com/autotelic/contentful-fetcher/cli/cmd",
    visibility = ["//visibility:public"],
    deps = [
        "//lib/contentful:go_default_library",
        "//lib/filesystem:go_default_library",
        "@com_github_spf13_cobra//:go_default_library",
        "@com_github_spf13_viper//:go_default_library",
    ],
)
