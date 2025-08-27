package controls

import (
	"fmt"
	"json/encoding"
	"net/http"
	"github.com/WDaily/go-RBAC-Booking-api/models"
)

func (c*Controllers) GetBookings(w http.ResponseWriter, r*http.Request){
	bookings,err:= c.db.GetBookings()

	if err != nil{
		fmt.Fprintf(w,"Error while retrieving bookings: %s",err)
	}

	resp,_ := json.Marshal(bookings)
	w.Header().Set("content-type","applicaton/json")
	w.WriteHeader(http.StatusOK)
	w.write(resp)

}