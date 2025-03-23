package core

import (
	"fmt"
)

type Movie struct {
	Id     string   `json:"id"`
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Genres []string `json:"genres"`
}

func (movie Movie) FormatMovie() string {
	if len(movie.Genres) < 1 {
		panic("movie has to be at least one genre")
	}

	var genres string = movie.GetGenres("|")

	return fmt.Sprintf("%s,%s(%d),%s", []any{movie.Id, movie.Title, movie.Year, genres}...)
}

func (movie Movie) GetGenres(delimiter string) string {
	var formatted = ""

	for index, genre := range movie.Genres {
		if index == len(movie.Genres)-1 {
			formatted = formatted + genre
		} else {
			formatted = formatted + genre + delimiter
		}
	}

	return formatted

}
