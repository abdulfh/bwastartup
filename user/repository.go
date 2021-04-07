package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FinById(ID int) (User, error)
	Update(user User) (User, error)
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
func (repository *repository) FindByEmail(email string) (User, error) {
	var user User

	err := repository.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user,nil
}

func(repository *repository) FinById(ID int) (User, error) {
	var user User

	err := repository.db.Where("id = ?", ID).Find(&user).Error
	if err != nil {
		return user, err
	}

	return user,nil
}

func(repository *repository) Update(user User) (User, error) {
	err := repository.db.Save(&user).Error
	
	if err != nil {
		return user, err
	}
	
	return user, nil
}