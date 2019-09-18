#!/bin/bash

# Build for linux and MacOS
bazel build --platforms=@io_bazel_rules_go//go/toolchain:darwin_amd64 -- //cli
bazel build --platforms=@io_bazel_rules_go//go/toolchain:linux_amd64 -- //cli

platforms=("linux_amd64" "darwin_amd64")

project_name="contentful-fetcher"

cd releases

for p in ${platforms[@]}; do

  binary_name="${p}_pure_stripped"

  release_binary_name="${project_name}_${p}";

  cp ../bazel-bin/cli/$binary_name/cli $release_binary_name

  zip -r ${release_binary_name}.zip . -i $release_binary_name

  rm -f $release_binary_name

  # cd ..
done

# Create the release and upload it to github.
# The GITHUB_TOKEN environment variable must be set for this to succeed.
ghr --prerelease --soft v0.0.1-alpha .
