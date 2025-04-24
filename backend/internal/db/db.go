package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	"backend/internal/config"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	client     *mongo.Client
	clientOnce sync.Once
)

func СonnectCollection(cfg *config.Config) (*mongo.Collection, error) {
	var err error
	clientOnce.Do(func() {
		connectionString := fmt.Sprintf("mongodb+srv://%s:%s@%s", cfg.DBuser, cfg.DBpassword, cfg.DBmongoname)
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		client, err = mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
		if err != nil {
			err = fmt.Errorf("ошибка подключения к MongoDB: %w", err)
			return
		}
	})

	if err != nil {
		return nil, err
	}

	collection := client.Database(cfg.DBname).Collection(cfg.DBcollection)
	return collection, nil
}
