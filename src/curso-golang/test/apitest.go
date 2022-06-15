package testunitario

import (
	"gomssql/config"
	"gomssql/models"
	"testing"
)

func TestGetList(t *testing.T) {
	db, err := config.GetDb()
	if err != nil {
		t.Errorf("Ocurrió un error conectando a la base de datos")
	} else {
		permisoModel := models.PermisoModel{
			Db: db,
		}
		permisos, err2 := permisoModel.FindAll()
		if err2 != nil {
			t.Errorf("Ocurrió un error consultando los permisos.")
		} else {
			if permisos != nil {
				t.Errorf("Ocurrió un error consultando los permisos.")
			}
		}
	}
}
