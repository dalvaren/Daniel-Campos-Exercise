package tasks

import (
	"fmt"
	"time"
	"strconv"
	"strings"
	"config"
	"github.com/gin-gonic/gin"
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

func (this *Task) Add() {
	config.DB.NewRecord(this)
	config.DB.Create(&this)
}

func CreateSampleTask() {
	rightNow := time.Now()
	task := Task{
		Title: "Title of the task",
    Description: "Description of the task",
    Priority: 1,
    CreatedAt: &rightNow,
    UpdatedAt: &rightNow,
    CompletedAt: &rightNow,
    IsDeleted: false,
    IsCompleted: false,
	}
	config.DB.Create(&task)
	// task.Add()
}

// CreateTasksTable is migration to generate the table (still using GORM)
// TODO: change to goose
func CreateTasksTable() {
	config.DB.AutoMigrate(&Task{})
}

// DropTasksTable is migration to drop the table (still using GORM)
// TODO: change to goose
func DropTasksTable() {
	config.DB.DropTable(&Task{})
}

// SetRoutes sets the crud and migration (if them exists) routes
func SetRoutes(r *gin.Engine){
	taskRoute := r.Group("/task")

	// taskRoute.GET("/migration/create", func(c *gin.Context) {
	// 	CreateTasksTable()
	// 	c.JSON(200, gin.H{"message": "task table created"})
	// })
	//
	// taskRoute.GET("/migration/drop", func(c *gin.Context) {
	// 	DropTasksTable()
	// 	c.JSON(200, gin.H{"message": "task table droped"})
	// })
	//
	// taskRoute.GET("/migration/sample", func(c *gin.Context) {
	// 	CreateSampleTask()
	// 	c.JSON(200, gin.H{"message": "sample task created"})
	// })

	taskRoute.GET("/*id", func(c *gin.Context) {
		id := c.Params.ByName("id")
		if id != "/" {
			var task Task
			taskID := strings.Replace(id, "/", "", -1)
			config.DB.First(&task, taskID)
			c.JSON(200, gin.H{"task": task})
			return
		}

		var tasks []Task
		config.DB.Where("is_deleted = ?", false).Find(&tasks)
		c.JSON(200, gin.H{"items": tasks})
	})

	type TaskJSON struct {
    Title     string `form:"title" json:"title" binding:"required"`
    Description string `json:"description" binding:"required"`
    IsCompleted bool `json:"completed" binding:"required"`
		Priority string `json:"priority"`
	}

	taskRoute.POST("/", func(c *gin.Context) {
		var json TaskJSON
		c.Bind(&json)

		if json.Title == "" {
			c.JSON(500, gin.H{"message": "You need to send the title"})
			return
		}
		title := json.Title

		if json.Description == "" {
			c.JSON(500, gin.H{"message": "You need to send the description"})
			return
		}
		description := json.Description

		priority := 1
		if json.Priority != "" {
			var err error
			priority, err = strconv.Atoi(json.Priority)
			if err != nil {
				c.JSON(500, gin.H{"message": "Priority shall be a number"})
				return
			}
		}

		rightNow := time.Now()
		task := Task{
			Title: title,
	    Description: description,
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

	taskRoute.PUT("/:id", func(c *gin.Context) {
		id := c.Params.ByName("id")

		var json TaskJSON
		c.Bind(&json)

		if json.Title == "" {
			c.JSON(500, gin.H{"message": "You need to send the title"})
			return
		}

		if json.Description == "" {
			c.JSON(500, gin.H{"message": "You need to send the description"})
			return
		}

		priority := 1
		if json.Priority != "" {
			var err error
			priority, err = strconv.Atoi(json.Priority)
			if err != nil {
				c.JSON(500, gin.H{"message": "Priority shall be a number"})
				return
			}
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

// PrintTest prints a test string in console
func PrintTest() {
	fmt.Println("Package Tasks loaded!")
}

// SumNumbers sums a and b
func SumNumbers(a, b int) int{
	return a + b
}
