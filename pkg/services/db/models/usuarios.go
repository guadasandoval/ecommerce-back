package models

// Constantes de la tabla
const TABLE_USUARIOS string = "usuarios"
const PREFIX_USUARIOS string = "usr"

// Usuario model struct
type Usuario struct {
	ID            int    `gorm:"column:usr_id"`
	RolID         int    `gorm:"column:usrrol_id"`
	GrupoID       int    `gorm:"column:usrgrp_id"`
	Email         string `gorm:"column:usr_email"`
	Nombres       string `gorm:"column:usr_nombres"`
	Apellidos     string `gorm:"column:usr_apellidos"`
	Password      string `gorm:"column:usr_password"`
	ResetPassword string `gorm:"column:usr_resetpassword"`
	Habilitado    bool   `gorm:"column:usr_habilitado"`
}
