/*
Copyright 2018 The Kubernetes Authors.

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

package entrypoint

import (
	"testing"

	"sigs.k8s.io/prow/pkg/pod-utils/wrapper"
)

func TestOptions_Validate(t *testing.T) {
	var testCases = []struct {
		name        string
		input       Options
		expectedErr bool
	}{
		{
			name: "all ok",
			input: Options{
				Options: &wrapper.Options{
					Args:       []string{"/usr/bin/true"},
					ProcessLog: "output.txt",
					MarkerFile: "marker.txt",
				},
			},
			expectedErr: false,
		},
		{
			name: "missing args",
			input: Options{
				Options: &wrapper.Options{
					ProcessLog: "output.txt",
					MarkerFile: "marker.txt",
				},
			},
			expectedErr: true,
		},
	}

	for _, testCase := range testCases {
		err := testCase.input.Validate()
		if testCase.expectedErr && err == nil {
			t.Errorf("%s: expected an error but got none", testCase.name)
		}
		if !testCase.expectedErr && err != nil {
			t.Errorf("%s: expected no error but got one: %v", testCase.name, err)
		}
	}
}
