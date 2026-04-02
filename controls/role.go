package controls

import (
	"net/http"
	"strconv"
	"github.com/WDaily/go-RBAC-Booking-api/database"
	"github.com/WDaily/go-RBAC-Booking-api/utils"

	"github.com/gorilla/mux"
)


type Controls struct{
	db database.DatabaseAccess
}

func NewControls(db database.DatabaseAccess) Controls{
	return Controls{db}
}

func (c Controls) CreateRole(w http.ResponseWriter, r*http.Request){
	var role database.Role 

	if err := utils.ChangeFormat(r, &role); err != nil{
		Response(w,err)
	}

	if err := c.db.CreateRole(&role); err != nil {
		Response(w,err)
	}

	Output(w, role)
}

func (c Controls)GetRoles(w http.ResponseWriter,r*http.Request){
	param := mux.Vars(r)
	id,err := strconv.ParseUint(param["ID"],0,0)
	if err != nil {
		Response(w,err)
	}
	roles,err := c.db.GetRoles(uint(id))

	if err != nil {
		Response(w,err)
	}

	Output(w, roles)
}