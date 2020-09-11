package services

import (
	models "github.com/antonioalfa22/go-rest-template/internal/pkg/models/users"
	"github.com/antonioalfa22/go-rest-template/internal/pkg/persistence"
)

type userDAO interface {
	Get(id string) (*models.User, error)
	GetByUsername(username string) (*models.User, error)
	All() (*[]models.User, error)
	Add(user *models.User) error
	Update(user *models.User) error
	Delete(user *models.User) error
	Query(q *models.User) (*[]models.User, error)
}

type UserService struct {
	dao userDAO
}

var userService *UserService

func GetUserService() *UserService {
	if userService == nil {
		userService = &UserService{persistence.NewUserDAO()}
	}
	return userService
}

func (s *UserService) Get(id string) (*models.User, error) { return s.dao.Get(id) }
func (s *UserService) GetByUsername(username string) (*models.User, error) {
	return s.dao.GetByUsername(username)
}
func (s *UserService) All() (*[]models.User, error)   { return s.dao.All() }
func (s *UserService) Add(user *models.User) error    { return s.dao.Add(user) }
func (s *UserService) Update(user *models.User) error { return s.dao.Update(user) }
func (s *UserService) Delete(user *models.User) error { return s.dao.Delete(user) }
func (s *UserService) Query(q *models.User) (*[]models.User, error) {
	return s.dao.Query(q)
}
