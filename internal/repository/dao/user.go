package dao

import (
	"context"
	"time"
)
import "gorm.io/gorm"

type UserDAO struct {
	db *gorm.DB
}

func NewUserDAO(db *gorm.DB) *UserDAO {
	return &UserDAO{
		db: db,
	}
}

func (dao *UserDAO) Insert(ctx context.Context, u User) error {
	now := time.Now().UnixMilli() // 毫秒
	u.Utime = now
	u.Ctime = now
	return dao.db.WithContext(ctx).Create(&u).Error
}

// User 与数据库表结构对应。叫法比较多，如：entity、model、PO(Persistent object)
type User struct {
	Id       int64  `gorm:"primaryKey,autoIncrement"` // 自增，主键
	Email    string `gorm:"unique"`                   // 唯一索引
	Password string
	Ctime    int64 // 创建时间，毫秒数
	Utime    int64 // 更新时间，毫秒数
}
