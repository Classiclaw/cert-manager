/*
Copyright 2020 The Jetstack cert-manager contributors.

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

package certificate

import (
	cmapi "github.com/jetstack/cert-manager/pkg/apis/certmanager/v1alpha2"
	cmmeta "github.com/jetstack/cert-manager/pkg/apis/meta/v1"
	"strings"
	"testing"
)

func TestFormatStringSlice(t *testing.T) {
	tests := map[string]struct {
		slice     []string
		expOutput string
	}{
		// Newlines are part of the expected output
		"Empty slice returns empty string": {
			slice:     []string{},
			expOutput: ``,
		},
		"Slice with one element returns string with one line": {
			slice: []string{"hello"},
			expOutput: `- hello
`,
		},
		"Slice with multiple elements returns string with multiple lines": {
			slice: []string{"hello", "World", "another line"},
			expOutput: `- hello
- World
- another line
`,
		},
	}

	for _, test := range tests {
		actualOutput := formatStringSlice(test.slice)
		if actualOutput != test.expOutput {
			t.Errorf("Unexpected output; expected: \n%s\nactual: \n%s", test.expOutput, actualOutput)
		}
	}
}

func TestCRInfoString(t *testing.T) {
	tests := map[string]struct {
		cr        *cmapi.CertificateRequest
		expOutput string
	}{
		// Newlines are part of the expected output
		"Nil pointer output correct": {
			cr: nil,
			expOutput: `No CertificateRequest found for this Certificate
`,
		},
		"CR with no condition output correct": {
			cr: &cmapi.CertificateRequest{Status: cmapi.CertificateRequestStatus{Conditions: []cmapi.CertificateRequestCondition{}}},
			expOutput: `CertificateRequest:
  Name: 
  Namespace: 
  Conditions:
    No Conditions set
`,
		},
		"CR with conditions output correct": {
			cr: &cmapi.CertificateRequest{
				Status: cmapi.CertificateRequestStatus{
					Conditions: []cmapi.CertificateRequestCondition{
						{Type: cmapi.CertificateRequestConditionReady, Status: cmmeta.ConditionTrue, Message: "example"},
					}}},
			expOutput: `CertificateRequest:
  Name: 
  Namespace: 
  Conditions:
    Ready: True, Reason: , Message: example
`,
		},
	}

	for _, test := range tests {
		actualOutput := crInfoString(test.cr)
		if strings.TrimSpace(actualOutput) != strings.TrimSpace(test.expOutput) {
			t.Errorf("Unexpected output; expected: \n%s\nactual: \n%s", test.expOutput, actualOutput)
		}
	}
}
