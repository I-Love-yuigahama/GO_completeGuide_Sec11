package main

import (
	// "net/http"
	// "strconv"

	"example.com/tut/db"
	"example.com/tut/route"
	// "example.com/tut/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.INitDB()

	server := gin.Default()

	route.RegisterRoutes(server)

	server.Run(":8080")

}
