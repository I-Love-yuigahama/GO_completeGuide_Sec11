package route

import (
	"example.com/tut/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine){

	server.GET("/events", getEvents)

	server.GET("/events/:id", getEvent)

	auth:=server.Group("/")

	auth.Use(middlewares.Authenticate)

	auth.POST( "/events",createEvent)

	auth.PUT( "/events/:id",updateEvent)

	auth.DELETE( "/events/:id",deleteEvent)

	auth.POST( "/events", middlewares.Authenticate,createEvent)

	auth.POST("/events/:id/register",registerForEvent)

	auth.DELETE("/events/:id/register",cancelRehistration)

	// server.POST("/events", middlewares.Authenticate,createEvent)

	// server.PUT("/events/:id",updateEvent)

	// server.DELETE("/events/:id",deleteEvent)

	server.POST("/signup",signUp)

	server.POST("/login",signUp)
}