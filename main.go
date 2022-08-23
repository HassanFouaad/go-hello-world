package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
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

var upgrader = websocket.Upgrader{
	//check origin will check the cross region source (note : please not using in production)
	CheckOrigin: func(r *http.Request) bool {

		//Here we just allow the chrome extension client accessable (you should check this verify accourding your client source)
		return true
	},
}

func routes() {
	router := gin.Default()
	router.GET("/jobs", getJObs)
	router.POST("/jobs", createJob, getJObs)

	router.GET("/socket", func(c *gin.Context) {
		//upgrade get request to websocket protocol
		ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer ws.Close()
		for {
			//Read Message from client
			mt, message, err := ws.ReadMessage()

			if err != nil {
				fmt.Println(err)
				break
			}
			//If client message is ping will return pong
			if string(message) == "ping" {
				message = []byte("pong")
			}
			//Response message to client
			err = ws.WriteMessage(mt, message)
			println("Sent " + string(message) + " to The Client")
			if err != nil {
				fmt.Println(err)
				break
			}
		}
	})

	router.Run("localhost:8080")
}

func websocketServer(ws *websocket.Conn) {

}

func main() {
	routes()

}
