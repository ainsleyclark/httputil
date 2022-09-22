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

// Is2xx determines if a response status code is flagged as OK.
func (r *Response) Is2xx() bool {
	return statusXXX(r.Status, http.StatusOK, 299)
}

func statusXXX(status, high, low int) bool {
	if status >= high && status <= low {
		return true
	}
	return false
}
