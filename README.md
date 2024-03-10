# Golang API Application: 5, 6, 7 tasks in one

## API Routes

### Authentication

- `POST /signup`: Register a new user.
- `POST /signin`: Login with existing credentials.
- `POST /refresh-token`: Refresh JWT token.

### User Profile

- `POST /edit`: Update user profile.

### Project Management

- `POST /create-project`: Create a new project.
- `PUT /update-project/:id`: Update an existing project. Requires admin privileges.

#### Admin data
```json
{
  "email": "chingizkhan@gmail.com",
  "password": "123123"
}
```

## Running the Application

To run the application, execute the following command in your terminal:

```
go run ./cmd
```

Make sure to set the necessary environment variables in a `.env` file before running the application.

## Dependencies

- [Gin](https://github.com/gin-gonic/gin): HTTP web framework.
- [MongoDB Go Driver](https://github.com/mongodb/mongo-go-driver): Official MongoDB driver for Go.
- [JWT Go](https://github.com/dgrijalva/jwt-go): Library for JSON Web Tokens.

## Environment Variables

The application requires the following environment variables to be set:

- `MONGO_URI`: MongoDB connection URI.
- `PORT`: Port number for the server.
- `JWT_SECRET`: Secret key for JWT token generation.
- `GIN_MODE`: Gin mode (`debug` or `release`).

Ensure these environment variables are correctly set in a `.env` file before running the application.
