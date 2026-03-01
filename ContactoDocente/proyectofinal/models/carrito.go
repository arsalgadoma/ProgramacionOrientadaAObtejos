package models

import (
	"database/sql"
	"log"
	"proyectofinal/db"
)

type Carrito struct {
	CarritoID int
	UsuarioID int
}

func GetCarritoByUsuarioID(usuarioID int) (Carrito, error) {
	var c Carrito
	DB, err := db.Connect()
	if err != nil {
		return c, err
	}
	defer DB.Close()

	err = DB.QueryRow(`
        SELECT CarritoID, UsuarioID
        FROM carrito
        WHERE UsuarioID = ?`,
		usuarioID,
	).Scan(&c.CarritoID, &c.UsuarioID)

	return c, err
}

func CreateCarrito(usuarioID int) (Carrito, error) {
	var c Carrito
	DB, err := db.Connect()
	if err != nil {
		return c, err
	}
	defer DB.Close()

	res, err := DB.Exec(`INSERT INTO carrito (UsuarioID) VALUES (?)`, usuarioID)
	if err != nil {
		return c, err
	}
	id, _ := res.LastInsertId()
	c.CarritoID = int(id)
	c.UsuarioID = usuarioID
	return c, nil
}

func GetOrCreateCarrito(usuarioID int) (Carrito, error) {
	c, err := GetCarritoByUsuarioID(usuarioID)
	if err == nil {
		return c, nil
	}
	if err != sql.ErrNoRows {
		return c, err
	}
	return CreateCarrito(usuarioID)
}

// Suma de cantidades de todos los items del carrito del usuario
func GetCartCountByUsuarioID(usuarioID int) (int, error) {
	DB, err := db.Connect()
	if err != nil {
		return 0, err
	}
	defer DB.Close()

	var count sql.NullInt64
	err = DB.QueryRow(`
        SELECT COALESCE(SUM(ic.Cantidad), 0) AS total
        FROM carrito c
        LEFT JOIN itemcarrito ic ON ic.CarritoID = c.CarritoID
        WHERE c.UsuarioID = ?`,
		usuarioID,
	).Scan(&count)
	if err != nil {
		return 0, err
	}
	if !count.Valid {
		return 0, nil
	}
	return int(count.Int64), nil
}

// Total del carrito (sumatoria Cantidad * Precio)
func GetCartTotalByCarritoID(carritoID int) (float64, error) {
	DB, err := db.Connect()
	if err != nil {
		return 0, err
	}
	defer DB.Close()

	var total sql.NullFloat64
	err = DB.QueryRow(`
        SELECT COALESCE(SUM(ic.Cantidad * p.Precio), 0)
        FROM itemcarrito ic
        INNER JOIN producto p ON p.ProductoID = ic.ProductoID
        WHERE ic.CarritoID = ?`,
		carritoID,
	).Scan(&total)
	if err != nil {
		return 0, err
	}
	if !total.Valid {
		return 0, nil
	}
	return total.Float64, nil
}

func DebugLogCarrito(usuarioID int) {
	count, err := GetCartCountByUsuarioID(usuarioID)
	if err != nil {
		log.Println("Error obteniendo count carrito:", err)
		return
	}
	log.Printf("[Debug] Usuario %d -> items en carrito: %d\n", usuarioID, count)
}
