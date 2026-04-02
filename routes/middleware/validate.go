package routes

import (
	"fmt"
    "net/http"
)

func ValidateMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if  err := validateToken(r); err != nil{
            fmt.Fprintf(w,"Error: %s", err)
            return
        }
        next.ServeHTTP(w, r)
    })
}

func AdminMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if  err := AdminValidate(r); err != nil{     
            fmt.Fprintf(w,"Error: %s", err)
            return
        }
        next.ServeHTTP(w, r)
    })
}