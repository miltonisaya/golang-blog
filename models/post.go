package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Title  string `json:"title"`
	Body   string `json:"body"`
	Likes  int    `json:"likes"`
	Draft  bool   `json:"draft"`
	Author string `json:"author"`
	UserID uint
}
