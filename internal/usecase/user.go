package usecase

import (
	"errors"
	"time"

	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/domain"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/repository"
	"github.com/OctavianoRyan25/belajar-pattern-code-go/internal/util"
)

type UserUseCase interface {
	GetUserByEmail(email string) (*domain.User, error)
	CreateUser(user *domain.User) error
	LoginUser(user *domain.User) (string, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (uu *userUseCase) GetUserByEmail(email string) (*domain.User, error) {
	if email == "" {
		return nil, errors.New("email is required")
	}
	return uu.repo.GetUserByEmail(email)
}

func (uu *userUseCase) CreateUser(user *domain.User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if user.Email == "" {
		return errors.New("email is required")
	}
	if user.Password == "" {
		return errors.New("password is required")
	}
	hashed, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashed

	return uu.repo.CreateUser(user)
}

func (uu *userUseCase) LoginUser(user *domain.User) (string, error) {
	if user.Email == "" {
		return "", errors.New("email is required")
	}
	if user.Password == "" {
		return "", errors.New("password is required")
	}
	res, err := uu.repo.GetUserByEmail(user.Email)
	if err != nil {
		return "", err
	}
	err = util.CheckPassword(res.Password, user.Password)
	if err != nil {
		return "", err
	}
	jwt, err := util.GenerateJWT(res.ID)
	if err != nil {
		return "", err
	}
	return jwt, nil
}

func (uu *userUseCase) CreateAPIKey(userID uint) (*domain.APIKey, error) {

	key, err := util.GenerateAPIKey()
	if err != nil {
		return nil, err
	}

	apiKey := &domain.APIKey{
		UserID:       userID,
		Key:          key,
		MonthlyLimit: 100,
		UsageCount:   0,
		ResetAt:      firstDayNextMonth(),
	}

	err = uu.repo.CreateAPIKey(apiKey)
	return apiKey, err
}

func firstDayNextMonth() time.Time {
	now := time.Now()

	firstDayNextMonth := time.Date(
		now.Year(),
		now.Month()+1,
		1,
		0, 0, 0, 0,
		now.Location(),
	)

	return firstDayNextMonth
}
