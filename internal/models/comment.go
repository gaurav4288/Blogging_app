package models

type Comment struct {
	CommentId string `bson:"comment_id" json:"comment_id"`
	PostId    string `bson:"post_id" json:"post_id"`
	Content   string `bson:"content" json:"content"`
	Author    string `bson:"author" json:"author"`
}
