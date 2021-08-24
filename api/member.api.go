package api

import (
	"main/db"
	"main/interceptor"
	"main/model"

	// "net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Setupmember(router *gin.Engine) {
	memberAPI := router.Group("/api/v2")
	{
		memberAPI.GET("/member", interceptor.JwtVerify, getMember)
		memberAPI.GET("/member/:id", interceptor.JwtVerify, getMemberByID)
		memberAPI.POST("/member", interceptor.JwtVerify, createMember)
		// memberAPI.PUT("/product" /*interceptor.JwtVerify,*/, editProduct)
	}
}

func getMember(c *gin.Context) {
	var member []model.Member
	db.GetDB().Find(&member)
	c.JSON(200, member)

}

func createMember(c *gin.Context) {

	member := model.Member{}
	member.Username = c.PostForm("username")
	member.Password = c.PostForm("password")
	member.Email = c.PostForm("email")
	// product.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	member.CreatedAt = time.Now()
	result := db.GetDB().Create(&member)

	if result.Error != nil {
		c.JSON(400, gin.H{"result": "Error"})
		return
	}

	c.JSON(200, gin.H{"result": member})

}

func getMemberByID(c *gin.Context) {
	var member model.Member
	db.GetDB().Where("id = ?", c.Param("id")).First(&member)
	c.JSON(200, member)
}
