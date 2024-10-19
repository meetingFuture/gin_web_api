package v1

import (
	"gin_web_api/pkg/app"
	"gin_web_api/pkg/e"
	"gin_web_api/vo/v1"
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorld(c *gin.Context) {
	appG := app.Gin{C: c}
	hello := v1.Hello{
		Hello: "Hello",
		Name:  "World",
	}
	appG.Response(http.StatusOK, e.SUCCESS, hello)
	return
}
