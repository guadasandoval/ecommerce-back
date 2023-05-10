package usuarios

import (
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/libs/password"
	"ecommerce/pkg/services/db"
	"ecommerce/pkg/services/db/services"
)

// FormInputPassword struct
type FormInputPassword struct {
	Password string `json:"Password"`
}

// Update function
func (input *FormInputPassword) UpdatePassword(id int) *errors.Error {
	const op errors.Operation = "pkg.usuarios.FormInputPassword.Update"

	conn := db.GetDB()

	usuario, eerr := BuscarUsuarioPorID(id)
	if eerr != nil {
		return errors.NewError(op, errors.KindNotFound, eerr)
	}

	usuario.Password = password.Encode(input.Password)
	eerr = services.UsuarioUpdatePassword(conn, usuario.ID, usuario.Password)
	if eerr != nil {
		return errors.NewError(op, errors.KindNotFound, eerr)
	}

	return nil
}
