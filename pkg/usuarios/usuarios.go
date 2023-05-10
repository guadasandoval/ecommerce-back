package usuarios

import (
	"ecommerce/pkg/libs/errors"
	"ecommerce/pkg/libs/password"
	"ecommerce/pkg/services/db"
	"ecommerce/pkg/services/db/models"
	"ecommerce/pkg/services/db/services"
	"ecommerce/pkg/services/log"
	"strconv"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// FormInputUsuario struct
type FormInputUsuario struct {
	ID         int    `json:"ID"`
	RolID      int    `json:"RolId"`
	GrupoID    int    `json:"GrupoId"`
	Email      string `json:"Email"`
	Nombres    string `json:"Nombres"`
	Apellidos  string `json:"Apellidos"`
	Password   string `json:"Password"`
	Habilitado bool   `json:"Habilitado"`
}

// Create function
func (input *FormInputUsuario) CrearUsuario() (int, error) {
	const op errors.Operation = "pkg.usuarios.usuarios.CreateUsuario"
	usuario, err := input.getUsuarioModel(0)
	if err != nil {
		log.Error2(op, err)
		return 0, err
	}

	usuario.Password = password.Encode(input.Password)

	conn := db.GetDB()
	err = services.UsuarioCreate(conn, &usuario)
	if err != nil {
		log.Error2(op, err)
		return 0, errors.NewError(op, errors.KindDBQuery, err)
	}

	return usuario.ID, nil
}

// Update function
func (input *FormInputUsuario) Update(id int) *errors.Error {
	const op errors.Operation = "pkg.usuarios.usuarios.FormInputUsuario.Update"

	conn := db.GetDB()

	oldUsuario, eerr := BuscarUsuarioPorID(id)
	if eerr != nil {
		return errors.NewError(op, errors.KindNotFound, eerr)
	}
	usuario, eerr := input.getUsuarioModel(id)
	if eerr != nil {
		return errors.NewError(op, errors.KindUnexpected, eerr)
	}

	usuario.Password = oldUsuario.Password

	eerr = services.UsuarioUpdate(conn, &usuario)
	if eerr != nil {
		return errors.NewError(op, errors.KindUnexpected, eerr)
	}

	return nil
}

// GetFormInputUsuario function
func GetFormInputUsuario(id int) (FormInputUsuario, *errors.Error) {
	const op errors.Operation = "pkg.usuarios.usuarios.GetFormInputUsuario"

	if id == 0 {
		return FormInputUsuario{}, nil
	}

	conn := db.GetDB()
	usuarix, err := services.BuscarUsuarioPorID(conn, id)
	if err != nil {
		return FormInputUsuario{}, errors.NewError(op, errors.KindUnexpected, err)
	}
	if usuarix.ID == 0 {
		err := errors.Msg("Usuario ID %d not found", id)
		return FormInputUsuario{}, errors.NewError(op, errors.KindUnexpected, err)
	}

	var input FormInputUsuario
	eerr := input.setUsuarioModel(conn, usuarix)
	if eerr != nil {
		return FormInputUsuario{}, eerr.New(op)
	}

	return input, nil
}

//BuscarPersonaPorID Busca una persona dado el ID
func BuscarUsuarioPorID(ID int) (*models.Usuario, error) {
	const op errors.Operation = "pkg.usuarios.usuarios.BuscarUsuarioPorID"
	conn := db.GetDB()
	usuario, err := services.BuscarUsuarioPorID(conn, ID)
	if err != nil {
		log.Error2(op, err, zap.String("usuarixID", strconv.Itoa(ID)))
		return nil, errors.NewError(op, errors.KindNotFound, err)
	} else {
		log.Debug("Busqueda de usuario finalizada correctamente", log.Int("usuarixID", ID))
	}

	return usuario, nil
}

func ListarUsuariosTable() (*[]models.Usuario, error) {
	const op errors.Operation = "pkg.usuarios.usuarios.ListarUsuariosTable"
	conn := db.GetDB()
	usuariosTable, err := services.UsuariosTable(conn)
	if err != nil {
		err = errors.Msg("Ocurrió un error en ListarPersonasDatosCompletos")
		log.Error2(op, err)
		return nil, errors.NewError(op, errors.KindUnexpected, err)
	}
	log.Debug("Se listaron los usuarios", log.Int("cantidadDeUsuarios", len(*usuariosTable)))

	return usuariosTable, nil
}

func (input *FormInputUsuario) getUsuarioModel(id int) (models.Usuario, error) {

	usuario := models.Usuario{
		ID:         input.ID,
		RolID:      input.RolID,
		GrupoID:    input.GrupoID,
		Email:      input.Email,
		Nombres:    input.Nombres,
		Apellidos:  input.Apellidos,
		Password:   input.Password,
		Habilitado: input.Habilitado,
	}

	return usuario, nil
}

func (input *FormInputUsuario) setUsuarioModel(conn *gorm.DB, usuario *models.Usuario) *errors.Error {
	const op errors.Operation = "pkg.core.usuarios.FormInputUsuario.setUsuarioModel"

	input.ID = usuario.ID
	input.RolID = usuario.RolID
	input.GrupoID = usuario.GrupoID
	input.Email = usuario.Email
	input.Nombres = usuario.Nombres
	input.Apellidos = usuario.Apellidos
	input.Password = usuario.Password
	input.Habilitado = usuario.Habilitado

	return nil
}
