package registro

import (
	"regexp"
)

func ValidaNombre(nombre string) (bool, string) {

	mensaje := ""
	regexName := regexp.MustCompile(`^[a-zA-ZáéíóúÁÉÍÓÚñÑ\s'-]+$`)

	if nombre == "" {
		mensaje = "El nombre no se puede omitir ..."
		return true, mensaje
	}

	if len(nombre) < 3 {
		mensaje = "Ingresa un nombre valido."
		return true, mensaje
	}

	if !regexName.MatchString(nombre) {
		mensaje = "El nombre tiene caracteres no permitido."
		return true, mensaje
	}

	return false, mensaje
}

func ValidaEmail(email string) (bool, string) {
	mensaje := ""
	pattern := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)

	if email == "" {
		mensaje = "Favor de ingresar un correo."
		return true, mensaje
	}

	if !pattern.MatchString(email) {
		mensaje = "ingresar un correo valido."
		return true, mensaje
	}

	return false, mensaje
}

func ValidaPassword(psw string) (bool, string) {

	message := ""
	if psw == "" {
		message = "La contraseña es obligatoria."
		return true, message
	}

	if len(psw) < 8 {
		message = "La contraseña debe tener minimo 8 caracteres."
		return true, message
	}

	return false, message
}

func ValidaDatos(data map[string]string) (bool, string) {

	msg := ""
	name := data["nombre"]
	email := data["email"]
	passw := data["password"]

	errName, msgName := ValidaNombre(name)
	errEmail, msgEmail := ValidaEmail(email)
	errPass, msgPass := ValidaPassword(passw)

	if errName {
		msg = msgName
	}

	if errEmail {
		msg = msg + " | " + msgEmail
	}

	if errPass {
		msg = msg + " | " + msgPass
	}

	return (msg != ""), msg
}
