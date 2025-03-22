package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "myuser"
	password = "mypassword"
	dbname   = "absolute-cinema"
)

func GetDatabaseUsingDefaultLibrary() *sql.DB {

	// TODO: add to environment variables
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		log.Fatalln("error connecting to database", connectionString, err)
	}

	if error := db.Ping(); error != nil {
		log.Fatalln("error pinging database", connectionString, error)
	}

	log.Println("successfully connected to database")

	return db
}

func GetDatabaseUsingSqlX() *sqlx.DB {
	//TODO: Add database values to environment variables
	connectionString := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", connectionString)

	if err != nil {
		log.Fatalln("error connecting to database", connectionString, err)
	}

	log.Println("successfully connected to database")

	return db

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
