package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const secretkey = "YuigahamaYuiMyWife"

func GenerateToken(email string, userId int64)(string,error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,jwt.MapClaims{
		"email":email,
		"userId":userId,
		"exp": time.Now().Add(time.Hour*2).Unix(),
	})

	return token.SignedString([]byte(secretkey))

}

func VerifyToken(token string)(int64, error){
	parseToken,err:=jwt.Parse(token,func(token *jwt.Token)(interface{},error){
		_,ok := token.Method.(*jwt.SigningMethodHMAC)

		if !ok{
			return nil,errors.New("Unexpected signing methhod")
		}
		return []byte(secretkey),nil
	})

	if err != nil{
			return 0,errors.New("Unexpected signing methhod")
	}

	tokenIsValid := parseToken.Valid

	if !tokenIsValid{
		return 0,errors.New("Unvalid token")
	}

	claims,ok := parseToken.Claims.(jwt.MapClaims)

	if !ok{
		return 0,errors.New("Unvalid token")
	} 
	
	// email := claims["email"].(string)
	userId := int64(claims["userId"].(float64))

	return userId,nil
}