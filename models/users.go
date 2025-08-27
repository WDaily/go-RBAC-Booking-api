package models

import (
	"gorm.io/gorm"
	"html"
	"strings"
	"golang.org/x/crypto/bcrypt"
)


type User struct{
	gorm.model
	RoleID unit `gorm:"size:255;DEFAULT:3" json:"role_id"`
	Username string `gorm:"size:255;not null;unique" json:"username"`
	Password string `gorm:"size:255;notnull;" json:"-"`
	Role Role `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

func (database*DatabaseAccess) Create() (*User,error){
	var user User
	err := database.DB.Create(&user).Error
	if err != nil{
		return &user{},err
	}
	return user, nil
}

//before saving encrypt the password

func (user *User) BeforeSave (*gorm.DB) error{
	
	//unfinished
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.Default)
	if err != nil{
		return err
	}
	user.Password = string(passwordHash)
	// find out escape string
	user.Username = html.EscapeString(strings.TrimSpace(user.Username))
	return nil
}

func(database*DatabaseAccess) GetUserByName(name string)(User,error){
	var user User
	err:= database.DB.Where("username=?",name).Find(&user).Error
	if err != nil {
		return User{},err
	}
	return user, nil
}

func (user*User) ValidatePassword(password string) error{
	return bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
}


func (database*DatabaseAccess) GetUsers() []User{
	var users []User 
	database.DB.Find(&users)
	return users
}


//get user by id
func(database*DatabaseAccess) GetUser(Id int64) (*User,error){
	var user User 
	err:= database.DB.Where("id=?",Id).Find(&user).Error
	if err != nil{
		return nil,err
	}
	return &user,nil
}

func(database*DatabaseAccess) UpdateUser(user*User) (*User,error){
	err:= database.DB.Omit("password").Updates(&user).Error

	if err != nil{
		return nil,err
	}

	return &user,nil
}