package global

import (
	"github.com/gin-gonic/gin"
	"github.com/my.blogs/config"
	"gorm.io/gorm"
)

var (
	GORM_DB        *gorm.DB
	GLOBAL_CONFIG  *config.Server
	GLOBAL_ROUTERS *gin.Engine
)
