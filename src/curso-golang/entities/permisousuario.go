package entities

type PermisoUsuario struct {
	PermisoUsuarioId int64  `json:"permisousuarioid"`
	UsuarioId        string `json:"usuarioid"`
	PermisoId        int    `json:"permisoid"`
}
