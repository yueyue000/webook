package repository

import (
	"context"
	"github.com/yueyue000/webook/internal/domain"
	"github.com/yueyue000/webook/internal/repository/dao"
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

func (r *UserRepository) FindById(int64) {

}
