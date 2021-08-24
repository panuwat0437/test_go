package api
import (
	"main/db"
	"main/model"
	"github.com/gin-gonic/gin"
)
func Setupmember(router *gin.Engine) {
	memberAPI := router.Group("/api/v2")
	{
		memberAPI.GET("/member" /*interceptor.JwtVerify,*/, getMember)
		// memberAPI.GET("/product/:id" /*interceptor.JwtVerify,*/, getProductByID)
		// memberAPI.POST("/product" /*interceptor.JwtVerify,*/, createProduct)
		// memberAPI.PUT("/product" /*interceptor.JwtVerify,*/, editProduct)
	}
}

func getMember(c *gin.Context) {
	var member []model.Member
	db.GetDB().Find(&member)
	c.JSON(200, member)

}
