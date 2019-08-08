package generated

import (
	"github.com/gin-gonic/gin"
)

type IHello interface {
	getHello()
}

type HelloGenerated struct {
	IPing
}

func (hello HelloGenerated) getHello(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "Not Implemented",
	})
}



