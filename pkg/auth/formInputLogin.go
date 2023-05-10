package auth

import (
	"ecommerce/pkg/libs/password"
	"ecommerce/pkg/services/db"
	"ecommerce/pkg/services/db/models"
	"ecommerce/pkg/services/db/services"
	"ecommerce/pkg/services/keyvalue"
	"ecommerce/pkg/services/log"
	"fmt"
	"strings"

	"github.com/google/uuid"
)

// FormInputLogin struct
type FormInputLogin struct {
	Email    string `valid:"Required|Email|MaxLength:100" json:"email"`
	Password string `valid:"Required|MinLength:6|MaxLength:100" json:"password"`
}

type UsuarioToken struct {
	Token   string         `json:"token"`
	Usuario models.Usuario `json:"usuario"`
}

// Login function
func (input *FormInputLogin) Login() (LoginResponse, error) {

	conn := db.GetDB()

	usuario, err := services.UsuarioFindByEMail(conn, input.Email)

	if err != nil {
		log.Error("Ocurrio un error en UsuarioFindByEMail()", log.String("error", err.Error()))
		return LoginResponse{}, err
	}
	if usuario.ID == 0 {
		return LoginResponse{}, ErrUserEMailNotFound
	}

	if !password.CheckValid(input.Password, usuario.Password) {
		return LoginResponse{}, ErrUserInvalidPassword
	}

	if !usuario.Habilitado {
		return LoginResponse{}, ErrUserDisabled
	}

	token := fmt.Sprintf("userlogin-%s-%d-%s", uuid.New().String(), usuario.ID, uuid.New().String())

	err = keyvalue.SetValue(token, usuario.ID)
	if err != nil {
		return LoginResponse{}, err
	}

	usuarioNombre := usuario.Nombres + " " + usuario.Apellidos
	usuarioNombre = strings.TrimSpace(usuarioNombre)

	loginResponse := LoginResponse{
		Token: token,
		Usuario: LoginResponseUsuario{
			ID:     usuario.ID,
			Nombre: usuarioNombre,
			Email:  usuario.Email,
		},
	}

	return loginResponse, nil
}
