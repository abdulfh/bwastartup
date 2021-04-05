package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
}

type repository struct{
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository  {
	return &repository{db}
}

func (repository *repository) Save(user User) (User, error) {
	err := repository.db.Create(&user).Error
	if err != nil {
		return user,err
	}

	return user, nil
}