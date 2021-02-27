package users

import (
	"go.mongodb.org/mongo-driver/bson"
	"pdf-backend/context"
	"pdf-backend/utils"
	"time"
)

type IUserStore interface {
	GetAllUser() ([]bson.M, error)
	FindUser(userID string) (bson.M, error)
	CreateUser(username string, password string, firstName string, lastName string, role string) (bson.M, error)
	UpdateUser(userID string, firstName string, lastName string, role string) (bson.M, error)
	DeleteUser(userID string) error
	GetToken(token string) (bson.M, error)
}

type UserStore struct {
	ctx context.IContext
}

func NewUserStore(ctx context.IContext) IUserStore {
	return &UserStore{
		ctx: ctx,
	}
}

func (s *UserStore) GetAllUser() ([]bson.M, error) {
	m := s.ctx.MongoDB(s.ctx.Config().MongoConfig())
	results, err := m.Find("authentication", bson.M{
		"is_enabled": true,
	})
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *UserStore) FindUser(userID string) (bson.M, error) {
	m := s.ctx.MongoDB(s.ctx.Config().MongoConfig())
	result, err := m.FindOne("authentication", bson.M{
		"user_id": userID,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UserStore) CreateUser(username string, password string, firstName string, lastName string, role string) (bson.M, error) {
	userID := utils.GetUUID()
	hashedPassword, hErr := utils.HashPassword(password)
	if hErr != nil {
		return nil, hErr
	}
	m := s.ctx.MongoDB(s.ctx.Config().MongoConfig())
	_, err := m.InsertOne("authentication", bson.M{
		"user_id":    userID,
		"username":   username,
		"password":   hashedPassword,
		"first_name": firstName,
		"last_name":  lastName,
		"role":       role,
		"is_enabled": true,
		"created_at": time.Now().Unix(),
		"updated_at": time.Now().Unix(),
		"delete_at":  nil,
	})
	if err != nil {
		return nil, err
	}

	return bson.M{
		"user_id":    userID,
		"username":   username,
		"password":   hashedPassword,
		"first_name": firstName,
		"last_name":  lastName,
		"role":       role,
		"is_enabled": true,
		"created_at": time.Now().Unix(),
		"updated_at": time.Now().Unix(),
		"delete_at":  nil,
	}, nil
}
func (s *UserStore) UpdateUser(userID string, firstName string, lastName string, role string) (bson.M, error) {
	m := s.ctx.MongoDB(s.ctx.Config().MongoConfig())
	_, err := m.UpdateOne("authentication", bson.M{
		"user_id": userID,
	}, bson.M{
		"$set": bson.M{
			"first_name": firstName,
			"last_name":  lastName,
			"role":       role,
			"updated_at": time.Now().Unix(),
		},
	})
	if err != nil {
		return nil, err
	}

	return bson.M{
		"user_id":    userID,
		"first_name": firstName,
		"last_name":  lastName,
		"role":       role,
		"updated_at": time.Now().Unix(),
	}, nil
}

func (s *UserStore) DeleteUser(userID string) error {
	m := s.ctx.MongoDB(s.ctx.Config().MongoConfig())
	_, err := m.UpdateOne("authentication", bson.M{
		"user_id": userID,
	}, bson.M{
		"$set": bson.M{
			"is_enabled": false,
			"updated_at": time.Now().Unix(),
			"deleted_at": time.Now().Unix(),
		},
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *UserStore) GetToken(token string) (bson.M, error) {
	m := s.ctx.MongoDB(s.ctx.Config().MongoConfig())
	result, err := m.FindOne("tokes", bson.M{
		"token": token,
	})
	if err != nil {
		return nil, err
	}

	return result, nil
}
