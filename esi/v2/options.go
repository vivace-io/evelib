package esi

import "net/http"

// Options configures the ESI client.
type Options struct {
	// UserAgent (required) the user agent sent in requests to ESI. This gives
	// CCP developers the ability to identify and contact you if you're doing
	// something wrong, or your application grows sentient. Try to include an
	// e-mail and/or character name.
	UserAgent string
	// HTTPClient (optional) the HTTP Client to be used in executing requests to
	// ESI resources and, ideally, handling caching.
	// NOTE: While this library does not provide any sort of caching, you might
	// check out the httpcache library. (https://github.com/gregjones/httpcache)
	HTTPClient *http.Client
	// APIRoot (optional) the root address of the ESI API without a trailing slash.
	// Ideally, this should be left blank unless using a proxy or cache service.
	// When left blank, defaults to 'https://esi.tech.ccp.is'.
	APIRoot string
}
