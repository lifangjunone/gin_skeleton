package models

import "gin_skeleton/SkeletonExample/form"

type User struct {
	Id int `form:"id" json:"id"`
	*form.UserForm
	Hobby string `form:"hobby" json:"hobby"`
	Sex   string `form:"sex" json:"sex"`
}

func (u *User) TableName() string {
	return "user"
}

func NewDefault() *User {
	return &User{}
}
