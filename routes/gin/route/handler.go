package route

import (
	"Routers/routes/gin/database"
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

func get(users database.Getter) gin.HandlerFunc {
	return func(c *gin.Context) {
		result := users.GetAll()
		// c.String(http.StatusOK, "You are at the get route")
		c.JSON(http.StatusOK, result)
	}
}

func home() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.String(http.StatusOK, "Welcome to the home route")
	}
}

func post(users database.AddOrUpdate) gin.HandlerFunc {
	return func(c *gin.Context) {
		requestBodyPost := database.User{}
		id := uuid.NewV4()
		if err := c.ShouldBindJSON(&requestBodyPost); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		users.Add(id, requestBodyPost)
		// TODO: Convert uuid to string send
		c.String(http.StatusOK, "User successfully created, User id is ", id)
	}
}

func put(users database.AddOrUpdate) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err = users.Check(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		requestBody := database.User{}
		c.ShouldBindJSON(&requestBody)
		users.Add(id, requestBody)
		c.JSON(http.StatusOK, requestBody)
	}
}

func delete(users database.Delete) gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := uuid.FromString(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err = users.Check(id); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		users.Delete(id)
		c.String(http.StatusOK, "User deleted successfully")
	}
}
