package urlshort

import (
	"encoding/json"
	"log"
	"net/http"

	// urlDB "./levelDB"
	urlDB "./sqlDB"
	"gopkg.in/yaml.v2"
)

// yml is a byte array of strings.
// Each element in the array is a dictionary with 2 elements path: " " and a corresponding url: " "
// yml := `
// - path: /urlshort
// url: https://github.com/gophercises/urlshort
// - path: /urlshort-final
// url: https://github.com/gophercises/urlshort/tree/solution
// `
type urlPathMap struct {
	URLOriginal string `yaml:"url_original" json:"url_original"`
	URLRedirect string `yaml:"url_redirect" json:"url_redirect"`
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	//log.Printf("DEBUG: MapHandler(): Creating a map[string]string based Handler Function.\n")
	//log.Printf("DEBUG: MapHandler(): Received pathsToUrls: %+v\n", pathsToUrls)

	tmpHandler := func(resp http.ResponseWriter, reqPtr *http.Request) {

		receivedURL := reqPtr.URL.Path
		//log.Printf("DEBUG: tmpHandler(): reqPtr.URL.Path: %q\n", reqPtr.URL.Path)

		// if the recieved path is in pathsToUrls, redirect the server from key URL to Value URL
		if value, ok := pathsToUrls[receivedURL]; ok {
			redirectURL := value
			//log.Printf("DEBUG: tmpHandler(): Redirecting: %q to %q\n", receivedURL, redirectURL)
			http.Redirect(resp, reqPtr, redirectURL, http.StatusFound)

		} else if receivedURL == "/google" { // Redirect to Google.com for fun
			redirectURL := "https://www.google.com"
			//log.Printf("DEBUG: tmpHandler(): Redirecting for fun: %q to %q\n", receivedURL, redirectURL)
			http.Redirect(resp, reqPtr, redirectURL, http.StatusFound)
		} else { // if the URL is not in PathsToURL[] return the fallback handler func
			//log.Printf("DEBUG: tmpHandler(): Falling back to the fallback() function.\n")
			fallback.ServeHTTP(resp, reqPtr)
		}

	}

	return tmpHandler
}

// MapHandlerDB will return an http.HandlerFunc that will attempt to map any paths in the LevelDB to their corresponding URL (values
// that each key in the map points to, in string format). If the path is not provided in the LevelDB, then the fallback http.Handler will be called instead.
func MapHandlerDB(fallback http.Handler) http.HandlerFunc {

	//log.Printf("DEBUG: MapHandler(): Creating a map[string]string based Handler Function.\n")
	//log.Printf("DEBUG: MapHandler(): Received pathsToUrls: %+v\n", pathsToUrls)

	// For Debug Purpose only
	//PrintMapData()

	tmpHandler := func(resp http.ResponseWriter, reqPtr *http.Request) {

		receivedURL := reqPtr.URL.Path
		//log.Printf("DEBUG: tmpHandler(): reqPtr.URL.Path: %q\n", reqPtr.URL.Path)

		// if the recieved path is in pathsToUrls, redirect the server from key URL to Value URL
		if value, ok := urlDB.GetURLRedirect(receivedURL); ok {
			redirectURL := value
			//log.Printf("DEBUG: tmpHandler(): Redirecting: %q to %q\n", receivedURL, redirectURL)
			http.Redirect(resp, reqPtr, redirectURL, http.StatusFound)

		} else if receivedURL == "/google" { // Redirect to Google.com for fun
			redirectURL := "https://www.google.com"
			//log.Printf("DEBUG: tmpHandler(): Redirecting for fun: %q to %q\n", receivedURL, redirectURL)
			http.Redirect(resp, reqPtr, redirectURL, http.StatusFound)
		} else { // if the URL is not in PathsToURL[] return the fallback handler func
			//log.Printf("DEBUG: tmpHandler(): Falling back to the fallback() function.\n")
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

	//log.Printf("DEBUG: YAMLHandler(): Creating a YAML based Handler Function\n")

	//log.Printf("DEBUG: YAMLHandler(): Received yml: %s\n", yml)

	ymlArray, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	//log.Printf("DEBUG: YAMLHandler(): ymlArray: %+v\n", ymlArray)

	pathsToUrls := buildMAP(ymlArray)
	//log.Printf("DEBUG: YAMLHandler(): pathsToUrls: %+v\n", pathsToUrls)

	mapHandler := MapHandler(pathsToUrls, fallback)

	return mapHandler, nil

}

func parseYAML(data []byte) ([]urlPathMap, error) {
	//log.Printf("DEBUG: parseYAML(): Received yml: %s\n", yml)

	var yamlArray []urlPathMap

	err := yaml.Unmarshal(data, &yamlArray)
	if err != nil {
		log.Printf("ERROR: parseYAML(): \n%v\n while Unmarshaling the following byte array: %v\n", err, data)
		return nil, err
	}
	//log.Printf("DEBUG: parseYAML(): Created ymlArray: %+v\n", ymlArray)
	return yamlArray, nil

}

// JSONHandler will parse the provided JSON and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
func JSONHandler(data []byte, fallback http.Handler) (http.HandlerFunc, error) {

	//log.Printf("DEBUG: JSONHandler(): Creating a JSON based Handler Function\n")

	//log.Printf("DEBUG: JSONHandler(): Received byte array: %s\n", data)

	jsonArray, err := parseJSON(data)
	if err != nil {
		return nil, err
	}
	//log.Printf("DEBUG: JSONHandler(): jsonArray: %+v\n", jsonArray)

	pathsToUrls := buildMAP(jsonArray)
	//log.Printf("DEBUG: JSONHandler(): pathsToUrls: %+v\n", pathsToUrls)

	mapHandler := MapHandler(pathsToUrls, fallback)

	return mapHandler, nil

}

func parseJSON(data []byte) ([]urlPathMap, error) {
	//log.Printf("DEBUG: parseYAML(): Received jsn: %s\n", yml)

	var jsonArray []urlPathMap

	err := json.Unmarshal(data, &jsonArray)
	if err != nil {
		log.Printf("ERROR: parseJSON(): \n%v\n while Unmarshaling the following byte array: %v\n", err, data)
		return nil, err
	}
	//log.Printf("DEBUG: parseYAML(): Created ymlArray: %+v\n", ymlArray)
	return jsonArray, nil

}

func buildMAP(urlArray []urlPathMap) map[string]string {
	//log.Printf("DEBUG: buildMAP((): Received urlArray: %+v\n", urlArray)
	pathsToUrls := make(map[string]string, len(urlArray))
	for _, value := range urlArray {
		pathsToUrls[value.URLOriginal] = value.URLRedirect
	}
	//log.Printf("DEBUG: buildMAP(): Created pathsToUrls: %+v\n", pathsToUrls)
	return pathsToUrls
}
