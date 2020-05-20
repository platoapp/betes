package restful

// Copyright 2014 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.
import (
	"github.com/emicklei/go-restful/log"
)

var trace bool = false
var traceLogger log.StdLogger

func init() {
	traceLogger = log.Logger // use the package logger by default
}

// TraceLogger enables detailed logging of H