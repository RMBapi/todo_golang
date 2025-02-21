package models

import (
	"example.com/todo/db"
)

type TodoList struct {
	Id          int64
	Task        string `binding:"required"`
	Description string `binding:"required"`
	Datetime    string `binding:"required"`
	Status      string `binding:"required"`
}

func (u *TodoList) Save() error {
	query := `INSERT INTO todoList(task,description,datetime,status)
              VALUES(?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(u.Task, u.Description, u.Datetime, u.Status)
	if err != nil {
		return err
	}
	task_id, err := result.LastInsertId()
	u.Id = task_id
	return err
}

func (u *TodoList) ViewTask() ([]TodoList, error) {

	rows, err := db.DB.Query("SELECT id, task, description, datetime ,status FROM todoList")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var tasks []TodoList

	for rows.Next() {
		var todo TodoList
		err := rows.Scan(&todo.Id, &todo.Task, &todo.Description, &todo.Datetime, &todo.Status)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, todo)
	}

	return tasks, nil
}

func (u *TodoList) Update() error {
	query := `
	UPDATE todoList
	SET task = ?, description = ?, datetime = ?,status = ?
	WHERE id = ?
	`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(u.Task, u.Description, u.Datetime, u.Status, u.Id)

	return err
}
