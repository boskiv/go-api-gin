package generated

import (
	"github.com/gin-gonic/gin"
)

type IPing interface {
	getPing()
}

type PingGenerated struct {
	IPing
}

func (ping PingGenerated) getPing(c *gin.Context)  {
	c.JSON(200, gin.H{
		"message": "Not Implemented",
	})
}



