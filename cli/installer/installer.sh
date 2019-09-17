#!/bin/bash

relative_cli_path=$(find . -type l -perm -u+x -name 'cli')
current_dir=$(pwd)
absolute_cli_path="${current_dir}/${relative_cli_path}"
dest_path="/usr/local/bin/${1}"

# Ask for sudo provileges.
[ "$UID" -eq 0 ] || exec sudo "$0" "$@"

# Remove symlink only if it already exists in the destination.
if [ -L $dest_path ]; then
  sudo unlink $dest_path
fi

# Symlink the compiled cli.
printf "Symlinking 'bazel-bin/%s' into '%s'... " "${relative_cli_path:2}" $dest_path
sudo ln -s "$absolute_cli_path" "$dest_path"
printf "ok\\n"
