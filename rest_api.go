package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type todoStruct struct {

	//Use Caps for field names to enable exporting to the main function and use json tags to ensure one to one mappint
	UserID    int    `jason:"userId"`
	ID        int    `jason:"id"`
	Title     string `jason:"title"`
	Completed bool   `json:"completed"`
}

func get() {

	//API Endpoint for Get
	var apiEndPoint = "https://jsonplaceholder.typicode.com/todos/1"

	//Call the Endpoint.
	apiResponse, err := http.Get(apiEndPoint)
	if err != nil {
		fmt.Printf("get(%v): %v", apiEndPoint, err.Error())
		panic(err)
	}

	//Close the reponse Body at the end of the section, in this case the get() function.
	defer apiResponse.Body.Close()

	//Reponse Code 200 means Okay
	if apiResponse.StatusCode == 200 {
		body, err := ioutil.ReadAll(apiResponse.Body)
		if err != nil {
			fmt.Printf("ioutil.ReadAll(%v): %v", apiEndPoint, err.Error())
			panic(err)
		}

		fmt.Printf("body from API Endpoint: \n%v\n", string(body))

		var todo todoStruct
		err = json.Unmarshal(body, &todo)
		if err != nil {
			fmt.Printf("Error: json.Unmarshal(body): %v", err.Error())
			panic(err)
		}
		fmt.Printf("todo from body: %+v\n", todo)

	} else {
		fmt.Printf("Unable to extract data, response Status from %v: %v\n", apiEndPoint, apiResponse.StatusCode)
	}

}
