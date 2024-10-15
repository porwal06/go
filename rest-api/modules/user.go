package modules

import (
	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	//Save event in sqllite database
	query := `INSERT INTO users (email, password)
	VALUES (?, ?)
	`
	stmt, err := db.DB.Prepare(query)
	//Error handling
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashPassword := utils.GenerateHashPassword(u.Password)
	result, err := stmt.Exec(u.Email, hashPassword)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}
	u.ID = id
	//events = append(events, e)
	return err
}

func (u User) Login() error {
	return nil

}
