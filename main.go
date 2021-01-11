// Exercise #2: URL Shortener: https://github.com/gophercises/urlshort
//
// The goal of this exercise is to create an http.Handler that will look at the path of any incoming web request and determine if it should redirect
// the user to a new page, much like URL shortener would.
//
// For instance, if we have a redirect setup for /dogs to https://www.somesite.com/a-story-about-dogs we would look for any incoming web requests
// with the path /dogs and redirect them.
//
// To complete this exercises you will need to implement the stubbed out methods in handler.go. There are a good bit of comments explaining what each
// method should do, and there is also a main/main.go source file that uses the package to help you test your code and get an idea of what your
// program should be doing.
//
// I suggest first commenting out all of the code in main.go related to the YAMLHandler function and focusing on implementing the MapHandler
// function.
//
// Once you have that working, focus on parsing the YAML using the gopkg.in/yaml.v2 package.
// Note: You will need to 'go get' this package if you don’t have it already.

// After you get the YAML parsing down, try to convert the data into a map and then use the MapHandler to finish the YAMLHandler implementation. Eg
// you might end up with some code like this:

// func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
//   parsedYaml, err := parseYAML(yaml)
//   if err != nil {
//     return nil, err
//   }
//   pathMap := buildMap(parsedYaml)
//   return MapHandler(pathMap, fallback), nil
// }
//
// But in order for this to work you will need to create functions like parseYAML and buildMap on your own. This should give you ample experience
// working with YAML data.
//
// Bonus
// As a bonus exercises you can also…
// 	Update the main/main.go source file to accept a YAML file as a flag and then load the YAML from a file rather than from a string.
// 	Build a JSONHandler that serves the same purpose, but reads from JSON data.
// 	Build a Handler that doesn’t read from a map but instead reads from a database:  LevelDB, Maria MySQL

package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	urlshort "./gophercises/urlshort"
)

var debugLog *log.Logger

func init() {

	//log.Printf("\nDEBUG: Main(): Init()\n\n")

	//Setting up my DEBUG Logger
	debugLog = log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	//debugLog.Printf("Debug logger is set\n")

}

func main() {

	fmt.Printf("\n########## Starting Exercise ##########\n\n")

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	// pathsToUrls := map[string]string{
	// 	"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
	// 	"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	// }
	// mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	mapHandler := urlshort.MapHandlerDB(mux)

	// Build the YAMLHandler using the mapHandler as the fallback
	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	var mapFileName string
	flag.StringVar(&mapFileName, "f", "", "Name of the file with URL Mapping with .yml or .json extention")
	flag.Parse()

	data := []byte(yaml)
	var myHandler http.HandlerFunc
	var err error

	if mapFileName == "" {
		//debugLog.Printf("DEBUG: main:() No MAP file provided.\n  Parsed default mapping data: %s\n", data)
		myHandler, err = urlshort.YAMLHandler(data, mapHandler)
		if err != nil {
			log.Printf("ERROR: main(): urlshort.YAMLHandler([]byte(yaml), mapHandler): %v\n", err)
			panic(err)
		}
	} else {
		log.Printf("INFO: Initializing Server with URLs from map file: %s\n", mapFileName)
		data = parseMapFile(mapFileName)
		//debugLog.Printf("DEBUG: Parsed URL Mapping data: %s\n", data)

		if strings.Contains(mapFileName, "yml") {
			//debugLog.Printf("DEBUG: main:() MAP file provided in YAML format.\n")
			myHandler, err = urlshort.YAMLHandler(data, mapHandler)
			if err != nil {
				log.Printf("ERROR: main(): urlshort.YAMLHandler([]byte(yaml), mapHandler): %v\n", err)
				panic(err)
			}
		} else if strings.Contains(mapFileName, "json") {
			//log.Printf("DEBUG: main:() MAP file provided in JSON format.\n")
			myHandler, err = urlshort.JSONHandler(data, mapHandler)
			if err != nil {
				log.Printf("ERROR: main(): urlshort.JSONHandler([]byte(yaml), mapHandler): %v\n", err)
				panic(err)
			}
		} else {
			err := fmt.Sprintf("Unable to process file %v, Program can only accept file name with .yml or .jsn extention \n", mapFileName)
			log.Printf("ERROR: main(): %v", err)
			panic(err)
		}
	}
	runWebServer(data, myHandler)
	fmt.Printf("\n########## Ending Exercise ##########\n")
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func parseMapFile(mapFileName string) []byte {

	//Env independent quizfile path
	dir1 := "gophercises"
	dir2 := "urlshort"
	path := filepath.Join(dir1, dir2, mapFileName)
	//debugLog.Printf("DEBUG: parseMapFile(): path: %v\n", path)

	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("ERROR: parseYAMLFile(): ioutil.ReadFile(%v): %v\n", path, err)
		panic(err)
	}

	return data
}

func runWebServer(data []byte, myHandler http.Handler) {
	//debugLog.Printf("DEBUG: runWebServer(): Initilize\n")
	log.Println("INFO: Starting the server on port: 8080")
	http.ListenAndServe(":8080", myHandler)
}
