package handlers

import (
	"log"
	"net/http"
	"proyectofinal/models"
)

// Mostrar login por metodo GET
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, r, []string{
		"templates/base.html",
		"templates/login.html",
	}, ViewData{
		Title: "Iniciar sesión",
		// Usuario y CartCount los inyecta Render()
	})
}

// Se procesa el logeo del usuario
func LoginPostHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar formulario", http.StatusBadRequest)
		return
	}

	email := r.FormValue("email")
	password := r.FormValue("password")

	usuario, err := models.GetUsuarioByEmail(email)
	if err != nil {
		http.Error(w, "Correo o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	// Se compara la contraseña del usuario con la del formulario
	if usuario.Password != password {
		http.Error(w, "Correo o contraseña incorrectos", http.StatusUnauthorized)
		return
	}

	// Crear sesión
	sess, _ := getSession(r)
	sess.Values["userID"] = usuario.UsuarioID
	sess.Values["nombre"] = usuario.Nombre
	sess.Values["email"] = usuario.Email
	_ = sess.Save(r, w)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Mostrar registro para la creación de un nuevo usuario
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	Render(w, r, []string{
		"templates/base.html",
		"templates/registro.html",
	}, ViewData{
		Title: "Crear cuenta",
	})
}

// Se crea un nuevo usuario
func RegisterPostHandler(w http.ResponseWriter, r *http.Request) {
	// Procesar el formulario
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error al procesar formulario", http.StatusBadRequest)
		return
	}

	nombre := r.FormValue("nombre")
	email := r.FormValue("email")
	password := r.FormValue("password")

	// Verificar que todos los campos sean obligatorios
	if nombre == "" || email == "" || password == "" {
		http.Error(w, "Todos los campos son obligatorios", http.StatusBadRequest)
		return
	}

	// Crear el nuevo usuario
	if err := models.CreateUsuario(nombre, email, password); err != nil {
		log.Println("Error al crear usuario:", err)
		http.Error(w, "Error al crear usuario", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Cerrar sesión
func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	sess, err := getSession(r)
	if err == nil {
		sess.Options.MaxAge = -1
		_ = sess.Save(r, w)
	}
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
