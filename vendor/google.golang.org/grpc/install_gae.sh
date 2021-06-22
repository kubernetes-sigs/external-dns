#!/bin/bash

TMP=$(mktemp -d /tmp/sdk.XXX) \
&& curl -o $TMP.zip "https://storage.googleapis.com/appengine-sdks/featured/go_appengine_sdk_linux_amd64-1.9.68.zip" \
&& unzip -q $TMP.zip -d $TMP \
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
&& export PATH="$PATH:$TMP/go_appengine"
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
&& export PATH="$PATH:$TMP/go_appengine"
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
&& export PATH="$PATH:$TMP/go_appengine"
=======
&& export PATH="$PATH:$TMP/go_appengine"
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
&& export PATH="$PATH:$TMP/go_appengine"
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
&& export PATH="$PATH:$TMP/go_appengine"
>>>>>>> 4a9b15dc1 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
&& export PATH="$PATH:$TMP/go_appengine"
>>>>>>> b60b08dfc (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
