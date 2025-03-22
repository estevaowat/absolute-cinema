package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/estevaowat/absolute-cinema/core"
)

func GetMoviesUsingGoRoutines(length int) {
	start := time.Now()

	url := "http://localhost:8080/movies?length=" + strconv.Itoa(length)
	response, error := http.Get(url)

	if error != nil {
		log.Fatal("error to get movies", error)
	}

	defer response.Body.Close()
	log.Println("decoding response body")
	var movies []Movie

	if err := json.NewDecoder(response.Body).Decode(&movies); err != nil {
		log.Fatal("error decoding response body", err)
	}

	homeFolder, err := os.UserHomeDir()

	if err != nil {
		log.Fatal("error getting user home dir", err)
	}

	pathFolder := homeFolder + "/Desktop"
	file := createCsv(pathFolder, strconv.Itoa(length))
	defer file.Close()
	writer := csv.NewWriter(file)

	moviesChan := make(chan []string, length)

	var wg sync.WaitGroup
	defer wg.Wait()

	var moviesToWrite [][]string = make([][]string, length)

	for _, movie := range movies {
		wg.Add(1)
		go parseMovie(moviesChan, &movie, &wg)
		formatted := <-moviesChan
		moviesToWrite = append(moviesToWrite, formatted)
	}

	writer.WriteAll(moviesToWrite)

	defer writer.Flush()

	elapsed := time.Since(start)
	log.Println("to generate using goroutines took ", elapsed)

}

func parseMovie(channel chan []string, movie *Movie, wg *sync.WaitGroup) {
	movieFormatted := make([]string, 3)
	movieFormatted[0] = movie.Id
	movieFormatted[1] = fmt.Sprintf("%s(%d)", movie.Title, movie.Year)
	movieFormatted[2] = core.GetGenres(movie.Genres)

	channel <- movieFormatted
	wg.Done()

}
