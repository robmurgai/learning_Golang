// Create an HTTP Server to understand how it works.

package main

import (
	"fmt"
	"log"
	"os"

	httpServer "./http_server"
)

var debugLog *log.Logger

func init() {

	//Setting up my DEBUG Logger
	debugLog = log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {

	fmt.Printf("\n########## Starting Exercise ##########\n\n")

	httpServer.Start()

	fmt.Printf("\n########## Ending Exercise ##########\n")
}
