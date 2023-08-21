package domain

// User 领域对象,是DDD中的entity，有的叫BO(Business Object)
type User struct {
	ID          int64  `json:"id,omitempty"`
	Email       string `json:"email,omitempty"`
	Password    string `json:"password,omitempty"`
	Nick        string `json:"nick"`
	Birthday    string `json:"birthday"`
	Description string `json:"description"`
}

// Encrypt 加密, 是领域对象的事情
func (u User) Encrypt() {

}

// ComparePassword 比较密码是否正确
func (u User) ComparePassword() {

}

type Address struct{}
