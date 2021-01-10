package urlshort

import (
	"fmt"
	"net/http"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	fmt.Printf("DEBUG: MapHandler(): Creating a Handler Function.\n")

	tmpHandler := func(resp http.ResponseWriter, reqPtr *http.Request) {

		receivedURL := reqPtr.URL.Path
		fmt.Printf("DEBUG: tmpHandler(): reqPtr.URL.Path: %q\n", reqPtr.URL.Path)

		// if the recieved path is in pathsToUrls, redirect the server from key URL to Value URL
		if value, ok := pathsToUrls[receivedURL]; ok {
			redirectURL := value
			fmt.Printf("DEBUG: tmpHandler(): Redirecting: %q to %q\n", receivedURL, redirectURL)
			http.Redirect(resp, reqPtr, redirectURL, http.StatusFound)

		} else if receivedURL == "/google" { // Redirect to Google.com for fun
			redirectURL := "https://www.google.com"
			fmt.Printf("DEBUG: tmpHandler(): Redirecting for fun: %q to %q\n", receivedURL, redirectURL)
			http.Redirect(resp, reqPtr, redirectURL, http.StatusFound)
		} else { // if the URL is not in PathsToURL[] return the fallback handler func
			fmt.Printf("DEBUG: tmpHandler(): Falling back to the fallback() function.\n")
			fallback.ServeHTTP(resp, reqPtr)
		}

	}

	return tmpHandler
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// TODO: Implement this...
	return nil, nil
}
