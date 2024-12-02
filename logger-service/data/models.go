package data

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

var ctx = context.Background()

type Models struct {
	LogEntry LogEntry
}

type LogEntry struct {
	ID        string    `bson:"_id,omitempty" json:"id,omitempty"`
	Name      string    `bson:"name" json:"name"`
	Data      string    `bson:"data" json:"data"`
	CreatedAt time.Time `bson:"created_at" json:"created_at"`
	UpdatedAt time.Time `bson:"updated_at" json:"updated_at"`
}

func New(mongo *mongo.Client) Models {
	client = mongo

	return Models{
		LogEntry: LogEntry{},
	}
}

func (l *LogEntry) Insert(entry LogEntry) error {
	collection := client.Database("logs").Collection("logs")

	_, err := collection.InsertOne(ctx, LogEntry{
		Name:      entry.Name,
		Data:      entry.Data,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		log.Println("Error inserting entry: ", err)
		return err
	}

	return nil
}

func (l *LogEntry) All() ([]*LogEntry, error) {

	collection := client.Database("logs").Collection("logs")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	opts := options.Find()
	opts.SetSort(bson.D{{"created_at", -1}})

	cursor, err := collection.Find(ctx, bson.D{}, opts)
	if err != nil {
		log.Println("Finding docs failed: ", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var logs []*LogEntry

	for cursor.Next(ctx) {
		var items LogEntry
		err := cursor.Decode(&items)
		if err != nil {
			log.Print("Error decoding logs:", err)
			return nil, err
		} else {
			logs = append(logs, &items)
		}
	}

	return logs, nil
}

func (l *LogEntry) GetOne(id string) (*LogEntry, error) {
	collection := client.Database("logs").Collection("logs")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid document id: ", id)
		return nil, err
	}

	var entry LogEntry
	filter := bson.D{primitive.E{Key: "_id", Value: docID}}

	err = collection.FindOne(ctx, filter).Decode(&entry)
	if err != nil {
		return nil, err
	}

	return &entry, nil
}

func (l *LogEntry) Drop() error {
	collection := client.Database("logs").Collection("logs")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	err := collection.Drop(ctx)
	if err != nil {
		log.Println("Failed to drop collection: ", err)
		return err
	}

	return nil
}

func (l *LogEntry) Update(id string) (*mongo.UpdateResult, error) {
	collection := client.Database("logs").Collection("logs")

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	docID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println("Invalid document id: ", id)
		return nil, err
	}

	filter := bson.M{"id": docID}

	result, err := collection.UpdateOne(ctx, filter, bson.D{{Key: "$set", Value: bson.D{{Key: "name", Value: l.Name}, {Key: "data", Value: l.Data}, {Key: "updated_at", Value: time.Now()}}}})
	if err != nil {
		return nil, err
	}

	return result, nil
}
