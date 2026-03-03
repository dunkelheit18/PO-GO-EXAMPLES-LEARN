package models

type Cliente struct {
	Id       int
	Nombre   string
	Correo   string
	Telefono string
}

type Clientes []Cliente

type ClientesResponse struct {
	Css     string
	Mensaje string
	Data    Clientes
}

type ClienteResponse struct {
	Css     string
	Mensaje string
	Data    Cliente
}

type Usuario struct {
	Id       int
	Nombre   string
	Correo   string
	Password string
}

type Usuarios []Usuario
