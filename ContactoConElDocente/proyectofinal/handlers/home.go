package handlers

import (
	"net/http"
)

// HomeHandler muestra la página de inicio
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	// Mostrar la página de inicio
	Render(w, r, []string{
		"templates/base.html",
		"templates/index.html",
	}, ViewData{
		Title: "Inicio",
	})

}
