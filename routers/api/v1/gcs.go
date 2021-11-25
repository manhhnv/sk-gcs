package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/manhhnv/sk-gcs/pkg/gcs"
)

func ListFiles(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	tag := c.DefaultQuery("type", "all")
	files := gcs.ListFiles(name, tag)
	c.JSON(200, gin.H{
		"items": files,
	})
}
