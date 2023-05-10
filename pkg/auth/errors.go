package auth

import "errors"

// Errors
var (
	ErrUserDisabled        = errors.New("ErrUserDisabled")
	ErrUserInvalid         = errors.New("ErrUserInvalid")
	ErrRolDisabled         = errors.New("ErrRolDisabled")
	ErrGrupoDisabled       = errors.New("ErrGrupoDisabled")
	ErrUserEMailNotFound   = errors.New("ErrUserEMailNotFound")
	ErrUserInvalidPassword = errors.New("ErrUserInvalidPassword")
)
