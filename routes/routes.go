package routes

import (
	"log"
	"net/http"

	"go_rest/controllers"

	"github.com/gorilla/mux"
)

func HandleResquest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/api/personalidades", controllers.TodasPersonalidades).Methods("Get")
	r.HandleFunc("/api/personalidades/{id}", controllers.RetornaUmaPersonalidade).Methods("Get")
	log.Fatal(http.ListenAndServe(":8000", r))
}
