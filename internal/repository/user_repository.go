package repository

import (
	"github.com/celpung/clean-gin-architecture/internal/entity"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (ur *UserRepository) Create(user *entity.User) error {
	return ur.DB.Create(user).Error
}

func (ur *UserRepository) FindByID(id uint) (*entity.User, error) {
	user := &entity.User{}
	err := ur.DB.First(user, id).Error
	return user, err
}

func (ur *UserRepository) GetAllUser() ([]entity.User, error) {
	var user []entity.User
	err := ur.DB.Find(&user).Error
	return user, err
}

func (ur *UserRepository) Update(user *entity.User) error {
	return ur.DB.Save(user).Error
}

func (ur *UserRepository) Delete(id uint) error {
	return ur.DB.Delete(&entity.User{}, id).Error
}
