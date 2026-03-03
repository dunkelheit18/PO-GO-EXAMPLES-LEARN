package main

import (
	Configuratios "PO-GO-Examples-Learn/internal/config"
	Repository "PO-GO-Examples-Learn/internal/repository"
	Services "PO-GO-Examples-Learn/internal/services"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var muxS *mux.Router
var clientService *Services.ServiceClients

func main() {
	dbConfig, _ := Configuratios.LoadDBConfig()
	dbConnection := Configuratios.GetInstance(dbConfig)
	defer dbConnection.Close()

	userRepo := Repository.NewUserRepository(dbConnection)
	clientService = Services.NewServiceClients(userRepo)

	StaticFilesMuxServer()
}

func StaticFilesMuxServer() {
	EnvGet()
	muxS = mux.NewRouter()
	defer ServerConfig()

	muxS.HandleFunc("/", Services.LoadTemplate)
	muxS.HandleFunc("/Formularios", Services.LoadRegistrotemplate)
	muxS.HandleFunc("/Registro", Services.Regitro).Methods("POST")
	muxS.HandleFunc("/Upload", Services.LoadTemplateUpload)
	muxS.HandleFunc("/Upload-file", Services.Uploadfiles).Methods("POST")
	muxS.HandleFunc("/Resources", Services.LoadTemplateResources)
	muxS.HandleFunc("/Resources/Generate-pdf", Services.GeneratePDF)
	muxS.HandleFunc("/Resources/Generate-excel", Services.GenerateExcel)
	muxS.HandleFunc("/Clientes", Services.LoadTemplateClientes)
	muxS.HandleFunc("/Get-clientes", clientService.ListClients)
	muxS.HandleFunc("/Save-clientes", Services.LoadTemplateNewClient)
	muxS.HandleFunc("/Registrar-clientes", clientService.RegistreClient).Methods("POST")
	muxS.HandleFunc("/Actualizar-clientes/{id:.*}", clientService.LoadTemplateUpdateClient)
	muxS.HandleFunc("/Update-cliente", clientService.UpdateClient).Methods("POST")

	//CONFIGURACIÓN PARA QUE MUX PUEDA CARGAR LOS ARCHIVOS PUBLICOS
	handlerPrefix := http.StripPrefix("/ui/", http.FileServer(http.Dir("./ui/")))
	muxS.PathPrefix("/ui/").Handler(handlerPrefix)
}

func ServerConfig() {

	serverConfig := &http.Server{
		Addr:         os.Getenv("SERVER_HOST") + ":" + os.Getenv("SERVER_PORT"),
		Handler:      muxS,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(serverConfig.ListenAndServe())
}

func EnvGet() {
	error := godotenv.Load()

	if error != nil {
		log.Panicln("Error al cargar las veriables de entorno ... ", error)
	}
}
