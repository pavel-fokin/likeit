package likes

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestLikesCount(t *testing.T) {
	// Setup.
	db, mock, _ := sqlmock.New()
	rows := sqlmock.NewRows(
		[]string{"count"},
	).AddRow(1)
	mock.ExpectQuery("SELECT count FROM likes;").WillReturnRows(rows)

	likes := New(db)

	// Test.
	count, err := likes.Count(context.Background())

	assert.Equal(t, count, 1)
	assert.NoError(t, err)
}

func TestLikesIncrement(t *testing.T) {
	// Setup.
	db, mock, _ := sqlmock.New(
		sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual),
	)
	mock.ExpectExec("UPDATE likes SET count = count + 1;").
		WillReturnResult(sqlmock.NewResult(1, 1))

	likes := New(db)

	// Test.
	err := likes.Increment(context.Background())

	assert.NoError(t, err)
}
