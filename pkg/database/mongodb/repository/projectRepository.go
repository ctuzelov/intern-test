package repository

import (
	"auth-service/pkg/database/mongodb/models"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var projectCollection *mongo.Collection = OpenCollection(Client, "project")
var projectCountCollection *mongo.Collection = OpenCollection(Client, "project_count")

// Function that creates a project in the database and updates the project count
func CreateProject(project models.Project) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	// Increment the project ID in the project count collection
	_, err = projectCountCollection.UpdateOne(
		ctx,
		bson.M{},
		bson.M{"$inc": bson.M{"last_id": 1}},
		options.Update().SetUpsert(true),
	)
	if err != nil {
		return err
	}

	// Get the incremented ID from the project count collection
	var projectCount models.ProjectCount
	err = projectCountCollection.FindOne(ctx, bson.M{}).Decode(&projectCount)
	if err != nil {
		return err
	}

	project.ID = projectCount.Lastid

	// Insert the project into the project collection
	_, err = projectCollection.InsertOne(ctx, project)
	if err != nil {
		return err
	}

	return nil
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
	updatedProject.ID = oldOne.ID
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

// Function that returns all projects from the database
func GetAllProjects() ([]models.Project, error) {
	// Define the context for the operation
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Define options for the find operation
	findOptions := options.Find()

	// Perform the find operation to retrieve all project documents
	cursor, err := projectCollection.Find(ctx, bson.M{}, findOptions)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	// Iterate over the results and construct project objects
	var projects []models.Project
	for cursor.Next(ctx) {
		var project models.Project
		if err := cursor.Decode(&project); err != nil {
			return nil, err
		}
		projects = append(projects, project)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return projects, nil
}
