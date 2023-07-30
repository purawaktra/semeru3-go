package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/purawaktra/semeru3-go/modules"
	"github.com/purawaktra/semeru3-go/utils"
)

func main() {
	utils.InitConfig()
	utils.InitLog()

	utils.Log("============= SEMERU3 RUNTIME STARTED =============")
	// create gin engine
	engine := gin.New()

	// create dsn
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		utils.MysqlUser,
		utils.MysqlPassword,
		utils.MysqlHost,
		utils.MysqlPort,
		utils.MysqlScheme)

	// create mysql instance
	gormInstance, err := utils.CreateGorm(dsn)
	if err != nil {
		utils.Error(err, "main", "")
		panic(err)
	}

	// declare architecture
	repository := modules.CreateSemeru3Repo(gormInstance)
	usecase := modules.CreateSemeru3Usecase(repository)
	controller := modules.CreateSemeru3Controller(usecase)
	requestHandler := modules.CreateSemeru3RequestHandler(controller)
	router := modules.CreateSemeru3Router(engine, requestHandler)

	// init routing
	router.Init("/semeru3/api/v1")

	utils.Log("============= SEMERU3 ENGINE STARTING =============")
	utils.Log(fmt.Sprintf("App Address = %s:%s", utils.AppHost, utils.AppPort))

	// init routing to swagger
	swaggerRouter := utils.CreateSwaggerRouter(engine)
	swaggerRouter.InitRouter("/semeru3/api/v1/swagger")

	// start http api engine
	err = engine.Run(fmt.Sprintf("%s:%s", utils.AppHost, utils.AppPort))
	if err != nil {
		utils.Fatal(err, "main", "")
		panic(err)
	}

	utils.Log("============= SEMERU3 ENGINE FAILED =============")
}
