package main

import (
	"fmt"
	"net/http"

	"training/rest-api/docs"
	"training/rest-api/helpers"
	"training/rest-api/models"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary Get Users
// @Schemes
// @Description Get users
// @Tags Users
// @Success 200 {object} models.User
// @Router /users [get]
func getUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Users)
}

// @Summary Get Users
// @Schemes
// @Description Get user
// @Tags Users
// @Param id path int true "User ID"
// @Failure 400
// @Failure 404
// @Success 200
// @Router /users/{id} [get]
func getUser(c *gin.Context) {
	id := helpers.ParseIntegerParam("id", c)
	if id == -1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error",
			"detail": fmt.Sprintf("Invalid ID '%s'", c.Params.ByName("id"))})
		return
	}
	for _, user := range models.Users {
		if user.ID == id {
			c.IndentedJSON(http.StatusOK, user)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"status": "error", "detail": fmt.Sprintf("Could not find user with id %d", id)})
}

// @Summary Create User
// @Schemes
// @Description Create a new user
// @Tags Users
// @Accept json
// @Produce json
// @Param user body models.UserCreate true "User object"
// @Success 201 {object} models.User
// @Failure 400
// @Router /users [post]
func createUser(c *gin.Context) {
	var userInput models.UserCreate
	if err := c.BindJSON(&userInput); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "detail": "Invalid request body"})
		return
	}
	newUser := models.User{
		ID:       helpers.GetNewObjectID(models.Users, func(user models.User) int { return user.ID }),
		Username: userInput.Username,
		FullName: userInput.FullName,
	}
	models.Users = append(models.Users, newUser)
	c.IndentedJSON(http.StatusCreated, newUser)
}

// @Summary Create Task
// @Schemes
// @Description Create a new task
// @Tags Tasks
// @Accept json
// @Produce json
// @Param task body models.TaskCreate true "Task object"
// @Success 201 {object} models.Task
// @Failure 400
// @Router /tasks [post]
func createTask(c *gin.Context) {
	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error", "detail": "Invalid request body"})
		return
	}

	models.Tasks = append(models.Tasks, newTask)
	c.IndentedJSON(http.StatusCreated, newTask)
}

// @Summary Get Tasks
// @Schemes
// @Description Get Tasks
// @Tags Tasks
// @Success 200 {object} models.Task
// @Router /tasks [get]
func getTasks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, models.Tasks)
}

// @Summary Get Tasks
// @Schemes
// @Description Get task
// @Tags Tasks
// @Param id path int true "Task ID"
// @Failure 400
// @Failure 404
// @Success 200
// @Router /tasks/{id} [get]
func getTask(c *gin.Context) {
	id := helpers.ParseIntegerParam("id", c)
	if id == -1 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"status": "error",
			"detail": fmt.Sprintf("Invalid ID '%s'", c.Params.ByName("id"))})
		return
	}
	for _, task := range models.Tasks {
		if task.ID == id {
			c.IndentedJSON(http.StatusOK, task)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"status": "error", "detail": fmt.Sprintf("Could not find task with id %d", id)})
}

func getReadiness(c *gin.Context) {
	status := make(map[string]string)
	status["status"] = "Operational"
	c.IndentedJSON(http.StatusOK, status)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	// Swagger info
	docs.SwaggerInfo.Title = "Swagger Example API"
	docs.SwaggerInfo.Description = "My thing."
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http"}

	v1 := r.Group("/api/v1")

	// Swagger docs route
	v1.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v1.GET("/readiness", getReadiness)
	users := v1.Group("/users")
	{
		users.GET("/", getUsers)
		users.GET("/:id", getUser)
		users.POST("/", createUser)
	}
	tasks := v1.Group("/tasks")
	{
		tasks.GET("/", getTasks)
		tasks.GET("/:id", getTask)
		tasks.POST("/", createTask)
	}

	return r
}

// @version 1
// @Description description

// @contact.name Itai Melamed
// @contact.url https://github.com/ItaiMelamed

// @securityDefinitions.apikey bearerToken
// @in header
// @name No Need

// @host localhost:8080
// @BasePath /api/v1
func main() {
	r := setupRouter()

	r.Run(":8080")
}
