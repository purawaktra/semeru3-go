package modules

import "github.com/gin-gonic/gin"

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
