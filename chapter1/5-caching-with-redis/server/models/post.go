package models

import (
	"github.com/jinzhu/gorm"
)

type Post struct {
	gorm.Model
	Name string `gorm:"type:varchar(255);" json:"name"`
	Age  int    `gorm:"type:smallint" json:"age"`
}

func GetPost() Post {
	var post Post
	return post
}

func GetPosts() []Post {
	var posts []Post
	return posts
}
