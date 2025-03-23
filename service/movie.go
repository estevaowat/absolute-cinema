package service

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/estevaowat/absolute-cinema/database"
)

type Movie struct {
	Id     string `db:"id"`
	Title  string `db:"title"`
	Year   string `db:"year"`
	Genres string `db:"genres"`
}

func SaveInDatabaseSequentially(fileName string) {
	log.Println("starting save in database sequentially")
	start := time.Now()

	//get file in "tmp folder"
	homeFolder, err := os.UserHomeDir()
	if err != nil {
		log.Fatal("could not get home dir", err)
	}

	log.Println("getting path")
	path := fmt.Sprintf("%s/%s/%s", homeFolder, "/Desktop", fileName)
	file, err := os.Open(path)

	defer file.Close()

	reader := csv.NewReader(file)

	var records []Movie

	for {
		record, err := reader.Read()

		if err == io.EOF {
			if len(records) > 0 {
				_, error := database.DbSqlX.NamedExec("insert into movies (id, title, year, genres) values (:id, :title, :year, :genres);", records)
				if error != nil {
					log.Fatal("error saving last movies", error)
				}
			}

			log.Println("finish reading file")
			break
		}

		if err != nil {
			log.Println("error reading file ", path, " error ", err)
			break
		}

		id := record[0]
		title := strings.Split(record[1], "(")[0]
		year := strings.Split(strings.Split(record[1], "(")[1], ")")[0]
		genres := record[2]
		movie := Movie{Id: id, Title: title, Year: year, Genres: genres}
		records = append(records, movie)

		if len(records) == 10000 {
			log.Println("saving movies in database")
			_, error := database.DbSqlX.NamedExec("insert into movies (id, title, year, genres) values (:id, :title, :year, :genres);", records)

			if error != nil {
				log.Fatalln("error saving in database", error)
			}

			records = nil
		}
	}

	elapsed := time.Since(start)
	log.Println("to save in database sequentially, in chunks took ", elapsed)
}

func saveInDatabaseUsingGoRoutines() {

}
