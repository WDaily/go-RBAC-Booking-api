package database



import (
	"os"
	
	"log"
	"gorm.io/gorm"
)

func Initialise() *gorm.DB{
	var err error

	host:= os.Getenv("HOST")
	username := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("NAME")
	port := os.Getenv("PORT")
	//incomplete

	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",username,password,host,port,dbName)

	database, err:= gorm.open(mysql.Open(conn),&gorm.Config{})

	if err !=nil {
		log.Fatalf("Failed to open Database: %v" ,err)
		return nil
	}
	log.Println("Database opened")
	
	return database
}

