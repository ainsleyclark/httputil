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
	"github.com/Danny-Dasilva/CycleTLS/cycletls"
	"github.com/ainsleyclark/errors"
	"github.com/ainsleyclark/httputil/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"testing"
)

func TestNewClient(t *testing.T) {
	got := NewClient()
	assert.NotNil(t, got)
}

func TestHttpClient_Do(t *testing.T) {
	url := "https://google.com"

	tt := map[string]struct {
		mock func(m *mocks.CycleTLS)
		want any
	}{
		"Error": {
			func(m *mocks.CycleTLS) {
				m.On("Do", mock.Anything, mock.Anything, mock.Anything).
					Return(cycletls.Response{}, errors.New("error"))
			},
			"Error performing client request",
		},
		"Redirect": {
			func(m *mocks.CycleTLS) {
				m.On("Do", mock.Anything, mock.Anything, mock.Anything).
					Return(cycletls.Response{Status: http.StatusMovedPermanently, Headers: map[string]string{"Location": "test"}}, nil).
					Once()
				m.On("Do", mock.Anything, mock.Anything, mock.Anything).
					Return(cycletls.Response{Body: "test"}, nil).
					Once()
			},
			"test",
		},
		"Location Error": {
			func(m *mocks.CycleTLS) {
				m.On("Do", mock.Anything, mock.Anything, mock.Anything).
					Return(cycletls.Response{Status: http.StatusMovedPermanently, Headers: map[string]string{"Location": badURL}}, nil)
			},
			"Error parsing relative URI",
		},
		"Same Location": {
			func(m *mocks.CycleTLS) {
				m.On("Do", mock.Anything, mock.Anything, mock.Anything).
					Return(cycletls.Response{Status: http.StatusMovedPermanently, Headers: map[string]string{"Location": url}, Body: "test"}, nil)
			},
			"test",
		},
		"OK": {
			func(m *mocks.CycleTLS) {
				m.On("Do", mock.Anything, mock.Anything, mock.Anything).
					Return(cycletls.Response{Body: "test"}, nil)
			},
			"test",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			cycle := &mocks.CycleTLS{}
			if test.mock != nil {
				test.mock(cycle)
			}
			client := httpClient{cycle: cycle}
			_, err := client.Do(url, http.MethodGet)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
		})
	}
}
