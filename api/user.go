package api

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"sync"
)

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

// to replace with a database
var (
	users = map[int]*User{}
	seq   = 1
	lock  = sync.Mutex{}
)

func CreateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	u := &User{
		ID: seq,
	}

	if err := c.Bind(u); err != nil {
		return err
	}

	if u.Email == "" || u.Name == "" {
		return HandleError(c, "missing parameter(s)")
	}

	users[u.ID] = u
	seq++

	return c.JSON(http.StatusCreated, u)
}

func GetUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, users[id])
}

func UpdateUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	users[id].Name = u.Name

	return c.JSON(http.StatusOK, users[id])
}

func DeleteUser(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	delete(users, id)

	return c.NoContent(http.StatusNoContent)
}

func GetAllUsers(c echo.Context) error {
	lock.Lock()
	defer lock.Unlock()

	return c.JSON(http.StatusOK, users)
}
