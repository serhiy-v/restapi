package db

import "github.com/serhiy-v/restapi/models"

func ShowCol(id int) (models.Column,error){
	var col models.Column
	query := "SELECT id, name, project_id, position FROM columns WHERE id = $1"
	err := db.QueryRow(query,id).Scan(&col.ID,&col.Name,&col.ProjectID,&col.Position)
	if err != nil {
		panic(err)
	}
	return col, err
}

func CreateCol(col models.Column) error {
	query := "INSERT INTO columns(name,project_id,position ) VALUES ($1,$2,$3)"
	_, err := db.Exec(query, col.Name,col.ProjectID,col.Position)
	if err != nil {
		panic(err)
	}
	return err
}

func UpdateCol(col models.Column, id int) error {
	query := "UPDATE columns SET name = $1, project_id = $2, position = $3 WHERE id = $4"
	_, err := db.Exec(query, col.Name,col.ProjectID,col.Position,id)
	if err != nil {
		panic(err)
	}
	return err
}

func DeleteCol(id int) error{
	query := "DELETE FROM columns WHERE id = $1"
	_,err := db.Exec(query,id)
	if err != nil {
		panic(err)
	}
	return err

}
