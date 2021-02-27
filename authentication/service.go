package authentication

import (
	"go.mongodb.org/mongo-driver/bson"
	"pdf-backend/context"
	"time"
)

type ILoginService interface {
	Login(username string, password string) (bson.M, error)
	AssignToken(userId string, token string, expireAt time.Time) (bson.M, error)
}

type LoginService struct {
	ctx   context.IContext
	store ILoginStore
}

func NewLoginService(ctx context.IContext, store ILoginStore) ILoginService {
	return &LoginService{
		ctx:   ctx,
		store: store,
	}
}

func (s *LoginService) Login(username string, password string) (bson.M, error) {
	result, err := s.store.Login(username, password)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}

	return result, nil
}

func (s *LoginService) AssignToken(userId string, token string, expireAt time.Time) (bson.M, error) {
	result, err := s.store.AssignToken(userId, token, expireAt)
	if err != nil {
		return nil, err
	}
	if result == nil {
		return nil, nil
	}

	return result, nil
}
