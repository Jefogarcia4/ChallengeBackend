package handlers

import (
	"encoding/json"
	"fmt"
	"gomssql/config"
	"gomssql/entities"
	"gomssql/models"
	"net/http"
)

func GetPermisos(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	db, err := config.GetDb()
	if err != nil {
		fmt.Println(err)
	} else {
		permisoModel := models.PermisoModel{
			Db: db,
		}
		permisos, err2 := permisoModel.FindAll()
		if err2 != nil {
			output, _ := json.Marshal(err2)
			fmt.Fprintln(rw, string(output))
		} else {
			output, _ := json.Marshal(permisos)
			fmt.Fprintln(rw, string(output))
		}
	}
}

func CreatePermiso(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	db, err := config.GetDb()
	if err != nil {
		fmt.Println(err)
	} else {
		permisoModel := models.PermisoModel{
			Db: db,
		}

		permiso := entities.Permiso{}
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&permiso); err != nil {
			fmt.Fprintln(rw, http.StatusUnprocessableEntity)
		} else {
			err2 := permisoModel.Create(&permiso)
			if err2 != nil {
				output, _ := json.Marshal(err2)
				fmt.Fprintln(rw, string(output))
			} else {
				output, _ := json.Marshal(&permiso)
				fmt.Fprintln(rw, string(output))
			}
		}

	}
}

func UpdatePermiso(rw http.ResponseWriter, r *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	db, err := config.GetDb()
	if err != nil {
		fmt.Println(err)
	} else {
		permisoModel := models.PermisoModel{
			Db: db,
		}

		permiso := entities.Permiso{}
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&permiso); err != nil {
			fmt.Fprintln(rw, http.StatusUnprocessableEntity)
		} else {
			rowsAffected, err2 := permisoModel.Update(&permiso)
			if err2 != nil {
				output, _ := json.Marshal(err2)
				fmt.Fprintln(rw, string(output))
			} else {
				output, _ := json.Marshal(&rowsAffected)
				fmt.Fprintln(rw, string(output))
			}
		}

	}
}
