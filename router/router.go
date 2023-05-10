package router

import (
	"ecommerce/pkg/services/log"
	"ecommerce/router/auth"
	"ecommerce/router/categorias"
	"ecommerce/router/productos"
	"ecommerce/router/importador"
	"ecommerce/router/usuarioactual"
	"ecommerce/router/usuarios"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// App router
type App struct {
	r    *httprouter.Router
	port string
}

// Initialize router
func Initialize(puerto string) {
	app := App{
		r:    httprouter.New(),
		port: puerto,
	}
	app.start()
}

func (a *App) start() {

	// Auth
	//a.r.HandleFunc("/api/login", auth.Login).Methods("POST")
	a.r.POST("/api/login-nuevo", MiddlewareLog(auth.LoginNuevo))
	a.r.POST("/api/check-login", MiddlewareUsuario(auth.CheckLogin))
	a.r.POST("/api/logout", MiddlewareUsuario(auth.Logout))

	a.r.POST("/api/altausuario", MiddlewareUsuario(usuarios.RegistrarNuevoUsuario))
	a.r.GET("/api/usuarios/roles-select", MiddlewareUsuario(usuarios.GetUsuarioRolesSelect))
	a.r.GET("/api/usuarios/grupos-select", MiddlewareUsuario(usuarios.GetUsuarioGruposSelect))
	a.r.POST("/api/usuarios/email-available", MiddlewareUsuario(usuarios.IsUsuarioEMailAvailable))
	a.r.GET("/api/usuarios/usuarios-table", MiddlewareUsuario(usuarios.ListarUsuariosTable))
	a.r.GET("/api/usuarios/usuarios/:id", MiddlewareUsuario(usuarios.GetUsuario))
	a.r.POST("/api/usuarios/usuarios/:id", MiddlewareUsuario(usuarios.UpdateUsuario))
	a.r.POST("/api/usuarios/usuarios-password/:id", MiddlewareUsuario(usuarios.UpdateUsuarioPassword))

	// Usuario actual
	a.r.POST("/api/usuario-actual/cambiar-contrasena", MiddlewareUsuario(usuarioactual.CambiarContrasena))
	a.r.GET("/api/usuario-actual/mis-datos", MiddlewareUsuario(usuarioactual.GetMisDatos))

	a.r.GET("/api/productos", MiddlewareUsuario(productos.ListarProductos))
	a.r.POST("/api/importar-productos", MiddlewareUsuario(importador.ImportarProductos))


	// Categorias
	a.r.GET("/api/categorias", MiddlewareUsuario(categorias.CategoriasProductos))


	
	err := http.ListenAndServe(":"+a.port, a.r)
	if err != nil {
		log.Fatal("Error al inicializar el router", log.String("error", err.Error()))
	}

}
