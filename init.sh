#!/bin/bash 

export VERSION=0.5.7.${BUILD_NUMBER}

docker rmi -f external-dns:${VERSION} || echo "No Image"
docker rmi -f cevalogistics/external-dns:${VERSION} || echo "No Image"

echo  ${docker_token} | docker --debug login --password-stdin --username ${docker_login} 
docker build . -t external-dns:${VERSION} --no-cache
docker tag  external-dns:${VERSION} cevalogistics/external-dns:${VERSION}
docker tag  external-dns:${VERSION} cevalogistics/external-dns:experimental
docker push cevalogistics/external-dns:${VERSION} 
docker push cevalogistics/external-dns:experimental
