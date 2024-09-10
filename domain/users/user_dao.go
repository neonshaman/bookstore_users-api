package users

import (
	"fmt"
	"github.com/neonshaman/bookstore_users-api/datasources/mysql/bookstore_users_db"
	"github.com/neonshaman/bookstore_users-api/utils/date_utils"
	"github.com/neonshaman/bookstore_users-api/utils/errors"
)

var usersDb = make(map[int64]*User)

// Save performs final input validation, attempts to write to database, then returns
func (user *User) Save() *errors.RestErr {
	current := usersDb[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("email %s already registered", user.Email))
		}
		return errors.NewBadRequestError(fmt.Sprintf("user %d already exists", user.Id))
	}

	now := date_utils.GetNowString()
	user.DateCreated = now
	usersDb[user.Id] = user
	return nil
}

// Get performs final input validation, attempts to read from database, then returns
func (user *User) Get() *errors.RestErr {
	if err := bookstore_users_db.Client.Ping(); err != nil {
		panic(err)
	}

	result := usersDb[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("user %d not found", user.Id))
	}

	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}
