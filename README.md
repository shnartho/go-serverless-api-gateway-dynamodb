# go-serverless-api-gateway-dynamodb
This repository contains a serverless web API project implemented in Golang, utilizing AWS API Gateway for endpoint management and DynamoDB for data storage. 

## Function Call Flow

The program is structured as follows:

### Main.go

- The program starts in `main.go`.
- In `main.go`, it checks the HTTP method of the incoming request.
  - If the HTTP method is "GET," it calls `handlers.GetUser`.
  - If the HTTP method is "POST," it calls `handlers.CreateUser`.
  - If the HTTP method is "PUT," it calls `handlers.UpdateUser`.
  - If the HTTP method is "DELETE," it calls `handlers.DeleteUser`.
  - For any other method, it calls `handlers.UnhandledMethod`.

### Handlers Package

- In the `handlers` package:
  - `handlers.GetUser` calls functions in the `user` package to fetch user data.
  - `handlers.CreateUser` calls functions in the `user` package to create a new user.
  - `handlers.UpdateUser` calls functions in the `user` package to update a user's data.
  - `handlers.DeleteUser` calls functions in the `user` package to delete a user.
  - `handlers.UnhandledMethod` is a utility function to handle unsupported HTTP methods.

### User Package

- In the `user` package:
  - `user.FetchUser` retrieves user data from a DynamoDB table.
  - `user.FetchUsers` retrieves a list of users from the DynamoDB table.
  - `user.CreateUser` creates a new user in the DynamoDB table.
  - `user.UpdateUser` updates user data in the DynamoDB table.
  - `user.DeleteUser` deletes a user from the DynamoDB table.

### Validators Package

- Additionally, the `validators` package is used to check if an email is valid or not.

### ApiResponse Package

- The `apiResponse` package is used to generate API responses with appropriate status codes and JSON bodies.

This structure outlines the flow of function calls in the program, starting from `main.go` and branching out to the functions in the `handlers`, `user`, `validators`, and `apiResponse` packages.

![golang](https://fgp.dev/static/media/GolangDevelopmentBanner.aba7a1d6.jpg)