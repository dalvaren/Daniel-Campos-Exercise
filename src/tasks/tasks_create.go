package tasks

import (
	"time"
	"config"
	"github.com/gin-gonic/gin"
)

func taskCreateRoute(c *gin.Context) {
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
}
