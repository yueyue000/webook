package repository

import (
	"context"
	"github.com/yueyue000/webook/internal/domain"
	"github.com/yueyue000/webook/internal/repository/dao"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrUserNotFound
)

type UserRepository struct {
	dao *dao.UserDAO
}

// NewUserRepository 依赖注入的方式初始化，不要在内部初始化
func NewUserRepository(dao *dao.UserDAO) *UserRepository {
	return &UserRepository{
		dao: dao,
	}
}

func (r *UserRepository) Create(ctx context.Context, u domain.User) error {
	daoUser := dao.User{
		Email:    u.Email,
		Password: u.Password,
	}
	return r.dao.Insert(ctx, daoUser)
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (domain.User, error) {
	u, err := r.dao.SelectByEmail(ctx, email)
	if err != nil {
		return domain.User{}, err
	}
	return domain.User{
		ID:       u.ID,
		Email:    u.Email,
		Password: u.Password,
	}, nil
}

func (r *UserRepository) UpdateByID(ctx context.Context, userDomain domain.User) error {
	user := dao.User{
		ID:          userDomain.ID,
		Nick:        userDomain.Nick,
		Birthday:    userDomain.Birthday,
		Description: userDomain.Description,
	}
	err := r.dao.UpdateByID(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
