package main

import (
	Endpints "WEB/Routs"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

var muxS *mux.Router

func main() {

	// Servidor usando net/htpp
	//ServerTest()

	// Servidor usando Gorilla Mux
	//GorillaMuxServer()

	// Carga de archivos estaticos en Gorilla Mux
	StaticFilesMuxServer()
}

func ServerTest() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "Hello word!")
	})

	fmt.Println("URL del servidor de prueba: http://localhost:8081")
	log.Fatal(http.ListenAndServe("localhost:8081", nil))
}

func StaticFilesMuxServer() {
	EnvGet()
	muxS = mux.NewRouter()
	defer ServerConfig()

	muxS.HandleFunc("/", Endpints.LoadTemplate)

	//CONFIGURACIÓN PARA QUE MUX PUEDA CARGAR LOS ARCHIVOS PUBLICOS
	handlerPrefix := http.StripPrefix("/assets/templates/staticfiles/", http.FileServer(http.Dir("./assets/templates/staticfiles/")))
	muxS.PathPrefix("/assets/templates/staticfiles/").Handler(handlerPrefix)
}

func GorillaMuxServer() {
	EnvGet()
	muxS = mux.NewRouter()
	muxS.HandleFunc("/", Endpints.Init)
	muxS.HandleFunc("/home", Endpints.Home)
	muxS.HandleFunc("/version", Endpints.ServiceVersion)
	//URL con Parametros en el PATH
	muxS.HandleFunc("/saludo/{name:.*}/{apellido:.*}", Endpints.Saludo)
	// Ejemplo: http://localhost:8080/despedida?name=Benito&apellido=Juarez
	muxS.HandleFunc("/despedida", Endpints.Despedida)
	muxS.HandleFunc("/estructuras/variables", Endpints.Estructuras)

	ServerConfig()
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
