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

func (s *Sql) Create(title string) (int, error) {
	res, err := s.db.Exec("insert into todos (title) values (?)", title)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *Sql) Update(id int, is_done bool) (int, error) {
	_, err := s.db.Exec("update todos set is_done = ? where id = ?", is_done, id)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *Sql) Delete(id int) (int, error) {
	_, err := s.db.Exec("delete from todos where id = ?", id)
	if err != nil {
		return 0, err
	}
	return int(id), nil
}
