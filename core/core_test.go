package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	movie := Movie{
		title:  "spiderman",
		year:   2002,
		genres: []string{"comedy", "action", "superhero"}}

	result := FormatMovie(movie)

	assert.Equal(t, "spiderman(2002),comedy|action|superhero", result, "they should be equal")
}
