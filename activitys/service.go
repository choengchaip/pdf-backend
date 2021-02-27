package activitys

import (
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"pdf-backend/context"
	"pdf-backend/users"
)

type IActivityService interface {
	GetAllActivities() ([]bson.M, error)
}

type ActivityService struct {
	ctx         context.IContext
	store       IActivityStore
	userService users.IUserService
}

func NewActivityService(ctx context.IContext, store IActivityStore, userService users.IUserService) IActivityService {
	return &ActivityService{
		ctx:         ctx,
		store:       store,
		userService: userService,
	}
}

func (s *ActivityService) GetAllActivities() ([]bson.M, error) {
	results, err := s.store.GetAllActivities()
	if err != nil {
		return nil, err
	}

	activities := make([]bson.M, len(results))
	for i, result := range results {
		if ok := result["user_id"]; ok != nil {
			activities[i] = result

			user, err := s.userService.FindUser(result["user_id"].(string))
			if err != nil {
				log.Println(err)
				continue
			}
			activities[i]["first_name"] = user["first_name"]
			activities[i]["last_name"] = user["last_name"]
		}
	}

	return activities, nil
}
