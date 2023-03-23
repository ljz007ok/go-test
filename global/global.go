package global

import (
	"sync"

	"github.com/gin-gonic/gin"
)

var Once sync.Once

type Option func(group *gin.RouterGroup)

var Options = []Option{}
