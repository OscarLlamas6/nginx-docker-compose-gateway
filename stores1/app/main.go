package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type store struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location string `json:"location"`
}

var stores = []store{
	{ID: "1", Name: "Platinium", Location: "Plaza Fontabella z.10, Ciudad de Guatemala"},
	{ID: "2", Name: "Office Depot", Location: "CC Rambla 10, Boulevar Los Proceres 25-76, Cdad. de Guatemala"},
	{ID: "3", Name: "Cemaco", Location: "12 Avenida 15-76, zona 10. Ciudad de Guatemala"},
}

func getStores(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, stores)
}

func getHello(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Hello World from stores 1",
	})
}

func main() {
	router := gin.Default()
	router.GET("/stores", getStores)
	router.GET("/", getHello)
	router.Run(":3001")
}
