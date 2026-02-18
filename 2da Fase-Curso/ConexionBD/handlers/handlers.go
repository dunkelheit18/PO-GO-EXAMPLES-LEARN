package handlers

import (
	DbConnection "CONEXIONBD/connection"
	"CONEXIONBD/models"
	"fmt"
	"log"
	"time"
)

func GetClientsAll() {
	DbConnection.Conexion()

	query := "SELECT ID, NOMBRE, CORREO, TELEFONO FROM CLIENTES;"

	clientsAll, errorQuery := DbConnection.DbMysql.Query(query)

	if errorQuery != nil {
		log.Panicln("Ocurrió un error al consultar los clientes ... ", errorQuery)
	}
	defer DbConnection.CloseConnection()

	// OPCION UNO
	//clientes := models.Clientes{}

	for clientsAll.Next() {
		/* OPCION UNO PARA OBTENER UNA LISTA DE OBJETOS TIPO CLIENTE
		 		cliente := models.Cliente{}
				clientsAll.Scan(&cliente.Id, &cliente.Nombre, &cliente.Correo, &cliente.Telefono)
				clientes = append(clientes, cliente) */

		//OPCION DOS PARA OBTENER UNA LISTA DE OBJETOS TIPO CLIENTE
		var cliente = models.Cliente{}

		errorClient := clientsAll.Scan(&cliente.Id, &cliente.Nombre, &cliente.Correo, &cliente.Telefono)

		if errorClient != nil {
			log.Fatalln("Error en la extracción de los clientes ... ", errorClient)
		}

		//OPCION DOS
		fmt.Printf("Id: %v | Nombre: %v | E-mail: %v | Telefono: %v \n", cliente.Id, cliente.Nombre, cliente.Correo, cliente.Telefono)
		fmt.Println("------------------------------------------------------------------------------")
	}
	// OPCION UNO
	// fmt.Println(clientes)
}

func SaveOrUpdate(cliente models.Cliente) {

	if (cliente != models.Cliente{}) {
		var query string
		var errorQ error

		userExist := ExistEmail(cliente.Correo)
		DbConnection.Conexion()
		defer DbConnection.CloseConnection()
		if cliente.Id > 0 || userExist {
			fmt.Println("Actualizar cliente existente")
			query = "UPDATE CLIENTES SET NOMBRE = ?, CORREO = ?, TELEFONO = ?, FECHA = ? WHERE ID = ? ;"
			_, errorQ = DbConnection.DbMysql.Exec(query, cliente.Nombre, cliente.Correo, cliente.Telefono, time.Now(), cliente.Id)
		} else {
			fmt.Println("Insertar nuevo cliente")
			query = "INSERT INTO CLIENTES(NOMBRE, CORREO, TELEFONO, FECHA) VALUES( ?,?,?,?);"
			_, errorQ = DbConnection.DbMysql.Exec(query, cliente.Nombre, cliente.Correo, cliente.Telefono, time.Now())
		}

		if errorQ != nil {
			log.Fatalln("Error al actualizar o insertar registro ... ", errorQ)
		}

		log.Println("Se actualizo o inserto correctamente ...")
	}

	GetClientsAll()
}

func ExistEmail(email string) bool {

	if email != "" {

		var cliente = models.Cliente{}

		DbConnection.Conexion()
		defer DbConnection.CloseConnection()

		query := "SELECT ID, NOMBRE, CORREO, TELEFONO FROM CLIENTES WHERE CORREO = ? LIMIT 1;"
		errorQ := DbConnection.DbMysql.QueryRow(query, email).Scan(&cliente.Id, &cliente.Nombre, &cliente.Correo, &cliente.Telefono)

		if errorQ != nil {
			log.Fatalln("El usuario no existe ...", errorQ)
		}

		if cliente.Correo == email {
			log.Printf("El correo: %v ya existe ...", email)
			return true
		}

		return false
	}

	return false
}

func DeleteClientByEmail(email string) {

	if email != "" {

		DbConnection.Conexion()
		defer DbConnection.CloseConnection()

		query := "DELETE FROM CLIENTE WHERE CORREO = ? ;"
		_, errorQ := DbConnection.DbMysql.Exec(query, email)

		if errorQ != nil {
			log.Fatalln("El usuario no existe ...", errorQ)
		} else {
			log.Printf("El usuario con email: %v fué eliminado ...", email)
		}
	}

	log.Println("Es necesario ingresar un correo ...")
}
