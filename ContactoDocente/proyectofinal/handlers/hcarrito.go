package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"proyectofinal/models"
	"strconv"
	"strings"
)

// GET /carrito
func CarritoViewHandler(w http.ResponseWriter, r *http.Request) {
	user := CurrentUser(r)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	carrito, err := models.GetOrCreateCarrito(user.UsuarioID)
	if err != nil {
		http.Error(w, "No se pudo obtener el carrito", http.StatusInternalServerError)
		return
	}
	items, err := models.GetItemsByCarrito(carrito.CarritoID)
	if err != nil {
		http.Error(w, "No se pudieron obtener los items", http.StatusInternalServerError)
		return
	}

	Render(w, r, []string{
		"templates/base.html",
		"templates/carrito.html",
	}, ViewData{
		Title: "Mi Carrito",
		Data:  items,
	})
}

// POST /carrito/agregar (form o AJAX)
// body: productoID, cantidad
func CarritoAgregarHandler(w http.ResponseWriter, r *http.Request) {
	user := CurrentUser(r)
	if user == nil {
		// Si es AJAX, retorna 401; si es form, redirige a login
		if isAjax(r) {
			http.Error(w, "Necesitas iniciar sesión", http.StatusUnauthorized)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}
	productoID, _ := strconv.Atoi(r.FormValue("productoID"))
	cantidad, _ := strconv.Atoi(r.FormValue("cantidad"))
	if cantidad <= 0 {
		cantidad = 1
	}

	carrito, err := models.GetOrCreateCarrito(user.UsuarioID)
	if err != nil {
		http.Error(w, "No se pudo obtener/crear carrito", http.StatusInternalServerError)
		return
	}

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

// POST /carrito/actualizar
// body: itemID, cantidad
func CarritoActualizarHandler(w http.ResponseWriter, r *http.Request) {
	user := CurrentUser(r)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}
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

// POST /carrito/eliminar
// body: itemID
func CarritoEliminarHandler(w http.ResponseWriter, r *http.Request) {
	user := CurrentUser(r)
	if user == nil {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Solicitud inválida", http.StatusBadRequest)
		return
	}
	itemID, _ := strconv.Atoi(r.FormValue("itemID"))
	if err := models.RemoveItem(itemID); err != nil {
		http.Error(w, "No se pudo eliminar el item", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/carrito", http.StatusSeeOther)
}

// POST /carrito/vaciar
func CarritoVaciarHandler(w http.ResponseWriter, r *http.Request) {
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
	if err := models.VaciarCarrito(carrito.CarritoID); err != nil {
		http.Error(w, "No se pudo vaciar carrito", http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/carrito", http.StatusSeeOther)
}

// (Opcional) GET /api/carrito/count -> devuelve {cartCount:n}
func CarritoCountAPI(w http.ResponseWriter, r *http.Request) {
	user := CurrentUser(r)
	count := 0
	if user != nil {
		c, _ := models.GetCartCountByUsuarioID(user.UsuarioID)
		count = c
	}
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(map[string]int{"cartCount": count})
}

func isAjax(r *http.Request) bool {
	xrw := r.Header.Get("X-Requested-With")
	acc := r.Header.Get("Accept")
	return strings.EqualFold(xrw, "XMLHttpRequest") || strings.Contains(acc, "application/json")
}
