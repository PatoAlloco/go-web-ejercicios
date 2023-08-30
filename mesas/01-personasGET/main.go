package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type Producto struct {
	Id            int       `jason:"id"`
	Nombre        string    `jason:"nombre"`
	Precio        float64   `jason:"precio"`
	Stock         int       `jason:"stock"`
	Codigo        string    `jason:"codigo"`
	Publicado     bool      `jason:"publicado"`
	FechaCreacion time.Time `jason:"fecha_de_creacion"`
}

const (
	fileName = "./productos.json"
)

func main() {

	//Crear un router con gin
	router := gin.Default()

	//Creo un endpoint para comprobar conecxion
	router.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	//Creo endpoint que devuelvo los porductos
	router.GET("/productos", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"data": readFile(fileName),
		})
	})

	//Hago correr en el puerto 8080 por defecto
	router.Run()

}

func readFile(fileName string) []Producto {
	file, err := os.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	var productos []Producto
	err = json.Unmarshal(file, &productos) //desempaquetar los datos que vienen de un json y va con puntero a productos
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	return productos
}
