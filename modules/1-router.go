package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/semeru3-go/utils"
)

type Semeru3Router struct {
	engine *gin.Engine
	rh     Semeru3RequestHandler
}

func CreateSemeru3Router(engine *gin.Engine, rh Semeru3RequestHandler) Semeru3Router {
	return Semeru3Router{
		engine: engine,
		rh:     rh,
	}
}

func (r Semeru3Router) Init(path string) {
	pathGroup := r.engine.Group(path, utils.CheckBasicAuth)
	pathGroup.POST("/select/profile/id", r.rh.SelectProfile)
	pathGroup.POST("/insert/profile", r.rh.InsertProfile)
	pathGroup.POST("/update/profile/id", r.rh.UpdateProfile)
	pathGroup.POST("/delete/profile/id", r.rh.DeleteProfile)
}
