package usecase

import (
	"github.com/celpung/clean-gin-architecture/infrastructure"
	"github.com/celpung/clean-gin-architecture/internal/entity"
	"github.com/celpung/clean-gin-architecture/internal/repository"
)

type UserUseCase struct {
	UserRepo    *repository.UserRepository
	PasswordSrv *infrastructure.PasswordService
	Jwtsrv      *infrastructure.JwtService
}

func NewUserUseCase(repo *repository.UserRepository, passwordSrv *infrastructure.PasswordService, jwtSrv *infrastructure.JwtService) *UserUseCase {
	return &UserUseCase{
		UserRepo:    repo,
		PasswordSrv: passwordSrv,
		Jwtsrv:      jwtSrv,
	}
}

func (uc *UserUseCase) CreateUser(user *entity.User) error {
	hashedPassword, err := uc.PasswordSrv.HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
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

func (uc *UserUseCase) SignIn(email string, password string) (string, error) {
	user, err := uc.UserRepo.FindByEmail(email)
	if err != nil {
		return "", err
	}

	err = uc.PasswordSrv.VerifyPassword(user.Password, password)
	if err != nil {
		return "", err
	}

	token, err := uc.Jwtsrv.JWTGenerator(*user)
	if err != nil {
		return "", err
	}

	return token, nil
}
