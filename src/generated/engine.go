package generated

import "github.com/gin-gonic/gin"

type ApiGenerated struct {
	PingGenerated
	HelloGenerated
}

type Api struct {
	ApiGenerated
}

func GinEngine() *gin.Engine {
	api := Api{}
	r := gin.Default()
	r.GET("/ping", api.getPing)
	r.GET("/hello", api.getHello)
	return r
}