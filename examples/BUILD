# gazelle:ignore

load("@io_bazel_rules_go//go:def.bzl", "go_binary")

go_binary(
    name = "example",
    data = [":manifest.json"],
    embed = ["//cli:go_default_library"],
    visibility = ["//visibility:public"],
)
