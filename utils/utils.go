package unmarshal

import(
	"net/http"
	"io/ioutil"
	"encoding/json"
)


func ChangeFmt(r*http.Request,t interface{}){
	body, err := ioutil.ReadAll(r.Body)
	if err == nil{
		if err := json.Unmashal([]byte(body),t); err != nil{
			return
		}
	}
}