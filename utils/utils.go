package utils

import(
	"net/http"
	"encoding/json"
)


func ChangeFormat(r*http.Request,t interface{}) error{
	return json.NewDecoder(r.Body).Decode(t)
}