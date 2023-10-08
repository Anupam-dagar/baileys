# Writing controllers

In your application, a controller file serves as the container for the code responsible for handling API endpoints. It functions as the entry point where incoming requests are received and responses are generated for the client.

## Controllers in baileys
Baileys' controllers are designed to work seamlessly with entity structs, effortlessly constructing APIs for CRUD (Create, Read, Update, Delete) operations without any need for additional configuration. The base service in Baileys offers the following essential functions:
1. GetById: This function retrieves a database row based on its id. It returns an error if the row is not found. 
2. Create: Accepting an entity struct as input, this function creates a new database row to store the provided data. 
3. Update: Given an entity struct for an id, this function updates the corresponding row in the database and returns the updated data entity.
4. Delete: This function performs a soft delete on the database row associated with the provided id. 
5. Search: The search functionality is capable of processing search queries written in bql (Baileys Query Language). It returns a list of entities that match the specified query conditions. For detailed information on the search API, refer to the search documentation.

## Creating a controller by extending baileys controller
To create a controller powered by Baileys, both the controller's interface and struct should embed `BaseControllerInterface` from Baileys `controller` package. Below is an example demonstrating how to create a controller for the poll entity, as outlined in the tutorial:
```go
package controller

import (
	"github.com/Anupam-dagar/baileys/controller"
	"github.com/gin-gonic/gin"
	"polls/entity"
	"polls/service"
)

type PollControllerInterface interface {
	controller.BaseControllerInterface
}

type pollController struct {
	controller.BaseControllerInterface
	PollService service.PollServiceInterface
}

func NewPollController(rg *gin.RouterGroup) PollControllerInterface {
	pc := new(pollController)
	pc.BaseControllerInterface = controller.NewBaseController[entity.PollPtr](rg)

	return pc
}
```