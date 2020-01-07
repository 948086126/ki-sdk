package action

import (
	g "github.com/gin-gonic/gin"
	Kiapi "ki-sdk/controllers/api"
)

func initRouter() *g.Engine {

	router := g.Default()

	api := router.Group("/api")

	v1 := api.Group("/v1")

	Test := v1.Group("/test")

	{
		Test.POST("/getCert", Kiapi.GenerateRSAKey)
		Test.POST("/getCert2", Kiapi.GenerateRSAKeyforPem)
		Test.POST("/jiami", Kiapi.Encryption)
		Test.POST("/jiemi", Kiapi.GenerateRSAKey)
		Test.GET("/testget", Kiapi.TestGet)
		Test.GET("/testpost", Kiapi.TestPost)
	}

	//
	return router
}
