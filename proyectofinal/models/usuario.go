package models

type Usuario struct {
	UsuarioID int
	Nombre    string
	Email     string
	Carrito   *Carrito
	Ordenes   []Orden
}
