package entities

type Permiso struct {
	PermisoId   int64  `json:"permisoid"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
	Activo      bool   `json:"activo"`
}
