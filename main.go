// Create an HTTP Server to understand how it works.

package main

import (
	"fmt"
	"log"
	"os"

	hr "./hackerrank"
)

var debugLog *log.Logger

func init() {

	//Setting up my DEBUG Logger
	debugLog = log.New(os.Stderr, "DEBUG: ", log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {

	fmt.Printf("\n########## Starting Exercise ##########\n\n")

	debugLog.Println("Get()")
	hr.Get()

	debugLog.Println("Post()")
	hr.Post()

	debugLog.Println("PostForm()")
	hr.PostForm()

	debugLog.Println("Put()")
	hr.Put()

	debugLog.Println("Delete")
	hr.Delete()

	fmt.Printf("\n########## Ending Exercise ##########\n")
}
