package restful

// Copyright 2013 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"bufio"
	"compress/gzip"
	"compress/zlib"
	"errors"
	"io"
	"net"
	"net/http"
	"strings"
)

// OBSOLETE : use restful.DefaultContainer.EnableContentEncoding(true) to change this setting.
var EnableContentEncoding = false

// CompressingResponseWriter is a http.ResponseWriter that can perform content encoding (gzip and zlib)
type CompressingResponseWriter struct {
	writer     http.ResponseWriter
	compressor io.WriteCloser
	encoding   string
}

// Header is part of http.ResponseWriter interface
func (c *CompressingResponseWriter) Header() http.Header {
	return c.writer.Header()
}

// WriteHeader is part of http.ResponseWriter interface
func (c *CompressingResponseWriter) WriteHeader(status int) {
	c.writer.WriteHeader(status)
}

// Write is part of http.ResponseWriter interface
// It is passed through the compressor
func (c *CompressingResponseWriter) Write(bytes []byte) (int, error) {
	if c.isCompressorClosed() {
		return -1, errors.New("Compressing error: tried to write data using closed compressor")
	}
	return c.compressor.Write(bytes)
}

// CloseNotify is part of http.CloseNotifier interface
func (c *CompressingResponseWriter) CloseNotify() <-chan bool {
	return c.writer.(http.CloseNotifier).CloseNotify()
}

// Close the underlying compressor
func (c *CompressingResponseWriter) Close() error {
	if c.isCompressorClosed() {
		return errors.New("Compressing error: tried to close already closed compressor")
	}

	c.compressor.Close()
	if ENCODING_GZIP == c.encoding {
		currentCompressorProvider.ReleaseGzipWriter(c.compressor.(*gzip.Writer))
	}
	if ENCODING_DEFLATE == c.encoding {
		currentCompressorProvider.ReleaseZlibWriter(c.compressor.(*zlib.Writer))
	}
	// gc hint needed?
	c.compressor = nil
	return nil
}

func (c *CompressingResponseWriter) isCompressorClosed() bool {
	return nil == c.compressor
}

// Hijack implements the Hijacker interface
// This is especially useful when combining Container.EnabledContentEncoding
// in combination with websockets (for instance gorilla/websocket)
func (c *CompressingResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := c.writer.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("ResponseWriter doesn't support Hijacker interface")
	}
	return hijacker.Hijack()
}

// WantsCompressedResponse reads the Accept-Encoding header to see if and which encoding is requested.
func wantsCompressedResponse(httpRequest *http.Request) (bool, string) {
	header := httpRequest.Header.Get(HEADER_AcceptEncoding)
	gi := strings.Index(header, ENCODING_GZIP)
	zi := strings.Index(header, ENCODING_DEFLATE)
	// use in order of appearance
	if gi == -1 {
		return zi != -1, ENCODING_DEFLATE
	} els