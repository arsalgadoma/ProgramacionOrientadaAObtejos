package models

import (
	"log"
	"proyectofinal/db"
)

type Usuario struct {
	UsuarioID int
	Nombre    string
	Email     string
	Password  string
}

// Obtener usuario por email (texto plano)
func GetUsuarioByEmail(email string) (Usuario, error) {
	var usuario Usuario

	DB, err := db.Connect()
	if err != nil {
		return usuario, err
	}
	defer DB.Close()

	query := "SELECT UsuarioID, Nombre, Email, Password FROM usuario WHERE Email = ?"
	err = DB.QueryRow(query, email).Scan(
		&usuario.UsuarioID,
		&usuario.Nombre,
		&usuario.Email,
		&usuario.Password,
	)

	if err != nil {
		log.Println("Usuario no encontrado:", err)
		return usuario, err
	}

	return usuario, nil
}

// Crear usuario nuevo (texto plano)
func CreateUsuario(nombre, email, password string) error {

	DB, err := db.Connect()
	if err != nil {
		log.Println("Error al conectar BD:", err)
		return err
	}
	defer DB.Close()

	query := "INSERT INTO usuario (Nombre, Email, Password) VALUES (?, ?, ?)"
	_, err = DB.Exec(query, nombre, email, password)
	if err != nil {
		log.Println("Error al insertar usuario:", err)
		return err
	}

	return nil
}
