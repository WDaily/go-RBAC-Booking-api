package models

import(
	"gorm.io/gorm"
)

type Book struct{
	gorm.model
	Status string `json:"status"`
	User uint `json:"-"`
	RoomID uint `json:"-"`
}

func (database*DatabaseAccess) GetBookings() ([]Book,error){
	var bookings [] Book
	err := database.DB.Find(&bookings).Error
	if err != nil{
		return nil,err
	}
	return bookings,nil
}