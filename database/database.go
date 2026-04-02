package database

import ( 
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseAccess struct{
	DB *gorm.DB
}

func Newdb(db* gorm.DB) DatabaseAccess{
	return DatabaseAccess{db}
}

func Setup(url string) (*gorm.DB, error){
	database, err:= gorm.Open(mysql.Open(url),&gorm.Config{})

	if err !=nil {
		return nil, err
	}
	return database, nil
}

func DataBaseAutoMigrate(db * gorm.DB){
	db.AutoMigrate(&Bookings{}, &Role{}, &Room{}, &User{})
}

func LoadData(db*gorm.DB){
	var room Room

	room.Number = 100
	room.Floor = 1
	db.Create(&room)

	userDetails := User{RoleName:"admin", UserName:"username", Password:"password"}

	db.Create(&userDetails)
}