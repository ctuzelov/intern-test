package repository

import (
	"auth-service/pkg/database/mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func UpdateUser(user models.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{"_id": user.ID} // Use the user's unique ID as the filter

	update := bson.M{
		"$set": bson.M{
			"name":          user.Name,
			"number":        user.Number,
			"date_of_birth": user.DateOfBirth,
		},
	}

	_, err := userCollection.UpdateOne(ctx, filter, update)

	return err
}
