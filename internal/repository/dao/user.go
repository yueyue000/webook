package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"time"
)
import "gorm.io/gorm"

var (
	ErrUserDuplicateEmail = errors.New("邮箱冲突")
	ErrUserNotFound       = gorm.ErrRecordNotFound
)

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
	err := dao.db.WithContext(ctx).Create(&u).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		const uniqueConflictsErrNo = 1062
		if mysqlErr.Number == 1062 {
			return ErrUserDuplicateEmail
		}
	}
	return err
}

func (dao *UserDAO) SelectByEmail(ctx context.Context, email string) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Where("email=?", email).First(&user).Error
	return user, err
}

func (dao *UserDAO) SelectByID(ctx context.Context, id int64) (User, error) {
	var user User
	err := dao.db.WithContext(ctx).Where("id=?", id).First(&user).Error
	return user, err
}

func (dao *UserDAO) UpdateByID(ctx context.Context, user User) error {
	user.Utime = time.Now().UnixMilli()
	err := dao.db.WithContext(ctx).Where("id=?", user.ID).Updates(&user).Error
	return err
}

// User 与数据库表结构对应。叫法比较多，如：entity、model、PO(Persistent object)
type User struct {
	ID          int64  `gorm:"primaryKey,autoIncrement"` // 自增，主键
	Email       string `gorm:"unique"`                   // 唯一索引
	Password    string
	Nick        string
	Birthday    int64
	Description string
	Ctime       int64 // 创建时间，毫秒数
	Utime       int64 // 更新时间，毫秒数
}
