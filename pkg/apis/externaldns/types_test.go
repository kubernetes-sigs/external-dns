/*
Copyright 2017 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package externaldns

import "testing"

func TestParseFlags(t *testing.T) {
	// extra := "one-extra-argument"
	// args := []string{
	// 	"-bool",
	// 	"-bool2=true",
	// 	"--int", "22",
	// 	"--int64", "0x23",
	// 	"-uint", "24",
	// 	"--uint64", "25",
	// 	"-string", "hello",
	// 	"-float64", "2718e28",
	// 	"-duration", "2m",
	// 	extra,
	// }
}

type arg struct {
	raw string
	key string
	val string
}
