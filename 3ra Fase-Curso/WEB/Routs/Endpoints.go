package routs

import (
	"fmt"
	"log"
	"net/http"
	"text/template"

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

func Init(response http.ResponseWriter, request *http.Request) {

	html, error := template.ParseFiles("assets/templates/index.html")

	if error != nil {
		log.Fatal("Error al cargar la plantilla HTML ... ", error)
	}

	data := map[string]string{
		"Id":     "1",
		"nombre": "Juan Lopez",
		"correo": "juan.lopez@micorreo.edu",
	}

	html.Execute(response, data)
}

func Estructuras(response http.ResponseWriter, request *http.Request) {

	type Skill struct {
		Nombre string
	}

	type Persona struct {
		Nombre string
		Edad   int
		Genero string
		Skills []Skill
	}
	html, error := template.ParseFiles("assets/templates/estructuras/estruct.html")

	if error != nil {
		log.Fatal("Error al cargar la plantilla HTML ... ", error)
	}

	html.Execute(response,
		Persona{
			"Roberto Gómez",
			55, "Masculino",
			[]Skill{
				{Nombre: "Java"},
				{Nombre: "Golang"},
				{Nombre: "SQL"},
				{Nombre: "HTML"},
				{Nombre: "CSS"},
				{Nombre: "TypeScript"},
			},
		},
	)
}

func LoadTemplate(response http.ResponseWriter, request *http.Request) {

	html, error := template.ParseFiles("assets/templates/staticfiles/index.html")

	if error != nil {
		log.Fatal("Error al cargar la plantilla HTML ... ", error)
	}

	html.Execute(response, nil)
}
