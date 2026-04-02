package controls

import(
	"net/http"
	"strconv"

	"github.com/WDaily/go-RBAC-Booking-api/database"
	"github.com/WDaily/go-RBAC-Booking-api/utils"

	"github.com/gorilla/mux"
)

func (c Controls) AddRoom(w http.ResponseWriter, r*http.Request){
	var room database.Room 

	if err := utils.ChangeFormat(r,room); err != nil{
		Response(w,err)
	}

	if err := c.db.CreateRoom(&room); err != nil{
		Response(w,err)
	}

	Output(w, room)
}

func (c Controls)UpdateRoom(w http.ResponseWriter, r*http.Request){
	var updateRoom database.Room

	param := mux.Vars(r)
	id,err := strconv.ParseUint(param["ID"],0,0)
	if err != nil{
		Response(w,err)
	}
	//check whether its possible if not present not to update
	//err := c.db.GetRoom(id)
	//if err != nil {
	//	fmt.Fprintf(w,"Error while retrieving room: %s",err)
	//}

	if err := utils.ChangeFormat(r,&updateRoom); err != nil{
		Response(w,err)
	}

	updateRoom.ID = uint(id)

	if err := c.db.UpdateRoom(&updateRoom); err != nil {
		Response(w,err)
	}

	Output(w, updateRoom)
}

func (c Controls) GetRooms(w http.ResponseWriter, r*http.Request){
	rooms, err:= c.db.GetRooms()

	if err != nil{
		Response(w,err)
	}

	Output(w, rooms)	
}

func (c Controls) GetRoom(w http.ResponseWriter, r*http.Request){
	param := mux.Vars(r)
	id,err := strconv.ParseUint(param["ID"],0,0)
	if err != nil{
		Response(w,err)
	}
	room, err:= c.db.GetRoom(uint(id))

	if err != nil {
		Response(w,err)
	}

	Output(w, room)
}