package controls

import(
	"net/http"
	"encoding/json"
)

func Output(w http.ResponseWriter, data interface{}){

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(data); err != nil{
		return
	}
}

func Response(w http.ResponseWriter, err error){
	newData := map[string]interface{}{"error":err.Error()}

	Output(w,newData)
}