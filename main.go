package main

import (
	"net/http"
	"strconv"

	"example.com/tut/db"
	"example.com/tut/models"
	"github.com/gin-gonic/gin"
)

func main() {

	db.INitDB()

	server := gin.Default()

	server.GET("/events", getEvents)

	server.GET("/events/:id", getEvent)

	server.POST("/events", createEvent)

	server.Run(":8080")

}

func getEvent(context *gin.Context) {
	eventId,err:= strconv.ParseInt(context.Param("id"),10, 64)

	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "could not parse data"}) 
	}

	event,err:=models.GetEventByID(eventId)
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "could not parse data"}) 
	}

	context.JSON(http.StatusOK,event)
}

func getEvents(context *gin.Context) {
	// models.GetEvents()
	events, err := models.GetEvents()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "could not parse data"}) 
	}
	context.JSON(http.StatusOK, events)
}

func createEvent(context *gin.Context) {

	var event models.Event // This declares an empty Event with every field at its zero value ("", 0, and so on). It's the container the request data will be poured into.
	err := context.ShouldBindJSON(&event)

	if err != nil {
		println(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	event.ID = 1
	event.UserID = 1
	event.Email = "testing@gmail.com"

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "Event": event})
}
