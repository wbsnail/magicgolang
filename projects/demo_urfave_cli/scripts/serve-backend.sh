#!/bin/sh

set -e

# change working directory to project root
wd=$(dirname "$(cd "$(dirname "$0")" && pwd)")
cd "${wd}" || echo "Cannot get into working directory: ${wd}"

# create default config file is not exist
mkdir -p ./local
cp -n ./config/config.tmpl.yaml ./local/config.yaml || true

# go!
go run ./cmd/server -c ./local/config.yaml $*
