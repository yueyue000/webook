package repository

import (
	"context"
	"github.com/yueyue000/webook/internal/domain"
	"github.com/yueyue000/webook/internal/repository/cache"
	"github.com/yueyue000/webook/internal/repository/dao"
)

var (
	ErrUserDuplicateEmail = dao.ErrUserDuplicateEmail
	ErrUserNotFound       = dao.ErrUserNotFound
)

type UserRepository struct {
	dao   *dao.UserDAO
	cache *cache.UserCache
}

// NewUserRepository 依赖注入的方式初始化，不要在内部初始化
func NewUserRepository(dao *dao.UserDAO, c *cache.UserCache) *UserRepository {
	return &UserRepository{
		dao:   dao,
		cache: c,
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

func (r *UserRepository) FindByID(ctx context.Context, id int64) (domain.User, error) {
	// 从缓存取数据
	u, err := r.cache.Get(ctx, id)
	switch err {
	case nil:
		return u, nil
	case cache.ErrKeyNotExist:
		ue, err := r.dao.SelectByID(ctx, id) // redis如果异常或者挂掉了，缓存击穿到MySQL，考虑加布隆过滤器或者本地缓存。
		if err != nil {
			return domain.User{}, err
		}
		u = domain.User{
			ID:          ue.ID,
			Email:       ue.Email,
			Password:    ue.Password,
			Nick:        ue.Nick,
			Birthday:    ue.Birthday,
			Description: ue.Description,
		}
		go func() {
			err = r.cache.Set(ctx, u)
			if err != nil {
				// todo 非核心逻辑，打错误日志，做好监控
			}
		}()
	default:
		return domain.User{}, err
	}
	return domain.User{}, nil
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
