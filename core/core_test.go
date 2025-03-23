package core

import (
	"testing"

	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/assert"
)

func TestSomething(t *testing.T) {
	id := uuid.NewV4().String()

	movie := Movie{
		Id:     id,
		Title:  "spiderman",
		Year:   2002,
		Genres: []string{"comedy", "action", "superhero"}}

	result := movie.FormatMovie()

	assert.Equal(t, id+",spiderman(2002),comedy|action|superhero", result, "they should be equal")
}
