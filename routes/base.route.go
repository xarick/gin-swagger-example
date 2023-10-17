package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/xarick/gin-swagger-example/controllers"
	_ "github.com/xarick/gin-swagger-example/docs"
	"github.com/xarick/gin-swagger-example/middlewares"
)

func ApiRoutes() *gin.Engine {

	route := gin.Default()
	// docs.SwaggerInfo.BasePath = "/api"

	route.GET("/", func(c *gin.Context) {
		c.Redirect(http.StatusFound, "/docs/index.html")
	})

	api := route.Group("/api").Use(middlewares.AuthMiddleware())
	{
		api.GET("/msg", controllers.MSGController)
	}

	route.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return route
}
