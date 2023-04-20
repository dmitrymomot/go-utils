package utils

import (
	"net/http"
	"strings"
)

// Is request a json request?
func IsJsonRequest(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("Content-Type")), "application/json")
}

// Is request a xml request?
func IsXmlRequest(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("Content-Type")), "application/xml")
}

// Is request a multipart request?
func IsMultipartRequest(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("Content-Type")), "multipart/form-data")
}

// Is request a form request?
func IsFormRequest(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("Content-Type")), "application/x-www-form-urlencoded")
}

// Is request a text request?
func IsTextRequest(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("Content-Type")), "text/plain")
}
