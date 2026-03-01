package handlers

import (
	"log"
	"net/http"
	"proyectofinal/models"
	"strconv"

	"github.com/gorilla/mux"
)

// CategoriasHandler muestra todas las categorías
func CategoriasHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener todas las categorías de la base de datos
	categorias, err := models.GetAllCategorias()
	if err != nil {
		log.Println("Error al obtener categorías:", err)
		http.Error(w, "Error al obtener categorías", http.StatusInternalServerError)
		return
	}

	Render(w, r, []string{
		"templates/base.html",
		"templates/categorias.html",
	}, ViewData{
		Title: "Categorías",
		Data:  categorias,
	})
}

// ProductosByCategoriaHandler muestra los productos de una categoría
func ProductosByCategoriaHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener el ID de la categoría
	vars := mux.Vars(r)
	categoriaIDStr := vars["id"]
	categoriaID, err := strconv.Atoi(categoriaIDStr)
	if err != nil {
		log.Println("ID de categoría inválido: ", err)
		http.Error(w, "ID de categoría inválido", http.StatusBadRequest)
		return
	}

	// Obtener los productos de la categoría
	productos, err := models.GetProductosByCategoria(strconv.Itoa(categoriaID))
	if err != nil {
		log.Println("Error al obtener productos por categoría: ", err)
		http.Error(w, "Error al obtener productos por categoría", http.StatusInternalServerError)
		return
	}

	Render(w, r, []string{
		"templates/base.html",
		"templates/categoria_producto.html",
	}, ViewData{
		Title: "Productos",
		Data:  productos, // Se accederá como .Data en la vista
	})
}
