package services

import (
	"github.com/neonshaman/bookstore_users-api/domain/users"
	"github.com/neonshaman/bookstore_users-api/utils/errors"
)

// GetUser attempts to get user through DAO, then return
func GetUser(userId int64) (*users.User, *errors.RestErr) {
	result := &users.User{Id: userId}
	if err := result.Get(); err != nil {
		return nil, err
	}
	return result, nil
}

// CreateUser attempts to validate user struct through DTO, save to persistence later through DAO, then return
func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func FindUser() {

}
