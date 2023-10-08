# Writing services

In your application, a service file plays a pivotal role in housing the core business logic. This is where you have the freedom to craft custom code to implement the specific functionality your application demands.

## Services in baileys
Baileys provides a fundamental service that comes equipped with essential CRUD (Create, Read, Update, Delete) operations. You have the flexibility to extend this service and seamlessly integrate your custom code. The base service in Baileys offers the following critical functions:
1. GetById: This function retrieves a database row based on its id. It returns an error if the row is not found.
2. Create: Accepting entity structs as input, this function orchestrates the creation of a new database row to store the provided data. 
3. Update: Given entity struct and an id, this function expertly manages the process of updating the corresponding row in the database and returns the updated data entity. 
4. Delete: This function carries out a soft delete operation on the database row associated with the provided id. 
5. Search: The search functionality is designed to accept search queries written in bql (Baileys Query Language). It then furnishes a list of entities that align with the specified query criteria. For a comprehensive guide on the search API, please refer to the search documentation.

## Creating a service extending baileys service
To create a service powered by Baileys, both the service's interface and struct should embed `BaseServiceInterface` from Baileys `service` package. Below is an example showcasing the creation of a service tailored for the `poll` entity, as demonstrated in the tutorial:

```go
package service

import (
	"github.com/Anupam-dagar/baileys/service"
	"polls/entity"
)

type PollServiceInterface interface {
	service.BaseServiceInterface[entity.PollPtr]
}

type pollService struct {
	service.BaseServiceInterface[entity.PollPtr]
}

func NewPollService() PollServiceInterface {
	ps := new(pollService)
	ps.BaseServiceInterface = service.NewBaseService[entity.PollPtr]()

	return ps
}
```
