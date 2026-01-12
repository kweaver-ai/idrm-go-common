package impl

import (
	"context"
	"errors"

	"github.com/kweaver-ai/idrm-go-common/d_session"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/log"
	"github.com/kweaver-ai/idrm-go-frame/core/telemetry/trace"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

type session struct {
	redisClient *redis.Client
}

func NewSession(redisClient *redis.Client) d_session.Session {
	return &session{redisClient: redisClient}
}

func (r *session) GetRawRedisClient() *redis.Client {
	return r.redisClient
}

func (r *session) GetSession(ctx context.Context, sessionId string) (res *d_session.SessionInfo, err error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()
	result, err := r.redisClient.Get(ctx, sessionId).Result()
	if err != nil {
		log.WithContext(ctx).Error("session Get Session  err", zap.Error(err))
		return nil, err
	}
	if result == "" {
		log.WithContext(ctx).Error("session Get Session  empty")
		return nil, errors.New("session empty")
	}
	var sessionInfo d_session.SessionInfo
	log.Infof("GetSession result :%v", result)
	err = sessionInfo.Deserialization(result)
	if err != nil {
		log.WithContext(ctx).Error("session Session Deserialization err", zap.Error(err))
		return nil, err
	}
	return &sessionInfo, nil
}
func (r *session) DelSession(ctx context.Context, sessionId string) (err error) {
	ctx, span := trace.StartInternalSpan(ctx)
	defer func() { trace.TelemetrySpanEnd(span, err) }()
	result, err := r.redisClient.Del(ctx, sessionId).Result()
	if err != nil {
		log.WithContext(ctx).Error("session DelSession err", zap.Error(err))
		return err
	}
	log.Infof("result :%v", result)
	return nil
}
