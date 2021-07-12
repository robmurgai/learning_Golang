package practicefundamentals

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"path/filepath"
	"time"

	"gopkg.in/yaml.v2"
)

// Todo struct
type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

// - name: keyDomain
//   pub_key: <your key>
//	 prvt_key: <your private key>

type Apikeys struct {
	Name    string `yaml:"name"`
	PubKey  string `yaml:"pub_key"`
	PrvtKey string `yaml:"prvt_key"`
}

//API Domain
var apiDomain = ""

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

	// Read the response body as a byte array
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error in ioutil.ReadAll(resp.Body): %v", err.Error())
		panic(err)
	}

	// Convert byte array to string for printing.
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert byte array to a pretty JSON string for printing so you know what kind of structs you want to create
	var understandableBodyBuffer bytes.Buffer
	err = json.Indent(&understandableBodyBuffer, bodyBytes, "", "  ")
	if err != nil {
		fmt.Printf("Error in json.Indent(): %v", err.Error())
		panic(err)
	}
	fmt.Printf("%s\n", understandableBodyBuffer.Bytes())

	// Convert byte array to an effective struct using the Unmarshall function.
	// Unmarshal parses the JSON-encoded data in the first argument bodyBytes and stores the result in the value pointed to by the second argument v. If v is nil or not
	// a pointer, Unmarshal returns an InvalidUnmarshalError.  To unmarshal JSON into a struct, Unmarshal matches incoming object keys to the keys used by either the
	// struct field name or its tag, preferring an exact match but also accepting a case-insensitive match. By default, object keys which don't have a corresponding
	// struct field are ignored. By knowing the structure of the JSON response that the above API endpoint gives us we can crreate a series of structs to map our
	// data to.

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

func getAPIKey(keyDomain string) (Apikeys, error) {

	var apikey Apikeys

	if keyDomain == "" {
		return apikey, nil
	}

	var myKeys []Apikeys

	dir := "config"
	fileName := "api_keys.yaml"
	filePath := filepath.Join(dir, fileName)
	fmt.Printf("DEBUG: filePath: %v\n", filePath)
	bytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return apikey, err
	}

	err = yaml.Unmarshal(bytes, &myKeys)
	if err != nil {
		return apikey, err
	}

	for _, apikey = range myKeys {
		if apikey.Name == keyDomain {
			break
		}
	}

	return apikey, nil
}

func gethash() time.Time {
	return time.Now()
}

//This function is incomplete
func putAPIKeys(keyDomain string) (bool, error) {

	if keyDomain == "" {
		return true, nil
	}

	// apikeys
	// -
	//	 name: keyDomain
	//   pub_key: <your key>

	type apikeys struct {
		Name   string `yaml:"name"`
		PubKey string `yaml:"pub_key"`
	}

	type configuration struct {
		Apikey []apikeys `yaml:"apikeys"`
	}

	dir := "config"
	fileName := "api_keys.yaml"
	filePath := filepath.Join(dir, fileName)
	fmt.Printf("DEBUG: filePath: %v\n", filePath)

	var apikey apikeys
	apikey.Name = "marvel"
	apikey.PubKey = "gibbrishkey"

	var config configuration
	config.Apikey = append(config.Apikey, apikey)

	bytes, err := yaml.Marshal(&config)
	if err != nil {
		return false, err
	}
	fmt.Println(bytes)

	return true, nil
}
