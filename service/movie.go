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

	for {
		record, err := reader.Read()

		if err == io.EOF {
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
		log.Println("printing id = ", id, ",title = ", title, ",year = ", year, ",genres = ", genres)

		_, error := database.DbSqlX.NamedExec("insert into movies (id, title, year, genres) values (:id, :title, :year, :genres);",
			map[string]any{
				"id":     id,
				"title":  title,
				"year":   year,
				"genres": genres,
			})

		if error != nil {
			log.Println("error saving movie in database", error)
		}

	}

	elapsed := time.Since(start)
	log.Println("to save in database sequentially took ", elapsed)
}

func saveInDatabaseUsingGoRoutines() {

}
