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
	"github.com/ainsleyclark/errors"
	"github.com/stretchr/testify/assert"
	"net/url"
	"testing"
)

var (
	// badURL is a dirty URL used for testing.
	badURL = "@@@£$%££"
)

func TestIs2xx(t *testing.T) {
	tt := map[string]struct {
		input int
		want  bool
	}{
		"200": {
			200,
			true,
		},
		"201": {
			201,
			true,
		},
		"300": {
			300,
			false,
		},
		"299": {
			200,
			true,
		},
		"199": {
			199,
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Is2xx(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestIs3xx(t *testing.T) {
	tt := map[string]struct {
		input int
		want  bool
	}{
		"300": {
			300,
			true,
		},
		"301": {
			301,
			true,
		},
		"400": {
			400,
			false,
		},
		"399": {
			399,
			true,
		},
		"299": {
			299,
			false,
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got := Is3xx(test.input)
			assert.Equal(t, test.want, got)
		})
	}
}

func TestGetAbsoluteURL(t *testing.T) {
	tt := map[string]struct {
		fullURL  string
		relative string
		want     any
	}{
		"Full URL Error": {
			badURL,
			"",
			"Error parsing full URI",
		},
		"Relative URL Error": {
			"",
			badURL,
			"Error parsing relative URI",
		},
		"Relative": {
			"https://www.comparethemarket.com/",
			"./img/test.jpg",
			"https://www.comparethemarket.com/img/test.jpg",
		},
		"Forward Slash": {
			"https://www.comparethemarket.com/home-insurance/content/green-eco-friendly-renovations",
			"/home-insurance/content/green-eco-friendly-renovations/",
			"https://www.comparethemarket.com/home-insurance/content/green-eco-friendly-renovations/",
		},
		"Absolute Full URL": {
			"https://www.comparethemarket.com/home-insurance/content/green-eco-friendly-renovations",
			"https://www.comparethemarket.com/home-insurance/content/green-eco-friendly-renovations/",
			"https://www.comparethemarket.com/home-insurance/content/green-eco-friendly-renovations/",
		},
		"Absolute Full URL 2": {
			"https://www.autotrader.co.uk/cars/electric/ev-drivers-with-disabilities",
			"/cars/electric/ev-drivers-with-disabilities/",
			"https://www.autotrader.co.uk/cars/electric/ev-drivers-with-disabilities/",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			got, err := GetAbsoluteURL(test.fullURL, test.relative)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
				return
			}
			assert.Equal(t, test.want, got)
		})
	}
}

func TestStartEndTimes(t *testing.T) {
	tt := map[string]struct {
		input url.Values
		want  any
	}{
		"Success": {
			url.Values{
				"start": []string{"2006-01-02"},
				"end":   []string{"2006-01-01"},
			},
			nil,
		},
		"Default Success": {
			url.Values{
				"end": []string{"2006-01-01"},
			},
			nil,
		},
		"Invalid Start": {
			url.Values{
				"start": []string{"hello"},
			},
			"Error, start date is invalid",
		},
		"Invalid End": {
			url.Values{
				"start": []string{"2006-01-02"},
				"end":   []string{"wrong"},
			},
			"Error, end date is invalid",
		},
	}

	for name, test := range tt {
		t.Run(name, func(t *testing.T) {
			_, _, err := GetStartEndTimes(test.input)
			if err != nil {
				assert.Contains(t, errors.Message(err), test.want)
			}
		})
	}

	t.Run("Parse Error", func(t *testing.T) {
		orig := defaultStartTime
		defer func() {
			defaultStartTime = orig
		}()
		defaultStartTime = "wrong"
		_, _, err := GetStartEndTimes(nil)
		assert.Error(t, err)
	})
}
