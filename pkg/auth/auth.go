package auth

import (
	"ecommerce/pkg/services/keyvalue"
	"errors"
	"fmt"
)

// InputLogin login data
type InputLogin struct {
	User     string `json:"user"`
	Password string `json:"password"` // TODO: poner restricciones a la forma de la password
}


// GetLogin function
func GetLogin(token string) (int, error) {
	const op = "core.auth.GetLogin"

	has, err := keyvalue.HasValue(token)

	if err != nil {
		return 0, errors.New(op)
	}

	if !has {
		err := fmt.Sprintf("Invalid token %s", token)
		return 0, errors.New(op + ", " + err)
	}

	usuarioID, err := keyvalue.GetInt(token)

	if err != nil {
		return 0, errors.New(op)
	}

	err = keyvalue.SetValue(token, usuarioID)
	if err != nil {
		return 0, errors.New(op)
	}

	return usuarioID, nil
}

// Logout function
func Logout(token string) error {
	const op = "core.auth.Logout"

	has, err := keyvalue.HasValue(token)
	if err != nil {
		return errors.New(op)
	}

	if !has {
		return nil
	}

	err = keyvalue.DeleteValue(token)
	if err != nil {
		return errors.New(op)
	}

	return nil
}
