package repository

import (
	"auth-service/pkg/database/mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var projectCollection *mongo.Collection = OpenCollection(Client, "project")
var projectCountCollection *mongo.Collection = OpenCollection(Client, "project_count")

// Function that creates a project in the database and updates the project count
func CreateProject(project models.Project) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	project.Id = primitive.NewObjectID()

	_, err = projectCollection.InsertOne(ctx, project)
	if err != nil {
		return
	}

	// Increment the last project ID in the project count collection
	update := bson.M{"$inc": bson.M{"id": 1}}
	_, err = projectCountCollection.UpdateOne(ctx, bson.M{}, update, options.Update().SetUpsert(true))
	if err != nil {
		return
	}

	return
}

// Function that updates a project in the database
func UpdateProject(projectId int, updatedProject models.Project) error {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{"id": projectId}
	oldOne, err := GetProjectByID(projectId)
	if err != nil {
		return err
	}
	updatedProject.Id = oldOne.Id
	update := bson.M{"$set": updatedProject}

	_, err = projectCollection.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}

	return nil
}

// Function that retrieves a project from the database by its custom ID
func GetProjectByID(projectId int) (foundProject models.Project, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	filter := bson.M{"id": projectId}
	err = projectCollection.FindOne(ctx, filter).Decode(&foundProject)
	if err != nil {
		return
	}

	return
}
