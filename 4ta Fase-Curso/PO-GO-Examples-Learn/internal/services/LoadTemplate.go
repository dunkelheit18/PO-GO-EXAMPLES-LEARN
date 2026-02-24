package services

import (
	RegistroVal "PO-GO-Examples-Learn/internal/services/registro"
	Rutas "PO-GO-Examples-Learn/internal/utils"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/jung-kurt/gofpdf"
	excelize "github.com/xuri/excelize/v2"
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
