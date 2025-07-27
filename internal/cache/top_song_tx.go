package cache

import (
	"PracticeCrud/internal/domain"
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func (r *RedisCache) InvalidateTopSongWithTx(ctx context.Context, song *domain.TopSong) error {
	key := "top_songs:list"
	songKey := fmt.Sprintf("top_song:%d", song.ID)

	data, err := json.Marshal(song)
	if err != nil {
		return fmt.Errorf("marshal song: %w", err)
	}

	fmt.Println("Redis Tx: start InvalidateTopSongsWithTx")

	_, err = r.client.TxPipelined(ctx, func(pipe redis.Pipeliner) error {
		pipe.Del(ctx, key)
		pipe.Set(ctx, songKey, data, 20*time.Minute)
		return nil
	})
	if err != nil {
		fmt.Println("Redis Tx failed:", err)
		return err
	}
	fmt.Println("Redis Tx succeeded")
	return nil
}
