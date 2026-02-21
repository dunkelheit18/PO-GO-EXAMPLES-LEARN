package utils

import (
	"net/http"

	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte("session-name"))

func CallMessage(response http.ResponseWriter, request *http.Request) {

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
