package likes

import (
	"testing"
)


func TestLikesCount(t *testing.T) {
	// setup
	likes := New()

	// test
	likes.Count()
}

func TestLikesIncrement(t *testing.T) {
	// setup
	likes := New()

	// test
	likes.Increment()
}
