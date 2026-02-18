package main

import (
	Interactive "Curso_Golang/modules"
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func main() {

	fmt.Println("======================================== CURSO GOLANG  ============================================")
	fmt.Println("[] Ejercicios para comprender como funciona Go.")
	fmt.Println("")
	fmt.Println("Hola mundo!")
	fmt.Println("-----------------------------------------------------------------------------------------------------")

	// VARIABLES & CONSTANTES
	//variablesYConstantes()
	// PUNTEROS
	//puneros()
	// ESTRUCTURAS
	//estructuras()
	//INTERFACES
	//interfaces()
	// TIME
	//fechas()
	//ARGUMENTOS
	//argumentos()
	//LOGS
	//tipoLogs()
	//MODULOS
	//callModules()
}

func callModules() {
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("############################################ MODULOS ################################################")
	fmt.Println("")
	fmt.Println("[Este saludo se ejecuta desde un modulo] : ", Interactive.Saludar())
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("")
}

func tipoLogs() {
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("############################################### LOGS ################################################")
	fmt.Println("")
	//log.Fatal("Error Fatal --> Detiene la ejecución, después de esto no se ejecuta nada más...")
	log.Println("La aplicación esta iniciada ...")
	//log.Panic("pAnic: Esto es un mensaje de error de panic")

	fichero, execption := os.OpenFile("logs.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)

	if execption != nil {
		log.Fatal(execption)
	}
	//Cierra el fichero cunado termine la ejecución.
	defer fichero.Close()
	log.SetOutput(fichero)
	log.Printf("Error en linea %v", 1)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("")
}

func argumentos() {
	// Ejecutar: go run main.go -Nombre Aaron -Edad 10
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("######################################### ARGUMENTOS ################################################")
	fmt.Println("")
	nombre := flag.String("Nombre", "", "¿Cual es tu nombre?")
	edad := flag.Int("Edad", 1, "¿Cual es tu edad?")

	flag.Parse()

	fmt.Println("Tu nombre es: ", *nombre)
	fmt.Println("Tu edad es: ", *edad)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("")
}

func fechas() {
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("############################################### TIME ################################################")
	fmt.Println("")
	hoy := time.Now()
	fmt.Println("La hora actual es: ", hoy)
	fmt.Println("El año es: ", hoy.Year())
	fmt.Println("El mes en texto es: ", hoy.Month())
	fmt.Println("El mes digito es: ", int(hoy.Month()))
	fmt.Println("La hora es: ", hoy.Hour())
	fmt.Println("Los minutos son: ", hoy.Minute())
	fmt.Println("Los segundos son: ", hoy.Second())
	fmt.Println("La fecha de hoy mas 30 días: ", hoy.Add((time.Hour*24)*30))
	fmt.Println("La fecha de hoy menos 30 días: ", hoy.Add((time.Hour*24)*-30))
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("")
}

func interfaces() {
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("######################################### INTERFACES ################################################")
	fmt.Println("")
	miCoche := coche{}

	miCoche.avanzar()
	miCoche.frenar()
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("")
}

type coche struct{}

type automovil interface {
	avanzar() bool
	frenar() bool
}

func (*coche) frenar() bool {
	fmt.Println("Detenido...")
	return true
}

func (*coche) avanzar() bool {
	fmt.Println("En marcha...")
	return true
}

func estructuras() {

	type Usuario struct {
		idUsuario int
		nombre    string
		apellido  string
		correo    string
	}

	type Direccion struct {
		idDireccion  int
		calle        string
		colonia      string
		municipio    string
		ciudad       string
		pais         string
		codigoPostal string
	}

	type Cliente struct {
		idCLiente   int
		idUsuario   Usuario
		idDireccion Direccion
	}

	direccionCLiente := Direccion{
		idDireccion:  1,
		calle:        "Miguel Hidalgo #108",
		colonia:      "Vicente Segura",
		municipio:    "Ciudad de México",
		ciudad:       "CDMX",
		pais:         "México",
		codigoPostal: "00201",
	}

	usuarioJuan := Usuario{
		idUsuario: 1,
		nombre:    "Juan",
		apellido:  "Lopez",
		correo:    "juan.lopez@micorreo.com",
	}

	nuevoCliente := Cliente{
		idCLiente:   1,
		idUsuario:   usuarioJuan,
		idDireccion: direccionCLiente,
	}
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("######################################## ESTRUCTURAS ################################################")
	fmt.Println("")
	fmt.Println("[Esto es una estructura]: ", nuevoCliente)
	fmt.Println("")
}

func puneros() {
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("######################################## PUNTEROS ###################################################")
	fmt.Println("")
	puntero := "Aaron"
	println("El espacio en memoria de la variable: ", puntero, "es --> ", &puntero)
}

func variablesYConstantes() {

	var numeroEntero int = 66
	var numeroEntero8 int8 = 8
	var numeroEntero16 int16 = 16
	var numeroEntero32 int32 = 32
	var numeroEntero64 int64 = 64
	var numeroUint uint = 1
	var numeroUint8 uint8 = 8
	var numeroUint16 uint16 = 16
	var numeroUint32 uint32 = 32
	var numeroUint64 uint64 = 64

	var decimales32 float32 = 32.32
	var decimales64 float64 = 64.64
	var estadoBinario bool = true
	var texto string = "Var de tipo texto"
	textoGrande := `Esto es un curso de golang por lo que estamos probando una variable de tipo string para almacenar un texo demasiado grande, por que se decide almacenar de esta manera y demostrar la aseberación.`
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("##################################### ALGUNOS TIPOS DE DATO #########################################")
	fmt.Println("")
	fmt.Println("[ Esto es un texto corto ] : ", texto)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("[ Esto es un texto grande ] : ", textoGrande)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("[ Esto es una variable de tipo boleano (true/false) ] : ", estadoBinario)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("[ Esto es una variable de tipo float32 ] : ", decimales32)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("[ Esto es una variable de tipo float64 ] : ", decimales64)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("[ Esto es una variable de tipo int ] : ", numeroEntero)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("[ Esto es una variable de tipo int8 (-128 a 127) ] : ", numeroEntero8)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("[ Esto es una variable de tipo int16 (-2^15 a 2^15 -1) ] : ", numeroEntero16)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("[ Esto es una variable de tipo int32 (-2^31 a 2^31 -1) ] : ", numeroEntero32)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("[ Esto es una variable de tipo int64 (-2^63 a 2^63 -1) ] : ", numeroEntero64)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("[ Esto es una variable de tipo uint ] : ", numeroUint)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("[ Esto es una variable de tipo uint8 (0 a 255) ] : ", numeroUint8)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("[ Esto es una variable de tipo uint16 (0 a 2^16 -1) ] : ", numeroUint16)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("[ Esto es una variable de tipo uint32 (0 a 2^32 -1) ] : ", numeroUint32)
	fmt.Println("-----------------------------------------------------------------------------------------------------")
	fmt.Println("[ Esto es una variable de tipo uint64 (0 a 2^64 -1) ] : ", numeroUint64)
}
