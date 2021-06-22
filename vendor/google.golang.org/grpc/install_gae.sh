#!/bin/bash

TMP=$(mktemp -d /tmp/sdk.XXX) \
&& curl -o $TMP.zip "https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.68.zip" \
&& unzip -q $TMP.zip -d $TMP \
<<<<<<< HEAD
&& export PATH="$PATH:$TMP/go_appengine"
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
&& export PATH="$PATH:$TMP/go_appengine"
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
