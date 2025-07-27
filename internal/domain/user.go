package domain

type User struct {
	ID       int    `json:"ID"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type UserRepository interface {
	FindByUsername(username string) (*User, error)
}
type UserUsecase interface {
	FindByUsername(username string) (*User, error)
}
