package domain

// User 领域对象,是DDD中的entity，有的叫BO(Business Object)
type User struct {
	Email    string
	Password string
}

type Address struct{}
