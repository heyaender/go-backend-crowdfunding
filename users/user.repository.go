package users

import (
	"gorm.io/gorm"
)

// Repository is an interface that defines the methods that a repository must implement.
// @property Save - This is a function that takes a user and returns a user and an error.
type Repository interface {
	Save(user Users) (Users, error)
}

// A repository is a struct that contains a pointer to a gorm.DB object.
// @property db - This is the database connection.
type repository struct {
	db *gorm.DB
}

// `NewRepository` is a function that takes a `*gorm.DB` and returns a `*repository`
func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

// Saving the user to the database.
func (r *repository) Save(user Users) (Users, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
