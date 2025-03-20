package api

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/estevaowat/absolute-cinema/core"
	"github.com/satori/go.uuid"
)

type Status struct {
	Status string `json:"status"`
}

type Movie struct {
	Id     string   `json:"id"`
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Genres []string `json:"genres"`
}

func GetMovies(length int) {
	start := time.Now()

	url := "http://localhost:8080/movie?length=" + strconv.Itoa(length)
	log.Println(url)
	response, error := http.Get(url)

	if error != nil {
		log.Println("error to get movies")
	}

	defer response.Body.Close()

	log.Println("decoding response body")

	var movies []Movie

	err := json.NewDecoder(response.Body).Decode(&movies)

	if err != nil {
		log.Fatal("printing error", err)
	}

	homeFolder, e := os.UserHomeDir()

	if e != nil {
		log.Println("error getting user home dir")
	}
	pathFolder := homeFolder + "/Desktop"
	file := createCsv(pathFolder, strconv.Itoa(length))

	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	for _, movie := range movies {
		columns := 3
		formatted := make([]string, columns)

		formatted[0] = movie.Id
		formatted[1] = fmt.Sprintf("%s(%d)", movie.Title, movie.Year)
		formatted[2] = core.GetGenres(movie.Genres)

		err := writer.Write(formatted)

		if err != nil {
			log.Println("could not write movie ", err, "movie", movie)
		}
	}

	log.Println("finished writting movies into csv file")
	log.Println("go check dir=", pathFolder, "file.name=", file.Name())
	elapsed := time.Since(start)
	log.Println("to generate the file took", elapsed)
}

func createCsv(pathFolder string, prefix string) *os.File {
	log.Println("creating csv file")
	path := fmt.Sprintf("%s/%s-%s.csv", pathFolder, prefix, uuid.NewV4().String())
	file, err := os.Create(path)

	if err != nil {
		log.Println("error to create file", err)
	}

	log.Println("created csv file")

	return file
	//do i need to check if pathFolder exists?

	// create file inside the path folder using a random uuid
}

func save(movies []Movie) {

}
