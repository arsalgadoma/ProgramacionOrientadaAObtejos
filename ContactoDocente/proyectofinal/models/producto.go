package models

import (
	"database/sql"
	"log"
	"proyectofinal/db"
)

type Producto struct {
	ProductoID  int
	Nombre      string
	Descripcion string
	Precio      float64
	Stock       int
	CategoriaID string
}

// Función para obtener todos los productos de la base de datos
func GetAllProductos() ([]Producto, error) {
	var productos []Producto
	DB, err := db.Connect()
	if err != nil {
		log.Println("Error al conectar a la base de datos: ", err)
		return productos, err
	}
	defer DB.Close()

	//Consultar dentro de la base de datos y obtener los productos
	filas, err := DB.Query("SELECT ProductoID, Nombre, Descripcion, Precio, Stock, CategoriaID FROM producto")
	if err != nil {
		log.Println("Error al consultar productos: ", err)
		return productos, err
	}
	defer filas.Close()

	//Formatear la salida de la consultya y agregar a la lista de productos
	for filas.Next() {
		var producto Producto
		err := filas.Scan(&producto.ProductoID, &producto.Nombre, &producto.Descripcion, &producto.Precio, &producto.Stock, &producto.CategoriaID)
		if err != nil {
			log.Println("Error al escanear fila de productos: ", err)
			return productos, err
		}
		productos = append(productos, producto)
	}

	// Verificar si hubo errores durante la iteración de las filas
	if err = filas.Err(); err != nil {
		log.Println("Error durante la iteración de filas: ", err)
		return productos, err
	}
	log.Println("Productos obtenidos exitosamente", productos)
	return productos, nil
}

// Función para obtener un producto por su ID
func GetProductoByID(id int) (Producto, error) {
	//Inicializar el producto a retornar
	var producto Producto
	DB, err := db.Connect()
	if err != nil {
		log.Println("Error al conectar a la base de datos: ", err)
		return producto, err
	}
	defer DB.Close()

	//Consultar dentro de la base de datos y obtener el producto por su ID
	smtp, err := DB.Prepare("SELECT * FROM producto WHERE ProductoID = ?")
	if err != nil {
		log.Println("Error al preparar la consulta", err)
		return producto, err
	}
	defer smtp.Close()

	//ejecutar la consulta
	row := smtp.QueryRow(id)

	//Completar consulta
	err = row.Scan(&producto.ProductoID, &producto.Nombre, &producto.Descripcion, &producto.Precio, &producto.Stock, &producto.CategoriaID)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Printf("No se encontró el producto con ID %d", id)
			return producto, nil // Retornar un producto vacío sin error
		}
		log.Println("Error al escanear fila de producto: ", err)
		return producto, err
	}
	log.Println("Producto obtenido exitosamente: ", producto)
	return producto, nil
}

// Función para obtener productos por categoría
func GetProductosByCategoria(categoriaID string) ([]Producto, error) {
	var productos []Producto
	DB, err := db.Connect()
	if err != nil {
		log.Println("Error al conectar a la base de datos: ", err)
		return productos, err
	}
	defer DB.Close()

	//Consultar dentro de la base de datos y obtener los productos por categoría
	filas, err := DB.Query("SELECT * FROM producto WHERE CategoriaID = ?", categoriaID)
	if err != nil {
		log.Println("Error al consultar productos por categoría: ", err)
		return productos, err
	}
	defer filas.Close()

	//Formatear la salida de la consultya y agregar a la lista de productos
	for filas.Next() {
		var producto Producto
		err := filas.Scan(&producto.ProductoID, &producto.Nombre, &producto.Descripcion, &producto.Precio, &producto.Stock, &producto.CategoriaID)
		if err != nil {
			log.Println("Error al escanear fila de productos: ", err)
			return productos, err
		}
		productos = append(productos, producto)
	}

	// Verificar si hubo errores durante la iteración de las filas
	if err = filas.Err(); err != nil {
		log.Println("Error durante la iteración de filas: ", err)
		return productos, err
	}
	log.Println("Productos obtenidos exitosamente por categoría", categoriaID, productos)
	return productos, nil
}

// Función para actualizar el stock de un producto
func ActualizarStock(productoID int, nuevoStock int) error {
	DB, err := db.Connect()
	if err != nil {
		log.Println("Error al conectar a la base de datos: ", err)
		return err
	}
	defer DB.Close()

	// Actualizar el stock del producto
	_, err = DB.Exec("UPDATE producto SET Stock = ? WHERE ProductoID = ?", nuevoStock, productoID)
	if err != nil {
		log.Println("Error al actualizar el stock del producto: ", err)
		return err
	}
	log.Println("Stock actualizado exitosamente para el producto con ID:", productoID)
	return nil
}
