#!/bin/bash 

export VERSION=0.91

export  http_proxy="http://LAB-SVC-Inception:Welcome1@nadevproxy.logistics.corp:3128"
export https_proxy="http://LAB-SVC-Inception:Welcome1@nadevproxy.logistics.corp:3128"

if [ ! -x /usr/local/bin/bzr ]; then 
  brew install bzr
fi

docker rmi -f external-dns:${VERSION} || echo "No Image"
docker rmi -f cevalogistics/external-dns:${VERSION} || echo "No Image"

echo  dbc55c0b-6992-4503-a02a-b934247f6e42 | docker --debug login --password-stdin --username rsliotta 
docker build . -t external-dns:${VERSION} --no-cache
docker tag  external-dns:${VERSION} cevalogistics/external-dns:${VERSION}
docker push cevalogistics/external-dns:${VERSION} 
