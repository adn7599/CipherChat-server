package model

import (
	db "advait/chatServer/model/database"
)

type User struct {
	ID string `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func CreateUser(user *User) error{
	err := db.Db.Create(user).Error

	if err != nil {
		return err
	}

	return nil
}
