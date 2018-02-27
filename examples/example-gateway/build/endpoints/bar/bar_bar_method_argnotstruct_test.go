// Code generated by zanzibar
// @generated

// Copyright (c) 2018 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package barendpoint

import (
	"bytes"
	"net/http"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	testGateway "github.com/uber/zanzibar/test/lib/test_gateway"
)

func getDirNameArgNotStructSuccessfulRequest() string {
	_, file, _, _ := runtime.Caller(0)

	return filepath.Dir(file)
}

func TestArgNotStructSuccessfulRequestOKResponse(t *testing.T) {
	var counter int

	gateway, err := testGateway.CreateGateway(t, nil, &testGateway.Options{
		KnownHTTPBackends: []string{"bar"},
		TestBinary: filepath.Join(
			getDirNameArgNotStructSuccessfulRequest(),
			"../../..",
			"build", "services", "example-gateway",
			"main", "main.go",
		),
		ConfigFiles: []string{
			filepath.Join(
				getDirNameArgNotStructSuccessfulRequest(),
				"../../..",
				"config", "test.json",
			),
			filepath.Join(
				getDirNameArgNotStructSuccessfulRequest(),
				"../../..",
				"config", "example-gateway", "test.json",
			),
		},
	})
	if !assert.NoError(t, err, "got bootstrap err") {
		return
	}
	defer gateway.Close()

	fakeArgNotStruct := func(w http.ResponseWriter, r *http.Request) {

		w.WriteHeader(200)

		payload := []byte(`null`)

		// TODO(zw): generate client response.
		if _, err := w.Write(payload); err != nil {
			t.Fatal("can't write fake response")
		}
		counter++
	}

	gateway.HTTPBackends()["bar"].HandleFunc(
		"POST", "/arg-not-struct-path", fakeArgNotStruct,
	)

	headers := map[string]string{}

	endpointRequest := []byte(`{"request":"foo"}`)

	res, err := gateway.MakeRequest(
		"POST",
		"/bar/arg-not-struct-path",
		headers,
		bytes.NewReader(endpointRequest),
	)
	if !assert.NoError(t, err, "got http error") {
		return
	}

	assert.Equal(t, 200, res.StatusCode)

	assert.Equal(t, 1, counter)
}
