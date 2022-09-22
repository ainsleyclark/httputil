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

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"testing"
)

func UtilTestResponseStatus(t *testing.T, status int, fn func(r Response) bool) {
	t.Helper()
	for i := status; i < status+99; i++ {
		r := Response{Status: i}
		got := fn(r)
		assert.Equal(t, true, got)
	}
	assert.False(t, fn(Response{Status: status - 1}))
	assert.False(t, fn(Response{Status: status + 100}))
}

func TestResponse_Is1xx(t *testing.T) {
	UtilTestResponseStatus(t, http.StatusContinue, func(r Response) bool {
		return r.Is1xx()
	})
}

func TestResponse_Is2xx(t *testing.T) {
	UtilTestResponseStatus(t, http.StatusOK, func(r Response) bool {
		return r.Is2xx()
	})
}

func TestResponse_Is3xx(t *testing.T) {
	UtilTestResponseStatus(t, http.StatusMultipleChoices, func(r Response) bool {
		return r.Is3xx()
	})
}

func TestResponse_Is4xx(t *testing.T) {
	UtilTestResponseStatus(t, http.StatusBadRequest, func(r Response) bool {
		return r.Is4xx()
	})
}

func TestResponse_Is5xx(t *testing.T) {
	UtilTestResponseStatus(t, http.StatusInternalServerError, func(r Response) bool {
		return r.Is5xx()
	})
}
