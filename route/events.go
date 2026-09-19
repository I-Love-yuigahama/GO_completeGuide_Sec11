package route

import (
	"net/http"
	"strconv"

	// "example.com/tut/db"
	"example.com/tut/models"
	// "example.com/tut/utils"
	"github.com/gin-gonic/gin"
)

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

	userId:=context.GetInt16("userId")

	// event.ID = 1
	event.UserID = int64(userId)

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message": "Event created", "Event": event})
}

func updateEvent(context *gin.Context){

	eventId,err:= strconv.ParseInt(context.Param("id"),10, 64)

	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "could get data"}) 
	}

	_,err = models.GetEventByID(eventId)

	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "could get id"}) 
	}

	var updateEvent models.Event
	err = context.ShouldBindJSON(&updateEvent)

	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "could not Update data"}) 
	}

	updateEvent.ID = eventId
	err = updateEvent.Update()
	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "could not Update data"}) 
	}
	
	context.JSON(http.StatusOK, gin.H{"message": "Updated data"})
}

func deleteEvent(context *gin.Context){

	eventId,err:= strconv.ParseInt(context.Param("id"),10, 64)

	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "could get data"}) 
	}

	event ,err := models.GetEventByID(eventId)

	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "could get id"}) 
	}

	event.Delete()

	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "could not delete id"}) 
	}

	context.JSON(http.StatusOK, gin.H{"message": "Deleted data"})

}
