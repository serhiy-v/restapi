package models

type Project struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
}

type Column struct {
	ID int `json:"id"`
	Name string `json:"name"`
	ProjectID int `json:"project_id"`
	Position int `json:"position"`
}

type Task struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	ColumnId int `json:"column_id"`
	Position int `json:"position"`
}

type Comment struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	TaskId int `json:"task_id"`
	Position int `json:"position"`
}
