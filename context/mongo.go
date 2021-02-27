package context

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"pdf-backend/config"
	"time"
)

type IMongo interface {
	Initial() error
	Find(collectionName string, filter bson.M) ([]bson.M, error)
	FindOne(collectionName string, filter bson.M) (bson.M, error)
	InsertOne(collectionName string, data bson.M) (interface{}, error)
	UpdateOne(collectionName string, filter bson.M, data bson.M) (interface{}, error)
	DeleteOne(collectionName string, filter bson.M) (interface{}, error)
	logActivity(collectionName string, action string, actionDate time.Time)
}

type Mongo struct {
	mongoClient  *mongo.Client
	mongoConfig  config.IMongoConfig
	userID       string
	databaseName string
}

func NewMongo(userID string, config config.IMongoConfig) IMongo {
	return &Mongo{
		userID: userID,
		databaseName: "pdf-backend",
		mongoConfig:  config,
	}
}

func (m Mongo) Initial() error {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx,
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", m.mongoConfig.Endpoint(), m.mongoConfig.Port())),
		options.Client().SetAuth(options.Credential{
			Username: m.mongoConfig.Username(),
			Password: m.mongoConfig.Password(),
		}),
	)
	if err != nil {
		return err
	}
	m.mongoClient = client
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}

	return nil
}

func (m *Mongo) Find(collectionName string, filters bson.M) ([]bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, cerr := mongo.Connect(ctx,
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", m.mongoConfig.Endpoint(), m.mongoConfig.Port())),
		options.Client().SetAuth(options.Credential{
			Username: m.mongoConfig.Username(),
			Password: m.mongoConfig.Password(),
		}),
	)
	if cerr != nil {
		return nil, cerr
	}
	m.mongoClient = client
	cerr = client.Ping(ctx, readpref.Primary())
	if cerr != nil {
		return nil, cerr
	}

	collection := m.mongoClient.Database(m.databaseName).Collection(collectionName)
	ctx, cancel = context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, filters)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)
	items := make([]bson.M, 0)
	for cursor.Next(ctx) {
		result := bson.M{}
		err := cursor.Decode(&result)
		if err != nil {
			return nil, err
		}
		items = append(items, result)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	m.logActivity(collectionName, "Find", time.Now())
	return items, nil
}
func (m *Mongo) FindOne(collectionName string, filters bson.M) (bson.M, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, cerr := mongo.Connect(ctx,
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", m.mongoConfig.Endpoint(), m.mongoConfig.Port())),
		options.Client().SetAuth(options.Credential{
			Username: m.mongoConfig.Username(),
			Password: m.mongoConfig.Password(),
		}),
	)
	if cerr != nil {
		return nil, cerr
	}
	m.mongoClient = client
	cerr = client.Ping(ctx, readpref.Primary())
	if cerr != nil {
		return nil, cerr
	}

	result := bson.M{}
	collection := m.mongoClient.Database(m.databaseName).Collection(collectionName)
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := collection.FindOne(ctx, filters).Decode(&result)
	if err != nil {
		return nil, err
	}

	m.logActivity(collectionName, "FindOne", time.Now())
	return result, nil
}
func (m *Mongo) InsertOne(collectionName string, data bson.M) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, cerr := mongo.Connect(ctx,
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", m.mongoConfig.Endpoint(), m.mongoConfig.Port())),
		options.Client().SetAuth(options.Credential{
			Username: m.mongoConfig.Username(),
			Password: m.mongoConfig.Password(),
		}),
	)
	if cerr != nil {
		return nil, cerr
	}
	m.mongoClient = client
	cerr = client.Ping(ctx, readpref.Primary())
	if cerr != nil {
		return nil, cerr
	}

	collection := m.mongoClient.Database(m.databaseName).Collection(collectionName)
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return "", err
	}

	m.logActivity(collectionName, "InsertOne", time.Now())
	return res.InsertedID, nil
}

func (m *Mongo) UpdateOne(collectionName string, filter bson.M, data bson.M) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, cerr := mongo.Connect(ctx,
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", m.mongoConfig.Endpoint(), m.mongoConfig.Port())),
		options.Client().SetAuth(options.Credential{
			Username: m.mongoConfig.Username(),
			Password: m.mongoConfig.Password(),
		}),
	)
	if cerr != nil {
		return nil, cerr
	}
	m.mongoClient = client
	cerr = client.Ping(ctx, readpref.Primary())
	if cerr != nil {
		return nil, cerr
	}

	collection := m.mongoClient.Database(m.databaseName).Collection(collectionName)
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.UpdateOne(ctx, filter, data)
	if err != nil {
		return "", err
	}

	m.logActivity(collectionName, "UpdateOne", time.Now())
	return res.UpsertedID, nil
}

func (m *Mongo) DeleteOne(collectionName string, filter bson.M) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, cerr := mongo.Connect(ctx,
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", m.mongoConfig.Endpoint(), m.mongoConfig.Port())),
		options.Client().SetAuth(options.Credential{
			Username: m.mongoConfig.Username(),
			Password: m.mongoConfig.Password(),
		}),
	)
	if cerr != nil {
		return nil, cerr
	}
	m.mongoClient = client
	cerr = client.Ping(ctx, readpref.Primary())
	if cerr != nil {
		return nil, cerr
	}

	collection := m.mongoClient.Database(m.databaseName).Collection(collectionName)
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return "", err
	}

	m.logActivity(collectionName, "DeleteOne", time.Now())
	return res.DeletedCount, nil
}

func (m *Mongo) logActivity(collectionName string, action string, actionDate time.Time) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	client, cerr := mongo.Connect(ctx,
		options.Client().ApplyURI(fmt.Sprintf("mongodb://%s:%s", m.mongoConfig.Endpoint(), m.mongoConfig.Port())),
		options.Client().SetAuth(options.Credential{
			Username: m.mongoConfig.Username(),
			Password: m.mongoConfig.Password(),
		}),
	)
	if cerr != nil {
		log.Println(cerr)
	}
	m.mongoClient = client
	cerr = client.Ping(ctx, readpref.Primary())
	if cerr != nil {
		log.Println(cerr)
	}

	collection := m.mongoClient.Database(m.databaseName).Collection("activity_logs")
	ctx, cancel = context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err := collection.InsertOne(ctx, bson.M{
		"user_id":         m.userID,
		"database_name":   m.databaseName,
		"collection_name": collectionName,
		"action":          action,
		"action_date":     actionDate.Unix(),
	})
	if err != nil {
		log.Println(err)
	}
}
