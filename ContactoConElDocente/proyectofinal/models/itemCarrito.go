package models

import (
	"database/sql"
	"log"
	"proyectofinal/db"
)

type ItemCarrito struct {
	ItemCarID  int
	Cantidad   int
	CarritoID  int
	ProductoID int

	// Campos enriquecidos para mostrar
	Nombre   string
	Precio   float64
	Subtotal float64
}

func GetItemsByCarrito(carritoID int) ([]ItemCarrito, error) {
	// Obtener todos los items del carrito
	items := []ItemCarrito{}

	DB, err := db.Connect()
	if err != nil {
		return items, err
	}
	defer DB.Close()

	rows, err := DB.Query(`
        SELECT ic.ItemCarID, ic.Cantidad, ic.CarritoID, ic.ProductoID,
               p.Nombre, p.Precio
        FROM itemcarrito ic
        INNER JOIN producto p ON p.ProductoID = ic.ProductoID
        WHERE ic.CarritoID = ?`,
		carritoID,
	)
	if err != nil {
		return items, err
	}
	defer rows.Close()

	// Iterar sobre los items y agregarlos a la lista
	for rows.Next() {
		var it ItemCarrito
		if err := rows.Scan(
			&it.ItemCarID,
			&it.Cantidad,
			&it.CarritoID,
			&it.ProductoID,
			&it.Nombre,
			&it.Precio,
		); err != nil {
			log.Println("Error scan item carrito:", err)
			return items, err
		}
		it.Subtotal = float64(it.Cantidad) * it.Precio
		items = append(items, it)
	}
	return items, rows.Err()
}

// Agregar un item al carrito de un usuario
func AddToCarrito(carritoID, productoID, cantidad int) error {
	DB, err := db.Connect()
	if err != nil {
		return err
	}
	defer DB.Close()

	// Verificar si el item ya existe en el carrito
	var itemID sql.NullInt64
	err = DB.QueryRow(`
        SELECT ItemCarID
        FROM itemcarrito
        WHERE CarritoID = ? AND ProductoID = ?`,
		carritoID, productoID,
	).Scan(&itemID)

	// Si no existe, agregarlo al carrito
	if err == sql.ErrNoRows || !itemID.Valid {
		_, err = DB.Exec(`
            INSERT INTO itemcarrito (Cantidad, CarritoID, ProductoID)
            VALUES (?, ?, ?)`,
			cantidad, carritoID, productoID,
		)
		return err
	}
	if err != nil {
		return err
	}

	// Si ya existe, actualizar la cantidad
	_, err = DB.Exec(`
        UPDATE itemcarrito
        SET Cantidad = Cantidad + ?
        WHERE ItemCarID = ?`,
		cantidad, itemID.Int64,
	)
	return err
}

// Actualizar la cantidad de un item en el carrito
func UpdateItemCantidad(itemCarID, cantidad int) error {
	DB, err := db.Connect()
	if err != nil {
		return err
	}
	defer DB.Close()

	_, err = DB.Exec(`
        UPDATE itemcarrito
        SET Cantidad = ?
        WHERE ItemCarID = ?`,
		cantidad, itemCarID,
	)
	return err
}

// Eliminar un item del carrito
func RemoveItem(itemCarID int) error {
	DB, err := db.Connect()
	if err != nil {
		return err
	}
	defer DB.Close()

	_, err = DB.Exec(`DELETE FROM itemcarrito WHERE ItemCarID = ?`, itemCarID)
	return err
}

// Eliminar todos los items del carrito
func VaciarCarrito(carritoID int) error {
	DB, err := db.Connect()
	if err != nil {
		return err
	}
	defer DB.Close()

	_, err = DB.Exec(`DELETE FROM itemcarrito WHERE CarritoID = ?`, carritoID)
	return err
}
