package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type Status struct {
	Status string `json:"status"`
}

type Movie struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
	// Genres []string `json:"genres"`
}

func GetMovies(length int) {
	url := "http://localhost:8080/movie?length=" + strconv.Itoa(length)
	fmt.Println(url)
	response, error := http.Get(url)

	if error != nil {
		fmt.Println("error to get movies")
	}

	defer response.Body.Close()

	log.Println("decoding response body")

	//TODO: change ioutil implementation
	var movies []Movie

	err := json.NewDecoder(response.Body).Decode(&movies)

	if err != nil {
		log.Fatal("printing error", err)
	}

	fmt.Println("printing status")

	for _, movie := range movies {
		fmt.Printf("printing movie %v \n", movie)
	}

}
