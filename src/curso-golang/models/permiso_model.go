package models

import (
	"database/sql"
	"gomssql/entities"
)

type PermisoModel struct {
	Db *sql.DB
}

func (permisoModel PermisoModel) FindAll() ([]entities.Permiso, error) {
	rows, err := permisoModel.Db.Query("select * from Permisos")
	if err != nil {
		return nil, err
	} else {
		var permisos []entities.Permiso
		for rows.Next() {
			var permisoId int64
			var nombre string
			var descripcion string
			var activo bool
			err2 := rows.Scan(&permisoId, &nombre, &descripcion, &activo)
			if err2 != nil {
				return nil, err2
			} else {
				permiso := entities.Permiso{
					PermisoId:   permisoId,
					Nombre:      nombre,
					Descripcion: descripcion,
					Activo:      activo,
				}
				permisos = append(permisos, permiso)
			}
		}
		return permisos, nil

	}
}

func (permisoModel PermisoModel) Create(permiso *entities.Permiso) error {
	row, err := permisoModel.Db.Query("INSERT INTO Permisos VALUES(?,?,?) select convert(bigint, scope_identity());",
		permiso.Nombre, permiso.Descripcion, permiso.Activo)
	if err != nil {
		return err
	} else {
		var newId int64
		row.Next()
		row.Scan(&newId)
		permiso.PermisoId = newId
		return nil
	}
}

func (permisoModel PermisoModel) Update(permiso *entities.Permiso) (int64, error) {
	result, err := permisoModel.Db.Exec("UPDATE Permisos set Nombre = ?, Descripcion = ?, Activo = ? WHERE PermisoId = ?  ",
		permiso.Nombre, permiso.Descripcion, permiso.Activo, permiso.PermisoId)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
