package models

import "gin_skeleton/SkeletonExample/form"

type User struct {
	*form.UserForm
	Id    int    `form:"id" json:"id" bson:"id"`
	Hobby string `form:"hobby" json:"hobby" bson:"hobby"`
	Sex   string `form:"sex" json:"sex" bson:"sex"`
}

func (u *User) TableName() string {
	return "user"
}

func NewDefaultUser() *User {
	formUser := new(form.UserForm)
	user := new(User)
	user.UserForm = formUser
	return user
}
