package controls

import (
	"strconv"
	"net/http"

	"github.com/WDaily/go-RBAC-Booking-api/utils"
	user "github.com/WDaily/go-RBAC-Booking-api/routes/middleware"
	"github.com/WDaily/go-RBAC-Booking-api/database"

	"github.com/gorilla/mux"
)

func (c Controls) GetBookings(w http.ResponseWriter, r*http.Request){
	bookings,err:= c.db.GetBookings()

	if err != nil{
		Response(w,err)
	}

	Output(w, bookings)
}

func (c Controls) BookRoom (w http.ResponseWriter, r*http.Request){
	var booking database.Bookings

	param := mux.Vars(r)
	roomId,err := strconv.ParseUint(param["ID"],0,0)
	if err != nil{
		Response(w,err)
	}

	if err := utils.ChangeFormat(r, &booking); err != nil{
		Response(w,err)
	}

	userId, err := user.CurrentUser(r)

	if err != nil {
		Response(w,err)
	}


	booking.UserID = userId
	booking.RoomID = uint(roomId)

	if err := c.db.BookRoom(&booking); err != nil{
		Response(w,err)
	}

	w.WriteHeader(http.StatusOK)
}