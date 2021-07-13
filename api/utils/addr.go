/*
Copyright 2021 Gravitational, Inc.

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

package utils

import (
	"net"
	"net/url"
	"strings"

	"github.com/gravitational/trace"
)

// ExtractHostPort takes addresses like "tcp://host:port/path" and returns "host:port".
func ExtractHostPort(addr string) (string, error) {
	if addr == "" {
		return "", trace.BadParameter("missing parameter address")
	}
	if !strings.Contains(addr, "://") {
		addr = "tcp://" + addr
	}
	u, err := url.Parse(addr)
	if err != nil {
		return "", trace.BadParameter("failed to parse %q: %v", addr, err)
	}
	switch u.Scheme {
	case "tcp", "http", "https":
		return u.Host, nil
	default:
		return "", trace.BadParameter("'%v': unsupported scheme: '%v'", addr, u.Scheme)
	}
}

// ExtractHost takes addresses like "tcp://host:port/path" and returns "host".
func ExtractHost(addr string) (ra string, err error) {
	parsed, err := ExtractHostPort(addr)
	if err != nil {
		return "", trace.Wrap(err)
	}
	host, _, err := net.SplitHostPort(parsed)
	if err != nil {
		if strings.Contains(err.Error(), "missing port in address") {
			return addr, nil
		}
		return "", trace.Wrap(err)
	}
	return host, nil
}

// ExtractPort takes addresses like "tcp://host:port/path" and returns "port".
func ExtractPort(addr string) (string, error) {
	parsed, err := ExtractHostPort(addr)
	if err != nil {
		return "", trace.Wrap(err)
	}
	_, port, err := net.SplitHostPort(parsed)
	if err != nil {
		return "", trace.Wrap(err)
	}
	return port, nil
}
