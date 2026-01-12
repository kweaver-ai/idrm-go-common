package d_session

import (
	"context"
	"encoding/json"

	"github.com/redis/go-redis/v9"
)

type Session interface {
	GetRawRedisClient() *redis.Client
	GetSession(ctx context.Context, sessionId string) (*SessionInfo, error)
	DelSession(ctx context.Context, sessionId string) error
}

type SessionInfo struct {
	RefreshToken string
	Token        string
	IdToken      string
	Userid       string
	UserName     string
	State        string
	Platform     int32
}

func (s *SessionInfo) Serialize() string {
	marshal, _ := json.Marshal(s)
	return string(marshal)
}

func (s *SessionInfo) Deserialization(str string) error {
	return json.Unmarshal([]byte(str), s)
}
