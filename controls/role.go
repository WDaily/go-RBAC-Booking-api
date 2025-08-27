package controls

import (
	"fmt"
	"json/encoding"
	"net/http"
	"github.com/WDaily/go-RBAC-Booking-api/Database"
	"github.com/gorilla/mux"
)


type Controllers struct{
	db models.DatabaseAccess
}

func (c*Controllers) CreateRole(w http.ResponseWriter, r*http.Request){
	var role models.Role 

	utils.ChangeFmt(r,&role)

	roleNew, err := role.CreateRole()

	if err != nil {
		fmt.Fprintf(w,"Error while creating role: %s",err)
	}

	resp,_ := json.Marshal(roleNew)

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func (c*Controllers)GetRoles(w http.ResponseWriter,r*http.Request){
	roles,err := c.db.GetRoles()
	resp ,_ := json.Marshal(roles)

	if err != nil {
		fmt.Fprintf(w,"Error while retrieving roles: %s",err)
	}

	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

func (c*Controllers)UpdateRoles(w http.ResponseWriter,r*http.Request){
	var updateRole models.Role
	param := mux.Vars(r)
	Id,err := strconv.ParseInt(param["ID"],0,0)
	if err != nil{
		fmt.Fprintf(w,"Parsing error: %s",err)
	}

	err := c.db.GetRole(Id)
	if err != nil {
		fmt.Fprintf(w,"Error while retrieving role: %s",err)
	}

	utils.ChangeFmt(r,&updateRole)

	role,err := models.UpdateRole(&updateRole)
	if err != nil {
		fmt.Fprintf(w,"Error while updating role: %s",err)
	}

	resp,_ := json.Marshal(role)
	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}