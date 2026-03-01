package handlers

import (
	"log"
	"net/http"
	"proyectofinal/models"
)

// Lista de productos (página pública)
func ProductoHandler(w http.ResponseWriter, r *http.Request) {
	productos, err := models.GetAllProductos()
	if err != nil {
		log.Println("Error al obtener productos:", err)
		http.Error(w, "Error al obtener productos", http.StatusInternalServerError)
		return
	}
	log.Printf("Productos obtenidos: %d\n", len(productos))

	Render(w, r, []string{
		"templates/base.html",
		"templates/productos.html", // <- usa el template correcto
	}, ViewData{
		Title: "Productos",
		Data:  productos, // Se accede como .Data en el template
	})
}
