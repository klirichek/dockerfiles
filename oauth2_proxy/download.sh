#!/bin/bash

set -eu

user=$1
package=$2
sha=$3

curl -L https://github.com/${user}/${package}/archive/${sha}.tar.gz | tar -xzf -
# The archive contains a directory with the sha, move it to a path without it
mv $package-$sha* $package

