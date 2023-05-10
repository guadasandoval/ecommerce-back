package models

// Constantes de la tabla
const (
	TABLE_USUARIOSGRUPOS  string = "usuarios_grupos"
	PREFIX_USUARIOSGRUPOS string = "grp"
)

// UsuariosGrupo struct
type UsuariosGrupo struct {
	ID          int    `gorm:"column:grp_id"`
	Nombre      string `gorm:"column:grp_nombre"`
	Descripcion string `gorm:"column:grp_descripcion"`
	Habilitado  int    `gorm:"column:grp_habilitado"`
}
