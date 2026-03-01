package handlers

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"proyectofinal/models"

	"github.com/gorilla/sessions"
)

var sessionStore *sessions.CookieStore

// Inicializar el gestor de sesiones
func InitSessionStore() {
	secret := os.Getenv("SESSION_KEY")
	if secret == "" {
		secret = "cambia-esta-clave-super-secreta"
	}
	sessionStore = sessions.NewCookieStore([]byte(secret))
	sessionStore.Options = &sessions.Options{
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		// Secure: true, // habilitar solo si usas HTTPS
		MaxAge: 60 * 60 * 24 * 7,
	}
}

// Obtener la sesión actual
func getSession(r *http.Request) (*sessions.Session, error) {
	return sessionStore.Get(r, "app_session")
}

// Obtener el usuario actual
func CurrentUser(r *http.Request) *models.Usuario {
	sess, err := getSession(r)
	if err != nil || sess.Values["userID"] == nil {
		return nil
	}
	var id int
	switch v := sess.Values["userID"].(type) {
	case int:
		id = v
	case int64:
		id = int(v)
	case float64:
		id = int(v)
	case string:
		if n, err := strconv.Atoi(v); err == nil {
			id = n
		}
	default:
		id = 0
	}
	if id == 0 {
		return nil
	}
	nombre := fmt.Sprint(sess.Values["nombre"])
	email := fmt.Sprint(sess.Values["email"])
	return &models.Usuario{UsuarioID: id, Nombre: nombre, Email: email}
}
