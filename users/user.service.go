package users

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Service is an interface that defines the RegisterUser method.
// @property RegisterUser - This is the name of the method that will be exposed to the outside world.
type Service interface {
	RegisterUser(input RegisterUserInput) (Users, error)
}

// A service is a struct that has a repository.
// @property {Repository} repository - This is the repository that the service will use to store and
// retrieve data.
type service struct {
	repository Repository
}

// NewService returns a new service with the given repository.
func NewService(repository Repository) *service {
	return &service{repository}
}

// A function that takes in a RegisterUserInput struct and returns a Users struct and an error.
func (s *service) RegisterUser(input RegisterUserInput) (Users, error) {
	user := Users{}
	user.Name = input.Name
	user.Occupation = input.Occupation
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	user.Created_at = time.Now()
	user.Updated_at = time.Now()

	if err != nil {
		return user, err
	}

	user.Password_hash = string(passwordHash)
	user.Role = "user"

	newUser, err := s.repository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
