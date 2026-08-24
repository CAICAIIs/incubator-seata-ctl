/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package seata

import (
	"net/http"
	"strings"
	"testing"
)

func TestBuildPostRequestRejectsInvalidURL(t *testing.T) {
	oldAuth := auth
	auth = Auth{token: "test-token"}
	defer func() { auth = oldAuth }()

	tests := []struct {
		name string
		fn   func() (*http.Request, error)
	}{
		{
			name: "array data",
			fn: func() (*http.Request, error) {
				return BuildPostRequestWithArrayData("http://%zz", []string{"key"})
			},
		},
		{
			name: "map data",
			fn: func() (*http.Request, error) {
				return BuildPostRequestWithMapData("http://%zz", map[string]string{"key": "value"})
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request, err := test.fn()
			if err == nil || !strings.Contains(err.Error(), "invalid URL escape") {
				t.Fatalf("error = %v, want invalid URL escape", err)
			}
			if request != nil {
				t.Fatalf("request = %v, want nil", request)
			}
		})
	}
}

func TestTryCommandsRejectInvalidURLWithoutPanic(t *testing.T) {
	oldAuth := auth
	oldClient := defaultHTTPClient
	auth = Auth{ServerIP: "%zz", ServerPort: 0, token: "test-token"}
	defaultHTTPClient = &http.Client{Transport: roundTripFunc(func(*http.Request) (*http.Response, error) {
		t.Fatal("HTTP client should not be called when request creation fails")
		return nil, nil
	})}
	defer func() {
		auth = oldAuth
		defaultHTTPClient = oldClient
	}()

	tests := []struct {
		name string
		run  func()
	}{
		{name: "begin txn", run: func() { BeginTxn(3000) }},
		{name: "commit txn", run: func() { CommitTxn("xid-1") }},
		{name: "rollback txn", run: func() { RollbackTxn("xid-1") }},
		{name: "reload config", run: ReloadConfiguration},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer func() {
				if recovered := recover(); recovered != nil {
					t.Fatalf("%s panicked: %v", test.name, recovered)
				}
			}()
			test.run()
		})
	}
}
