package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/manhhnv/sk-gcs/pkg/gcs"
	"strings"
)

func ListFiles(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	tag := c.DefaultQuery("type", "")
	tags := strings.Split(tag, ",")
	for i, _ := range tags {
		tags[i] = strings.TrimSpace(strings.ToLower(tags[i]))
	}
	files := gcs.ListFiles(name, tags)
	c.JSON(200, gin.H{
		"items": files,
	})
}
