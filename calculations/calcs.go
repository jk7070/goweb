package calculations

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Result struct {
	Input    int     `json:"input"`
	Output   int     `json:"output"`
	Duration float64 `json:"duration"`
}

var results = []Result{}

func CalcResults(context *gin.Context) {
	if len(results) == 0 {
		context.IndentedJSON(http.StatusOK, "No results added yet!")
	} else {
		context.IndentedJSON(http.StatusOK, results)
	}
}

func Calculate(context *gin.Context) {
	n := context.Param("n")
	fmt.Print(n)

	N, _ := strconv.Atoi(n)

	start := time.Now()

	res := FibonacciRecursion(N)

	end := time.Now()
	elapsed := end.Sub(start).Seconds()

	results = append(results, Result{Input: N, Output: res, Duration: elapsed})

	context.IndentedJSON(http.StatusOK, results)

}

func FibonacciRecursion(n int) int {
	if n <= 1 {
		return n
	}
	return FibonacciRecursion(n-1) + FibonacciRecursion(n-2)
}
