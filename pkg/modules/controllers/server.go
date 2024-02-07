package controllers

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/martirosharutyunyan/axxon-test-task/pkg/config"
	"log"
)

var Router *gin.Engine
var ApiRouter *gin.RouterGroup

func RunServer() {
	Router = gin.Default()
	if config.GetEnv() == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		if err := Router.SetTrustedProxies(nil); err != nil {
			log.Fatal(err.Error())
		}
	}
	Router.Use(cors.Default())
	ApiRouter = Router.Group("/api")
	InitTaskController()
	pprof.Register(Router)
	log.Printf("server is running on http://127.0.0.1:%s", config.GetPort())
	err := Router.Run(fmt.Sprintf("127.0.0.1:%s", config.GetPort()))

	if err != nil {
		log.Fatal(err)
	}
}
