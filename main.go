package main

import (
	"log"

	"github.com/estevaowat/absolute-cinema/cmd"
	"github.com/estevaowat/absolute-cinema/database"
)

func main() {
	database.InitDB()
	database.InitDBSqlX()
	cmd.Execute()

	defer database.DbSqlX.Close()
	defer log.Println("closing database sqlx")

	defer database.Db.Close()
	defer log.Println("closing database standard library")
}
