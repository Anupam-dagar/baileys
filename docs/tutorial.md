# Polls App

To provide a more comprehensive understanding of how Baileys works, we will create a simple polls app. This app allows users to create polls, vote on them, and view all created polls. Throughout the development of the app, we will delve into the core components of Baileys.

## Setup a polls baileys project

1. Begin by creating a new directory named `polls`.
2. Initialize a Go module:

```bash
$ go mod init polls
```

3. Add baileys as a dependency

```bash
$ go get -u github.com/Anupam-dagar/baileys
```

## Database Setup

1. Create a new PostgreSQL database named `polls`.

```sql
CREATE DATABASE polls;
```

2. Create the required tables.
   - `polls` - Stores information about polls created by users.
   ```sql
   CREATE TABLE IF NOT EXISTS polls (
        id varchar(255) NOT NULL PRIMARY KEY,
        title varchar(255) NOT NULL,
        created_at timestamptz NOT NULL DEFAULT NOW(),
        updated_at timestamptz NOT NULL DEFAULT NOW(),
        deleted_at timestamptz NULL,
        created_by varchar(255) NOT NULL,
        updated_by varchar(255) NOT NULL,
        deleted_by varchar(255) NULL
   );
   ```
   - `poll_options` - Contains options for the polls.
   ```sql
   CREATE TABLE IF NOT EXISTS poll_options (
        id varchar(255) NOT NULL PRIMARY KEY,
        poll_id varchar(255) NOT NULL,
        title varchar(255) NOT NULL,
        created_at timestamptz NOT NULL DEFAULT NOW(),
        updated_at timestamptz NOT NULL DEFAULT NOW(),
        deleted_at timestamptz NULL,
        created_by varchar(255) NOT NULL,
        updated_by varchar(255) NOT NULL,
        deleted_by varchar(255) NULL
    );
   ```

   - `votes` - Records votes received for poll options.
   ```sql
   CREATE TABLE IF NOT EXISTS votes (
        id varchar(255) NOT NULL PRIMARY KEY,
        poll_id varchar(255) NOT NULL,
        poll_option_id varchar(255) NOT NULL,
        created_at timestamptz NOT NULL DEFAULT NOW(),
        updated_at timestamptz NOT NULL DEFAULT NOW(),
        deleted_at timestamptz NULL,
        created_by varchar(255) NOT NULL,
        updated_by varchar(255) NOT NULL,
        deleted_by varchar(255) NULL
    );
   ```

## Directory Structure

Set up the directory structure with the following commands:

```bash
mkdir config
mkdir controller
mkdir entity
mkdir service
mkdir route
mkdir repository
touch config/dev.yaml
touch main.go
```

## Configuration Setup

Create a `dev.yaml` file under the `config` directory and use the following configuration:

```yaml
server:
  port: "8080"
  base_api_path: "/api"

database:
  host: localhost
  port: 5432
  username: postgres
  password: postgres
  name: polls
```

## Entity Setup

Create the following files under the `entity` directory:

- `poll.go`
   ```go
   package entity

   import (
	  "github.com/Anupam-dagar/baileys/entity"
	  "github.com/oleiade/reflections"
   )

   type Poll struct {
	  entity.BaseModel
	  Title string `gorm:"column:title" json:"title"`
   }

   type PollPtr = *Poll

   func (p *Poll) GetModel() interface{} {
	  return &Poll{}
   }

   func (p *Poll) SetCol(field string, val interface{}) error {
	  return reflections.SetField(p, field, val)
   }
   ```
- `poll_option.go`
   ```go
   package entity

   import (
        "github.com/Anupam-dagar/baileys/entity"
        "github.com/oleiade/reflections"
   )

   type PollOption struct {
        entity.BaseModel
        PollId string `gorm:"column:poll_id" json:"pollId"`
        Title  string `gorm:"column:title" json:"title"`
   }

   type PollOptionPtr = *PollOption

   func (p *PollOption) GetModel() interface{} {
        return &PollOption{}
   }

   func (p *PollOption) SetCol(field string, val interface{}) error {
        return reflections.SetField(p, field, val)
   }
   ```
   
- `vote.go`
   ```go
  package entity

   import (
        "github.com/Anupam-dagar/baileys/entity"
        "github.com/oleiade/reflections"
   )

   type Vote struct {
        entity.BaseModel
        PollId       string `gorm:"column:poll_id" json:"pollId"`
        PollOptionId string `gorm:"column:poll_option_id" json:"pollOptionId"`
   }

   type VotePtr = *Vote

   func (v *Vote) GetModel() interface{} {
        return &Vote{}
   }

   func (v *Vote) SetCol(field string, val interface{}) error {
        return reflections.SetField(v, field, val)
   }
   ```

## Controller Setup
- `poll.go`
   
  ```go
   package controller

   import (
	  "github.com/Anupam-dagar/baileys/controller"
	  "github.com/gin-gonic/gin"
	  "polls/entity"
   )

   type PollControllerInterface interface {
	  controller.BaseControllerInterface
   }

   type pollController struct {
	  controller.BaseControllerInterface
   }

   func NewPollController(rg *gin.RouterGroup) PollControllerInterface {
	  pc := new(pollController)
	  pc.BaseControllerInterface = controller.NewBaseController[entity.PollPtr](rg)

	  return pc
   }
   ```

- `poll_option.go`
   ```go
   package controller

   import (
	  "github.com/Anupam-dagar/baileys/controller"
	  "github.com/gin-gonic/gin"
	  "polls/entity"
   )

   type PollOptionControllerInterface interface {
	  controller.BaseControllerInterface
   }

   type pollOptionController struct {
	  controller.BaseControllerInterface
   }

   func NewPollOptionController(rg *gin.RouterGroup) PollOptionControllerInterface {
	  poc := new(pollOptionController)
	  poc.BaseControllerInterface = controller.NewBaseController[entity.PollOptionPtr](rg)

	  return poc
   }
   ```

- `vote.go`
   ```go
   package controller

   import (
	  "github.com/Anupam-dagar/baileys/controller"
	  "github.com/gin-gonic/gin"
	  "polls/entity"
   )

   type VoteControllerInterface interface {
	  controller.BaseControllerInterface
   }

   type voteController struct {
	  controller.BaseControllerInterface
   }

   func NewVoteController(rg *gin.RouterGroup) VoteControllerInterface {
	  vc := new(voteController)
	  vc.BaseControllerInterface = controller.NewBaseController[entity.VotePtr](rg)

	  return vc
   }
   ```

## Route Setup
-`poll.go`
   ```go
    package route

    import (
	    "github.com/Anupam-dagar/baileys/constant/types"
	    "github.com/gin-gonic/gin"
	    "polls/controller"
    )

    func PollRoutes(routerGroup *gin.RouterGroup) {
	    router := routerGroup.Group("/polls")
	    {
		    controller.NewPollController(router)
	    }
    }

    func BPollRoutes(routerGroup *gin.RouterGroup) types.RouteFunc {
	    return func() (*gin.RouterGroup, func(rg *gin.RouterGroup)) {
		    return routerGroup, PollRoutes
	    }
    }
   ```

- `poll_option.go`
   ```go
   package route

   import (
	  "github.com/Anupam-dagar/baileys/constant/types"
	  "github.com/gin-gonic/gin"
	  "polls/controller"
   )

   func PollOptionRoutes(routerGroup *gin.RouterGroup) {
	  router := routerGroup.Group("/poll-options")
	  {
		 controller.NewPollOptionController(router)
	  }
   }

   func BPollOptionRoutes(routerGroup *gin.RouterGroup) types.RouteFunc {
	  return func() (*gin.RouterGroup, func(rg *gin.RouterGroup)) {
		 return routerGroup, PollOptionRoutes
	  }
   }
   ```

- `vote.go`
   ```go
   package route

   import (
	  "github.com/Anupam-dagar/baileys/constant/types"
	  "github.com/gin-gonic/gin"
	  "polls/controller"
   )

   func VoteRoutes(routerGroup *gin.RouterGroup) {
	  router := routerGroup.Group("/votes")
	  {
		 controller.NewVoteController(router)
	  }
   }

   func BVoteRoutes(routerGroup *gin.RouterGroup) types.RouteFunc {
	  return func() (*gin.RouterGroup, func(rg *gin.RouterGroup)) {
		 return routerGroup, VoteRoutes
	  }
   }
   ```

- `route.go` - provides router groups and installs routes to the baileys
   ```go
   package route

   import "github.com/Anupam-dagar/baileys/server"

   func SetupRoutes() {
	  rootRouterGroup := server.GetGinEngine().GetRootRouterGroup()

	  server.AddRoute(BPollRoutes(rootRouterGroup))
	  server.AddRoute(BPollOptionRoutes(rootRouterGroup))
	  server.AddRoute(BVoteRoutes(rootRouterGroup))
   }
   ```

## Setting up the server
Create a `main.go` file in the root directory to run the server:
```go
package main

import (
	"github.com/Anupam-dagar/baileys/server"
	"polls/route"
)

func main() {
	server.NewGinEngine().InitGinApp(route.SetupRoutes).RunServer()
}
```

## Running the server
Open a terminal at the root of the directory and run the following command
```bash
go run main.go
```
You should see the following output in the terminal if everything is set up correctly
```bash
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/polls/search         --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func1 (3 handlers)
[GIN-debug] GET    /api/polls/:id            --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func2 (3 handlers)
[GIN-debug] POST   /api/polls                --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func3 (3 handlers)
[GIN-debug] PUT    /api/polls/:id            --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func4 (3 handlers)
[GIN-debug] DELETE /api/polls/:id            --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func5 (3 handlers)
[GIN-debug] POST   /api/poll-options/search  --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func1 (3 handlers)
[GIN-debug] GET    /api/poll-options/:id     --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func2 (3 handlers)
[GIN-debug] POST   /api/poll-options         --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func3 (3 handlers)
[GIN-debug] PUT    /api/poll-options/:id     --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func4 (3 handlers)
[GIN-debug] DELETE /api/poll-options/:id     --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func5 (3 handlers)
[GIN-debug] POST   /api/votes/search         --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func1 (3 handlers)
[GIN-debug] GET    /api/votes/:id            --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func2 (3 handlers)
[GIN-debug] POST   /api/votes                --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func3 (3 handlers)
[GIN-debug] PUT    /api/votes/:id            --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func4 (3 handlers)
[GIN-debug] DELETE /api/votes/:id            --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func5 (3 handlers)
{"time":"2023-10-01T17:10:21.156256+05:30","level":"INFO","prefix":"-","file":"server.go","line":"76","message":"Starting server on port: 8080"}
```

## Testing the polls app

### Creating a poll
To create a poll, use the `POST` endpoint `/api/polls` with the following request body:
```json
{
  "title": "Which is your favourite programming language"
}
```
You should receive the following response:
```json
{
   "status": {
      "code": 200,
      "message": "Successfully created",
      "type": "success"
   },
   "data": {
  "id": "f65ee8d5-57e8-4c6c-9ade-c65bad04b383",
  "title": "Which is your favourite programming language",
  "createdAt": "2021-10-01T17:14:21.156256+05:30",
  "updatedAt": "2021-10-01T17:14:21.156256+05:30",
  "createdBy": "",
  "updatedBy": "",
  "deletedBy": ""
}
}
```

### Creating a poll option
To create a poll option, use the `POST` endpoint `/api/poll-options` with the following request body:
```json
{
  "pollId": "f65ee8d5-57e8-4c6c-9ade-c65bad04b383",
  "title": "Go"
}
```
You should see the following response
```json
{
   "status": {
      "code": 200,
      "message": "Successfully created",
      "type": "success"
   },
   "data": {
  "id": "6bde533f-9b10-4ec4-87f8-dcc371b89948",
  "pollId": "f65ee8d5-57e8-4c6c-9ade-c65bad04b383",
  "title": "Go",
  "createdAt": "2021-10-01T17:14:21.156256+05:30",
  "updatedAt": "2021-10-01T17:14:21.156256+05:30",
  "createdBy": "",
  "updatedBy": "",
  "deletedBy": ""
}
}
```

### Creating a vote for the poll
To create a vote, use the `POST` endpoint `/api/votes` with the following request body:
```json
{
  "pollId": "f65ee8d5-57e8-4c6c-9ade-c65bad04b383",
  "pollOptionId": "6bde533f-9b10-4ec4-87f8-dcc371b89948"
}
```
You should receive the following response:
```json
{
   "status": {
      "code": 200,
      "message": "Successfully created",
      "type": "success"
   },
   "data": {
  "id": "419bdcf4-1236-46f0-864d-bb9291ea5dda",
  "pollId": "f65ee8d5-57e8-4c6c-9ade-c65bad04b383",
  "pollOptionId": "6bde533f-9b10-4ec4-87f8-dcc371b89948",
  "createdAt": "2021-10-01T17:14:21.156256+05:30",
  "updatedAt": "2021-10-01T17:14:21.156256+05:30",
  "createdBy": "",
  "updatedBy": "",
  "deletedBy": ""
}
}
```

### Retrieve a poll by id
To retrieve a poll by `id`, use the `GET` endpoint `/api/polls/:id` where `:id` will be replaced by a poll id (eg: `f65ee8d5-57e8-4c6c-9ade-c65bad04b383`)
You should receive the following response
```json
{
   "status": {
      "code": 200,
      "message": "Successfully fetched by Id",
      "type": "success"
   },
   "data": {
  "id": "f65ee8d5-57e8-4c6c-9ade-c65bad04b383",
  "title": "Which is your favourite programming language",
  "createdAt": "2021-10-01T17:14:21.156256+05:30",
  "updatedAt": "2021-10-01T17:14:21.156256+05:30",
  "createdBy": "",
  "updatedBy": "",
  "deletedBy": ""
}
}
```

### Retrieving all the polls
To retrieve all the polls, use the `POST` endpoint `/api/polls/search` with the following request body
```json
{
  "pagination": {
    "page": 1,
    "limit": 10
  }
}
```
You should receive the following response
```json
{
   "status": {
      "code": 200,
      "message": "Successfully searched",
      "type": "success",
      "totalCount": 0
   },
  "data": [
    {
      "id": "f65ee8d5-57e8-4c6c-9ade-c65bad04b383",
      "title": "Which is your favourite programming language",
      "createdAt": "2021-10-01T17:14:21.156256+05:30",
      "updatedAt": "2021-10-01T17:14:21.156256+05:30",
      "createdBy": "",
      "updatedBy": "",
      "deletedBy": ""
    }
  ]
}
```
To know more about search api, refer to [search api](#search-api).

### Update a poll
To update a poll, use the `PUT` endpoint `/api/polls/:id` where `:id` will be replaced by a poll id (eg: `f65ee8d5-57e8-4c6c-9ade-c65bad04b383`) with the following request body
```json
{
  "title": "Which is your favourite programming language?"
}
```
You should receive the following response
```json
{
   "status": {
      "code": 200,
      "message": "Successfully updated",
      "type": "success"
   },
   "data": {
  "id": "f65ee8d5-57e8-4c6c-9ade-c65bad04b383",
  "title": "Which is your favourite programming language?",
  "createdAt": "2021-10-01T17:14:21.156256+05:30",
  "updatedAt": "2021-10-01T17:14:21.156256+05:30",
  "createdBy": "",
  "updatedBy": "",
  "deletedBy": ""
}
}
```

### Delete a poll
To delete a poll, use the `DELETE` endpoint `/api/polls/:id` where `:id` will be replaced by a poll id (eg: `f65ee8d5-57e8-4c6c-9ade-c65bad04b383`)
You should receive the following response
```json
{
   "status": {
      "code": 200,
      "message": "Successfully deleted",
      "type": "success"
   },
   "data": null
}
```

## Postman Collection
You can find the postman collection for the polls app [here](https://github.com/Anupam-dagar/baileys-polls-app/blob/main/polls.postman_collection.json)

## Source Code
You can find the source code for the polls app [here](https://github.com/Anupam-dagar/baileys-polls-app)