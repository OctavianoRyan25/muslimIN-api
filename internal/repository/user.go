package repository

import (
	"errors"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByEmail(email string) (*domain.User, error)
	CreateUser(user *domain.User) error
	CreateAPIKey(apiKey *domain.APIKey) error
}

type userRepo struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepo{db: db}
}

func (ur *userRepo) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := ur.db.Where("email = ?", email).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func (ur *userRepo) CreateUser(user *domain.User) error {
	return ur.db.Create(user).Error
}

func (ur *userRepo) CreateAPIKey(apiKey *domain.APIKey) error {
	return ur.db.Create(apiKey).Error
}
