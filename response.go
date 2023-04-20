package utils

import (
	"net/http"
	"strings"
)

// is response accepted as json?
func IsJsonResponse(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("Accept")), "application/json")
}

// is response accepted as xml?
func IsXmlResponse(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("Accept")), "application/xml")
}

// is response accepted as text?
func IsTextResponse(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("Accept")), "text/plain")
}

// is response accepted as html?
func IsHtmlResponse(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("Accept")), "text/html")
}

// is response accepted as javascript?
func IsJavascriptResponse(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("Accept")), "text/javascript")
}

// is response accepted as css?
func IsCssResponse(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("Accept")), "text/css")
}

// is response accepted as image?
func IsImageResponse(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("Accept")), "image/")
}

// is response accepted as video?
func IsVideoResponse(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("Accept")), "video/")
}

// is response accepted as audio?
func IsAudioResponse(r *http.Request) bool {
	return strings.Contains(strings.ToLower(r.Header.Get("Accept")), "audio/")
}
