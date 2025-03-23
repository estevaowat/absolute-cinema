package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

var DbSqlX *sqlx.DB
var Db *sql.DB

func InitDB() {
	log.Println("initializing database using standard library")
	dbConfig := GetDbConfig()
	connection := dbConfig.getConnectionString()

	var err error
	Db, err = sql.Open("postgres", connection)

	if err != nil {
		log.Fatalln("error connecting to database", connection, err)
	}

	if error := Db.Ping(); error != nil {
		log.Fatalln("error pinging database", connection, error)
	}

	log.Println("successfully connected to database")

}

func InitDBSqlX() {
	log.Println("initializing database using sqlx")
	dbConfig := GetDbConfig()

	var err error
	DbSqlX, err = sqlx.Connect("postgres", dbConfig.getConnectionString())

	if err != nil {
		log.Fatalln("error connecting to database", dbConfig.getConnectionString(), err)
	}

	log.Println("successfully connected to database")
}

func GetDatabaseUsingDefaultLibrary() *sql.DB {
	if Db != nil {
		return Db
	}

	InitDB()

	return Db
}

func GetDatabaseUsingSqlX() *sqlx.DB {
	if DbSqlX != nil {
		return DbSqlX
	}

	InitDBSqlX()

	return DbSqlX
}

type User struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
	Age  int    `db:"age"`
	// CreatedAt time.Time  `db:"created_at"`
	//	UpdatedAt *time.Time `db:"updated_at"`
}

func GetMoviesFromDatabase() {
	log.Println("getting database")
	db := GetDatabaseUsingDefaultLibrary()
	log.Println("closing database")
	defer db.Close()

	rows, err := db.Query("select name, age from users where age = $1", 18)

	if err != nil {
		log.Fatal("could not execute query", err)
	}

	defer rows.Close()

	for rows.Next() {
		var name string
		var age int

		if err := rows.Scan(&name, &age); err != nil {
			log.Fatal("error getting names from database", err)
		}

		log.Println("printing", "name=", name, "age=", age)
	}

	error := rows.Err()
	if err != nil {
		log.Fatal("error during iteration", error)
	}

}

func GetMoviesFromDatabaseUsingSqlX() {
	log.Println("getting database using sqlx")
	db := GetDatabaseUsingSqlX()

	users := []User{}

	if err := db.Select(&users, "select id, name, age from users where age = $1 and name like $2", 18, "%zoro%"); err != nil {
		log.Fatalln("error getting users using sqlx", err)
	}

	for _, user := range users {
		fmt.Println("printing user ", user)
	}

	log.Println("closing database")
	db.Close()

}
