package models

import (
	"encoding/json"
	"fmt"
	"time"

	"proyectofinal/db"
)

type Orden struct {
	OrdenID   int
	Total     float64
	Fecha     time.Time
	UsuarioID int
	Detalle   string // JSON/TEXT
}

type OrdenItemDTO struct {
	ProductoID     int     `json:"productoId"`
	Nombre         string  `json:"nombre"`
	PrecioUnitario float64 `json:"precioUnitario"`
	Cantidad       int     `json:"cantidad"`
	Subtotal       float64 `json:"subtotal"`
}

// CREA orden directa guardando snapshot en Detalle (JSON) y vacía carrito
func CrearOrdenDirectaDesdeCarrito(usuarioID int) (int, error) {
	DB, err := db.Connect()
	if err != nil {
		return 0, err
	}
	defer DB.Close()

	tx, err := DB.Begin()
	if err != nil {
		return 0, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback()
		}
	}()

	carrito, err := GetOrCreateCarrito(usuarioID)
	if err != nil {
		return 0, err
	}
	items, err := GetItemsByCarrito(carrito.CarritoID)
	if err != nil {
		return 0, err
	}
	if len(items) == 0 {
		return 0, fmt.Errorf("carrito vacío")
	}

	total := 0.0
	snap := make([]OrdenItemDTO, 0, len(items))
	for _, it := range items {
		total += float64(it.Cantidad) * it.Precio
		snap = append(snap, OrdenItemDTO{
			ProductoID:     it.ProductoID,
			Nombre:         it.Nombre,
			PrecioUnitario: it.Precio,
			Cantidad:       it.Cantidad,
			Subtotal:       float64(it.Cantidad) * it.Precio,
		})
	}

	js, err := json.Marshal(snap)
	if err != nil {
		return 0, err
	}

	res, err := tx.Exec(`
        INSERT INTO orden (Total, Fecha, UsuarioID, Detalle)
        VALUES (?, ?, ?, ?)`,
		total, time.Now(), usuarioID, string(js),
	)
	if err != nil {
		return 0, err
	}
	id64, _ := res.LastInsertId()
	ordenID := int(id64)

	if _, err = tx.Exec(`DELETE FROM itemcarrito WHERE CarritoID = ?`, carrito.CarritoID); err != nil {
		return 0, err
	}

	if err = tx.Commit(); err != nil {
		return 0, err
	}
	return ordenID, nil
}

func GetOrdenByID(ordenID int) (Orden, error) {
	var o Orden
	DB, err := db.Connect()
	if err != nil {
		return o, err
	}
	defer DB.Close()

	err = DB.QueryRow(`
        SELECT OrdenID, Total, Fecha, UsuarioID, Detalle
        FROM orden
        WHERE OrdenID = ?`,
		ordenID,
	).Scan(&o.OrdenID, &o.Total, &o.Fecha, &o.UsuarioID, &o.Detalle)

	var items []OrdenItemDTO
	if err := json.Unmarshal([]byte(o.Detalle), &items); err != nil {
		return o, err
	}
	return o, err
}

func GetOrdenDetalleItems(o Orden) ([]OrdenItemDTO, error) {
	if o.Detalle == "" {
		return []OrdenItemDTO{}, nil
	}
	var items []OrdenItemDTO
	if err := json.Unmarshal([]byte(o.Detalle), &items); err != nil {
		return nil, err
	}
	return items, nil
}
