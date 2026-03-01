package handlers

import (
	"log"
	"net/http"
	"proyectofinal/models"
	"strconv"

	"github.com/gorilla/mux"
)

// GET /orden/checkout
func OrdenCheckoutHandler(w http.ResponseWriter, r *http.Request) {
	user := CurrentUser(r)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	carrito, err := models.GetOrCreateCarrito(user.UsuarioID)
	if err != nil {
		http.Error(w, "No se pudo obtener carrito", http.StatusInternalServerError)
		return
	}
	items, err := models.GetItemsByCarrito(carrito.CarritoID)
	if err != nil {
		http.Error(w, "No se pudieron obtener items", http.StatusInternalServerError)
		return
	}
	total, _ := models.GetCartTotalByCarritoID(carrito.CarritoID)

	Render(w, r, []string{
		"templates/base.html",
		"templates/checkout.html",
	}, ViewData{
		Title: "Confirmar compra",
		Data: struct {
			Items []models.ItemCarrito
			Total float64
		}{Items: items, Total: total},
	})
}

// POST /orden/crear
func OrdenCrearHandler(w http.ResponseWriter, r *http.Request) {
	user := CurrentUser(r)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	ordenID, err := models.CrearOrdenDirectaDesdeCarrito(user.UsuarioID)
	if err != nil {
		log.Println("[OrdenCrear] error:", err)
		http.Error(w, "No se pudo crear la orden", http.StatusInternalServerError)
		return
	}
	log.Printf("[OrdenCrear] ok user=%d ordenID=%d", user.UsuarioID, ordenID)

	http.Redirect(w, r, "/orden/"+strconv.Itoa(ordenID), http.StatusSeeOther)
}

// GET /orden/{id}
func OrdenDetalleHandler(w http.ResponseWriter, r *http.Request) {
	user := CurrentUser(r)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	vars := mux.Vars(r)
	ordenID, _ := strconv.Atoi(vars["id"])

	orden, err := models.GetOrdenByID(ordenID)
	log.Printf("[OrdenDetalle] user=%d ordenID=%d err=%v owner=%d",
		user.UsuarioID, ordenID, err, orden.UsuarioID)

	if err != nil || orden.UsuarioID != user.UsuarioID {
		http.Error(w, "Orden no encontrada", http.StatusNotFound)
		return
	}

	items, err := models.GetOrdenDetalleItems(orden)
	if err != nil {
		http.Error(w, "Error al leer detalle de orden", http.StatusInternalServerError)
		return
	}

	Render(w, r, []string{"templates/base.html", "templates/orden_detalle.html"}, ViewData{
		Title: "Orden #" + strconv.Itoa(ordenID),
		Data: struct {
			Orden models.Orden
			Items []models.OrdenItemDTO
		}{Orden: orden, Items: items},
	})
}
