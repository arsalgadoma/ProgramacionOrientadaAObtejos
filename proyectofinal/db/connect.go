package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

// Connect crea y valida la conexión a MySQL
func Connect() (*sql.DB, error) {

	// Cargar archivo .env
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("error cargando archivo .env: %w", err)
	}

	// Obtener variables de entorno
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")

	// Validar que no estén vacías
	if user == "" || password == "" || host == "" || port == "" || name == "" {
		return nil, fmt.Errorf("Faltan variables de entorno para la base de datos")
	}

	// Crear DSN
	dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		user, password, host, port, name,
	)

	// Abrir conexión
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	// Verificar conexión real
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
