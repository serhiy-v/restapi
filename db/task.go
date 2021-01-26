package db

import "github.com/serhiy-v/restapi/models"

func ShowTask(id int) (models.Task,error){
	query := "SELECT id,name,description,column_id, position FROM tasks WHERE id = $1"
	var task models.Task
	err := db.QueryRow(query,id).Scan(&task.ID,&task.Name,&task.Description,&task.ColumnId,&task.Position)
	if err != nil {
		panic(err)
	}
	return task, err
}

func CreateTask(t models.Task) error {
	query := "INSERT INTO tasks(name,description,column_id,position ) VALUES ($1,$2,$3,$4)"
	_, err := db.Exec(query, t.Name,t.Description,t.ColumnId,t.Position)
	if err != nil {
		panic(err)
	}
	return err
}

func UpdateTask(t models.Task,id int) error{
	query := "UPDATE tasks SET name = $1, description = $2,column_id = $3,position = $4 WHERE id = $5"
	_, err := db.Exec(query, t.Name,t.Description,t.ColumnId,t.Position,id)
	if err != nil {
		panic(err)
	}
	return err
}

func DeleteTask(id int) error {
	query := "DELETE FROM tasks WHERE id = $1"
	_,err := db.Exec(query,id)
	if err != nil {
		panic(err)
	}
	return err
}
