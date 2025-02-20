package models

import (
	"example.com/todo/db"
)

type TodoList struct {
	Id          int64
	Task        string `binding:"required"`
	Description string `binding:"required"`
	Datetime    string `binding:"required"`
}

func (u *TodoList) Save() error {
	query := `INSERT INTO todoList(task,description,datetime)
              VALUES(?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Task, u.Description, u.Datetime)
	if err != nil {
		return err
	}
	task_id, err := result.LastInsertId()
	u.Id = task_id
	return err
}
