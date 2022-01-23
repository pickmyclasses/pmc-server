package model

type CanvasUser struct {
	Base
	FirstName     string    `bson:"first_name"`
	LastName      string    `bson:"last_name"`
	Email         string    `bson:"email"`
	Avatar        string    `bson:"avatar"`
	AccessToken   string    `bson:"access_token"`
	TokenExpireAt string    `bson:"token_expire_at"`
	TakenCourses  []*Course `bson:"taken_courses"`
}
