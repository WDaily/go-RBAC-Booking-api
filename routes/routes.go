package routes

import(
	"github.com/gorilla/mux"

	"github.com/WDaily/go-RBAC-Booking-api/controls"
	"github.com/WDaily/go-RBAC-Booking-api/midddleware"
)

func AppRoutes(r*mux.router, db models.DatabaseAccess){

	appRoute := r.PathPrefix("/app/user").SubRouter()
	appRoute.HandleFunc("/input",controls.Input).Methods("POST")
	appRoute.HandleFunc("/login",db.Login).Methods("POST")

	//admin routes (can get all users/user, add/edit rooms, get all bookingd)

	adminRoute := r.PathPrefix("/admin").SubRouter()
	//middleware func

	adminRoute.Use(middleware.adminMiddleware)
	//users
	adminRoute.HandleFunc("/users",db.GetUsers).Methods("GET")
	adminRoute.HandleFunc("/users/{ID}",db.GetUser).Methods("GET")
	adminRoute.HandleFunc("/users/role",db.UpdateUser).Methods("POST")

	//roles
	adminRoute.HandleFunc("/users/role",db.CreateRole).Methods("POST")
	adminRoute.HandleFunc("/users/roles",db.GetRoles).Methods("GET")
	adminRoute.HandleFunc("/users/roles/{ID}",db.UpdateRole).Methods("POST")

	//rooms
	adminRoute.HandleFunc("/room/add",db.AddRoom).Methods("POST")
	adminRoute.HandleFunc("/room/edit",db.UpdateRoom).Methods("PUT")

	//bookings
	adminRoute.HandleFunc("/room/bookings",db.GetBookings).Methods("GET")


	publicRoute := r.PathPrefix("/public").SubRouter()
	publicRoute.Use(middleware.validateMiddleware)
	publicRoute.HandleFunc("/rooms",db.GetRooms).Methods("GET")

	//notyet
	publicRoute.HandleFunc("/rooms/{ID}",db.GetRoom).Methods("GET")
}