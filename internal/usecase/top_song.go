package usecase

import (
	"PracticeCrud/internal/cache"
	"PracticeCrud/internal/domain"
	"context"
	"encoding/json"
	"log"
	"time"
)

var ctx = context.Background()

type topSongUsecase struct {
	repo  domain.TopSongRepository
	cache *cache.RedisCache
}

func NewTopSongUsecase(r domain.TopSongRepository, c *cache.RedisCache) domain.TopSongUsecase {
	return &topSongUsecase{repo: r, cache: c}
}
func (u *topSongUsecase) Add(song *domain.TopSong) error {
	if err := u.repo.AddTopSong(song); err != nil {
		return err
	}
	if err := u.cache.InvalidateTopSongWithTx(ctx, song); err != nil {
		log.Printf("Redis transaction failed: %v", err)
	}
	return nil
}
func (u *topSongUsecase) GetAll() ([]domain.TopSong, error) {
	start := time.Now()
	cacheKey := "top_songs:list"

	cached, err := u.cache.Get(cacheKey)
	if err != nil {
		log.Println("Redis error", err)
	}

	if cached != "" {
		var songs []domain.TopSong
		err := json.Unmarshal([]byte(cached), &songs)
		if err != nil {
			log.Println("[ERROR] Failed to unmarshal from Redis:", err)
		} else {
			log.Println("[SOURCE] from Redis")
			log.Println("Duration:", time.Since(start))
			return songs, nil
		}
	}

	// если кэш пустой
	songs, err := u.repo.GetAllTopSong()
	if err != nil {
		return nil, err
	}
	data, _ := json.Marshal(songs)
	_ = u.cache.Set(cacheKey, string(data), 20*time.Minute)

	log.Println("[SOURCE] from DB")
	log.Println("Duration:", time.Since(start))
	return songs, nil
}
