package usecase

import (
	"github.com/celpung/clean-gin-architecture/internal/entity"
	"github.com/celpung/clean-gin-architecture/internal/repository"
)

type UserUseCase struct {
	UserRepo *repository.UserRepository
}

func NewUserUseCase(repo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{UserRepo: repo}
}

func (uc *UserUseCase) CreateUser(user *entity.User) error {
	// Add any business logic or validation here
	return uc.UserRepo.Create(user)
}

func (uc *UserUseCase) GetUserByID(id uint) (*entity.User, error) {
	return uc.UserRepo.FindByID(id)
}

func (uc *UserUseCase) GetAllUsers() ([]entity.User, error) {
	return uc.UserRepo.GetAllUser()
}

func (uc *UserUseCase) UpdateUser(user *entity.User) error {
	// Add any business logic or validation here
	return uc.UserRepo.Update(user)
}

func (uc *UserUseCase) DeleteUser(id uint) error {
	return uc.UserRepo.Delete(id)
}
