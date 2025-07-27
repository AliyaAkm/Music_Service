package domain

type TopSong struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Artist      string `json:"artist"`
	Genre       string `json:"genre"`
	Description string `json:"description"`
}

type TopSongRepository interface {
	AddTopSong(song *TopSong) error
	GetAllTopSong() ([]TopSong, error)
}

type TopSongUsecase interface {
	Add(song *TopSong) error
	GetAll() ([]TopSong, error)
}
