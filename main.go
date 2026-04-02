package main

import(
	"fmt"
	"log"
	"net/http"
	"github.com/WDaily/go-RBAC-Booking-api/controls"
	"github.com/WDaily/go-RBAC-Booking-api/database"
	router "github.com/WDaily/go-RBAC-Booking-api/routes"
)

func main(){

	r := router.CreateRouter()

	url := "booking.db"

	db, err := database.Setup(url)

	if err != nil{
		fmt.Printf("Database start up error:%s\n",err)
	}

	database.LoadData(db)

	databaseAccess :=database.Newdb(db)

	database.DataBaseAutoMigrate(db)

	controlsAccess := controls.NewControls(databaseAccess)

	router.AppRoutes(r,controlsAccess)

	if err := http.ListenAndServe(":8080",r); err != nil{
		log.Printf("http error:%s\n", err)
	}

}