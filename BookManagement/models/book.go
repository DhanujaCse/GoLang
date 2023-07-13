package models

type Book struct {
	Id int64 `json:"id" gorm:"primary_key"`

	Author string `json:"author"`
	Title  string `json:"title"`
}
