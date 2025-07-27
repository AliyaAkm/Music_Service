package postgres

import (
	"PracticeCrud/internal/domain"
	"database/sql"
)

type topSongRepo struct {
	db *sql.DB
}

func NewTopSongRepo(db *sql.DB) domain.TopSongRepository {
	return &topSongRepo{db: db}
}
func (r *topSongRepo) AddTopSong(song *domain.TopSong) error {
	return r.db.QueryRow(`
        INSERT INTO top_songs (title, artist, genre, description)
        VALUES ($1, $2, $3, $4)
        RETURNING id
    `, song.Title, song.Artist, song.Genre, song.Description).Scan(&song.ID)
}
func (r *topSongRepo) GetAllTopSong() ([]domain.TopSong, error) {
	rows, err := r.db.Query(`SELECT id, title, artist, genre, description FROM top_songs ORDER BY id`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []domain.TopSong
	for rows.Next() {
		var s domain.TopSong
		if err := rows.Scan(&s.ID, &s.Title, &s.Artist, &s.Genre, &s.Description); err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	return result, nil
}
