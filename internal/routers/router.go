package router

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := map[string]string{"name": "Armaan"}
		encoded_data, err := json.Marshal(data)
		if err != nil {
			w.WriteHeader(500)
			w.Write([]byte("Error while encoding json data"))
		} else {
			w.Write(encoded_data)
		}
	}).Methods("GET")
	return r
}
