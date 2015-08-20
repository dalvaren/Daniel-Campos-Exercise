package tasks

import (
	"errors"
	"time"
	"strconv"
	"config"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/jwt"
)

// Task structure
type Task struct {
    ID          int64 `gorm:"primary_key"`
    Title       string
    Description string
    Priority    int
    CreatedAt   *time.Time `json:"createdAt"`
    UpdatedAt   *time.Time `json:"updatedAt"`
    CompletedAt *time.Time `json:"completedAt"`
    IsDeleted   bool
    IsCompleted bool
}

// TaskJSON is the requested task json structure for POST and PUT
type TaskJSON struct {
	Title     string `form:"title" json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsCompleted bool `json:"completed" binding:"required"`
	Priority string `json:"priority"`
}

// SetRoutes sets the crud and migration (if them exists) routes
func SetRoutes(r *gin.Engine){
	taskRoute := r.Group("/api/task")

	taskRoute.Use(jwt.Auth(config.TokenSecret))

	// Task CRUD
	taskRoute.GET("/*id", taskRetrieveRoute)
	taskRoute.POST("/", taskCreateRoute)
	taskRoute.PUT("/:id", taskUpdateRoute)
	taskRoute.DELETE("/:id", taskDeleteRoute)
}

// validate if request JSON has the required fields
func hasRequiredFields(json TaskJSON) error {
	if json.Title == "" {
		return errors.New("You need to send the title")
	}
	if json.Description == "" {
		return errors.New("You need to send the description")
	}
	return nil
}

// validate and set Priority
func setPriority(json TaskJSON) (int, error){
	priority := 1
	if json.Priority != "" {
		var err error
		priority, err = strconv.Atoi(json.Priority)
		if err != nil {
			return 0, errors.New("Priority shall be a number")
		}
	}
	return priority, nil
}
