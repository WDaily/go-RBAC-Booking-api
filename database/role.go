package database

import (
	"time"
)

type Role struct{
	ID  		uint 		`gorm:"primaryKey" json:"Id"`
	name 		string 		`gorm:"size 50;not null;unique" json:"name"`
	CreatedAt	time.Time 	`gorm:"column:createdAt" json:"-"`
	UpdatedAt	time.Time 	`gorm:"column:updatedAt" json:"-"`
	UserID 		uint 		`json:"-"`
}

func (database DatabaseAccess) CreateRole(role*Role) error {
	if err:= database.DB.Create(role).Error; err != nil{
		return err
	}
	return nil
}

func (database DatabaseAccess) GetRoles(id uint) ([]Role,error){
	var roles []Role
	if err := database.DB.Where("ID = ?",id).Find(&roles).Error; err != nil{
		return []Role{},err
	}
	return roles,nil
}