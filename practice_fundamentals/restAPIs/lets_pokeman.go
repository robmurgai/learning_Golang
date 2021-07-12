package practicefundamentals

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// https://tutorialedge.net/golang/consuming-restful-api-with-go/
// In this tutorial, We are going to consume a RESTful API using go. There are currently hundreds upon thousands of open REST APIs out there and Todd Motto has put
// together quite an active repo on Github that lists all the public APIs that are available for consumption by us and he’s categorized them so that we can easily drill
// down to what we want to check out, you can find that here.
//
// For the purpose of this tutorial we wanted use an already live API that we can easily test to see if it works in our browser. We’ll be using the
// very popular pokeapi which is an API that exposes all the known information for Pokemon from the region of Kanto. A bit silly I know but it’s a fully fledged API that
// follows standard naming conventions and requires no authentication so there is no barrier to entry.

func LetsPokeman() {

	//API Endpoint for HTTP Calls
	var apiEndPoint = "http://pokeapi.co/api/v2/pokedex/kanto/"

	fmt.Println("Performing Http Get on API Endpoint: " + apiEndPoint)

	//Make a Get() request
	resp, err := http.Get(apiEndPoint)
	if err != nil {
		fmt.Printf("http.Get(%v): %v", apiEndPoint, err.Error())
		panic(err)
	}

	// Close the reponse Body to avoid resource leakage.
	defer resp.Body.Close()

	// Read the response body as a byte array
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error in ioutil.ReadAll(resp.Body): %v", err.Error())
		panic(err)
	}

	// Convert byte array to string for printing.
	// bodyString := string(bodyBytes)
	// fmt.Println(bodyString)

	// Convert byte array to a pretty JSON string for printing so you know what kind of structs you want to create
	var understandableBodyBuffer bytes.Buffer
	err = json.Indent(&understandableBodyBuffer, bodyBytes, "", "  ")
	if err != nil {
		fmt.Printf("Error in json.Indent(): %v", err.Error())
		panic(err)
	}
	//fmt.Printf("%s\n", understandableBodyBuffer.Bytes())

	// Convert byte array to an effective struct using the Unmarshall function.
	// Unmarshal parses the JSON-encoded data in the first argument bodyBytes and stores the result in the value pointed to by the second argument v. If v is nil or not
	// a pointer, Unmarshal returns an InvalidUnmarshalError.  To unmarshal JSON into a struct, Unmarshal matches incoming object keys to the keys used by either the
	// struct field name or its tag, preferring an exact match but also accepting a case-insensitive match. By default, object keys which don't have a corresponding
	// struct field are ignored
	// By knowing the structure of the JSON response that the above API endpoint gives us we can crreate a series of structs to map our data to.

	// Within our JSON response, we have a couple of key-value pairs,
	// the first key "descriptions", is an array of languages used by the API
	// 	"descriptions": [
	// 		{
	//  		"description": "Pokédex régional de Kanto dans Rouge/Bleu/Jaune",
	// 		  	"language": {
	// 		    "name": "fr",
	// 		    "url": "https://pokeapi.co/api/v2/language/5/"
	//		   }
	//		 },
	//
	// the next set of keys: "id", "is_main_series", "name" and "names" give us more information about the region where the Pokemon reside in:  kanto
	// 		 "id": 2,
	//  	 "is_main_series": true,
	// 		  "name": "kanto",
	// 		  "names": [
	// 		    {
	// 		        "language": {
	// 		 	       "name": "fr",
	// 			        "url": "https://pokeapi.co/api/v2/language/5/"
	//  		     },
	// 		 	     "name": "Kanto"
	// 		    },
	//
	// the next key "pokemon_entries" has an array of all the Pokemon .
	// 		"pokemon_entries": [
	// 			{
	//				"entry_number": 1,
	//		 		"pokemon_species": {
	// 					"name": "bulbasaur",
	//		 			"url": "https://pokeapi.co/api/v2/pokemon-species/1/"
	//		 		}
	//		 	},
	//		 	{
	// 			  	"entry_number": 2,
	// 			  	"pokemon_species": {
	// 					"name": "ivysaur",
	// 					"url": "https://pokeapi.co/api/v2/pokemon-species/2/"
	// 			  	}

	// If we wanted to create a struct for pokemon enttries

	type PokemonSpecies struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	}

	type PokemonEntries struct {
		EntryNumber int            `json:"entry_number"`
		Species     PokemonSpecies `json:"pokemon_species"`
	}

	type PokemonRegion struct {
		Name    string           `json:"name"`
		Pokemon []PokemonEntries `json:"pokemon_entries"`
	}

	var manyPokemon PokemonRegion
	err = json.Unmarshal(bodyBytes, &manyPokemon)
	if err != nil {
		fmt.Printf("Error: json.Unmarshal(body): %v", err.Error())
		panic(err)
	}
	//fmt.Printf("Pokemon from body: %+v\n", pokemon)

	fmt.Printf("Looking at Pokemon from the region of: %v\n", manyPokemon.Name)
	fmt.Printf("Found %v pokemon there\n", len(manyPokemon.Pokemon))

	for _, pokemon := range manyPokemon.Pokemon {
		fmt.Printf("#%v: %v at %v\n", pokemon.EntryNumber, pokemon.Species.Name, pokemon.Species.URL)
	}
}
