package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type todo_list struct {
	ID   int    `json:"id"`
	Todo string `json:"todo"`
}

var todoLists = []todo_list{
	{ID: 1, Todo: "Run"},
	{ID: 2, Todo: "Do Homeworks"},
}

func getTodo_list(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, todoLists)
}

const (
	s = iota
)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {

		fmt.Println(c.Request.Header)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, Origin, Cache-Control, X-Requested-With")
		//c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "PUT, DELETE")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
func postTask(c *gin.Context) {
	var newTodoList todo_list
	//	Call BindJson to bind the received JSON to toTodoList
	if err := c.BindJSON(&newTodoList); err != nil {
		return
	}
	// Add the new album to the slice
	todoLists = append(todoLists, newTodoList)
	c.Header("Access-Control-Allow-Origin", "*")

	c.IndentedJSON(http.StatusCreated, newTodoList)

}

func main() {
	router := gin.Default()
	router.Use(CORS())
	router.GET("/lists", getTodo_list)
	router.POST("/postTask", postTask)
	router.Run(":8080")
}
