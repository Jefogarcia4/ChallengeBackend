package models

import (
	"database/sql"
	"gomssql/entities"
)

type PermisoUsuarioModel struct {
	Db *sql.DB
}

func (permisoUsuarioModel PermisoUsuarioModel) Create(permisousuario *entities.PermisoUsuario) error {
	row, err := permisoUsuarioModel.Db.Query("INSERT INTO PermisosUsuarios VALUES(?,?) select convert(bigint, scope_identity());",
		permisousuario.UsuarioId, permisousuario.PermisoId)
	if err != nil {
		return err
	} else {
		var newId int64
		row.Next()
		row.Scan(&newId)
		permisousuario.PermisoUsuarioId = newId
		return nil
	}
}

func (permisoUsuarioModel PermisoUsuarioModel) Delete(id int64) (int64, error) {
	result, err := permisoUsuarioModel.Db.Exec("DELETE FROM PermisosUsuarios WHERE PermisoUsuarioId = ?  ", id)
	if err != nil {
		return 0, err
	} else {
		return result.RowsAffected()
	}
}
