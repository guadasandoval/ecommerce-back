package usuarioactual

import (
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/libs/password"
	"ecommerce/pkg/services/db"
	"ecommerce/pkg/services/db/models"
	"ecommerce/pkg/services/db/services"
)

// FormInputCambiarContrasena struct
type FormInputCambiarContrasena struct {
	CurrentPassword string `json:"CurrentPassword"`
	Password        string `json:"Password"`
}

// Update function
func (input *FormInputCambiarContrasena) Update(usuario models.Usuario) error {
	const op errors.Operation = "pkg.usuarieactual.cambiar-contrasena.FormInputCambiarContrasena.Update"

	conn := db.GetDB()

	if !password.CheckValid(input.CurrentPassword, usuario.Password) {
		return errors.NewError(op, errors.KindUnexpected, ErrInvalidPassword)
	}

	usuario.Password = password.Encode(input.Password)
	eerr := services.UsuarioUpdatePassword(conn, usuario.ID, usuario.Password)
	if eerr != nil {
		return errors.NewError(op, errors.KindUnexpected, eerr)
	}

	return nil
}
