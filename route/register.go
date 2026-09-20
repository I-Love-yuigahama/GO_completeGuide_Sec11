package route

import (
	// "context"

	"net/http"
	"strconv"

	"example.com/tut/models"
	"github.com/gin-gonic/gin"
	// "go.mongodb.org/mongo-driver/v2/event"
)

func registerForEvent(context *gin.Context){
	eventId,err:= strconv.ParseInt(context.Param("id"),10, 64)
	userId:=context.GetInt16("userId")


	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "could get data"}) 
		return
	}

	event,err:=models.GetEventByID(eventId)

	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "could get data"})
		return
	}

	err = event.Register(int64(userId))

	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "could get data"})
		return
	}
	
	context.JSON(http.StatusCreated,gin.H{"message": "Registered"})

}

func cancelRehistration(context *gin.Context){

	eventId,err:= strconv.ParseInt(context.Param("id"),10, 64)
	userId:=context.GetInt16("userId")

	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "could not get data"}) 
		return
	}

	var event models.Event
	event.ID = eventId
	
	err = event.CancelRehistration(int64(userId))

	if err != nil{
		context.JSON(http.StatusInternalServerError,gin.H{"message": "could not cancel data"}) 
		return
	}

	context.JSON(http.StatusCreated,gin.H{"message": "cancel sucessfully"})


}


