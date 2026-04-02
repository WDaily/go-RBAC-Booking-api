package database

import(
	"time"
)

type Room struct{
	ID  		uint 		`gorm:"primaryKey" json:"Id"`
	Number		uint 		`json:"room"`
	Floor		uint 		`json:"floor"`
	CreatedAt	time.Time 	`gorm:"column:createdAt" json:"-"`
	UpdatedAt	time.Time 	`gorm:"column:updatedAt" json:"-"`
	Bookings	Bookings 	`gorm:"constraint:OnUpdate:CASCADE;" json:"-"`
}


func (database DatabaseAccess) CreateRoom(room *Room) error{
	if err := database.DB.Create(room).Error; err != nil{
		return err
	}

	return nil
}

func (database DatabaseAccess) UpdateRoom(room*Room) error {
	if err:=database.DB.Save(room).Error; err != nil{
		return err
	}

	return nil

}


func (database DatabaseAccess) GetRoom(id uint) (Room,error){
	var room Room
	if err := database.DB.Where("ID = ?",id).Find(&room).Error; err != nil {
		return Room{}, err
	} 

	return room, nil
}

func(database DatabaseAccess) GetRooms() ([]Room,error){
	var rooms []Room
	if err := database.DB.Find(&rooms).Error; err != nil{
		return []Room{},err
	}
	return rooms, nil
}