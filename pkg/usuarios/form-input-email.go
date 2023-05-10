package usuarios

import (
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/services/db"
	"ecommerce/pkg/services/db/services"
	"ecommerce/pkg/services/log"

	"go.uber.org/zap"
)

// FormInputEMailAvailable struct
type FormInputEMailAvailable struct {
	ExcludeUserID int
	Email         string `json:"email"`
}

// Check function
func (input *FormInputEMailAvailable) Check() (bool, error) {
	const op errors.Operation = "pkg.usuarios.form-input-email.Check"
	conn := db.GetDB()

	usuario, err := services.UsuarioFindByEMail(conn, input.Email)
	if err != nil {
		log.Error2(op, err, zap.String("usuarixEmail: ", input.Email))
		return false, errors.NewError(op, errors.KindNotFound, err)
	} else {
		log.Debug("Busqueda de UsuarioFindByEMail finalizada correctamente")
	}

	mailAvailable := false
	if usuario.ID == 0 || usuario.ID == input.ExcludeUserID {
		mailAvailable = true
	}

	return mailAvailable, nil
}
