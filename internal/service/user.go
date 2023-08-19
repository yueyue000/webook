package service

import (
	"context"
	"errors"
	"github.com/yueyue000/webook/internal/domain"
	"github.com/yueyue000/webook/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

var ErrUserDuplicateEmail = repository.ErrUserDuplicateEmail
var ErrInvalidUserOrPassword = errors.New("账号或密码错误")

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hash)
	err = svc.repo.Create(ctx, u)
	return err
}

func (svc *UserService) Login(ctx context.Context, u domain.User) error {
	userInfo, err := svc.repo.FindByEmail(ctx, u.Email)
	if err == repository.ErrUserNotFound {
		return ErrInvalidUserOrPassword
	}
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(u.Password))
	if err != nil {
		// TODO 打日志
		return ErrInvalidUserOrPassword
	}
	return nil
}
