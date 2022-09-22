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
	"net/url"
	"strings"
	"time"
)

// Is2xx determines if a response status code is flagged as OK.
func Is2xx(status int) bool {
	if status < 200 || status >= 300 {
		return false
	}
	return true
}

// Is3xx determines if a response status code is a redirect.
func Is3xx(status int) bool {
	if status < 300 || status >= 400 {
		return false
	}
	return true
}

// GetAbsoluteURL retrieves the absolute URL of a full and
// relative URL.
// Returns errors.INVALID if the urls could not be parsed.
func GetAbsoluteURL(fullURL string, relative string) (string, error) {
	const op = "HTTPUtil.GetAbsoluteURL"
	full, err := url.Parse(fullURL)
	if err != nil {
		return "", errors.NewInvalid(err, "Error parsing full URI", op)
	}
	rel, err := url.Parse(relative)
	if err != nil {
		return "", errors.NewInvalid(err, "Error parsing relative URI", op)
	}
	if !strings.Contains(relative, "http") && !strings.HasPrefix(relative, "./") {
		return full.Scheme + "://" + full.Host + "/" + strings.TrimPrefix(relative, "/"), nil
	}
	if rel.IsAbs() {
		return relative, nil
	}
	return strings.TrimSuffix(fullURL, "/") + "/" + strings.TrimPrefix(strings.TrimPrefix(relative, "./"), "/"), nil
}

var (
	// defaultStartTime is the default start time when none is passed.
	defaultStartTime = "2022-08-01"
)

// GetStartEndTimes retrieves start and end times from a query.
// Returns errors.INVALID if the dates could not be passed.
func GetStartEndTimes(query url.Values) (start time.Time, end time.Time, err error) {
	const op = "HTTPUtil.StartEndTimes"

	defaultStart, err := time.Parse("2006-01-01", defaultStartTime)
	if err != nil {
		return start, end, err
	}

	// Services start time
	startStr := query.Get("start")
	if startStr != "" {
		startT, err := time.Parse("2006-01-02", startStr)
		if err != nil {
			return start, end, errors.NewInvalid(errors.New("invalid start date"), "Error, start date is invalid", op)
		}
		start = startT
	} else {
		start = defaultStart
	}

	// Services end time.
	end = time.Now()
	endStr := query.Get("end")
	if endStr != "" {
		endT, err := time.Parse("2006-01-02", endStr)
		if err != nil {
			return start, end, errors.NewInvalid(errors.New("invalid end date"), "Error, end date is invalid", op)
		}
		end = endT
	}

	return
}
