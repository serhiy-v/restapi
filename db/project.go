package db

import "github.com/serhiy-v/restapi/models"

func Projects() ([]*models.Project, error){
	rows, err := db.Query("SELECT * FROM projects")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	projects := make([]*models.Project, 0)
	for rows.Next() {
		prj := new(models.Project)
		err := rows.Scan(&prj.ID, &prj.Name, &prj.Description)
		if err != nil {
			return nil, err
		}
		projects = append(projects, prj)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return projects, nil
}

func Project(id int) (models.Project,error) {
	query := "SELECT id,name,description FROM projects WHERE id = $1"
	var prj models.Project
	err := db.QueryRow(query,id).Scan(&prj.ID,&prj.Name,&prj.Description)
	if err != nil {
		panic(err)
	}
	return prj, err

}

func Insert(prj models.Project) error {
	query := "INSERT INTO projects(name,description) VALUES ($1,$2)"
	_, err := db.Exec(query, prj.Name,prj.Description)
	if err != nil {
		panic(err)
	}
	return err
}

func Update(prj models.Project,id int) error{
	query := "UPDATE projects SET name = $1, description = $2 WHERE id = $3"
	_, err := db.Exec(query, prj.Name,prj.Description,id)
	if err != nil {
		panic(err)
	}
	return err
}

func Delete(id int) error {
	query := "DELETE FROM projects WHERE id = $1"
	_,err := db.Exec(query,id)
	if err != nil {
		panic(err)
	}
	return err
}
