package model

type CanvasUser struct {
	Base
	FirstName     string    `bson:"first_name" json:"firstName"`
	LastName      string    `bson:"last_name" json:"lastName"`
	Email         string    `bson:"email" json:"email"`
	Avatar        string    `bson:"avatar" json:"avatar"`
	AccessToken   string    `bson:"access_token" json:"accessToken"`
	TokenExpireAt string    `bson:"token_expire_at" json:"tokenExpireAt"`
	TakenCourses  []*Course `bson:"taken_courses" json:"takenCourses"`
}
