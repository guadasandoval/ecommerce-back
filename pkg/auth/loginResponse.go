package auth

// LoginResponse struct
type LoginResponse struct {
	Token   string               `json:"token"`
	Usuario LoginResponseUsuario `json:"usuario"`
}

// LoginResponseUsuario struct
type LoginResponseUsuario struct {
	ID     int    `json:"id"`
	Nombre string `json:"nombre"`
	Email  string `json:"email"`
}
