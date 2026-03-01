package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"proyectofinal/models"
	"strconv"
	"strings"
)

// Se utiliza la función GET carrito para mostrar el carrito
func CarritoViewHandler(w http.ResponseWriter, r *http.Request) {
	// Si no hay usuario, redirige a login
	user := CurrentUser(r)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Obtiene el carrito
	carrito, err := models.GetOrCreateCarrito(user.UsuarioID)
	if err != nil {
		http.Error(w, "No se pudo obtener el carrito", http.StatusInternalServerError)
		return
	}

	// Obtiene los items del carrito
	items, err := models.GetItemsByCarrito(carrito.CarritoID)
	if err != nil {
		http.Error(w, "No se pudieron obtener los items", http.StatusInternalServerError)
		return
	}
	// Se tiene conexió con la base de datos y se obtienen los items
	Render(w, r, []string{
		"templates/base.html",
		"templates/carrito.html",
	}, ViewData{
		Title: "Mi Carrito",
		Data:  items,
	})
}

// Se utiliza la función POST carrito/actualizar
func CarritoAgregarHandler(w http.ResponseWriter, r *http.Request) {
	// Si no hay usuario, redirige a login
	user := CurrentUser(r)
	if user == nil {
		if isAjax(r) {
			http.Error(w, "Necesitas iniciar sesión", http.StatusUnauthorized)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Procesar el formulario
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}
	productoID, _ := strconv.Atoi(r.FormValue("productoID"))
	cantidad, _ := strconv.Atoi(r.FormValue("cantidad"))
	if cantidad <= 0 {
		cantidad = 1
	}

	// Obtener el carrito
	carrito, err := models.GetOrCreateCarrito(user.UsuarioID)
	if err != nil {
		http.Error(w, "No se pudo obtener/crear carrito", http.StatusInternalServerError)
		return
	}

	// Agregar al carrito
	if err := models.AddToCarrito(carrito.CarritoID, productoID, cantidad); err != nil {
		log.Println("Error AddToCarrito:", err)
		http.Error(w, "No se pudo añadir al carrito", http.StatusInternalServerError)
		return
	}

	// Obtener nuevo contador
	count, _ := models.GetCartCountByUsuarioID(user.UsuarioID)

	if isAjax(r) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"ok":        true,
			"cartCount": count,
		})
		return
	}

	// flujo form: vuelve a la misma página o al carrito
	ref := r.Referer()
	if ref == "" {
		ref = "/carrito"
	}
	http.Redirect(w, r, ref, http.StatusSeeOther)
}

// Se utiliaza la función POST carrito/actualizar
func CarritoActualizarHandler(w http.ResponseWriter, r *http.Request) {
	// Si no hay usuario, redirige a login
	user := CurrentUser(r)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Procesar el formulario
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	// Actualizar la cantidad del item en el carrito
	itemID, _ := strconv.Atoi(r.FormValue("itemID"))
	cantidad, _ := strconv.Atoi(r.FormValue("cantidad"))
	if cantidad <= 0 {
		cantidad = 1
	}
	if err := models.UpdateItemCantidad(itemID, cantidad); err != nil {
		http.Error(w, "No se pudo actualizar la cantidad", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/carrito", http.StatusSeeOther)
}

// Se utiliza la función POST carrito/eliminar
func CarritoEliminarHandler(w http.ResponseWriter, r *http.Request) {
	// Si no hay usuario, redirige a login
	user := CurrentUser(r)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Procesar el formulario
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}

	// Eliminar el item del carrito
	itemID, _ := strconv.Atoi(r.FormValue("itemID"))
	if err := models.RemoveItem(itemID); err != nil {
		http.Error(w, "No se pudo eliminar el item", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/carrito", http.StatusSeeOther)
}

// Se usa la función POST carrito/vaciar
func CarritoVaciarHandler(w http.ResponseWriter, r *http.Request) {
	// Si no hay usuario, redirige a login
	user := CurrentUser(r)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	// Obtener el carrito
	carrito, err := models.GetOrCreateCarrito(user.UsuarioID)
	if err != nil {
		http.Error(w, "No se pudo obtener carrito", http.StatusInternalServerError)
		return
	}

	// Vaciar el carrito
	if err := models.VaciarCarrito(carrito.CarritoID); err != nil {
		http.Error(w, "No se pudo vaciar carrito", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/carrito", http.StatusSeeOther)
}

func isAjax(r *http.Request) bool {
	xrw := r.Header.Get("X-Requested-With")
	acc := r.Header.Get("Accept")
	return strings.EqualFold(xrw, "XMLHttpRequest") || strings.Contains(acc, "application/json")
}
