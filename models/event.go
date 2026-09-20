package models

import (
	// "testing/quick"
	"time"

	"example.com/tut/db"
	// "github.com/pelletier/go-toml/query"
	// "github.com/pelletier/go-toml/query"
	// "github.com/pelletier/go-toml/query"
)

type Event struct {
	ID          int64
	Name        string `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64
}
//slices
var event = []Event{}

// methods
func (data *Event) Save() error {
	query := `
	INSERT INTO events(name,description,location,dateTime, user_id) 
	VALUES (?,?,?,?,?)
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	//if have query that changes stuff is Exec
	//if have query fetch data is  Query
	// Prepare() + stmt.Exec() (when we inserted data into the database)

	defer stmt.Close()

	result, err := stmt.Exec(data.Name, data.Description, data.Location, data.DateTime, data.UserID)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	data.ID = id
	return nil
	// event = append(event, *data)
}

func GetEvents() ([]Event, error) {

	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)

		if err != nil {
			return nil, err
		}

		events = append(events, event)
	}

	return events, nil
}

func GetEventByID(id int64)(*Event,error){

	row := db.DB.QueryRow("SELECT * FROM events WHERE id = ?", id)
	
	var event Event
	err:= row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
			return nil, err
		}
	return  &event,nil
}

func (data *Event) Update()(error){

	query:=`
	UPDATE events
	SET	name =? , description = ? , location =? , dateTime=?
	WHERE id = ?
	`

	stmt,err:=db.DB.Prepare(query)

	if err != nil {
			return err
		}

		defer stmt.Close()

		_,err = stmt.Exec(data.Name, data.Description, data.Location, data.DateTime,data.ID)

		if err != nil {
			return err
		}

		return nil

}

func (data *Event) Delete()(error){
	query:=`
	DELETE FROM events WHERE id = ?
	`
	stmt,err:=db.DB.Prepare(query)
	
	if err != nil {
			return err
		}

		defer stmt.Close()

		_,err = stmt.Exec(data.ID)

		return err
}

func (data *Event) Register(userID int64)(error){
	query := "INTERT INTO registration(event_id, user_id) VALUES (?,?)"
 	stmt,err:=	db.DB.Prepare(query)

	if err!= nil{
		return err
	}

	defer stmt.Close()

 _,err =	stmt.Exec(data.ID,data.UserID)

 return err

}

func(data *Event) CancelRehistration(userId int64)(error){
	query := "DELETE FROM registration WHERE event_id =? AND user_id = ?)"
 	stmt,err:=	db.DB.Prepare(query)
	if err!= nil{
		return err
	}

	defer stmt.Close()

	_,err =	stmt.Exec(data.ID,data.UserID)

	return err
	}
