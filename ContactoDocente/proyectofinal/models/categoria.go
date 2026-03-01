package models

import (
	"log"
	"proyectofinal/db"
)

type Categoria struct {
	CategoriaID int
	Nombre      string
	Descripcion string
}

// Función para obtener todas las categorías de la base de datos
func GetAllCategorias() ([]Categoria, error) {
	var categorias []Categoria
	DB, err := db.Connect()
	if err != nil {
		log.Println("Error al conectar a la base de datos: ", err)
		return categorias, err
	}
	defer DB.Close()

	//Consultar dentro de la base de datos y obtener las categorías
	filas, err := DB.Query("SELECT CategoriaID, Nombre, Descripcion FROM categoria")
	if err != nil {
		log.Println("Error al consultar categorías: ", err)
		return categorias, err
	}
	defer filas.Close()

	//Formatear la salida de la consultya y agregar a la lista de categorías
	for filas.Next() {
		var categoria Categoria
		err := filas.Scan(&categoria.CategoriaID, &categoria.Nombre, &categoria.Descripcion)
		if err != nil {
			log.Println("Error al escanear fila de categorías: ", err)
			return categorias, err
		}
		categorias = append(categorias, categoria)
	}

	// Verificar si hubo errores durante la iteración de las filas
	if err = filas.Err(); err != nil {
		log.Println("Error durante la iteración de filas: ", err)
		return categorias, err
	}

	log.Println("Categorías obtenidas exitosamente", categorias)
	return categorias, nil
}
