package tasks

import (
  "time"
	"config"
	"github.com/gin-gonic/gin"
)

func taskUpdateRoute(c *gin.Context) {
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
}
