package tasks

import (
  "strings"
	"config"
	"github.com/gin-gonic/gin"
)

func taskRetrieveRoute(c *gin.Context) {
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
}
