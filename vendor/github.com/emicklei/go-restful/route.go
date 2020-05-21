package restful

// Copyright 2013 Ernest Micklei. All rights reserved.
// Use of this source code is governed by a license
// that can be found in the LICENSE file.

import (
	"net/http"
	"strings"
)

// RouteFunction declares the signature of a function that can be bound to a Route.
type RouteFunction func(*Request, *Response)

// RouteSelectionConditionFunction declares the signature of a function that
// can be used to add extra conditional logic when selecting whether the route
// matches the HTTP request.
type RouteSelectionConditionFunction func(httpRequest *http.Request) bool

// Route binds a HTTP Method,Path,Consumes combination to a RouteFunction.
type Route struct {
	Method   string
	Produces []string
	Consumes []string
	Path     string // webservice root path + described path
	Function RouteFunction
	Filters  []FilterFunction
	If       []RouteSelectionConditionFunction

	// cached values for dispatching
	relativePath string
	pathParts    []string
	pathExpr     *pathExpression // cached compilation of relativePath as RegExp

	// documentation
	Doc                     string
	Notes                   string
	Operation               string
	ParameterDocs           []*Parameter
	ResponseErrors          map[int]ResponseError
	ReadSample, WriteSample interface{} // structs that model an example request or response payload

	// Extra information used to store custom information about the route.
	Metadata map[string]interface{}

	// marks a route as deprecated
	Deprecated bool
}

// Initialize for Route
func (r *Route) postBuild() {
	r.pathParts = tokenizePath(r.Path)
}

// Create Request and Response from their http versions
func (r *Route) wrapRequestResponse(httpWriter http.ResponseWriter, httpRequest *http.Request, pathParams map[string]string) (*Request, *Response) {
	wrappedRequest := NewRequest(httpRequest)
	wrappedRequest.pathParameters = pathParams
	wrappedRequest.selectedRoutePath = r.Path
	wrappedResponse := NewResponse(httpWriter)
	wrappedResponse.requestAccept = httpRequest.Header.Get(HEADER_Accept)
	wrappedResponse.routeProduces = r.Produces
	return wrappedRequest, wrappedResponse
}

// dispatchWithFilters call the function after passing through its own filters
func (r *Route) dispatchWithFilters(wrappedRequest *Request, wrappedResponse *Response) {
	if len(r.Filters) > 0 {
		chain := FilterChain{Filters: r.Filters, Target: r.Function}
		chain.ProcessFilter(wrappedRequest, wrappedResponse)
	} else {
		// unfiltered
		r.Function(wrappedRequest, wrappedResponse)
	}
}

// Return whether the mimeType matches to what this Route can produce.
func (r Route) matchesAccept(mimeTypesWithQuality string) bool {
	parts := strings.Split(mimeTypesWithQuality, ",")
	for _, each := range parts {
		var withoutQuality string
		if strings.Contains(each, ";") {
			withoutQuality = strings.Split(each, ";")[0]
		} else {
			withoutQuality = each
		}
		// trim before compare
		withoutQuality = strings.Trim(withoutQuality, " ")
		if withoutQuality == "*/*" {
			return true
		}
		for _, producibleType := range r.Produces {
			if producibleType == "*/*" || producibleType == withoutQuality {
				return true
			}
		}
	}
	return false
}

// Return whether this Route can consume content with a type specified by mimeTypes (can be empty).
func (r Route) matchesContentType(mimeTypes string) bool {

	if len(r.Consumes) == 0 {
		// did not specify what it can consume ;  any media type (“*/*”) is assumed
		return true
	}

	if len(mimeType