package form

import (
	"fmt"
	"github.com/go-playground/validator/v10"
)

type UserForm struct {
	Username string `form:"username" json:"username" binding:"required" bson:"username"`
	Password string `form:"password" json:"password" binding:"required" bson:"password"`
}

func (u *UserForm) Validate() string {
	validate := validator.New()
	err := validate.Struct(u)
	fmt.Println(err)
	if err != nil {
		return err.Error()
	}
	return ""
}
