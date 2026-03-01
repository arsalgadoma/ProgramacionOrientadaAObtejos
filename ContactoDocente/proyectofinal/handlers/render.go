package handlers

import (
	"bytes"
	"html/template"
	"log"
	"net/http"
	"proyectofinal/models"
)

type ViewData struct {
	Title     string
	Usuario   *models.Usuario
	Data      interface{}
	CartCount int
}

// Render: base + vista, inyecta usuario/contador, ejecuta en buffer y escribe una sola vez.
func Render(w http.ResponseWriter, r *http.Request, files []string, data ViewData) {
	tpl, err := template.ParseFiles(files...)
	if err != nil {
		log.Printf("Error al cargar templates %v: %v", files, err)
		http.Error(w, "Error al cargar vistas", http.StatusInternalServerError)
		return
	}

	// Inyectar usuario
	if data.Usuario == nil {
		data.Usuario = CurrentUser(r)
	}
	// Inyectar contador del carrito si hay usuario
	if data.Usuario != nil {
		if c, err := models.GetCartCountByUsuarioID(data.Usuario.UsuarioID); err == nil {
			data.CartCount = c
		}
	}

	// Ejecutar en buffer para evitar escribir cabezales dos veces
	var buf bytes.Buffer
	if err := tpl.ExecuteTemplate(&buf, "base", data); err != nil {
		log.Printf("Error al ejecutar template base: %v", err)
		http.Error(w, "Error al renderizar vista", http.StatusInternalServerError)
		return
	}
	if _, err := buf.WriteTo(w); err != nil {
		log.Println("Error al escribir respuesta:", err)
	}
}
