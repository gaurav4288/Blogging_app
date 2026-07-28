package models

type Like struct {
	LikeId string `bson:"like_id" json:"like_id"`
	PostId string `bson:"post_id" json:"post_id"`
	UserId string `bson:"user_id" json:"user_id"`
}
