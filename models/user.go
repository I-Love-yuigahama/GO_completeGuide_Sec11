package models

import (
	"errors"

	"example.com/tut/db"
	"example.com/tut/utils"
	// "golang.org/x/tools/go/analysis/passes/defers"
	// "github.com/pelletier/go-toml/query"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (data *User) Save() (error){
	query := "INSERT INTO users (email.password) VALUES (?,?)"

	stmt,err := db.DB.Prepare(query)

	if err != nil{
		return err
	}

	defer stmt.Close()

	hashPassword,err := utils.HashPassword(data.Password)

	result,err := stmt.Exec(data.Email,hashPassword)

	if err != nil{
		return err
	}

	userId ,err := result.LastInsertId()

	if err != nil{
			return err
		}

	data.ID = userId

	return err
}

func (data *User) ValidateCredentials()(error){

	query := "SELECT id,password FROM users WHERE email = ? "
	row:=db.DB.QueryRow(query,data.Email)

	var retreivedPassword string
	err := row.Scan(&data.ID,&retreivedPassword)

	if err != nil{
		return  err
	}

	if !utils.CheckPasswordHash(data.Password,retreivedPassword){
		return errors.New("Credentials invalid")
	}

	return nil
}