package models

// Constantes de la tabla
const (
	TABLE_USUARIOSROLES  string = "usuarios_roles"
	PREFIX_USUARIOSROLES string = "rol"
)

// UsuariosRol struct
type UsuariosRol struct {
	ID            int    `gorm:"column:rol_id"`
	Nombre        string `gorm:"column:rol_nombre"`
	Descripcion   string `gorm:"column:rol_descripcion"`
	Habilitado    int    `gorm:"column:rol_habilitado"`
	ModUsuarios   int    `gorm:"column:rol_modusuarios"`
	ModDiversidad int    `gorm:"column:rol_moddiversidad"`
	ModMigracion  int    `gorm:"column:rol_modmigracion"`
}
