package routers

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/manhhnv/sk-gcs/routers/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/files", v1.ListFiles)
	}

	return r
}
