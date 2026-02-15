package models

type Producto struct {
	ProductoID  int
	Nombre      string
	Descripcion string
	Precio      float64
	Stock       int
	Categoria   Categoria
}
