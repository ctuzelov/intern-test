package controllers

import (
	"auth-service/pkg/database/mongodb/models"
	"auth-service/pkg/database/mongodb/repository"
	"errors"

	"github.com/gin-gonic/gin"
)

func EditPersonalData(c *gin.Context) error {
	var userData models.User
	email := c.PostForm("email") // Get the email from the request form data
	// Bind the updated user data from the request body
	if err := c.BindJSON(&userData); err != nil {
		return err
	}

	claims, ok := c.Get("user_email")
	if !ok {
		return errors.New("unable to retrieve user email from JWT claims")
	}
	userEmail := claims.(string)

	// Ensure that the user making the request is authorized to edit their own data
	if email != userEmail {
		return errors.New("unauthorized to edit data for this user")
	}

	// Check if the user with the provided email exists
	existingUser, err := repository.GetUserByEmail(&email)
	if err != nil {
		return err
	}
	if existingUser == (models.User{}) {
		return errors.New("user not found")
	}

	// Update the user's personal data
	if userData.Name != nil {
		existingUser.Name = userData.Name
	}
	if userData.Number != nil {
		existingUser.Number = userData.Number
	}
	if userData.DateOfBirth != nil {
		existingUser.DateOfBirth = userData.DateOfBirth
	}

	// Save the updated user data to the database
	if err := repository.UpdateUser(existingUser); err != nil {
		return err
	}

	return nil
}
