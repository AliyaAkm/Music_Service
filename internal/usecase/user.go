package usecase

import "PracticeCrud/internal/domain"

type userUsecase struct {
	repo domain.UserRepository
}

func NewUserUsecase(r domain.UserRepository) domain.UserUsecase {
	return &userUsecase{repo: r}
}
func (u *userUsecase) FindByUsername(username string) (*domain.User, error) {
	return u.repo.FindByUsername(username)
}
