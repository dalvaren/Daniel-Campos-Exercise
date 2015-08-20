package tasks

import (
	"errors"
	"time"
	"strconv"
	"strings"
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

	// Retrieve task(s)
	taskRoute.GET("/*id", func(c *gin.Context) {
		// if "id" param exists
		id := c.Params.ByName("id")
		if id != "/" {
			var task Task
			taskID := strings.Replace(id, "/", "", -1)
			config.DB.First(&task, taskID)
			c.JSON(200, gin.H{"task": task})
			return
		}

		// if is the main URI (list the tasks)
		var tasks []Task
		config.DB.Where("is_deleted = ?", false).Find(&tasks)
		c.JSON(200, gin.H{"items": tasks})
	})

	// Create new task
	taskRoute.POST("/", func(c *gin.Context) {
		var json TaskJSON
		c.Bind(&json)

		if err := hasRequiredFields(json); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		priority, err := setPriority(json)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		rightNow := time.Now()
		task := Task{
			Title: json.Title,
	    Description: json.Description,
	    Priority: priority,
	    CreatedAt: &rightNow,
	    UpdatedAt: &rightNow,
	    CompletedAt: &rightNow,
	    IsDeleted: false,
	    IsCompleted: false,
		}
		dbReturn := config.DB.Create(&task)
		if dbReturn.Error != nil {
			c.JSON(500, gin.H{"message": "Problem creating task."})
		}
		c.JSON(200, gin.H{"message": "task created", "task": dbReturn.Value})
	})

	// Update task
	taskRoute.PUT("/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")

		var json TaskJSON
		c.Bind(&json)

		if err := hasRequiredFields(json); err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		priority, err := setPriority(json)
		if err != nil {
			c.JSON(500, gin.H{"message": err.Error()})
			return
		}

		isCompleted := false
		if json.IsCompleted == true {
			isCompleted = true
		}

		rightNow := time.Now()
		var task Task
		config.DB.First(&task, id)

		completedAt := task.CompletedAt
		if json.IsCompleted == true {
			completedAt = &rightNow
		}

		task.Title = json.Title
		task.Description = json.Description
		task.Priority = priority
		task.UpdatedAt = &rightNow
		task.CompletedAt = completedAt
		task.IsCompleted = isCompleted
		dbReturn := config.DB.Save(&task)

		if dbReturn.Error != nil {
			c.JSON(500, gin.H{"message": "Problem updating task."})
		}
		c.JSON(200, gin.H{"message": "task updated", "task": dbReturn.Value, "json": json})
	})

	// Delete task (only set the value IsDeleted as true)
	taskRoute.DELETE("/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")

		var task Task
		config.DB.First(&task, id)
		task.IsDeleted = true
		dbReturn := config.DB.Save(&task)

		if dbReturn.Error != nil {
			c.JSON(500, gin.H{"message": "Problem deleting task."})
		}
		c.JSON(200, gin.H{"message": "task deleted", "task": dbReturn.Value})
	})
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
