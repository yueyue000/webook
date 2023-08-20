package domain

// User 领域对象,是DDD中的entity，有的叫BO(Business Object)
type User struct {
	ID       int64
	Email    string
	Password string
}

// Encrypt 加密, 是领域对象的事情
func (u User) Encrypt() {

}

// ComparePassword 比较密码是否正确
func (u User) ComparePassword() {

}

type Address struct{}
