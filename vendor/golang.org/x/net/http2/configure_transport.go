// Copyright 2015 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build go1.6

package http2

import (
	"crypto/tls"
	"fmt"
	"net/http"
)

func configureTransport(t1 *http.Transport) (*Transport, error) {
	connPool := new(clientConnPool)
	t2 := &Transport{
		ConnPool: noDialClientConnPool{connPool},
		t1:       t1,
	}
	connPool.t = t2
	if err := registerHTTPSProtocol(t1, noDialH2RoundTripper{t2}); err != nil {
		return nil, err
	}
	if t1.TLSClientCo