package core

import (
	"fmt"
)

func GetMoviesFromAPI(movies []Movie) []Movie {

	return movies
}

type Movie struct {
	title  string
	year   int
	genres []string
}

func FormatMovie(movie Movie) string {
	if len(movie.genres) < 1 {
		panic("movie has to be at least one genre")
	}

	var genres string = getGenres(movie.genres)

	return fmt.Sprintf("%s(%d),%s", []any{movie.title, movie.year, genres}...)
}

func getGenres(genres []string) string {

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
