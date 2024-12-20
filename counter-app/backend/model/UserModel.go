package models

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
}

type UserModel struct {
	Collection *mongo.Collection
}

// CreateUser creates a new user
func (um *UserModel) CreateUser(username, password string) (primitive.ObjectID, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return primitive.NilObjectID, err
	}

	user := User{
		Username: username,
		Password: string(hashedPassword),
	}

	res, err := um.Collection.InsertOne(context.Background(), user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

// FindUserByUsername finds a user by username
func (um *UserModel) FindUserByUsername(username string) (*User, error) {
	var user User
	err := um.Collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// ValidatePassword compares a hashed password with a plaintext one
func (um *UserModel) ValidatePassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
