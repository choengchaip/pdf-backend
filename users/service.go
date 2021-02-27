package users

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"pdf-backend/context"
)

type IUserService interface {
	GetAllUser() ([]bson.M, error)
	FindUser(userID string) (bson.M, error)
	CreateUser(username string, password string, firstName string, lastName string, role string) (bson.M, error)
	UpdateUser(userID string, firstName string, lastName string, role string) (bson.M, error)
	DeleteUser(userId string) error
	GetUserByToken(token string) (bson.M, error)
}

type UserService struct {
	ctx   context.IContext
	store IUserStore
}

func NewUserService(ctx context.IContext, store IUserStore) IUserService {
	return &UserService{
		ctx:   ctx,
		store: store,
	}
}

func (s *UserService) GetAllUser() ([]bson.M, error) {
	results, err := s.store.GetAllUser()
	if err != nil {
		return nil, err
	}

	return results, nil
}

func (s *UserService) FindUser(userID string) (bson.M, error) {
	result, err := s.store.FindUser(userID)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (s *UserService) CreateUser(username string, password string, firstName string, lastName string, role string) (bson.M, error) {
	result, err := s.store.CreateUser(username, password, firstName, lastName, role)
	if err != nil {
		return nil, err
	}

	return result, nil
}
func (s *UserService) UpdateUser(userID string, firstName string, lastName string, role string) (bson.M, error) {
	result, err := s.store.UpdateUser(userID, firstName, lastName, role)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return result, nil
}
func (s *UserService) DeleteUser(userId string) error {
	err := s.store.DeleteUser(userId)
	if err != nil {
		return err
	}

	return nil
}

func (s *UserService) GetUserByToken(token string) (bson.M, error) {
	tokenAttr, err := s.store.GetToken(token)
	if err != nil {
		return nil, err
	}

	user, err := s.FindUser(tokenAttr["user_id"].(string))
	if err != nil {
		return nil, err
	}

	return user, nil
}
