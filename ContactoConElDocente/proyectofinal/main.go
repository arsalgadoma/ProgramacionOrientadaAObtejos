package main

import (
	"fmt"
	"log"
	"net/http"
	"proyectofinal/db"
	"proyectofinal/handlers"

	"github.com/gorilla/mux"
)

func main() {

	db, err := db.Connect()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}

	defer db.Close()

	handlers.InitSessionStore()

	//Definir rutas de la aplicación
	//Vamos a generar las rutas utilziando un módulo
	r := mux.NewRouter()
	//Template para la base
	r.HandleFunc("/", handlers.HomeHandler)

	//Categorias
	r.HandleFunc("/categorias", handlers.CategoriasHandler)
	r.HandleFunc("/categorias/{id}", handlers.ProductosByCategoriaHandler)

	//Productos
	r.HandleFunc("/productos", handlers.ProductoHandler)

	//Registro
	r.HandleFunc("/registro", handlers.RegisterHandler).Methods("GET")
	r.HandleFunc("/registro", handlers.RegisterPostHandler).Methods("POST")

	// Login / Logout
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET")
	r.HandleFunc("/login", handlers.LoginPostHandler).Methods("POST")
	r.HandleFunc("/logout", handlers.LogoutHandler).Methods("POST")

	// Carrito
	r.HandleFunc("/carrito", handlers.CarritoViewHandler).Methods("GET")
	r.HandleFunc("/carrito/agregar", handlers.CarritoAgregarHandler).Methods("POST")
	r.HandleFunc("/carrito/actualizar", handlers.CarritoActualizarHandler).Methods("POST")
	r.HandleFunc("/carrito/eliminar", handlers.CarritoEliminarHandler).Methods("POST")
	r.HandleFunc("/carrito/vaciar", handlers.CarritoVaciarHandler).Methods("POST")

	// Orden
	r.HandleFunc("/orden/checkout", handlers.OrdenCheckoutHandler).Methods("GET")
	r.HandleFunc("/orden/crear", handlers.OrdenCrearHandler).Methods("POST")
	r.HandleFunc("/orden/{id}", handlers.OrdenDetalleHandler).Methods("GET")

	fmt.Println("Conexión exitosa a la base de datos")

	// iniciar el servidor HTTP

	if err := http.ListenAndServe(":8001", r); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
