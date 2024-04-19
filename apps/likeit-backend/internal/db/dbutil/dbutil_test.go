package dbutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomID(t *testing.T) {
	id, err := RandomID()
	assert.NoError(t, err)
	assert.NotEmpty(t, id)
	assert.Equal(t, maxLength, len(id))
}
