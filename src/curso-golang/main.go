package main

import (
	"gomssql/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//Rutas
	mux := mux.NewRouter()

	mux.HandleFunc("/api/permisos", handlers.GetPermisos).Methods("GET")
	mux.HandleFunc("/api/permisos", handlers.CreatePermiso).Methods("POST")
	mux.HandleFunc("/api/permisos", handlers.UpdatePermiso).Methods("PUT")

	mux.HandleFunc("/api/permisosusuarios", handlers.CreatePermisoUsuario).Methods("POST")
	mux.HandleFunc("/api/permisosusuarios/{id:[0-9]+}", handlers.DeletePermisoUsuario).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":3000", mux))
}
