package repositories

import (
	"gorm.io/gorm"
	"wimf-services/entities"
)

type UserRepository interface {
	Save(user entities.User) entities.User
	FindByUsername(username string) interface{}
	FindById(id string) interface{}
	Exists(username string) *gorm.DB
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) Save(user entities.User) entities.User {
	u.db.Save(&user)
	return user
}

func (u *userRepository) FindByUsername(username string) interface{} {
	var user entities.User
	u.db.First(&user, "username = ?", username)
	return user
}

func (u *userRepository) FindById(id string) interface{} {
	var user entities.User
	u.db.First(&user, "id = ?", id)
	return user
}

func (u *userRepository) Exists(username string) *gorm.DB {
	return u.db.Take(&entities.User{}, "username = ?", username)
}
