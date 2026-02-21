package registro

import "regexp"

func ValidaNombre(nombre string) (bool, string) {

	mensaje := ""
	regexName := regexp.MustCompile(`^[a-zA-ZáéíóúÁÉÍÓÚñÑ\s'-]+$`)

	if nombre == "" {
		mensaje = "El nombre no se puede omitir ..."
		return true, mensaje
	}

	if len(nombre) < 2 {
		mensaje = "Ingresa un nombre valido"
		return true, mensaje
	}

	if !regexName.MatchString(nombre) {
		mensaje = "El nombre tiene caracteres no permitidos"
		return true, mensaje
	}

	return false, mensaje
}
