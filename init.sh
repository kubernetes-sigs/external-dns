#!/bin/bash 

export VERSION=0.5.7.${BUILD_NUMBER}

#export  http_proxy="http://LAB-SVC-Inception:Welcome1@nadevproxy.logistics.corp:3128"
#export https_proxy="http://LAB-SVC-Inception:Welcome1@nadevproxy.logistics.corp:3128"

#if [ ! -x /usr/local/bin/bzr ]; then 
#  brew install bzr
#fi

docker rmi -f external-dns:${VERSION} || echo "No Image"
docker rmi -f cevalogistics/external-dns:${VERSION} || echo "No Image"

echo  ${docker_token} | docker --debug login --password-stdin --username ${docker_login} 
docker build . -t external-dns:${VERSION} --no-cache
docker tag  external-dns:${VERSION} cevalogistics/external-dns:${VERSION}
docker push cevalogistics/external-dns:${VERSION} 
docker push cevalogistics/external-dns:latest
