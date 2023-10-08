# Writing Repository

Within your application, a repository file plays a pivotal role by encapsulating the code responsible for interacting with the database.

## Repositories in baileys
Baileys offers a foundational repository that comes pre-equipped with essential CRUD (Create, Read, Update, Delete) operations. This repository can be extended to seamlessly integrate your custom code. The base repository in Baileys provides the following essential functions:
1. GetById: This function retrieves a database row based on a unique identifier (ID). It returns an error if the specified row is not found. 
2. Create: Accepting entity structs as input, this function orchestrates the creation of a new row in the database to store the provided data.
3. Update: Given entity structs, this function expertly manages the process of updating the corresponding row in the database and returns the updated data entity. 
4. Delete: This function carries out a soft delete operation on the database row associated with the provided ID. 
5. Search: The search functionality is designed to accept search queries written in bql (Baileys Query Language). It then provides a list of entities that match the specified query criteria. For a comprehensive guide on the search API, please refer to the search documentation.

## Creating a repository extending baileys repository
To create a repository powered by Baileys, both the repository's interface and struct should embed `BaseRepositoryInterface` from Baileys `repository` package. Below is an example illustrating the creation of a repository tailored for the `poll` entity, as demonstrated in the tutorial:
```go
package repository

import (
	"github.com/Anupam-dagar/baileys/repository"
	"polls/entity"
)

type PollRepositoryInterface interface {
	repository.BaseRepository[entity.PollPtr]
}

type pollRepository struct {
	repository.BaseRepository[entity.PollPtr]
}

func NewPollRepository() PollRepositoryInterface {
	pr := new(pollRepository)
	pr.BaseRepository = repository.NewBaseRepository[entity.PollPtr]()

	return pr
}
```