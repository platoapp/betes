
/*
Copyright 2016 The Kubernetes Authors.

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

package net

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"

	"github.com/golang/glog"
	"golang.org/x/net/http2"
)

// JoinPreservingTrailingSlash does a path.Join of the specified elements,
// preserving any trailing slash on the last non-empty segment
func JoinPreservingTrailingSlash(elem ...string) string {
	// do the basic path join
	result := path.Join(elem...)

	// find the last non-empty segment
	for i := len(elem) - 1; i >= 0; i-- {
		if len(elem[i]) > 0 {
			// if the last segment ended in a slash, ensure our result does as well
			if strings.HasSuffix(elem[i], "/") && !strings.HasSuffix(result, "/") {
				result += "/"
			}
			break
		}
	}

	return result
}

// IsProbableEOF returns true if the given error resembles a connection termination
// scenario that would justify assuming that the watch is empty.
// These errors are what the Go http stack returns back to us which are general
// connection closure errors (strongly correlated) and callers that need to
// differentiate probable errors in connection behavior between normal "this is
// disconnected" should use the method.
func IsProbableEOF(err error) bool {
	if uerr, ok := err.(*url.Error); ok {
		err = uerr.Err
	}
	switch {
	case err == io.EOF:
		return true
	case err.Error() == "http: can't write HTTP request on broken connection":
		return true
	case strings.Contains(err.Error(), "connection reset by peer"):
		return true
	case strings.Contains(strings.ToLower(err.Error()), "use of closed network connection"):
		return true
	}
	return false
}

var defaultTransport = http.DefaultTransport.(*http.Transport)

// SetOldTransportDefaults applies the defaults from http.DefaultTransport
// for the Proxy, Dial, and TLSHandshakeTimeout fields if unset
func SetOldTransportDefaults(t *http.Transport) *http.Transport {
	if t.Proxy == nil || isDefault(t.Proxy) {
		// http.ProxyFromEnvironment doesn't respect CIDRs and that makes it impossible to exclude things like pod and service IPs from proxy settings
		// ProxierWithNoProxyCIDR allows CIDR rules in NO_PROXY
		t.Proxy = NewProxierWithNoProxyCIDR(http.ProxyFromEnvironment)
	}
	if t.Dial == nil {
		t.Dial = defaultTransport.Dial
	}
	if t.TLSHandshakeTimeout == 0 {
		t.TLSHandshakeTimeout = defaultTransport.TLSHandshakeTimeout
	}
	return t
}

// SetTransportDefaults applies the defaults from http.DefaultTransport
// for the Proxy, Dial, and TLSHandshakeTimeout fields if unset
func SetTransportDefaults(t *http.Transport) *http.Transport {
	t = SetOldTransportDefaults(t)
	// Allow clients to disable http2 if needed.
	if s := os.Getenv("DISABLE_HTTP2"); len(s) > 0 {
		glog.Infof("HTTP2 has been explicitly disabled")
	} else {
		if err := http2.ConfigureTransport(t); err != nil {
			glog.Warningf("Transport failed http2 configuration: %v", err)
		}
	}
	return t
}

type RoundTripperWrapper interface {
	http.RoundTripper
	WrappedRoundTripper() http.RoundTripper
}

type DialFunc func(net, addr string) (net.Conn, error)

func DialerFor(transport http.RoundTripper) (DialFunc, error) {
	if transport == nil {
		return nil, nil
	}

	switch transport := transport.(type) {
	case *http.Transport:
		return transport.Dial, nil
	case RoundTripperWrapper:
		return DialerFor(transport.WrappedRoundTripper())
	default:
		return nil, fmt.Errorf("unknown transport type: %T", transport)
	}
}

type TLSClientConfigHolder interface {
	TLSClientConfig() *tls.Config
}

func TLSClientConfig(transport http.RoundTripper) (*tls.Config, error) {
	if transport == nil {
		return nil, nil
	}

	switch transport := transport.(type) {
	case *http.Transport:
		return transport.TLSClientConfig, nil
	case TLSClientConfigHolder:
		return transport.TLSClientConfig(), nil
	case RoundTripperWrapper:
		return TLSClientConfig(transport.WrappedRoundTripper())
	default:
		return nil, fmt.Errorf("unknown transport type: %T", transport)
	}
}

func FormatURL(scheme string, host string, port int, path string) *url.URL {
	return &url.URL{
		Scheme: scheme,
		Host:   net.JoinHostPort(host, strconv.Itoa(port)),
		Path:   path,
	}
}

func GetHTTPClient(req *http.Request) string {
	if userAgent, ok := req.Header["User-Agent"]; ok {
		if len(userAgent) > 0 {
			return userAgent[0]
		}
	}
	return "unknown"
}