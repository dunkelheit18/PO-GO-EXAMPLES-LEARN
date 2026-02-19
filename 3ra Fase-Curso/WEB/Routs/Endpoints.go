package routs

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func Home(response http.ResponseWriter, request *http.Request) {

	fmt.Fprintln(response, "Bienvenido al servicio web en Golang ...")
}

func ServiceVersion(response http.ResponseWriter, request *http.Request) {

	fmt.Fprintln(response, "1.0.0")
}

func Saludo(response http.ResponseWriter, request *http.Request) {

	vars := mux.Vars(request)

	fmt.Fprintln(response, "¡Hola "+vars["name"]+" "+vars["apellido"]+"! | Bienvenido a mi sitio")
}

func Despedida(response http.ResponseWriter, request *http.Request) {

	fmt.Fprintln(response, "Adios "+request.URL.Query().Get("name")+" "+request.URL.Query().Get("apellido")+"!")
}
