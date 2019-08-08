package src

import (
	"github.com/gin-gonic/gin"
	"interfaces/src/generated"
)

type Api struct {
	*generated.Api
}

func (api *Api) getPing(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}


