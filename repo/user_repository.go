package repo

import (
	"errors"
	"fossil/database"
	"fossil/model"
	"log"
	"sync"

	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

var lock = &sync.Mutex{}
var UserRepositoryInstance *UserRepository

func MakeUserRepository() *UserRepository {
	if UserRepositoryInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		UserRepositoryInstance = &UserRepository{
			DB: database.Conn.DB,
		}
	}

	return UserRepositoryInstance
}

func GetUserRepository() *UserRepository {
	return MakeUserRepository()
}

// Save user and populate saved result to a User struct
func (r *UserRepository) Create(user *model.User) error {
	result := r.DB.Create(user)

	if result.Error != nil {
		return result.Error
	}

	result = r.DB.Find(user, user.ID)

	// assuming created data cant be corrupted,
	// no error returned after data fetching
	// further action: consider db transaction
	if result.Error != nil {
		log.Println(result.Error)
	}

	return nil
}

func (r *UserRepository) Find(id uint) ([]model.User, error) {
	return nil, nil
}

func (r *UserRepository) Exists(condition *model.User) bool {
	user := new(model.User)
	result := r.DB.Where(condition).Find(user)

	if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return false
	}

	if result.RowsAffected < 1 {
		return false
	}

	return true
}
