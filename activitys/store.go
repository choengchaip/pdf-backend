package activitys

import (
	"go.mongodb.org/mongo-driver/bson"
	"pdf-backend/context"
)

type IActivityStore interface {
	GetAllActivities() ([]bson.M, error)
}

type ActivityStore struct {
	ctx context.IContext
}

func NewActivityStore(ctx context.IContext) IActivityStore {
	return &ActivityStore{
		ctx: ctx,
	}
}

func (s *ActivityStore) GetAllActivities() ([]bson.M, error) {
	m := s.ctx.MongoDB(s.ctx.Config().MongoConfig())
	activities, err := m.Find("activity_logs", bson.M{})
	if err != nil {
		return nil, err
	}

	return activities, nil
}
