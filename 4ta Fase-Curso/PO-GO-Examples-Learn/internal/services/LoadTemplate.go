package services

import (
	RegistroVal "PO-GO-Examples-Learn/internal/services/registro"
	Rutas "PO-GO-Examples-Learn/internal/utils"
	"fmt"
	"log"
	"net/http"
	"text/template"
)

func LoadTemplate(response http.ResponseWriter, request *http.Request) {

	html, error := template.ParseFiles(Rutas.INDEX_PATH, Rutas.TEMPLATE_PATH)

	if error != nil {
		log.Fatal("Error al cargar la plantilla HTML ... ", error)
	}

	html.Execute(response, nil)
}

func LoadRegistrotemplate(response http.ResponseWriter, request *http.Request) {

	html := template.Must(template.ParseFiles(Rutas.REGISTRO_PATH, Rutas.TEMPLATE_PATH))

	html.Execute(response, nil)
}

func Regitro(response http.ResponseWriter, request *http.Request) {

	nombre := request.FormValue("nombre")

	err, msg := RegistroVal.ValidaNombre(nombre)

	if err {
		log.Println(msg)
		Rutas.CreateMessage(response, request, "danger", msg)
		http.Redirect(response, request, "/Formularios", http.StatusSeeOther)
	}

	fmt.Fprintln(response, "Nombre: "+nombre)
}
