package handlers

import (
	"log"
	"net/http"
	"proyectofinal/models"
)

// Se enlistan todos los productos
func ProductoHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener todos los productos de la base de datos
	productos, err := models.GetAllProductos()
	if err != nil {
		log.Println("Error al obtener productos:", err)
		http.Error(w, "Error al obtener productos", http.StatusInternalServerError)
		return
	}
	log.Printf("Productos obtenidos: %d\n", len(productos))

	Render(w, r, []string{
		"templates/base.html",
		"templates/productos.html",
	}, ViewData{
		Title: "Productos",
		Data:  productos,
	})
}
