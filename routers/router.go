package routers

import (
	"gin_web_api/controllers/v1"
	"gin_web_api/middleware/jwt"
	"github.com/gin-gonic/gin"
)

// InitRouter initialize routing information
func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	//获取文章列表
	r.POST("/hello", v1.HelloWorld)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		apiv1.POST("/hello", v1.HelloWorld)
		apiv1.POST("/hello1", v1.HelloWorld)
	}
	return r
}
