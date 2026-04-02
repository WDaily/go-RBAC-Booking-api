package routes

import(
	"github.com/gorilla/mux"
	
	"github.com/WDaily/go-RBAC-Booking-api/controls"
	middleware "github.com/WDaily/go-RBAC-Booking-api/routes/middleware"
)

func CreateRouter() *mux.Router{
	r := mux.NewRouter()

	return r
}

func AppRoutes(r*mux.Router, db controls.Controls){

	appRoute := r.PathPrefix("/app/user").Subrouter()
	appRoute.HandleFunc("/input", db.Input).Methods("POST")
	appRoute.HandleFunc("/login", db.Login).Methods("POST")


	adminRoute := r.PathPrefix("/admin").Subrouter()
	adminRoute.Use(middleware.AdminMiddleware)
	//users
	adminRoute.HandleFunc("/users",db.GetUsers).Methods("GET")
	adminRoute.HandleFunc("/users/{ID}",db.GetUser).Methods("GET")

	//roles
	adminRoute.HandleFunc("/users/role",db.CreateRole).Methods("POST")
	adminRoute.HandleFunc("/users/roles/{ID}",db.GetRoles).Methods("GET")

	//rooms
	adminRoute.HandleFunc("/room/add",db.AddRoom).Methods("POST")
	adminRoute.HandleFunc("/room//{ID}/edit",db.UpdateRoom).Methods("PUT")

	//bookings
	adminRoute.HandleFunc("/room/bookings",db.GetBookings).Methods("GET")


	publicRoute := r.PathPrefix("/public").Subrouter()
	publicRoute.Use(middleware.ValidateMiddleware)
	publicRoute.HandleFunc("/rooms",db.GetRooms).Methods("GET")
	publicRoute.HandleFunc("/rooms/{ID}/Book", db.BookRoom).Methods("POST")
	
	publicRoute.HandleFunc("/rooms/{ID}",db.GetRoom).Methods("GET")
}