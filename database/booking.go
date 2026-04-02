package database

import(
	"time"
)

type Bookings struct{
	ID  		uint 		`gorm:"primaryKey" json:"-"`
	Status 		string 		`json:"status"`
	CreatedAt	time.Time 	`gorm:"column:createdAt" json:"-"`
	UpdatedAt	time.Time 	`gorm:"column:updatedAt" json:"-"`
	UserID 		uint 		`json:"-"`
	RoomID 		uint 		`json:"-"`
}

func (database DatabaseAccess) GetBookings() ([]Bookings,error){
	var bookings []Bookings
	err := database.DB.Find(&bookings).Error
	if err != nil{
		return []Bookings{},err
	}
	return bookings,nil
}

func (database DatabaseAccess) BookRoom(book *Bookings) error {
	if err := database.DB.Create(book).Error; err != nil{
		return err
	} 
	return nil
}