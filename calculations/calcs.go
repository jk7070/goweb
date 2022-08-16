package calculations

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Input  int `json:"input"`
	Output int `json:"output"`
}

var results = []Result{}

func CalcResults(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, results)
}

func Calculate(context *gin.Context) {
	n := context.Param("n")
	fmt.Print(n)

	N, _ := strconv.Atoi(n)

	res := FibonacciRecursion(N)
	results = append(results, Result{Input: N, Output: res})

	context.IndentedJSON(http.StatusOK, results)

}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
