# Contentful Fetcher

This is a command line tool for interacting with the Contentful public API.

## Compiling

### Dependencies

- [Bazel](https://docs.bazel.build/versions/master/install.html).

### The CLI

Run the installer:

```sh
bazel run //:install_cli
```

## Usage

In order to run the CLI, you will need to set the following environment variables:

- CONTENTFUL_ACCESS_TOKEN - The access token provided by Contentful under Settings > API keys.
- CONTENTFUL_API_URL - `https://cdn.contentful.com` or `https://preview.contentful.com`.

Then call the CLI binary with one or more manifests using the `--manifest` or `-m` flag.

### Example

From the project root, run (the CONTENTFUL_ACCESS_TOKEN below is the example token provided by
Contentful):

```sh
CONTENTFUL_ACCESS_TOKEN=fdb4e7a3102747a02ea69ebac5e282b9e44d28fb340f778a4f5e788625a61abe \
  CONTENTFUL_API_URL="https://cdn.contentful.com/" \
  contentul-fetcher-cli -m examples/manifest.json
```

The examples manifest will create a `.content` directory with the contents of the queries.
