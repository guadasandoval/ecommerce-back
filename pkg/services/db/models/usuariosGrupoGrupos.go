package models

// Constantes de la tabla
const (
	TableUsuariosGrupoGrupos  string = "usuarios_grupo_grupos"
	PrefixUsuariosGrupoGrupos string = "ugg"
)

// UsuariosGrupoGrupo struct
type UsuariosGrupoGrupo struct {
	ID         int `gorm:"column:ugg_id"`
	GrupoID    int `name:"column:ugggrp_id"`
	SubGrupoID int `name:"column:ugggrp_idchild"`
}
