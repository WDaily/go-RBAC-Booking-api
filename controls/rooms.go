package controls

import(
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/WDaily/go-RBAC-Booking-api/database"
	"github.com/WDaily/go-RBAC-Booking-api/utils"
	"github.com/gorilla/mux"
)

func AddRoom(w http.ResponseWriter, r*http.Request){
	var room models.Room 

	utils.ChangeFmt(r,&input)

	room,err := room.CreateRoom()
	if err != nil{
		fmt.Fprintf(w,"Error while adding room: %s",err)
	}

	resp,_:= json.Marshal(room)
	w.Header().Set("content-type","applicaton/json")
	w.WriteHeader(http.StatusOK)
	w.write(resp)
}

func (c*Controllers)UpdateRoom(w http.ResponseWriter, r*http.Request){
	var updateRoom models.Room

	param := mux.Vars(r)
	Id,err := strconv.ParseInt(param["ID"],0,0)
	if err != nil{
		fmt.Fprintf(w,"Parsing error: %s",err)
	}

	err := c.db.GetRoom(Id)
	if err != nil {
		fmt.Fprintf(w,"Error while retrieving room: %s",err)
	}

	utils.ChangeFmt(r,&updateRoom)

	room,err := c.db.UpdateRoom(&updateRoom)
	if err != nil {
		fmt.Fprintf(w,"Error while updating room: %s",err)
	}

	resp,_ := json.Marshal(room)
	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (c*Controllers) GetRooms(w http.ResponseWriter, r*http.Request){
	rooms:= c.db.GetRooms()

	resp :+ json.Marshal(rooms)
	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	
}

func (c*Controllers) GetRoom(w http.ResponseWriter, r*http.Request){
	param := mux.Vars(r)
	Id,err := strconv.ParseInt(param["ID"],0,0)
	if err != nil{
		fmt.Fprintf(w,"Parsing error: %s",err)
	}
	room:= c.db.GetRoom(Id)
	resp := json.Marshal(room)
	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
	
}