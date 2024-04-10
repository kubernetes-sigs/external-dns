#!/bin/bash

set -eo pipefail

cp CONTRIBUTING.md code-of-conduct.md ./docs/

cp LICENSE ./docs/LICENSE.md

cp README.md ./docs/index.md

sed -i -e 's#docs/##g' ./docs/index.md
