package controllers

import (
	"auth-service/pkg/database/mongodb/models"
	"auth-service/pkg/database/mongodb/repository"
	"errors"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateProject(c *gin.Context) (err error) {
	var project models.Project

	if err = c.BindJSON(&project); err != nil {
		return
	}

	// Create the project in the database
	err = repository.CreateProject(project)
	if err != nil {
		return errors.New("error occurred while creating project")
	}

	return
}

// Function that updates a project
func UpdateProject(c *gin.Context) (err error) {
	var project models.Project

	if err = c.BindJSON(&project); err != nil {
		return
	}

	// Extract project ID from the request
	projectID, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return
	}
	
	// Update the project in the database
	err = repository.UpdateProject(projectID, project)
	if err != nil {
		return errors.New("error occurred while updating project")
	}

	return
}
