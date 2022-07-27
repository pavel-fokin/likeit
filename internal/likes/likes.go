package likes

import (
	"database/sql"
)

type Likes struct {
	db *sql.DB
}

func New(db *sql.DB) *Likes {
	return &Likes{
		db: db,
	}
}

func (l *Likes) Count() (int, error) {
	var count int
	l.db.QueryRow("SELECT count FROM likes;").Scan(&count)
	return count, nil
}
func (l *Likes) Increment() error {
	l.db.Exec("UPDATE likes SET count = count + 1;")
	return nil
}
