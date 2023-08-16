package service

import (
	"context"
	"github.com/yueyue000/webook/internal/domain"
	"github.com/yueyue000/webook/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (svc *UserService) SignUp(ctx context.Context, u domain.User) error {
	err := svc.repo.Create(ctx, u)
	return err
}
