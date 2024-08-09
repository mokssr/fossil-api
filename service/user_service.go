package service

import (
	"errors"
	"fossil/api"
	"fossil/model"
	"fossil/repo"
	"sync"
)

type UserService struct {
	Repo *repo.UserRepository
}

var lock = &sync.Mutex{}
var UserServiceInstance *UserService

func MakeUserService() *UserService {
	if UserServiceInstance == nil {
		lock.Lock()
		defer lock.Unlock()

		UserServiceInstance = &UserService{
			Repo: repo.MakeUserRepository(),
		}
	}

	return UserServiceInstance
}

func (s *UserService) AddUser(data api.CreateUserRequest, user *model.User) error {
	// check mismatch password confirmation
	if data.Password != data.ConfirmPassword {
		return errors.New("Mismatch password confirmation")
	}

	// check used email
	exists := s.Repo.Exists(&model.User{Email: data.Email})
	if exists {
		return errors.New("Email is already registered")
	}

	// check used username
	exists = s.Repo.Exists(&model.User{Username: data.Username})
	if exists {
		return errors.New("Username is unavailable")
	}

	// contruct and populate result
	user.Username = data.Username
	user.Email = data.Email
	user.Password = data.Password + "+secret"

	err := s.Repo.Create(user)
	if err != nil {
		return errors.New("Failed saving user data")
	}

	return nil
}
