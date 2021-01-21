package hackerrank

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

// Todo struct
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

//API Endpoint for HTTP Calls
var apiEndPoint = "https://jsonplaceholder.typicode.com/todos/1"

//Get calls http.Get()
func Get() {

	fmt.Println("1. Performing Http Get...")

	//Make a Get() request
	resp, err := http.Get(apiEndPoint)
	if err != nil {
		fmt.Printf("get(%v): %v", apiEndPoint, err.Error())
		panic(err)
	}

	//Close the reponse Body to avoid resource leakage.
	defer resp.Body.Close()

	// Convert response body to a byte array
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Convert byte array to string for printing.
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert byte array to Todo struct
	var todo Todo
	err = json.Unmarshal(bodyBytes, &todo)
	if err != nil {
		fmt.Printf("Error: json.Unmarshal(body): %v", err.Error())
		panic(err)
	}
	fmt.Printf("todo from body: %+v\n", todo)

}

//Post calls http.Post()
func Post() {

	fmt.Println("2. Performing Http Post...")

	// Create the JSON byte array of key value mapping we want to Post
	req, err := json.Marshal(map[string]string{
		"username": "userName",
		"email":    "user@email.org",
	})
	if err != nil {
		fmt.Println(err)
	}

	// Make a Post Request
	resp, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(req))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// Convert response body to a byte array
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Convert byte array to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

}

//PostForm calls http.PostForm
func PostForm() {

	fmt.Println("3. Performing Http Post (with Form)...")

	//Create a URL encoded Key Value data to send in the Post Reqeust
	formData := url.Values{
		"username": {"userNameNew"},
	}

	// Make a Post Request
	resp, err := http.PostForm("https://httpbin.org/post", formData)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	// Convert the response Body to a map of whatever is returned by the API Endpoint.
	var data map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(data["form"])
}

//Put calls the http.NewRequest(http.MethodPut, ....
func Put() {

	fmt.Println("4. Performing Http Put...")

	// Create the object we will send in the Put request.
	reqTodo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	fmt.Printf("%+v\n", reqTodo)

	// Convert the struct to json
	jsonReq, err := json.Marshal(reqTodo)
	if err != nil {
		fmt.Println(err)
	}

	// Create a Put Request
	req, err := http.NewRequest(http.MethodPut, apiEndPoint, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	// Make the Put Request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// Convert response body to a byte array
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert byte array to Todo struct
	var resTodo Todo
	err = json.Unmarshal(bodyBytes, &resTodo)
	if err != nil {
		fmt.Printf("Error: json.Unmarshal(body): %v", err.Error())
		panic(err)
	}
	fmt.Printf("API Response as struct:\n%+v\n", resTodo)
}

//Delete calls the http.NewRequest(http.MethodDelete ...
func Delete() {

	fmt.Println("5. Performing Http Delete...")

	// Create the object we will send in the Delete request.
	reqTodo := Todo{1, 2, "lorem ipsum dolor sit amet", true}
	fmt.Printf("%+v\n", reqTodo)

	// Convert the struct to json
	jsonReq, err := json.Marshal(reqTodo)
	if err != nil {
		fmt.Println(err)
	}

	// Create a Delete Request
	req, err := http.NewRequest(http.MethodDelete, apiEndPoint, bytes.NewBuffer(jsonReq))
	if err != nil {
		log.Fatalln(err)
	}

	// Make the Delete Request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	// Convert response body to a byte array
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}
