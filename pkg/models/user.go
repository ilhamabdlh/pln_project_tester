package models

import (
	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)

type Users struct {
	gorm.Model         // adds ID, created_at etc.
	Name       string `gorm:"type:varchar(100);unique_index"json:"name"`
	Password      string `gorm:"type:varchar(100)"json:"password"`
	Previlage string `json:"previlage,omitempty"`
}

type Login struct {
	Name string `form:"name" json:"name" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (user *Users) HashPassword(password string) error{
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil{
		return err
	}
	user.Password = string(bytes)
	return nil
}

func (user *Users) CheckPassword(providedPassword string) error {
	hash, _ := bcrypt.GenerateFromPassword([]byte(providedPassword), 14)
	err := bcrypt.CompareHashAndPassword(hash, []byte(user.Password))
	if err != nil{
		return err
	}
	return nil
}