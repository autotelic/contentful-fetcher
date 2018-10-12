def cli_installer(
    name,
    alias,
):
    native.sh_binary(
        name = name,
        srcs = ["//cli/installer:installer.sh"],
        data = ["//cli"],
        args = [alias],
    )
