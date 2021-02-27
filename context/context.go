package context

import (
	"go.mongodb.org/mongo-driver/bson"
	"pdf-backend/config"
	"pdf-backend/utils"
)

type IContext interface {
	Config() config.IConfig
	MongoDB(config config.IMongoConfig) IMongo
	UserID() string
	SetUserID(userID string)
	BindString(filter bson.M, params map[string]interface{}, key string) bson.M
	BindTime(filter bson.M, params map[string]interface{}, key string) bson.M
}

type Context struct {
	userID string
}

func NewContext() IContext {
	return &Context{}
}

func (c *Context) Config() config.IConfig {
	return config.NewConfig()
}

func (c *Context) MongoDB(config config.IMongoConfig) IMongo {
	return NewMongo(c.userID, config)
}

func (c *Context) UserID() string {
	return c.userID
}

func (c *Context) SetUserID(userID string) {
	c.userID = userID
}

func (c *Context) BindString(filter bson.M, params map[string]interface{}, key string) bson.M {
	if params[key] != nil {
		filter[key] = params[key].(string)
	}
	return filter
}

func (c *Context) BindTime(filter bson.M, params map[string]interface{}, key string) bson.M {
	if params[key] != nil {
		t := utils.NewTimeStampT(params[key].(string))
		filter[key] = t
	}
	return filter
}
