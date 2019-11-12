package userProfileRepo

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"loyalty_go/pkg/src/domains/userProfile"
	"time"
)

type profileRepo struct {
	ctx        context.Context
	collection *mongo.Collection
}

func NewRepo() *profileRepo {
	return &profileRepo{}
}

func (pRepo profileRepo) SaveNewProfile(profile *userProfile.Profile) error {
	collection, ctx, err := pRepo.getConnection()
	if err != nil {
		log.Fatal(err, "Connection error")
	}
	res, err := collection.InsertOne(ctx, profile)
	id := res.InsertedID
	fmt.Println(id)

	return err
}
func (pRepo profileRepo) FindByID(userID string) (*userProfile.Profile, error) {

	collection, _, err := pRepo.getConnection()
	if err != nil {
		log.Fatal(err, "Connection error")
	}
	user := &userProfile.Profile{}
	filter := bson.D{{"userid", userID}}
	if err := collection.FindOne(context.Background(), filter).Decode(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (pRepo profileRepo) UpdateProfile(userID string, profile *userProfile.Profile) error {
	collection, ctx, err := pRepo.getConnection()
	if err != nil {
		log.Fatal(err, "Connection error")
	}
	filter := bson.M{"userid": bson.M{"$eq": userID}}
	update := bson.M{"$set": bson.M{"userlevel": profile.UserLevel, "experience": profile.Experience, "currentdiscount": profile.CurrentDiscount}}
	_, err = collection.UpdateOne(ctx, filter, update)
	return err
}

func (pRepo profileRepo) getConnection() (*mongo.Collection, context.Context, error) {
	var err error
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		return nil, ctx, err
	}
	err = client.Ping(ctx, readpref.Primary())
	collection := client.Database("user_profiles").Collection("profile")
	return collection, ctx, nil
}
