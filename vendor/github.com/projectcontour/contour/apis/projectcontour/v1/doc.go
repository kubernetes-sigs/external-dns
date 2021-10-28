<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
<<<<<<< HEAD
// Copyright Project Contour Authors
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +k8s:deepcopy-gen=package

// Package v1 holds the specification for the projectcontour.io Custom Resource Definitions (CRDs).
//
// In building this CRD, we've inadvertently overloaded the word "Condition", so we've tried to make
// this spec clear as to which types of condition are which.
//
// `MatchConditions` are used by `Routes` and `Includes` to specify rules to match requests against for either
// routing or inclusion.
//
// `DetailedConditions` are used in the `Status` of these objects to hold information about the relevant
// state of the object and the world around it.
//
// `SubConditions` are used underneath `DetailedConditions` to give more detail to errors or warnings.
//
||||||| parent of 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// Copyright © 2019 VMware
||||||| parent of 5ce8c7613 (update vendored files)
// Copyright © 2019 VMware
=======
// Copyright Project Contour Authors
>>>>>>> 5ce8c7613 (update vendored files)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +k8s:deepcopy-gen=package

<<<<<<< HEAD
// Package v1 is the v1 version of the API.
>>>>>>> 465fc751b (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 5ce8c7613 (update vendored files)
// Package v1 is the v1 version of the API.
=======
// Package v1 holds the specification for the projectcontour.io Custom Resource Definitions (CRDs).
//
// In building this CRD, we've inadvertently overloaded the word "Condition", so we've tried to make
// this spec clear as to which types of condition are which.
//
// `MatchConditions` are used by `Routes` and `Includes` to specify rules to match requests against for either
// routing or inclusion.
//
// `DetailedConditions` are used in the `Status` of these objects to hold information about the relevant
// state of the object and the world around it.
//
// `SubConditions` are used underneath `DetailedConditions` to give more detail to errors or warnings.
//
>>>>>>> 5ce8c7613 (update vendored files)
||||||| parent of 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
=======
// Copyright © 2019 VMware
||||||| parent of 6b7ce455e (update vendored files)
// Copyright © 2019 VMware
=======
// Copyright Project Contour Authors
>>>>>>> 6b7ce455e (update vendored files)
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// +k8s:deepcopy-gen=package

<<<<<<< HEAD
// Package v1 is the v1 version of the API.
>>>>>>> 2cb94ab58 (UPSTREAM: <carry>: openshift: OpenShift dockerfiles added)
||||||| parent of 6b7ce455e (update vendored files)
// Package v1 is the v1 version of the API.
=======
// Package v1 holds the specification for the projectcontour.io Custom Resource Definitions (CRDs).
//
// In building this CRD, we've inadvertently overloaded the word "Condition", so we've tried to make
// this spec clear as to which types of condition are which.
//
// `MatchConditions` are used by `Routes` and `Includes` to specify rules to match requests against for either
// routing or inclusion.
//
// `DetailedConditions` are used in the `Status` of these objects to hold information about the relevant
// state of the object and the world around it.
//
// `SubConditions` are used underneath `DetailedConditions` to give more detail to errors or warnings.
//
>>>>>>> 6b7ce455e (update vendored files)
// +groupName=projectcontour.io
package v1
