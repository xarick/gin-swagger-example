package main

import "github.com/xarick/gin-swagger-example/routes"

// @title           Swagger API
// @version         1.1.0
// @description     API for Developers

// @host      127.0.0.1:8777
// @BasePath  /api
// @schemes http https

// @securityDefinitions.basic BasicAuth
func main() {
	r := routes.ApiRoutes()
	r.Run(":8777")
}
