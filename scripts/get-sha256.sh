#!/bin/bash

IMAGE=$1

echo -n "image: "
crane digest "${IMAGE}"
echo "architecture"
crane manifest "${IMAGE}" | jq -r '.manifests.[] | .platform.architecture, .digest'
