package main

import (
	"context"
	Service "examen_final_febrero_golang_P1/Services"
	"examen_final_febrero_golang_P1/handlers"
	"examen_final_febrero_golang_P1/middlewares"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()
	r.Use(middlewares.AuthMiddleware())
	// Mongo directo
	client, err := mongo.Connect(context.TODO(),
		options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		panic(err)
	}

	collection := client.Database("examen").Collection("entidades")
	entidadService := Service.NewEntidadService(collection)

	handler := handlers.NewSucursalHandler(*entidadService)

	r.POST("/entidad", handler.Crear)
	r.GET("/entidad", handler.Listar)
	r.GET("/", handler.Calculos)
	r.GET("/proyeccion", handler.ObtenerTablaProyeccion)
	r.Run(":8080")
}
