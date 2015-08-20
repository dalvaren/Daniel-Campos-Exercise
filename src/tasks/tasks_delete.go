package tasks

import (
	"config"
	"github.com/gin-gonic/gin"
)

func taskDeleteRoute(c *gin.Context) {
  id := c.Params.ByName("id")

  var task Task
  config.DB.First(&task, id)
  task.IsDeleted = true
  dbReturn := config.DB.Save(&task)

  if dbReturn.Error != nil {
    c.JSON(500, gin.H{"message": "Problem deleting task."})
  }
  c.JSON(200, gin.H{"message": "task deleted", "task": dbReturn.Value})
}
