package tasks

import (
	"fmt"
	"time"
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

	taskRoute.GET("/migration/create", func(c *gin.Context) {
		CreateTasksTable()
		c.JSON(200, gin.H{"message": "task table created"})
	})

	taskRoute.GET("/migration/drop", func(c *gin.Context) {
		DropTasksTable()
		c.JSON(200, gin.H{"message": "task table droped"})
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
