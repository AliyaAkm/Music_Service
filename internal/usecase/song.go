package usecase

import "PracticeCrud/internal/domain"

// переменная, которая реализует интерфейс. чтобы брать с бд
type songUsecase struct {
	repo domain.SongRepository
}

// конструктор, создает новый объект, в котором поле repo = r
func NewSongUsecase(r domain.SongRepository) domain.SongUsecase {
	return &songUsecase{repo: r}
}

// GetAll Get Create Update Delete
func (u *songUsecase) GetAll() ([]domain.Song, error) {
	return u.repo.Fetch()
}
func (u *songUsecase) Get(id int) (*domain.Song, error) {
	return u.repo.GetByID(id)
}
func (u *songUsecase) Create(song *domain.Song) error {
	return u.repo.Create(song)
}
func (u *songUsecase) Update(song *domain.Song) error {
	return u.repo.Update(song)
}
func (u *songUsecase) Delete(id int) error {
	return u.repo.Delete(id)
}
