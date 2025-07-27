package postgres

import (
	"PracticeCrud/internal/domain"
	"database/sql"
	"log"
)

type songRepo struct {
	db *sql.DB
}

func NewSongRepo(db *sql.DB) domain.SongRepository {
	return &songRepo{db: db}
}

func (r *songRepo) Fetch() ([]domain.Song, error) {
	rows, err := r.db.Query("SELECT id,title,artist FROM songs")
	if err != nil {
		return nil, err
	}
	defer func() {
		err = rows.Close()
		if err != nil {
			log.Println("error", err)
		}
	}()

	var songs []domain.Song
	for rows.Next() {
		var s domain.Song
		err = rows.Scan(&s.ID, &s.Title, &s.Artist)
		if err != nil {
			return nil, err
		}
		songs = append(songs, s)
	}
	return songs, err
}

func (r *songRepo) GetByID(id int) (*domain.Song, error) {
	var s domain.Song
	err := r.db.QueryRow("SELECT id, title, artist FROM songs WHERE id = $1", id).Scan(&s.ID, &s.Title, &s.Artist)
	return &s, err
}

func (r *songRepo) Create(song *domain.Song) error {
	return r.db.QueryRow("INSERT INTO songs(title, artist) VALUES($1, $2) RETURNING id", song.Title, song.Artist).Scan(&song.ID)
}

func (r *songRepo) Update(song *domain.Song) error {
	_, err := r.db.Exec("UPDATE songs SET title = $1, artist = $2 WHERE id = $3", song.Title, song.Artist, song.ID)
	return err
}

func (r *songRepo) Delete(id int) error {
	_, err := r.db.Exec("DELETE FROM songs WHERE id = $1", id)
	return err
}
