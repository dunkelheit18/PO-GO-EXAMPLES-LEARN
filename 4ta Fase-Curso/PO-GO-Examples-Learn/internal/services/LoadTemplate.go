package services

import (
	Modelos "PO-GO-Examples-Learn/internal/models"
	Repository "PO-GO-Examples-Learn/internal/repository"
	RegistroVal "PO-GO-Examples-Learn/internal/services/registro"
	Rutas "PO-GO-Examples-Learn/internal/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/gorilla/mux"
	"github.com/jung-kurt/gofpdf"
	excelize "github.com/xuri/excelize/v2"
)

type ServiceClients struct {
	Repository *Repository.UserRepository
}

func NewServiceClients(repo *Repository.UserRepository) *ServiceClients {
	return &ServiceClients{
		Repository: repo,
	}
}

func LoadTemplate(response http.ResponseWriter, request *http.Request) {

	html, error := template.ParseFiles(Rutas.INDEX_PATH, Rutas.TEMPLATE_PATH)

	if error != nil {
		log.Fatal("Error al cargar la plantilla HTML ... ", error)
	}

	html.Execute(response, nil)
}

func LoadRegistrotemplate(response http.ResponseWriter, request *http.Request) {

	html := template.Must(template.ParseFiles(Rutas.REGISTRO_PATH, Rutas.TEMPLATE_PATH))

	css, msg := Rutas.CallMessage(response, request)

	data := map[string]string{
		"css":     css,
		"mensaje": msg,
	}

	html.Execute(response, data)
}

func Regitro(response http.ResponseWriter, request *http.Request) {

	nombre := request.FormValue("nombre")
	correo := request.FormValue("email")
	password := request.FormValue("password")

	data := map[string]string{
		"nombre":   nombre,
		"email":    correo,
		"password": password,
	}

	err, msg := RegistroVal.ValidaDatos(data)

	if err {
		log.Println(msg)
		Rutas.CreateMessage(response, request, "danger", msg)
		http.Redirect(response, request, "/Formularios", http.StatusSeeOther)
	}

	fmt.Fprintln(response, "Nombre: "+nombre)
}

func LoadTemplateUpload(response http.ResponseWriter, request *http.Request) {

	html, error := template.ParseFiles(Rutas.UPLOAD_PATH, Rutas.TEMPLATE_PATH)

	if error != nil {
		log.Fatal("Error al cargar la plantilla HTML ... ", error)
	}

	css, msg := Rutas.CallMessage(response, request)

	data := map[string]string{
		"css":     css,
		"mensaje": msg,
	}

	html.Execute(response, data)
}

func Uploadfiles(response http.ResponseWriter, request *http.Request) {

	file, handler, error := request.FormFile("image")

	if error != nil {
		Rutas.CreateMessage(response, request, "danger", "Error inesperado!")
	}

	extencion := strings.Split(handler.Filename, ".")[1]
	prefix := strings.Split(time.Now().String(), " ")

	fichero := string(prefix[4][6:14]) + "." + extencion

	archivo := "ui/resources/uploads/files/" + fichero

	f, err := os.OpenFile(archivo, os.O_WRONLY|os.O_CREATE, 0777)

	if err != nil {
		Rutas.CreateMessage(response, request, "danger", "Error al cargar el archivo!")
	}

	_, errCopy := io.Copy(f, file)

	if errCopy != nil {
		Rutas.CreateMessage(response, request, "danger", "Error al guardar el archivo!")
	}

	Rutas.CreateMessage(response, request, "success", "El archivo: "+fichero+" se cargo exitosamente!")

	http.Redirect(response, request, "/Upload", http.StatusSeeOther)
}

func LoadTemplateResources(response http.ResponseWriter, request *http.Request) {

	html, error := template.ParseFiles(Rutas.RESOURCE_PATH, Rutas.TEMPLATE_PATH)

	if error != nil {
		log.Fatal("Error al cargar la plantilla HTML ... ", error)
	}

	css, msg := Rutas.CallMessage(response, request)

	data := map[string]string{
		"css":     css,
		"mensaje": msg,
	}

	html.Execute(response, data)
}

func GeneratePDF(response http.ResponseWriter, request *http.Request) {

	msg := ""
	css := ""
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(40, 10, "Hello Word!")

	err := pdf.OutputFileAndClose("ui/resources/generate/pdf/MyPdf.pdf")

	if err != nil {
		msg = "Error al generar el archivo PDF."
		css = "danger"

	} else {
		msg = "El PDF se generó exitosamente"
		css = "success"
	}

	Rutas.CreateMessage(response, request, css, msg)
	http.Redirect(response, request, "/Resources", http.StatusSeeOther)
}

func GenerateExcel(response http.ResponseWriter, request *http.Request) {

	msg := ""
	css := ""

	exeFile := excelize.NewFile()
	defer func() {
		exeFile.Close()
	}()

	index, err := exeFile.NewSheet("Sheet1")

	if err != nil {
		msg = "Error al generar la hoja de excel."
		css = "danger"
	}

	exeFile.SetCellValue("Sheet1", "A1", "Id")
	exeFile.SetCellValue("Sheet1", "B1", "Nombre")
	exeFile.SetCellValue("Sheet1", "C1", "Correo")
	exeFile.SetActiveSheet(index)

	prefix := strings.Split(time.Now().String(), " ")
	fichero := "Excel-" + string(prefix[4][6:14]) + "." + "xlsx"

	if errSave := exeFile.SaveAs("ui/resources/generate/excel/" + fichero); errSave != nil {
		msg = "Error al guardar el archivo excel."
		css = "danger"
	} else {
		msg = "El archivo excel: " + fichero + " se generó correctamente."
		css = "success"
	}

	Rutas.CreateMessage(response, request, css, msg)
	http.Redirect(response, request, "/Resources", http.StatusSeeOther)
}

func LoadTemplateClientes(response http.ResponseWriter, request *http.Request) {

	html, error := template.ParseFiles(Rutas.CLIENTES_PATH, Rutas.TEMPLATE_PATH)

	if error != nil {
		log.Fatal("Error al cargar la plantilla HTML ... ", error)
	}

	css, msg := Rutas.CallMessage(response, request)

	data := map[string]string{
		"Css":     css,
		"Mensaje": msg,
	}

	html.Execute(response, data)
}

func (service *ServiceClients) ListClients(response http.ResponseWriter, request *http.Request) {

	msg := ""
	css := ""

	html, error := template.ParseFiles(Rutas.CLIENTES_PATH, Rutas.TEMPLATE_PATH)

	if error != nil {
		log.Fatal("Error al cargar la plantilla HTML ... ", error)
	}

	clients := service.Repository.GetAllClientes()

	if clients != nil {
		msg = "¡Búsqueda exitosa!"
		css = "success"
	} else {
		msg = "No se encontraron clientes"
		css = "danger"
	}

	Rutas.CreateMessage(response, request, css, msg)

	html.Execute(response,
		Modelos.ClientesResponse{
			Css:     css,
			Mensaje: msg,
			Data:    clients,
		},
	)
}

func LoadTemplateNewClient(response http.ResponseWriter, request *http.Request) {

	html, error := template.ParseFiles(Rutas.CLIENTES_REGISTRO_PATH, Rutas.TEMPLATE_PATH)

	if error != nil {
		log.Fatal("Error al cargar la plantilla HTML ... ", error)
	}

	css, msg := Rutas.CallMessage(response, request)

	data := map[string]string{
		"Css":     css,
		"Mensaje": msg,
	}

	html.Execute(response, data)
}

func (service *ServiceClients) RegistreClient(response http.ResponseWriter, request *http.Request) {

	css := ""

	html, error := template.ParseFiles(Rutas.CLIENTES_REGISTRO_PATH, Rutas.TEMPLATE_PATH)

	if error != nil {
		log.Fatal("Error al cargar la plantilla HTML ... ", error)
	}

	nombre := request.FormValue("nombre")
	correo := request.FormValue("email")
	telefono := request.FormValue("telefono")

	bandera, mensaje := service.Repository.SaveOrUpdate(
		Modelos.Cliente{
			Nombre:   nombre,
			Correo:   correo,
			Telefono: telefono,
		},
	)

	if bandera {
		css = "success"
	} else {
		css = "danger"
	}

	Rutas.CreateMessage(response, request, css, mensaje)

	data := map[string]string{
		"Css":     css,
		"Mensaje": mensaje,
	}

	html.Execute(response, data)
}

func (service *ServiceClients) LoadTemplateUpdateClient(response http.ResponseWriter, request *http.Request) {

	css := ""
	msg := ""

	html, error := template.ParseFiles(Rutas.CLIENTES_ACTUALIZAR_PATH, Rutas.TEMPLATE_PATH)

	if error != nil {
		log.Fatal("Error al cargar la plantilla HTML ... ", error)
	}

	vars := mux.Vars(request)
	idCliente, _ := strconv.Atoi(vars["id"])

	banCliente, datosCliente := service.Repository.FindById(idCliente)

	if banCliente {
		css = "success"
		msg = "Se consulto el cliente correctamente."
	} else {
		css = "danger"
		msg = "El cliente no fue encontrado"
	}

	Rutas.CreateMessage(response, request, css, msg)

	html.Execute(response,
		Modelos.ClienteResponse{
			Css:     css,
			Mensaje: msg,
			Data:    datosCliente,
		},
	)
}

func (service *ServiceClients) UpdateClient(response http.ResponseWriter, request *http.Request) {

	css := ""

	html, error := template.ParseFiles(Rutas.CLIENTES_ACTUALIZAR_PATH, Rutas.TEMPLATE_PATH)

	if error != nil {
		log.Fatal("Error al cargar la plantilla HTML ... ", error)
	}

	nombre := request.FormValue("nombre")
	correo := request.FormValue("email")
	telefono := request.FormValue("telefono")

	bandera, mensaje := service.Repository.SaveOrUpdate(
		Modelos.Cliente{
			Nombre:   nombre,
			Correo:   correo,
			Telefono: telefono,
		},
	)

	if bandera {
		css = "success"
	} else {
		css = "danger"
	}

	Rutas.CreateMessage(response, request, css, mensaje)

	html.Execute(response, Modelos.ClienteResponse{
		Css:     css,
		Mensaje: mensaje,
		Data: Modelos.Cliente{
			Nombre:   nombre,
			Correo:   correo,
			Telefono: telefono,
		},
	},
	)
}
