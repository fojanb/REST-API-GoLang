// We are going to implement a CRUD operation on a REST api using Go and a package name Gin
package main
import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strconv"
)
type task struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Completed bool `json:"completed"`
}
var tasks = []task{
	{ID:1,Title: "Cleaning",Completed: false},
	{ID:2,Title: "Study",Completed: false},
	{ID:10,Title: "Wlaking",Completed: false},
	{ID:20,Title: "Cooking",Completed: false},
}
// var index = `Welcome to-to-do list app`
// --------------------------------------------------------------------
// ⭐ Write a handler to return all tasks:
// getTasks responds with the list of all tasks as JSON.
// gin.Context is the most important part of Gin. 
// It carries request details, validates and serializes JSON, and more. (Despite the similar name, this is different from Go’s built-in context package.)
func getTasks(c *gin.Context){
	// Call Context.IndentedJSON to serialize the struct into JSON and add it to the response.
	// The function’s first argument is the HTTP status code you want to send to the client. Here, you’re passing the StatusOK constant from the net/http package to indicate 200 OK.
	// StatusOK is a 'constant' under the net/http package
	// You can find the rest of constants in this package here https://pkg.go.dev/net/http#pkg-constants
	c.IndentedJSON(http.StatusOK, tasks)
}
// --------------------------------------------------------------------
// ⭐ Write a handler to add a task:
func addTask(c *gin.Context){
	var newTask task
	err := c.BindJSON(&newTask)
	if err != nil {
		return
	}
	tasks = append(tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}
// --------------------------------------------------------------------
// ⭐ Write a handler to get a specific task:
func getTaskByID(c *gin.Context) {
    id := c.Param("id")
    // Loop over the list of tasks, looking for
    // an task whose ID value matches the parameter.
    for _, a := range tasks {
        if strconv.Itoa(a.ID) == id{
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
// --------------------------------------------------------------------
// ⭐ Write a handler to delete a specific task:
func deleteTaskByID(c *gin.Context) {
    id := c.Param("id")
    // Loop over the list of tasks, looking for
    // an tsak whose ID value matches the parameter.
    for i, task := range tasks {
        if strconv.Itoa(task.ID) == id{
			tasks = append(tasks[:i], tasks[i+1:]...)
            c.IndentedJSON(http.StatusOK,tasks)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}
// --------------------------------------------------------------------

func main(){
	//Initialize a Gin router using Default.
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			
		})
	})
	// Use the GET function to associate the GET HTTP method and /tasks path with a handler function.
	router.GET("/tasks", getTasks)
	// Use the POST function to associate the POST HTTP method and /tasks path with a handler function.
	router.POST("/tasks", addTask)
	// Use the GET function to associate the GET HTTP method and /tasks/:id path with a handler function.
	router.GET("/tasks/:id", getTaskByID)
	// Use the DELETE function to associate the DELETE HTTP method and /tasks/:id path with a handler function.
	router.DELETE("/tasks/:id", deleteTaskByID)
	//Use the Run function to attach the router to an http.Server and start the server.
	router.Run("localhost:8080")
}
