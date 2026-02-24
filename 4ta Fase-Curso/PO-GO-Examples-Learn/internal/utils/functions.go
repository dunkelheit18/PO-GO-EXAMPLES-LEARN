package utils

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("session-name"))

func CallMessage(response http.ResponseWriter, request *http.Request) (string, string) {

	cssSession := ""
	mensajeSession := ""
	session, _ := Store.Get(request, "flash-session")

	fmCss := session.Flashes("css")
	session.Save(request, response)

	if len(fmCss) > 0 {
		cssSession = fmCss[0].(string)
	}

	fmMensaje := session.Flashes("mensaje")
	session.Save(request, response)

	if len(fmMensaje) > 0 {
		mensajeSession = fmMensaje[0].(string)
	}

	return cssSession, mensajeSession
}

func CreateMessage(response http.ResponseWriter, request *http.Request, css string, mensaje string) {

	session, exception := Store.Get(request, "flash-session")

	if exception != nil {
		http.Error(response, exception.Error(), http.StatusInternalServerError)
		return
	}

	session.AddFlash(css, "css")
	session.AddFlash(mensaje, "mensaje")
	session.Save(request, response)
}
