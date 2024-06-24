package main

import (
	"net/http"
	"sync"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type (
	user struct {
		// json tag is used to map the json key to the struct field (serialization/deserialization)
		ID   int    `json:"id"`
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
)

var (
	userMap = map[int]*user{}
	seq     = 1
	lock    = sync.Mutex{}
)

// Handlers
func createUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	u := &user{
		ID: seq,
	}
	if err := c.Bind(u); err != nil {
		return err
	}
	userMap[u.ID] = u
	seq++
	return c.JSON(http.StatusCreated, u)
}

// func updateUser(c echo.Context) error {
// 	lock.Lock()
// 	defer lock.Unlock()
// }

func getAllUsers(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()
	return c.JSON(http.StatusOK, userMap)
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/users", getAllUsers)
	e.POST("/users", createUser)
	// e.GET("/users/:id", getUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
