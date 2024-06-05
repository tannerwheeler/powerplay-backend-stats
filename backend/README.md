# Power Play Backend

![image](https://github.com/jak103/powerplay/assets/16627408/4ec3df62-d760-40c6-aa57-fa63eaaaf61b)


[![Go](https://github.com/jak103/powerplay/actions/workflows/go.yml/badge.svg?branch=main)](https://github.com/jak103/powerplay/actions/workflows/go.yml)

## Table of Contents
- [Adding and Updating API Endpoints](#adding-and-updating-api-endpoints)
- [Adding and Updating Unit Testing for Database Model](#adding-and-updating-unit-testing-for-database-model)
- [Adding and Updating Unit Testing for API Endpoint](#adding-and-updating-unit-testing-for-api-endpoint)
## Adding and Updating API Endpoints
There are 5 major steps to creating or editing an endpoint.
- [Database Model](#ensure-model-is-created-and-up-to-date)
- [API Handlers](#add-or-edit-api-endpoint-handlers)
- [Database Methods](#add-or-edit-any-db-methods-used-in-the-api-endpoint)
- [Documentation](#update-the-open-api-docs)
- [Testing](#adding-and-updating-unit-testing-for-database-model)

### Ensure Model is Created and Up to Date
1. Navigate to the models directory:  
   `backend/internal/models`
2. Find the model for the endpoint you are working on.
   - If the model is not there, create it
   - If the model is not accurate, make changes to ensure the model will reflect the needs.

### Add or Edit API Endpoint Handlers
1. Navigate to the server directory:  
    `backend/internal/server/apis`
2. Create or find the .go file for the Endpoint you are working on.
3. Implement or edit the API Handlers. The most common handlers are GET and POST. 
   
#### **GET Handler**
The following is a generic code block for a GET Handler.
``` go
func getGenericsHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	db := db.GetSession(c)
	generics, err := db.GetGenerics()
	if err != nil {
		log.WithErr(err).Alert("Failed to get all generics from the database")
		return err
	}

	return responder.OkWithData(c, generics)
}
```

#### **POST Handler**
The following is a generic code block for a POST Handler.
``` go
func postGenericHandler(c *fiber.Ctx) error {
	log := locals.Logger(c)
	log.Debug("body: %q", c.Request().Body())

	// Parse generic
	genericRequest := &models.Generic{}
	err := c.BodyParser(GenericRequest)
	if err != nil {
		log.WithErr(err).Alert("Failed to parse generic request payload")
		return responder.BadRequest(c, "Failed to parse generic request payload")
	}

	db := db.GetSession(c)
	err = db.CreateGeneric(genericRequest)
	if err != nil {
		log.WithErr(err).Alert("Failed to save generic request")
		return responder.InternalServerError(c)
	}

	return responder.Ok(c)
}
```
4. Add in any handlers to the init function
``` go
func init() {
	apis.RegisterHandler(fiber.MethodGet, "/generics", auth.Public, getGenericsHandler)
	apis.RegisterHandler(fiber.MethodPost, "/generics", auth.Public, postGenericHandler)
}
```

### Add or Edit any DB methods used in the API endpoint
1. Navigate to the db directory:  
    `backend/internal/db`
2. Create or find the .go file for the endpoint you are working on.
3. Implement or edit the methods within the .go file. These can include Get, Create

#### **Get Method**
The following is a generic code block for a Get method. The preload can be used to load any needed database relation.
``` go
func (s session) GetGeneric() ([]models.Generic, error) {
	generic := make([]models.generic, 0)
	err := s.connection.Preload("GenericPreloadNeed").Find(&generic)
	return resultsOrError(generic, err)
}
```
#### **Create Method**  
The following is a generic code block for a Create method.
``` go
func (s session) CreateGeneric(request *models.Generic) error {
	result := s.connection.Create(request)
	return result.Error
}
```

### Update the Open API docs
1. Navigate to the open api spec directory:  
   `static/oas/v1`
2. Create or update the corresponding .yml file to correctly reflect any changes you have made.

## Adding and Updating Unit Testing for Database Model
Taken from an in class walkthrough of adding a unit test for a database model. These tests will be using a docker spin up of the actual database for testing the database interfacing functions.

Uses the dockertest go package

## Adding and Updating Unit Testing for API Endpoint
Taken from an in class walkthrough of adding a unit test for an API endpoint. These tests will be using a mock up of the database models involved.

