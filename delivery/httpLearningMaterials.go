package delivery

import (
	"github.com/gin-gonic/gin"
	"github.com/tyshkorostyslav/first_task/usecase"
)

func Endpoints(router *gin.Engine) {
	v1 := router.Group("api/v1")
	{
		v1.GET("/learn_m", usecase.AvailableLearmingMaterials)
		v1.GET("/books", usecase.AvailableBooks)
		v1.GET("/pages", usecase.AvailablePages)
		v1.POST("/commitment", usecase.Commitment)
		v1.POST("/user", usecase.AddUser)
		v1.POST("/learn_m", usecase.AddLearningMaterial)
		v1.POST("/book", usecase.AddBook)
		v1.POST("/page", usecase.AddPage)
	}
}
