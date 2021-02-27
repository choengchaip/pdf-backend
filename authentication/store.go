package authentication

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"pdf-backend/context"
	"pdf-backend/utils"
	"time"
)

type ILoginStore interface {
	Login(username string, password string) (bson.M, error)
	AssignToken(userId string, token string, expireAt time.Time) (bson.M, error)
}

type LoginStore struct {
	ctx context.IContext
}

func NewLoginStore(ctx context.IContext) ILoginStore {
	return &LoginStore{
		ctx: ctx,
	}
}

func (s *LoginStore) Login(username string, password string) (bson.M, error) {
	m := s.ctx.MongoDB(s.ctx.Config().MongoConfig())
	user, err := m.FindOne("authentication", bson.M{
		"username": username,
	})
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}
	if !user["is_enabled"].(bool) {
		return nil, nil
	}

	if !utils.CheckPasswordHas(password, user["password"].(string)) {
		log.Println("AdminOrEmailWrong")
		return nil, nil
	}

	tokenID := utils.GetUUID()
	token, expireAt := utils.GetJWTFromUser(user["user_id"].(string), tokenID, time.Now().Add(time.Hour*24), s.ctx.Config().Secret())

	tokenM, err := s.AssignToken(user["user_id"].(string), token, expireAt)
	if err != nil {
		return nil, err
	}

	tokenM["role"] = user["role"]
	return tokenM, nil
}

func (s *LoginStore) AssignToken(userId string, token string, expireAt time.Time) (bson.M, error) {
	m := s.ctx.MongoDB(s.ctx.Config().MongoConfig())
	_, err := m.InsertOne("tokens", bson.M{
		"token":     token,
		"user_id":   userId,
		"expire_at": expireAt.Unix(),
	})
	if err != nil {
		return nil, err
	}

	return bson.M{
		"user_id": userId,
		"token":   token,
	}, nil
}
