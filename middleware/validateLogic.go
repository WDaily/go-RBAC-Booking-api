package middleware

import(
	"net/http"
	"errors"
)

func Generate(user*models.User)




func validateToken(r * http.Request) error{
	token, err := token(r * http.Request)

	if err != nil{
		return err
	}

	_, ok := token.Claims(jwt.MapClaims)

	if ok && token.Valid {
		return nil
	}

	return errors.New("error")
}

func adminValidate(r * http.Request) error{

	token, err := token( r * http.Request)

	if err != nil {
		return err
	}

	claims, ok := token.Claims(jwt.MapClaims)

	role := uint(claims["role"].(float64))

	if ok && token.Valid && role == 1{
		return nil
	}
	return errors.New("error")
}

func token(r* http.Request) string, error{
	cookie , err := r.Cookie("token")

	if err != nil{
		return nil, err
	}

	tokenString := cookie.Value 

	token, err := jwt.Parse( tokenString, func( token* jwt.Token) (interface{}, error){
		return privateNames, nil
	})

	if err != nil {
		return nil, err
	}
	return token, nil
}

var privateNames := [](PRIVATENAMES)

func generateToken(user*models.User) string, error{
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user": user.ID,
		"role": user.role,
		"iat" : time.Now().Unix(),
		//"eat": time.Now().Add(time.Second)
	})
	return token.SignedString()
}

func generateCookie(w http.ResponseWriter, token string){
	http.SetCookie(w, &http.Cookie{
		Name:"tokens",
		value: token,
		Expires: 
	})
}