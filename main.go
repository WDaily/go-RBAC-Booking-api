package main

import(
	"net/http"
	"github.com/gorilla/mux"
	"github.com/WDaily/go-RBAC-Booking-api/models"
	"github.com/WDaily/go-RBAC-Booking-api/database"
	"github.com/WDaily/go-RBAC-Booking-api/routes"
)

func main(){
	r := mux.NewRouter()

	r.Use(auth.AdminValidate)

	db := database.Inisialise()

	DatabaseAccess := models.Newdb(db)

	routes.AppRoutes(r,DatabaseAccess)

	http.ListenAndServe(":8080",r)

}