package domain

// interface of usecase: GetAll,Get,Create, Update, Delete

type Song struct {
	ID     int    `json:"ID"`
	Title  string `json:"Title"`
	Artist string `json:"Artist"`
}

type SongRepository interface {
	Fetch() ([]Song, error)
	GetByID(id int) (*Song, error)
	Create(song *Song) error
	Update(song *Song) error
	Delete(id int) error
}

type SongUsecase interface {
	GetAll() ([]Song, error)
	Get(id int) (*Song, error)
	Create(song *Song) error
	Update(song *Song) error
	Delete(id int) error
}
