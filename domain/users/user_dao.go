package users

import (
	"fmt"
	"github.com/neonshaman/bookstore_users-api/datasources/mysql/bookstore_users_db"
	"github.com/neonshaman/bookstore_users-api/utils/date_utils"
	"github.com/neonshaman/bookstore_users-api/utils/errors"
	"strings"
)

// SQL queries
const (
	insertUserQuery = "INSERT INTO users(first_name, last_name, email, date_created) VALUES(?, ?, ?, ?)"
)

// Save performs final input validation, attempts to write to database, then returns
func (user *User) Save() *errors.RestErr {
	stmt, err := bookstore_users_db.Client.Prepare(insertUserQuery)
	if err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	defer stmt.Close()

	user.DateCreated = date_utils.GetNowString()

	result, err := stmt.Exec(
		user.FirstName,
		user.LastName,
		user.Email,
		user.DateCreated,
	)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.NewBadRequestError(
				fmt.Sprintf("user %s already exists", user.Email))
		}
		return errors.NewInternalServerError(
			fmt.Sprintf("could not write user to database: %s", err.Error()))
	}
	userId, err := result.LastInsertId()
	if err != nil {
		return errors.NewInternalServerError(fmt.Sprintf("could not get last inserted id: %s", err.Error()))
	}
	user.Id = userId
	return nil
}

// Get performs final input validation, attempts to read from database, then returns
func (user *User) Get() *errors.RestErr {
	if err := bookstore_users_db.Client.Ping(); err != nil {
		panic(err)
	}

	//result := usersDb[user.Id]
	//if result == nil {
	//	return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	//}
	//
	//user.Id = result.Id
	//user.FirstName = result.FirstName
	//user.LastName = result.LastName
	//user.Email = result.Email
	//user.DateCreated = result.DateCreated

	return nil
}
