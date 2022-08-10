package routes

import (
	"go-api-car/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandlerRequests() *gin.Engine {
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
		AllowHeaders:  []string{"content-type"},
		ExposeHeaders: []string{"X-total-Count"},
	}))
	r.GET("/car/", controllers.CarList)
	r.POST("/car/", controllers.CarCreate)
	r.GET("/car/:id", controllers.CarGetById)
	r.PATCH("/car/:id", controllers.CarUpdate)
	r.DELETE("/car/:id", controllers.CarDelete)
	return r
}
