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
)

var debugLog *log.Logger

func init() {

	//log.Printf("\nDEBUG: Main(): Init()\n\n")

	//Setting up my DEBUG Logger
	debugLog = log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)
	debugLog.Printf("Debug logger is set\n")

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

	//start the server
	fmt.Printf("Starting server on Port: 8080\n")
	http.ListenAndServe(":8080", nil)

	fmt.Printf("\n########## Ending Exercise ##########\n")
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
