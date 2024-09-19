package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb" // Importación del driver
)

var DB *sql.DB

func Connect() {
	// Configuración de conexión
	connString := "sqlserver://sa:12345678@localhost:1433?database=go"

	// Abrir la conexión
	var err error
	DB, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error creando la conexión: ", err)
	}

	// Verificar la conexión
	err = DB.Ping()
	if err != nil {
		log.Fatal("Error en la conexión con la base de datos: ", err)
	} else {
		fmt.Println("Conexión exitosa con SQL Server")
	}
}
