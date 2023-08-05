package model

import (
	db "advait/chatServer/model/database"

	"gorm.io/gorm"
)

//Database model
type User struct {
	gorm.Model
	ID string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

//Response models
type ChangePassword struct {
	Old_password string `json:"old_password" binding:"required"`
	New_password string `json:"new_password" binding:"required"`
}

type UpdateUser struct {
	Name string `json:"name" binding:"required"`
}

func CreateUser(user *User) error{
	err := db.Db.Create(user).Error

	if err != nil {
		return err
	}

	return nil
}
