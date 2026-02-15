package models

type Orden struct {
	OrdenID int
	Fecha   string
	Total   float64
	Items   []ItemOrden
}
