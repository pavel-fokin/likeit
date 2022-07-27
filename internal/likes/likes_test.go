package likes

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestLikesCount(t *testing.T) {
	// setup
	db, mock, _ := sqlmock.New()
	rows := sqlmock.NewRows(
		[]string{"count"},
	).AddRow(1)
	mock.ExpectQuery("SELECT count FROM likes;").WillReturnRows(rows)

	likes := New(db)

	// test
	count, err := likes.Count()

	assert.Equal(t, count, 1)
	assert.NoError(t, err)
}

func TestLikesIncrement(t *testing.T) {
	// setup
	db, mock, _ := sqlmock.New()
	mock.ExpectExec("UPDATE likes SET count = count + 1;").
		WillReturnResult(sqlmock.NewResult(1, 1))

	likes := New(db)

	// test
	err := likes.Increment()

	assert.NoError(t, err)
}
