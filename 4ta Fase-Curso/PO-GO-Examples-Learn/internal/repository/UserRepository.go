package repository

import (
	Modelos "PO-GO-Examples-Learn/internal/models"
	"database/sql"
	"fmt"
	"log"
	"time"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) GetAllClientes() Modelos.Clientes {
	var clientes Modelos.Clientes

	query := "SELECT ID, NOMBRE, CORREO, TELEFONO FROM CLIENTES;"

	clientsAll, errorQuery := r.DB.Query(query)

	if errorQuery != nil {
		log.Panicln("Ocurrió un error al consultar los clientes ... ", errorQuery)
	}

	for clientsAll.Next() {
		var cliente = Modelos.Cliente{}

		errorClient := clientsAll.Scan(&cliente.Id, &cliente.Nombre, &cliente.Correo, &cliente.Telefono)

		if errorClient != nil {
			log.Fatalln("Error en la extracción de los clientes ... ", errorClient)
		}
		clientes = append(clientes, cliente)
	}

	return clientes
}

func (r *UserRepository) SaveOrUpdate(cliente Modelos.Cliente) (bool, string) {

	msg := ""
	bandera := false

	if (cliente != Modelos.Cliente{}) {
		var query string
		var errorQ error

		userExist := ExistEmail(cliente.Correo, r)

		if cliente.Id > 0 || userExist {
			fmt.Println("Actualizar cliente existente")
			query = "UPDATE CLIENTES SET NOMBRE = ?, CORREO = ?, TELEFONO = ?, FECHA = ? WHERE ID = ? ;"
			_, errorQ = r.DB.Exec(query, cliente.Nombre, cliente.Correo, cliente.Telefono, time.Now(), cliente.Id)
		} else {
			fmt.Println("Insertar nuevo cliente")
			query = "INSERT INTO CLIENTES(NOMBRE, CORREO, TELEFONO, FECHA) VALUES( ?,?,?,?);"
			_, errorQ = r.DB.Exec(query, cliente.Nombre, cliente.Correo, cliente.Telefono, time.Now())
		}

		if errorQ != nil {
			log.Fatalln("Error al actualizar o insertar registro ... ", errorQ)
			msg = "Error al guardar el cliente"
		}

		log.Println("Se actualizo o inserto correctamente ...")
		msg = "¡Se actualizo correctamente el cliente: " + cliente.Nombre + " !"
		bandera = true
	}

	return bandera, msg
}

func ExistEmail(email string, r *UserRepository) bool {

	if email != "" {

		var cliente = Modelos.Cliente{}

		query := "SELECT ID, NOMBRE, CORREO, TELEFONO FROM CLIENTES WHERE CORREO = ? LIMIT 1;"
		errorQ := r.DB.QueryRow(query, email).Scan(&cliente.Id, &cliente.Nombre, &cliente.Correo, &cliente.Telefono)

		if errorQ != nil {
			log.Printf("El usuario no existe ...")
		}

		if cliente.Correo == email {
			log.Printf("El correo: %v ya existe ...", email)
			return true
		}
	}

	return false
}

func (r *UserRepository) FindById(id int) (bool, Modelos.Cliente) {

	var cliente = Modelos.Cliente{}

	if id > 0 {

		query := "SELECT ID, NOMBRE, CORREO, TELEFONO FROM CLIENTES WHERE ID = ? LIMIT 1;"
		errorQ := r.DB.QueryRow(query, id).Scan(&cliente.Id, &cliente.Nombre, &cliente.Correo, &cliente.Telefono)

		if errorQ != nil {
			log.Printf("El usuario no existe ...")

			return false, cliente
		}
	}

	return true, cliente
}
