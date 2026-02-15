// ======================================
// ARCHIVO: main.go
// RESPONSABILIDAD: Punto de entrada del sistema
// ======================================

package main

import (
	"log"
	"net/http"

	"proyectofinal/db"
)

func main() {

	// ==============================
	// Conectar a la base de datos
	// ==============================
	database, err := db.Connect()
	if err != nil {
		log.Fatal("Error al conectar a la base de datos: ", err)
	}
	defer database.Close()

	log.Println("Conexión exitosa a la base de datos")

	// ==============================
	// Iniciar servidor
	// ==============================
	if err := http.ListenAndServe(":8000", nil); err != nil {
		log.Fatal("Error al iniciar el servidor: ", err)
	}
}
