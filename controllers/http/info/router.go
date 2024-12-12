package info_router

import "net/http"

type infoRouter struct {
	
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) { 
	if r.Method == http.MethodGet{

	}
}