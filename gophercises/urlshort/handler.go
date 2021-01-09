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
	//	TODO: Implement this...
	fmt.Printf("DEBUG: MapHandler(): \n")

	fmt.Printf("DEBUG: MapHandler(): Mapping \"/\" to handler func notFoundHF()\n")
	http.HandleFunc("/", notFoundHF)

	fmt.Printf("DEBUG: MapHandler(): Mapping \"/google\" to handler func googleHF()\n")
	http.HandleFunc("/google", googleHF)

	//Map all URL values to tmpHF
	for key := range pathsToUrls {

		switch key {
		case "/urlshort-godoc":
			fmt.Printf("DEBUG: MapHandler(): Mapping %q to handler func urlHF()\n", key)
			http.HandleFunc(key, urlHF)
		default:
			fmt.Printf("DEBUG: MapHandler(): Mapping %q to handler func tmpHF()\n", key)
			http.HandleFunc(key, notFoundHF)
		}

	}

	return nil
}

func notFoundHF(resp http.ResponseWriter, reqPtr *http.Request) {

	fmt.Printf("DEBUG: notFoundHF() reqPtr.URL.Path: %v\n", reqPtr.URL.Path)

	// populate the resp
	fmt.Fprintf(resp, "Hello, We are not set up to handle this URL: %q\n", reqPtr.URL.Path)
}

func urlHF(resp http.ResponseWriter, reqPtr *http.Request) {

	receivedURL := reqPtr.URL
	redirectURL := "https://godoc.org/github.com/gophercises/urlshort"

	fmt.Printf("Received URL: %v\n", receivedURL)

	fmt.Printf("DEBUG: urlHF(): Redirecting: %q to %q\n", receivedURL, redirectURL)
	http.Redirect(resp, reqPtr, redirectURL, http.StatusFound)
}

func googleHF(resp http.ResponseWriter, reqPtr *http.Request) {

	receivedURL := reqPtr.URL
	redirectURL := "https://www.google.com"

	fmt.Printf("Received URL: %v\n", receivedURL)

	fmt.Printf("DEBUG: googleHF(): Redirecting: %q to %q\n", receivedURL, redirectURL)
	http.Redirect(resp, reqPtr, redirectURL, http.StatusFound)
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
