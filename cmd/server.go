package main

import (
	"poll-app/db"
	"poll-app/pkg/poll"

	"github.com/gin-gonic/gin"
)

func main() {
	// Configurar el modo release para producción
	gin.SetMode(gin.ReleaseMode)

	// Conectar a la base de datos
	db.Connect()

	// Crear una instancia del router de Gin
	r := gin.Default()

	// Rutas
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Bienvenido a la aplicación de encuestas!")
	})
	r.GET("/polls", poll.GetPollsHandler)
	r.POST("/polls/vote/:id", poll.VoteHandler)

	// Iniciar servidor en el puerto 8080
	r.Run(":8080")
}
