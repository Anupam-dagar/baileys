# Quick Start

## Requirements
- golang `1.20+`

## Installation
```bash
$ go get -u github.com/Anupam-dagar/baileys
```

### Additionally, you will also need a reflections package
```bash
$ go get -u github.com/oleiade/reflections
```

## Running a baileys app

### Init a go module
```bash
$ go mod init users
```

### Install baileys as a dependency
```bash
$ go get -u github.com/Anupam-dagar/baileys
```

### Setting up example Table
The quick start example makes use of `users` table which can be created with the following sql.
```sql
CREATE TABLE users (
	"name" varchar(255) NOT NULL,
	email varchar(255) NOT NULL,
	created_at timestamptz NULL,
	updated_at timestamptz NULL,
	deleted_at timestamptz NULL,
	created_by varchar NULL,
	updated_by varchar NULL,
	deleted_by varchar NULL,
	id varchar(255) NOT NULL,
	CONSTRAINT users_pk PRIMARY KEY (id)
);
```

### Basic Configuration Setup
Create a `dev.yaml` file under the directory `config` and use the following configuration
```yaml
server:
  port: "8080"
  base_api_path: "/api"

database:
  host: database_host
  port: database_port
  username: database_username
  password: database_password
  name: database_name
```

### Use the following code for `main.go`
```go
package main

import (
	"github.com/Anupam-dagar/baileys/constant/types"
	"github.com/Anupam-dagar/baileys/controller"
	"github.com/Anupam-dagar/baileys/entity"
	"github.com/Anupam-dagar/baileys/server"
	"github.com/gin-gonic/gin"
	"github.com/oleiade/reflections"
)

// Defining an entity

type User struct {
	entity.BaseModel
	Name  string `json:"name" gorm:"column:name"`
	Email string `json:"email" gorm:"column:email"`
}

type UserPtr = *User

func (u *User) SetCol(field string, val interface{}) error {
	return reflections.SetField(u, field, val)
}

func (u *User) GetStructData() (map[string]interface{}, error) {
	return reflections.Items(u)
}

// Defining routes

func SetupRoutes() {
	rootRouterGroup := server.GetGinEngine().GetRootRouterGroup()

	server.AddRoute(UserRoutesBaileys(rootRouterGroup))
}

func UserRoutes(routerGroup *gin.RouterGroup) {
	clientRouter := routerGroup.Group("/users")
	{
		NewClientController(clientRouter)
	}
}

func UserRoutesBaileys(routerGroup *gin.RouterGroup) types.RouteFunc {
	return func() (*gin.RouterGroup, func(rg *gin.RouterGroup)) {
		return routerGroup, UserRoutes
	}
}

// Defining a controller for the entity

type UserControllerInterface interface {
	controller.BaseControllerInterface
}

type userController struct {
	controller.BaseControllerInterface
}

func NewClientController(rg *gin.RouterGroup) UserControllerInterface {
	cc := new(userController)
	cc.BaseControllerInterface = controller.NewBaseController[UserPtr](rg)

	return cc
}

// Running the baileys app

func main() {
	server.NewGinEngine().InitGinApp(SetupRoutes).RunServer()
}
```

### Running the app
Run the app using
```bash
go run main.go
```

You should see the following output on successful run
```bash
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/user/search          --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func1 (3 handlers)
[GIN-debug] GET    /api/user/:id             --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func2 (3 handlers)
[GIN-debug] POST   /api/user                 --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func3 (3 handlers)
[GIN-debug] PUT    /api/user/:id             --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func4 (3 handlers)
[GIN-debug] DELETE /api/user/:id             --> github.com/Anupam-dagar/baileys/controller.NewBaseController[...].func5 (3 handlers)
{"time":"2023-07-22T22:51:56.285007+05:30","level":"INFO","prefix":"-","file":"server.go","line":"76","message":"Starting server on port: 8080"}
```

Congratulations, your baileys app is up and running. The above app has a single entity called `User` which interacts with `users` table in your provided database.
Baileys has exposed the CRUD apis along with a `search` api which you can use to
- Create new users - accepts entity as the payload with json keys defined under `json` tag.
- Soft delete a user
- Update a user - accepts entity as the payload with json keys defined under `json` tag. Only updates the keys sent in payload.
- Retrieve a user
- Search for users using `bql`

### Postman Collection
You can download the postman collection for the above example from [here](https://github.com/Anupam-dagar/baileys)