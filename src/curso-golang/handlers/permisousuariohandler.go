package handlers

import (
	"encoding/json"
	"fmt"
	"gomssql/config"
	"gomssql/entities"
	"gomssql/models"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func CreatePermisoUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	db, err := config.GetDb()
	if err != nil {
		fmt.Println(err)
	} else {
		permisoUsuarioModelModel := models.PermisoUsuarioModel{
			Db: db,
		}

		permisousuario := entities.PermisoUsuario{}
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&permisousuario); err != nil {
			fmt.Fprintln(rw, http.StatusUnprocessableEntity)
		} else {
			err2 := permisoUsuarioModelModel.Create(&permisousuario)
			if err2 != nil {
				output, _ := json.Marshal(err2)
				fmt.Fprintln(rw, string(output))
			} else {
				output, _ := json.Marshal(&permisousuario)
				fmt.Fprintln(rw, string(output))
			}
		}

	}
}

func DeletePermisoUsuario(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")

	db, err := config.GetDb()
	if err != nil {
		fmt.Println(err)
	} else {
		vars := mux.Vars(r)
		id := vars["id"]
		permisoUsuarioId, _ := strconv.ParseInt(id, 10, 64)
		permisoUsuarioModel := models.PermisoUsuarioModel{
			Db: db,
		}
		rowsAffected, err2 := permisoUsuarioModel.Delete(permisoUsuarioId)
		if err2 != nil {
			output, _ := json.Marshal(err2)
			fmt.Fprintln(rw, string(output))
		} else {
			output, _ := json.Marshal(rowsAffected)
			fmt.Fprintln(rw, string(output))
		}

	}
}
