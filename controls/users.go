package controls

import(
	"github.com/gorilla/mux"

	"github.com/WDaily/go-RBAC-Booking-api/application"
	"github.com/WDaily/go-RBAC-Booking-api/utils"
)

func Input(w http.ResponseWriter,r*http.Request){
	//if doesn't work fix here
	var input application.Input

	utils.ChangeFmt(r,&input)

	userNew := models.User{
		Username:input.Username,
		Password: input.Password,
		RoleID: 3,
	}
	//IF doesn't work fix here 
	user,err := userNew.Create()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

	}

	resp,_ := json.Marshal(user)

	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}


func(c*Controllers) Login(w http.ResponseWriter, r*http.Request){
	var loginInput application.Login

	utils.ChangeFmt(r, &loginInput)


	user, err := c.db.GetUserByName(loginInput.Username)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	//validate the entered password by comparing with the onein db
	if err := user.ValidatePassword(loginInput.Password); err != nil{
		w.WriteHeader(http.StatusBadRequest)
	}

	generated, err:= auth.Generate(user)
	if err != nil{
		fmt.Prinln(err)
	}

	//place generated in cookie

	w.WriteHeader(http.StatusOK)

}


func (c*Controllers) GetUsers(w http.ResponseWriter, r*http.Request){
	users := c.db.GetUsers()
	resp ,_ := json.Marshal(users)
	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

func (c*Controllers) GetUser(w http.ResponseWriter, r*http.Request){
	param := mux.Vars(r)
	Id,err := strconv.ParseInt(param["ID"],0,0)
	if err != nil{
		fmt.Prinln("str to Int conversion error")
	}

	user, err := c.db.GetUser(Id)
	if err != nil{
		fmt.Prinln()
	}

	resp,_ := json.Marshal(user)
	w.Header().Set("content-type","application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}