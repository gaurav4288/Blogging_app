package models

import "time"

type Post struct {
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
	Title     string    `bson:"title" json:"title"`
	PostId    string    `bson:"post_id" json:"post_id"`
	Content   string    `bson:"content" json:"content"`
	Author    string    `bson:"author" json:"author"`
}
