package api

import (
	"jiuxia/crowdfunding/service"
	"jiuxia/crowdfunding/utils"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	var useRegister service.UserService
	if err := c.ShouldBind(&useRegister); err == nil {
		res := useRegister.Register()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}

func UserLogin(c *gin.Context) {
	var userLogin service.UserService
	if err := c.ShouldBind(&userLogin); err == nil {
		res := userLogin.Login()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func UserUpdateEmail(c *gin.Context) {
	var userUpdateEmail service.UserService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdateEmail); err == nil {
		res := userUpdateEmail.UpdateEmail(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func UserUpdatePassword(c *gin.Context) {
	var userUpdatePassword service.UserService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&userUpdatePassword); err == nil {
		res := userUpdatePassword.UpdatePassword(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func SendEmail(c *gin.Context) {
	var sendEmail service.SendEmailService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&sendEmail); err == nil {
		res := sendEmail.Send(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func ValidEmail(c *gin.Context) {
	var validEmail service.ValidEmailService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&validEmail); err == nil {
		res := validEmail.Valid(validEmail.Token, claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func MakeContribution(c *gin.Context) {
	var makeContribution service.MakeContributionService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&makeContribution); err == nil {
		res := makeContribution.Contribution(claim.Id, c.Param("pid"))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
