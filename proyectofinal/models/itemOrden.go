package models

type ItemOrden struct {
	ProductoID int
	Cantidad   int
	Precio     float64
	Producto   Producto
}
