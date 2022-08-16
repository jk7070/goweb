package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"

	calc "github.com/jk7070/goweb/calculations"
	todo "github.com/jk7070/goweb/todos"
)

func getHome(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Welcome to this Golang Rest API")
}

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.Default() // this is the server

	router.GET("/", getHome)

	router.GET("/calc_results", calc.CalcResults)
	router.GET("/calc/:n", calc.Calculate)

	router.GET("/todos", todo.GetTodos)
	router.GET("/todos/:id", todo.GetTodo)
	router.PATCH("/todos/:id", todo.ToggleTodoStatus)
	router.POST("/todos", todo.AddTodo)

	// router.Run("localhost:7070")
	router.Run(":" + port)

	// setup workers

}
