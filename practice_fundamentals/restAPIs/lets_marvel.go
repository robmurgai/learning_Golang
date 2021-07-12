package practicefundamentals

import (
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

func LetsMarvel() {

	// Set the API EnPoint URL
	apiDomain := "marvel"
	apiEndPoint := "https://gateway.marvel.com/v1/public/characters"

	apiKey, err := getAPIKey(apiDomain)
	if err != nil {
		fmt.Printf("Error getAPIKey(%v): %v", apiDomain, err.Error())
		panic(err)
	}
	if apiKey.Name != "" {

		// Add Time Stamp
		ts := strconv.FormatInt(time.Now().Unix(), 10)
		apiEndPoint = apiEndPoint + "?ts=" + ts
		fmt.Printf("timestamp: %s\n", ts)

		//Add API Key
		apiEndPoint = apiEndPoint + "&apikey=" + apiKey.PubKey
		fmt.Printf("public key: %s\n", apiKey.PubKey)

		//Add Hash md5(ts+privateKey+publicKey)
		data := []byte(ts + apiKey.PrvtKey + apiKey.PubKey)
		hashData := md5.Sum(data)
		fmt.Printf("hashdata: %x\n", hashData)

		apiEndPoint = apiEndPoint + "&hash=" + fmt.Sprintf("%x", hashData)
	}

	fmt.Printf("API Enpoint: %v\n", apiEndPoint)

	// Make a REST API Get Call

	response, err := http.Get(apiEndPoint)
	if err != nil {
		fmt.Printf("Error http.Get(%v): %v\n", apiEndPoint, err.Error())
		panic(err)
	}

	defer response.Body.Close()

	responseInBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error trying to read Response: %v\n", err.Error())
		panic(err)
	}

	// fmt.Println(string(responseInBytes))

	// Preety Print the JSON response.

	var prettyResponse bytes.Buffer
	err = json.Indent(&prettyResponse, responseInBytes, "", "  ")
	if err != nil {
		fmt.Printf("Error: json.Indent(%v): %v", responseInBytes, err.Error())
		panic(err)
	}

	fmt.Printf("%s", prettyResponse)

	// Create a struct for the Response.
	// {
	// "code": 200,
	// "status": "Ok",
	// "copyright": "© 2021 MARVEL",
	// "attributionText": "Data provided by Marvel. © 2021 MARVEL",
	// "attributionHTML": "<a href=\"http://marvel.com\">Data provided by Marvel. © 2021 MARVEL</a>",
	// "etag": "66e10fa20cdf7ef9db54b4dc9ddc1dad71601e44",
	// "data": {
	//   "offset": 0,
	//   "limit": 20,
	//   "total": 1493,
	//   "count": 20,
	//   "results": [
	// 	{
	//    "id": 1017100,
	//    "name": "A-Bomb (HAS)",
	//    "description": "Rick Jones has been Hulk's best bud since day one, but now he's more than a friend...he's a teammate! Transformed by a Gamma energy explosion, A-Bomb's thick, armored skin is just as strong and powerful as it is blue. And when he curls into action, he uses it like a giant bowling ball of destruction! ",
	//    "modified": "2013-09-18T15:54:04-0400",
	//    "thumbnail": {
	//      "path": "http://i.annihil.us/u/prod/marvel/i/mg/3/20/5232158de5b16",
	//      "extension": "jpg"
	//    },
	//    "resourceURI": "http://gateway.marvel.com/v1/public/characters/1017100",
	//    "comics": {
	//      "available": 3,
	//      "collectionURI": "http://gateway.marvel.com/v1/public/characters/1017100/comics",
	//      "items": [
	//      {
	// 	       "resourceURI": "http://gateway.marvel.com/v1/public/comics/40632",
	// 	       "name": "Hulk (2008) #53"
	//        },
	//      {

	type items struct {
		Name string `json:"name"`
	}

	type comics struct {
		Available int     `json:"available"`
		Items     []items `json:"items"`
	}

	type results struct {
		Name        string `json:"name"`
		Description string `json:"description"`
		Comics      comics `json:"comics"`
	}

	type marvelData struct {
		Total   int       `json:"toal"`
		Count   int       `json:"count"`
		Results []results `json:"results`
	}

	type marvelCharecters struct {
		Code   int        `json:"code"`
		Status string     `json:"status"`
		Etag   string     `json:"etag"`
		Data   marvelData `json:"data"`
	}

	var mc marvelCharecters

	err = json.Unmarshal(responseInBytes, &mc)
	if err != nil {
		fmt.Printf("Error: json.Unmarshal(response): %v", err.Error())
		panic(err)
	}

	fmt.Printf("mc: %+v\n", mc)
}
