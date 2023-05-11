package models

import "gin_skeleton/SkeletonExample/form"

type User struct {
	Id    int            `form:"id" json:"id" bson:"id"`
	Form  *form.UserForm `bson:"form"`
	Hobby string         `form:"hobby" json:"hobby" bson:"hobby"`
	Sex   string         `form:"sex" json:"sex" bson:"sex"`
}

func (u *User) TableName() string {
	return "user"
}

func NewDefault() *User {
	return &User{}
}
