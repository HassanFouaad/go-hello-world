package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type job struct {
	Title  string `json:"title"`
	Salary int    `json:"salary"`
}

var jobs = []job{
	{Title: "Backend Developer", Salary: 10000},
	{Title: "Frontend Developer", Salary: 9000},
	{Title: "UX Designer", Salary: 8000},
}

func getJObs(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, jobs)
}

func returnok(c *gin.Context) {
	c.IndentedJSON(http.StatusCreated, jobs)
}

func createJob(c *gin.Context) {
	var newJob job

	// Call BindJSON to bind the received JSON to
	// newJob.
	if err := c.BindJSON(&newJob); err != nil {
		return
	}

	// Add the new job to the slice.
	jobs = append(jobs, newJob)

}

func routes() {
	router := gin.Default()
	router.GET("/jobs", getJObs)
	router.POST("/jobs", createJob, getJObs)
	router.Run("localhost:8080")
}

func main() {
	routes()
}
