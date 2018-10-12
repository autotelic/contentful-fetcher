load("@bazel_gazelle//:def.bzl", "gazelle")
load("//cli/installer:def.bzl", "cli_installer")

# gazelle:build_file_name BUILD,BUILD.bazel

gazelle(
    name = "gazelle",
    prefix = "github.com/autotelic/contentful-fetcher",
)

gazelle(
    name = "gazelle_diff",
    mode = "diff",
    prefix = "github.com/autotelic/contentful-fetcher",
)

cli_installer(
    name = "install_cli",
    alias = "contentful-fetcher-cli",
)
