package main

import (
	//For Printing to the terminal

	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
	"strconv"
)

/*
 * Complete the function below.
 * Base url: https://jsonmock.hackerrank.com/api/movies/search/?Title=
 */

type movieStruct struct {
	Title  string `json:"Title"`
	Year   int    `json:"Year"`
	ImdbID string `json:"imdbID"`
}
type movieReponse struct {
	Page       string        `json:"page"`
	PerPage    int           `json:"per_page"`
	Total      int           `json:"total"`
	TotalPages int           `json:"total_pages"`
	Data       []movieStruct `json:"data"`
}

// func main() {
// 	var substr string

// 	substr = "spiderman"
// 	res := getMovieTitles(substr)
// 	fmt.Printf("Sorted ")
// 	prettyPrintMovieArray(res)

// }

/*
 * Complete the function below.
 * Base url: https://jsonmock.hackerrank.com/api/movies/search/?Title=
 */
func getMovieTitles(substr string) []string {

	pageString := "1"

	var mr movieReponse

	baseURL := "https://jsonmock.hackerrank.com/api/movies/search/?Title="
	//fmt.Printf("Base URL: %v\n", baseURL)

	apiEndPoint := baseURL + substr + "&page=" + pageString
	//fmt.Printf("API Endpoint: %v\n", apiEndPoint)

	mr.getMovieInfo(apiEndPoint)

	//Array of Movie Titles, of the size of the titles.
	mt := make([]string, mr.Total)

	totalMovieCounter := 0

	for page := 1; page <= mr.TotalPages; page++ {

		//fmt.Printf("Going through movies on page %v\n", page)
		if page > 1 {
			apiEndPoint := baseURL + substr + "&page=" + strconv.Itoa(page)
			//fmt.Printf("API Endpoint: %v\n", apiEndPoint)
			mr.getMovieInfo(apiEndPoint)
		}

		for _, movies := range mr.Data {
			mt[totalMovieCounter] = movies.Title
			totalMovieCounter++
		}
		//fmt.Println("Movie Titles in Progress")
		//prettyPrintMovieArray(mt)

	}

	//Sort the movie Titles.
	sort.Strings(mt)
	//prettyPrintMovieArray(mt)

	return mt

}

func (mr *movieReponse) getMovieInfo(apiEndPoint string) {

	resp, err := http.Get(apiEndPoint)
	if err != nil {
		fmt.Printf("http.Get() for API Endpoint: %v returned Error: %v", apiEndPoint, err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("ioutil.ReadAll() for resp.Body: %v", err)
		panic(err)
	}
	//fmt.Printf("Response Body: %v\n\n", string(body))

	err = json.Unmarshal(body, &mr)
	if err != nil {
		fmt.Printf("ioutil.ReadAll() for resp.Body: %v", err)
		panic(err)
	}
	//fmt.Printf("JSON in struct: %+v\n\n", mr)
}

func prettyPrintMovieArray(mt []string) {
	fmt.Printf("Movie Titles:\n")
	for i, m := range mt {
		fmt.Printf("%v: %v\n", i, m)
	}
}
