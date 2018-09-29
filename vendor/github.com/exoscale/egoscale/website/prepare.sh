#!/bin/sh
set -xe

cp ../README.md content/_index.md
cp ../cmd/cs/README.md content/cs/_index.md
cp ../cmd/exo/README.md content/cli/_index.md
mkdir -p static
cp ../gopher.png static

cd ../
dep ensure -vendor-only

cd cmd/cs
dep ensure -vendor-only
go build
./cs gen-doc

cd ../exo
dep ensure -vendor-only
go run doc/main.go

set +xe
echo "we are now ready to: hugo serve"
