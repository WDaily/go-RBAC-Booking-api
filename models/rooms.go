package models

import(
	"gorm.io/gorm"
)

type Room struct{
	gorm.model 
	number	uint 	`json:"room"`
	Floor	uint 	`json:"floor"`
	Book Book `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}


func (database*DatabaseAccess) CreateRoom() (*Room,error){
	//database.DB.NewRecord(r) //check why its omitted
	var r Room
	err := database.DB.Create(&r).Error
	if err != nil{
		return nil,err
	}

	return r,nil
}

func (database*DatabaseAccess) UpdateRoom(room*Room) (*Room,error) {
	err:=database.DB.Updates(&room).Error

	if err != nil{
		return nil,err
	}

	return &room,nil

}


func (database*DatabaseAccess) GetRoom(Id int) error{
	var room Room

//check the appropriate return type
	err := database.DB.Where("ID = ?",Id).Find(&room).Error

	if err != nil {
		return err
	} 

	return nil
}

func(database*DatabaseAccess) GetRooms() []Room{
	var rooms []Room
	database.DB.Find(&rooms)
	return users
}