package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
	Name   string `json:"name"`
	Author string `json:"author"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
}
