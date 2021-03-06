package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	RegisterUser(input RegisterUserInput) (User, error)
	Login(input LoginUserInput) (User, error)
	IsEmailAvailable(input CheckEmailInput) (bool, error)
	SaveAvatar(ID int, fileLocation string) (User, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) RegisterUser(input RegisterUserInput) (User, error) {
	user := User{}
	user.Name = input.Name
	user.Email = input.Email
	user.Occupation = input.Occupation

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}

	user.PasswordHash = string(passwordHash)
	user.Role = "user"

	newUser, err := service.repository.Save(user)
	if err != nil {
		return newUser, err
	}

	return newUser, nil
}

func (service *service) Login(input LoginUserInput) (User, error) {
	email := input.Email
	password := input.Password

	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return user, err
	}

	if user.ID == 0 {
		return user, errors.New("User not registered")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return user, err
	}

	return user, nil
}

func (service *service) IsEmailAvailable(input CheckEmailInput) (bool, error) {
	email := input.Email

	user, err := service.repository.FindByEmail(email)
	if err != nil {
		return false, err
	}

	if user.ID == 0 {
		return true, nil
	}

	return false, nil
}

func (service *service) SaveAvatar(ID int, fileLocation string) (User, error) {
	user, err := service.repository.FinById(ID)
	if err != nil {
		return user, nil
	}

	user.AvatarFileName = fileLocation

	updatedUser, err := service.repository.Update(user)
	if err != nil {
		return user, nil
	}

	return updatedUser, nil
}
