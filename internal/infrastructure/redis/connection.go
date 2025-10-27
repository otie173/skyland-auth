package redis

import (
	"log"
	"strings"

	"github.com/otie173/skyland-auth/internal/config"
	"github.com/redis/go-redis/v9"
)

func NewClient(cfg *config.Config) *redis.Client {
	var sb strings.Builder
	sb.WriteString(cfg.RedisHost)
	sb.WriteString(":")
	sb.WriteString(cfg.RedisPort)

	addr := sb.String()
	log.Println(addr)

	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       0,
	},
	)
}
