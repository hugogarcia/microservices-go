package data

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func New(mongo *mongo.Client) Models {
	client = mongo

	return Models{
		LogEntry: LogEntry{},
	}
}

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        string `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string `bson:"name" json:"name"`
	Data      string `bson:"data" json:"data"`
	CreatedAt string `bson:"created_at" json:"created_at"`
	UpdatedAt string `bson:"updated_at" json:"updated_at"`
}

func (l *LogEntry) Insert(ctx context.Context, entry LogEntry) error {
	entry.CreatedAt = time.Now().Format(time.RFC3339)
	entry.UpdatedAt = time.Now().Format(time.RFC3339)

	_, err := client.Database("logs").Collection("logs").InsertOne(ctx, entry)
	return err
}

func (l *LogEntry) GetAll(ctx context.Context) ([]LogEntry, error) {
	entries := []LogEntry{}
	opts := options.Find().SetSort(bson.D{{Key: "created_at", Value: -1}})

	cursor, err := client.Database("logs").Collection("logs").Find(ctx, opts)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}

func (l *LogEntry) GetOne(ctx context.Context) (*LogEntry, error) {
	entry := LogEntry{}

	filter := bson.D{{Key: "_id", Value: l.ID}}

	err := client.Database("logs").Collection("logs").FindOne(ctx, filter).Decode(&entry)
	if err != nil {
		return &LogEntry{}, err
	}

	return &entry, nil
}

func (l *LogEntry) DropCollection(ctx context.Context) error {
	return client.Database("logs").Collection("logs").Drop(ctx)
}

func (l *LogEntry) Update(ctx context.Context) (*mongo.UpdateResult, error) {
	filter := bson.D{{Key: "_id", Value: l.ID}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: l.Name}, {Key: "data", Value: l.Data}, {Key: "updated_at", Value: time.Now().Format(time.RFC3339)}}}}

	return client.Database("logs").Collection("logs").UpdateOne(ctx, filter, update)

}
