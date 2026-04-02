package controls

import(
	"strconv"
	"net/http"

	"github.com/WDaily/go-RBAC-Booking-api/models"
	"github.com/WDaily/go-RBAC-Booking-api/utils"
	middleware "github.com/WDaily/go-RBAC-Booking-api/routes/middleware"
	"github.com/WDaily/go-RBAC-Booking-api/database"

	"github.com/gorilla/mux"
)

func (c Controls) Input(w http.ResponseWriter,r*http.Request){
	var input models.Input

	if err := utils.ChangeFormat(r, &input); err != nil{
		Response(w,err)
	}

	userNew := database.User{
		UserName:input.Username,
		Password: input.Password,
		RoleName: "user",
	}

	if err := c.db.Create(&userNew); err != nil {	
		Response(w,err)
	}

	Output(w, userNew)
}


func(c Controls) Login(w http.ResponseWriter, r*http.Request){
	var loginInput models.Login

	if err := utils.ChangeFormat(r,&loginInput); err != nil{
		Response(w,err)
	}


	user, err := c.db.GetUserByName(loginInput.Username)
	if err != nil {
		Response(w,err)
	}

	if err := database.ValidatePassword(loginInput.Password, user); err != nil{	
		Response(w,err)
	}

	generated, err:= middleware.GenerateToken(user)
	if err != nil{	
		Response(w,err)
	}

	middleware.SetInformation(w,generated)

	w.WriteHeader(http.StatusOK)

}


func (c Controls) GetUsers(w http.ResponseWriter, r*http.Request){
	users, err := c.db.GetUsers()

	if err != nil{
		Response(w,err)
	}

	Output(w, users)
}

func (c Controls) GetUser(w http.ResponseWriter, r*http.Request){
	param := mux.Vars(r)
	id,err := strconv.ParseUint(param["ID"],0,0)
	if err != nil{
		Response(w,err)
	}

	user, err := c.db.GetUser(uint(id))
	if err != nil{
		Response(w,err)
	}

	Output(w, user)
}