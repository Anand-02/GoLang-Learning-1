package models

type BlogItem struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	Author string `json:"author"`
}
