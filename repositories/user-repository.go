package repositories

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsoncodec"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"rakshit.dev/gin-rest-api-boilerplate/errors"
	"rakshit.dev/gin-rest-api-boilerplate/models"
)

type User = models.User
type UserReq = models.UserReq
type Error = errors.Error

// UserRepository ...
type UserRepository interface {
	GetUser(userID string) (*User, *Error)
	GetUsers() ([]User, *Error)
	DeleteUser(userID string) (*User, *Error)
	CreateUser(userReq UserReq) (*User, *Error)
}

type userRepository struct {
	usersCollection   *mongo.Collection
	bsonDecodeContext bsoncodec.DecodeContext
}

// UserRepository ...
func NewUserRepository(mongoDB *mongo.Database) UserRepository {
	return &userRepository{
		usersCollection: mongoDB.Collection("users"),
		bsonDecodeContext: bsoncodec.DecodeContext{
			Registry: bson.DefaultRegistry,
			Truncate: true,
		},
	}
}

func (u *userRepository) GetUser(userID string) (*User, *Error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		err := errors.CreateError(err.Error())
		return nil, &err
	}
	res := u.usersCollection.FindOne(context.Background(), bson.M{"_id": userObjectID})
	user := &User{}
	err = res.Err()
	if err == mongo.ErrNoDocuments {
		userError := errors.ErrUserNotFound
		err := errors.CreateError(userError.Message())
		return nil, &err
	}
	if err != nil {
		err := errors.CreateError(err.Error())
		return nil, &err
	}
	err = res.Decode(user)
	if err != nil {
		err := errors.CreateError(err.Error())
		return nil, &err
	}
	return user, nil
}

func (u *userRepository) GetUsers() ([]User, *Error) {
	cursor, err := u.usersCollection.Find(context.Background(), bson.M{})
	if err != nil {
		err := errors.CreateError(err.Error())
		return nil, &err
	}
	var users []User
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		curr := &User{}
		err := bson.UnmarshalWithContext(u.bsonDecodeContext, cursor.Current, curr)
		if err != nil {
			err := errors.CreateError(err.Error())
			return nil, &err
		}
		users = append(users, *curr)
	}
	return users, nil
}

func (u *userRepository) DeleteUser(userID string) (*User, *Error) {
	userObjectID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		err := errors.CreateError(err.Error())
		return nil, &err
	}
	user, userErr := u.GetUser(userID)
	if userErr != nil {
		return nil, userErr
	}
	_, err = u.usersCollection.
		DeleteOne(context.Background(), bson.M{"_id": userObjectID})
	if err != nil {
		err := errors.CreateError(err.Error())
		return nil, &err
	}
	return user, nil
}

func (u *userRepository) CreateUser(userReq UserReq) (*User, *Error) {
	currTime := time.Now()
	user := User{
		CreatedAt: currTime,
		UpdatedAt: currTime,
		UserReq:   userReq,
	}
	insertResult, err := u.usersCollection.
		InsertOne(context.Background(), user, options.InsertOne().SetBypassDocumentValidation(false))
	if err != nil {
		err := errors.CreateError(err.Error())
		return nil, &err
	}
	userID := insertResult.InsertedID.(primitive.ObjectID).Hex()
	return u.GetUser(userID)
}
