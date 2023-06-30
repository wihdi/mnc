package repository

import (
	"encoding/json"
	"log"
	"os"
	"github.com/wihdi/mnc/domain"
)

type UserRepository interface {
	FindByUsername(username string) (*domain.User, error)
	
}

type userRepository struct {
	filePath string
}

func NewUserRepository(filePath string) *userRepository {
	return &userRepository{filePath: filePath}
}
func (u *userRepository) FindByUsername(username string) (*domain.User, error) {
	users, err := u.readUsers()
	if err != nil {
		log.Println("Failed to read users from JSON file:", err)
		return nil, err
	}

	for _, user := range users {
		if user.Username == username {
			return &user, nil
		}
	}

	return nil, nil
}
func (u *userRepository) readUsers() ([]domain.User, error) {
	file, err := os.Open(u.filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var users []domain.User
	err = json.NewDecoder(file).Decode(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}


