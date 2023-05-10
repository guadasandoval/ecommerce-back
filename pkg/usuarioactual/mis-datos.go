package usuarioactual

import (
	"ecommerce/pkg/auth"
	"ecommerce/pkg/services/db/models"
)

// FormInputMisDatos struct
type FormInputMisDatos struct {
	Nombres   string `valid:"Required|MaxLength:100" json:"nombres"`
	Apellidos string `valid:"Required|MaxLength:100" json:"apellidos"`
	Email     string `valid:"Required|EMail|MaxLength:100" json:"email"`
}

// GetFormInputMisDatos function
func GetFormInputMisDatos(usuario *models.Usuario) (FormInputMisDatos, error) {

	if usuario.ID == 0 {
		return FormInputMisDatos{}, auth.ErrUserInvalid
	}

	misDatos := FormInputMisDatos{
		Nombres:   usuario.Nombres,
		Apellidos: usuario.Apellidos,
		Email:     usuario.Email,
	}

	return misDatos, nil
}
