package models

import (
	"gorm.io/gorm"
)



type Role struct{
	gorm.model
	name string `gorm:"size 50;not null;unique" json:"name"`
	UserID uint `json:"-"`
}

type DatabaseAccess struct{
	DB*gorm.DB
}

func Newdb (db*gorm.DB) *DatabaseAccess{
	return &DatabaseAccess{db}
}

func (database*DatabaseAccess)/*(role*Role)*/ CreateRole() (*Role,error) {
	var role Role
	err:= database.DB.Create(&role).Error
	if err != nil{
		return nil,err
	}
	return role,nil
}

func (database*DatabaseAccess) GetRoles(id int64) (*Role,error){
	var roles Role
	err := database.DB.Where("ID = ?",Id).Find(&roles).Error

	if err != nil{
		return nil,err
	}
	return &roles,nil
}

func (database*DatabaseAccess) UpdateRole(role*Role) (*Role,error) {
	err := database.DB.Updates(&role).Error

	if err != nil{
		return nil,err
	}

	return &role,nil

}


func (database*DatabaseAccess) GetRole(Id int) error{
	var role Role

	err := database.DB.Where("ID = ?",Id).Find(&role).Error

	if err != nil {
		return err
	} 

	return nil
}