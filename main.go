package main

import (
	"context"
	"log"

	service "examen_final_febrero_golang_P1/Services"
	handlers "examen_final_febrero_golang_P1/handlers"
	middlewares "examen_final_febrero_golang_P1/middlewares"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()

	// Conexión MongoDB
	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.TODO())

	collection := client.Database("examen").Collection("entidades")

	// Inyección de dependencias
	sucursalService := service.NewSucursalService(collection)
	handler := handlers.NewSucursalHandler(*sucursalService)

	// Rutas
	api := r.Group("/entidad")
	api.Use(middlewares.AuthMiddleware())

	api.POST("", handler.Crear)
	api.GET("", handler.Listar)
	api.GET("/calculos", handler.Calculos)
	api.GET("/proyeccion", handler.ObtenerTablaProyeccion)

	r.Run(":8080")
}
