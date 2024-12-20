package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Counter struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Value int                `bson:"value" json:"value"`
}

type CounterModel struct {
	Collection *mongo.Collection
}

// FetchAllCounters fetches all counters from MongoDB
func (cm *CounterModel) FetchAllCounters() ([]Counter, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := cm.Collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var counters []Counter
	if err := cursor.All(ctx, &counters); err != nil {
		return nil, err
	}
	return counters, nil
}

// CreateCounter inserts a new counter into MongoDB
func (cm *CounterModel) CreateCounter(counter Counter) (primitive.ObjectID, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	result, err := cm.Collection.InsertOne(ctx, counter)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

// UpdateCounter updates the value of a counter in MongoDB
func (cm *CounterModel) UpdateCounter(id primitive.ObjectID, value int) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	update := bson.M{"$set": bson.M{"value": value}}
	_, err := cm.Collection.UpdateOne(ctx, bson.M{"_id": id}, update)
	return err
}

// DeleteCounter deletes a counter from MongoDB
func (cm *CounterModel) DeleteCounter(id primitive.ObjectID) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := cm.Collection.DeleteOne(ctx, bson.M{"_id": id})
	return err
}
