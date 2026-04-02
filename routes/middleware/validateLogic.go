package routes

import(
	"time"
	"net/http"
	"errors"

	"github.com/WDaily/go-RBAC-Booking-api/database"
	
	"github.com/golang-jwt/jwt/v5"
)


var privateNames = []byte("privateNames")


func validateToken(r * http.Request) error{
	token, err := tokenFromUserRequest(r)

	if err != nil{
		return err
	}

	_, ok := token.Claims.(jwt.MapClaims)

	if ok && token.Valid {
		return nil
	}

	return errors.New("error: check user details")
}

func AdminValidate(r * http.Request) error{

	token, err := tokenFromUserRequest(r)

	if err != nil {
		return err
	}

	claims, ok := token.Claims.(jwt.MapClaims)


	role := claims["role"].(string)

	if ok && token.Valid && role == "admin"{
		return nil
	}
	return errors.New("Error: Must be an admin")
}

func tokenFromUserRequest(r* http.Request) (*jwt.Token, error){
	cookie , err := r.Cookie("token")

	if err != nil{
		return nil, err
	}

	tokenString := cookie.Value 

	token, err := jwt.Parse( tokenString, func( token* jwt.Token) (interface{}, error){
		return privateNames, nil
	})

	return token, err
}

func CurrentUser(r * http.Request) (uint,error){
	token,_ := tokenFromUserRequest(r)

	claims,_ := token.Claims.(jwt.MapClaims)

	user:= uint(claims["user"].(float64))
	return user, nil
}

func GenerateToken(user database.User) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.ID,
		"role": user.RoleName,
		"iat" : time.Now().Unix(),
		"eat": time.Now().Add(time.Minute*time.Duration(1)),
	})
	return token.SignedString(privateNames)
}

func SetInformation(w http.ResponseWriter, token string){
	http.SetCookie(w, &http.Cookie{
		Name:"token",
		Value: token,
		Expires:time.Now().Add(time.Minute*time.Duration(1)),
	}) 
}