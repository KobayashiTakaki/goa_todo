package todoapi

import (
	"database/sql"
	todo "todo/gen/todo"

	_ "github.com/mattn/go-sqlite3"
)

type Sql struct {
	db *sql.DB
}

func NewSqlDB(db *sql.DB) *Sql {
	return &Sql{db}
}

func (s *Sql) Find(id int) (*todo.Todo, error) {
	var t todo.Todo
	err := s.db.QueryRow("select id, title, is_done from todos where id = ?", id).Scan(&t.ID, &t.Title, &t.IsDone)
	if err != nil {
		return nil, err
	}
	return &t, nil
}
