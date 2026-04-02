package database

import (
	"html"
	"strings"
	"time"

	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
)


type User struct{
	ID  		uint 		`gorm:"primaryKey" json:"Id"`
	RoleName	string 		`gorm:"size:255;default:user;column:roleName" json:"roleName"`
	UserName 	string 		`gorm:"size:255;not null;unique;column:username" json:"username"`
	Password 	string 		`gorm:"size:255;notnull;" json:"-"`
	CreatedAt	time.Time 	`gorm:"column:createdAt" json:"-"`
	UpdatedAt	time.Time 	`gorm:"column:updatedAt" json:"-"`
	Role 		Role 		`gorm:"constraint:OnUpdate:CASCADE;" json:"-"`
}

func (database DatabaseAccess) Create(user*User) error{
	if err := database.DB.Create(user).Error; err != nil{
		return err
	}
	return nil
}

func (user *User)BeforeSave(*gorm.DB) error{
	
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost)
	if err != nil{
		return err
	}
	user.Password = string(passwordHash)

	user.UserName = html.EscapeString(strings.TrimSpace(user.UserName))
	return nil
}

func(database DatabaseAccess) GetUserByName(name string)(User,error){
	var user User
	if err := database.DB.Where("username=?",name).Find(&user).Error; err != nil {
		return User{},err
	}
	return user, nil
}

func ValidatePassword(password string, user User) error{
	return bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(password))
}

func (database DatabaseAccess) GetUsers() ([]User, error){
	var users []User 
	if err := database.DB.Find(&users).Error; err != nil{
		return []User{}, err
	}
	return users,nil
}

func(database DatabaseAccess) GetUser(id uint) (User,error){
	var user User 
	if err := database.DB.Where("id=?",id).Find(&user).Error; err != nil{
		return User{},err
	}
	return user,nil
}