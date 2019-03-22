package server

import (
	"github.com/ShawnRong/bento/config"
	"github.com/gin-gonic/gin"
)

func Init() {
	c := config.GetConfig()
	if !c.GetBool("site.debug") {
		gin.SetMode(gin.ReleaseMode)
	}
	r := NewRouter()
	r.Run(c.GetString("site.domain"))
}
