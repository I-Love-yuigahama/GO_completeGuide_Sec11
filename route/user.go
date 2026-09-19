package route

import (
	"net/http"

	"example.com/tut/models"
	"example.com/tut/utils"
	"github.com/gin-gonic/gin"
)

func signUp(context *gin.Context) {
	var user models.User 

	err := context.ShouldBindJSON(&user)

	if err != nil {
		println(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	err = user.Save()

	if err != nil{
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	context.JSON(http.StatusCreated,gin.H{"message": "saved data"})
}

func login(context *gin.Context)(){

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		println(err.Error())
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse data"})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		println(err.Error())
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Wrong password"})
		return
	} 

	token,err:=utils.GenerateToken(user.Email,user.ID)

	if err != nil {
		println(err.Error())
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Cant auth user"})
		return
	} 

	context.JSON(http.StatusOK,gin.H{"message": "Login sucess", "token": token})


}