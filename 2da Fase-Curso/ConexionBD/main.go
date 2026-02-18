package main

import (
	Handler "CONEXIONBD/handlers"
)

func main() {
	Handler.GetClientsAll()
	//nuevoCliente := models.Cliente{Nombre: "Maria Nieves", Correo: "maria.nieves@correo.com.mx", Telefono: "777777777"}
	//var otroCliente models.Cliente
	//Handler.SaveOrUpdate(nuevoCliente)
	//Handler.DeleteClientByEmail("maria.nieves@correo.com.mx")
}
