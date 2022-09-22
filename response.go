// Copyright 2022 Ainsley Clark. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package httputil

import "net/http"

type (
	// Response contains the client response data from
	// the request.
	Response struct {
		RequestID string
		Status    int
		Body      string
		Headers   map[string]string
		Location  string
	}
)

// Is1xx determines if the response status code is in between
// 100 and 199.
func (r *Response) Is1xx() bool {
	return r.statusXXX(http.StatusContinue)
}

// Is2xx determines if the response status code is in between
// 200 and 299.
func (r *Response) Is2xx() bool {
	return r.statusXXX(http.StatusOK)
}

// Is3xx determines if the response status code is in between
// 300 and 399.
func (r *Response) Is3xx() bool {
	return r.statusXXX(http.StatusMultipleChoices)
}

// Is4xx determines if the response status code is in between
// 400 and 499.
func (r *Response) Is4xx() bool {
	return r.statusXXX(http.StatusBadRequest)
}

// Is5xx determines if the response status code is above 500.
func (r *Response) Is5xx() bool {
	return r.statusXXX(http.StatusInternalServerError)
}

func (r *Response) statusXXX(high int) bool {
	if r.Status >= high && r.Status <= high+99 {
		return true
	}
	return false
}
