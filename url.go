// Practice Error Handling with type casting
// In the Go standard library, there’s a function to parse web addresses (see golang.org/pkg/net/url/#Parse). Display the error that occurs when url.Parse is used with
// an invalid web address, such as one containing a space: https://a b.com/.
//
// Use the %#v format verb with Printf to learn more about the error. Then perform a *url.Error type assertion to access and print the fields of the underlying structure.

package main

import (
	"fmt"
	"net/url"
)

func errParsing() {

	//apiEndPoint := "www.google.com"
	apiEndPoint := "https://a b.com/"

	value, err := url.Parse(apiEndPoint)
	if err != nil {
		//fmt.Printf("Error: %#v\n       %v\n", err, err)
		if e, ok := err.(*url.Error); ok { // This is type castimg where e is now of type url.Error, which is a struct woth members Op, URL, Err, etc
			fmt.Printf("Error: On Operation: %v on URL: %v \n     : %v\n", e.Op, e.URL, e.Err)
		}
		return
	}

	fmt.Println(value)
}
