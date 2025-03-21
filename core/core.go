package core

import (
	"fmt"
)

func GetMoviesFromAPI(movies []Movie) []Movie {

	return movies
}

type Movie struct {
	Id     string
	Title  string
	Year   int
	Genres []string
}

func FormatMovie(movie *Movie) string {
	if len(movie.Genres) < 1 {
		panic("movie has to be at least one genre")
	}

	var genres string = GetGenres(movie.Genres)

	return fmt.Sprintf("%s,%s(%d),%s", []any{movie.Id, movie.Title, movie.Year, genres}...)
}

func GetGenres(genres []string) string {

	var formatted = ""
	for index, genre := range genres {
		if index == len(genres)-1 {
			formatted = formatted + genre
		} else {
			formatted = formatted + genre + "|"
		}
	}

	return formatted

}
