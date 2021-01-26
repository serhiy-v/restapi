package db

import "github.com/serhiy-v/restapi/models"

func ShowComment(id int) (models.Comment, error) {
	query := "SELECT id,name,description,task_id, position FROM comments WHERE id = $1"
	var comm models.Comment
	err := db.QueryRow(query,id).Scan(&comm.ID,&comm.Name,&comm.Description,&comm.TaskId,&comm.Position)
	if err != nil {
		panic(err)
	}
	return comm, err
}
func CreateComment(comm models.Comment) error {
	query := "INSERT INTO comments(name,description,task_id,position ) VALUES ($1,$2,$3,$4)"
	_, err := db.Exec(query, comm.Name,comm.Description,comm.TaskId,comm.Position)
	if err != nil {
		panic(err)
	}
	return err
}

func UpdateComment(comm models.Comment,id int) error{
	query := "UPDATE comments SET name = $1, description = $2,task_id = $3,position = $4 WHERE id = $5"
	_, err := db.Exec(query, comm.Name,comm.Description,comm.TaskId,comm.Position,id)
	if err != nil {
		panic(err)
	}
	return err
}

func DeleteComment(id int) error {
	query := "DELETE FROM comments WHERE id = $1"
	_,err := db.Exec(query,id)
	if err != nil {
		panic(err)
	}
	return err
}
