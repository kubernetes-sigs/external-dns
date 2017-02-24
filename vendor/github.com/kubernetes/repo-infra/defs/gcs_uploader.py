#!/usr/bin/env python

# Copyright 2016 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

import argparse
import atexit
import os
import os.path
import shutil
import subprocess
import sys
import tempfile


def main(argv):
    scratch = tempfile.mkdtemp(prefix="bazel-gcs.")
    atexit.register(lambda: shutil.rmtree(scratch))

    with open(argv.manifest) as manifest:
        for artifact in manifest:
            artifact = artifact.strip()
            try:
                os.makedirs(os.path.join(scratch, os.path.dirname(artifact)))
            except (OSError):
                # skip directory already exists errors
                pass
            os.symlink(os.path.join(argv.root, artifact), os.path.join(scratch, artifact))

    sys.exit(subprocess.call(["gsutil", "-m", "rsync", "-C", "-r", scratch, argv.gcs_path]))


if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Upload build targets to GCS.')

    parser.add_argument("--manifest", required=True, help="path to manifest of targets")
    parser.add_argument("--root", required=True, help="path to root of workspace")
    parser.add_argument("gcs_path", help="path in gcs to push targets")

    main(parser.parse_args())
