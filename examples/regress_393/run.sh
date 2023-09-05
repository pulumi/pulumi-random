#!/bin/bash

set -eu -o pipefail

export PULUMI_CONFIG_PASSPHRASE=""

mkdir -p state
pulumi login --cloud-url "file://$PWD/state"
pulumi stack select organization/test/test --create
PULUMI_DEBUG_GRPC=grpc.json pulumi up       --yes #--debug -v=9 --logtostderr --logflow
                            pulumi destroy  --yes
                            pulumi stack rm --yes
pulumi logout
rm "$PWD/state"
