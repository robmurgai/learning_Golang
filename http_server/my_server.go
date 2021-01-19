// Create an HTTP Server to understand how it works.
// Test the function with the following CURL commands
// 		curl -v http://localhost:8080/handler_struct
// 		curl -v http://localhost:8080/handler_function

package httpserver

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

var debugLog *log.Logger
var publicDir = "public"
var httpServerDir = "http_server"

func init() {

	//log.Printf("\nDEBUG: Main(): Init()\n\n")

	//Setting up my DEBUG Logger
	debugLog = log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	//debugLog.Printf("Debug logger is set\n")

}

// Start starts a simple webserver and registers two addresses
// "/handler_function" with urlHandlerFunc
// "/handler_struct" with urlHandlerStruct
func Start() {

	// Register the URL pattern and the handler pairing

	// Register "/handler_function" with urlHandlerFunc
	http.HandleFunc("/handler_function", urlHandlerFunc)

	// Register "/handler_struct" with urlHandlerStruct
	urlhs := urlHandlerStruct{}
	http.Handle("/handler_struct", urlhs)

	// Register "/" with queryHandler()
	http.HandleFunc("/", queryHandlerFunc)

	//start the server
	fmt.Printf("Starting server on Port: 8080\n")
	http.ListenAndServe(":8080", nil)

}

// CustomStart starts a custome HTTP server and registers two addresses
// "/handler_function" with urlHandlerFunc
// "/handler_struct" with urlHandlerStruct
func CustomStart() {

	//Create a custom multiplexer
	mux := http.NewServeMux()

	// Register the URL pattern and the handler pairing

	// Register "/" with statis files
	filePath := filepath.Join(httpServerDir, publicDir)
	fileServer := http.FileServer(http.Dir(filePath))
	mux.Handle("/", fileServer)

	// Register "/" with processGetOrPostHandlerFunc()
	mux.HandleFunc("/form", processGetOrPost)

	//Create a custom HTTP Listener
	s := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	//start the server
	fmt.Printf("Starting Custom HTTP server on Port: 8080\n")
	s.ListenAndServe()

}

//Create the handlers
func urlHandlerFunc(res http.ResponseWriter, reqPtr *http.Request) {
	data := []byte("This is my flight song with the urlHandlerfunc()")
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

type urlHandlerStruct struct{}

func (urlhs urlHandlerStruct) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	data := []byte("This is my flight song with the urlHandlerstruct")
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}

func queryHandlerFunc(res http.ResponseWriter, reqPtr *http.Request) {
	name := "guest"
	keys, ok := reqPtr.URL.Query()["name"]
	if ok {
		name = keys[0]
	}
	fmt.Fprintf(res, "Hello %s!\n", name)
}

func processGetOrPost(w http.ResponseWriter, r *http.Request) {

	//debugLog.Printf("processGetOrPost(): URL Path: %v\n", r.URL.Path)

	if r.URL.Path != "/form" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		//debugLog.Printf("processGetOrPost(): r.Method: %v\n", r.Method)
		filePath := filepath.Join(httpServerDir, publicDir, "form.html")
		//debugLog.Printf("processGetOrPost(): filePath: %v\n", filePath)
		http.ServeFile(w, r, filePath)
	case "POST":
		//debugLog.Printf("processGetOrPost(): r.Method: %v\n", r.Method)
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		name := r.FormValue("name")
		loc := r.FormValue("location")
		fmt.Fprintf(w, "%s is at %s\n", name, loc)
	default:
		fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	}
}
