package handlers

import (
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {

	Render(w, r, []string{
		"templates/base.html",
		"templates/index.html",
	}, ViewData{
		Title: "Inicio",
		// Usuario y CartCount los inyecta Render()
	})

}
