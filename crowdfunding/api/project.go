package api

import (
	"jiuxia/crowdfunding/service"
	"jiuxia/crowdfunding/utils"

	"github.com/gin-gonic/gin"
)

func ProjectUpload(c *gin.Context) {
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(500, err)
		c.Abort()
		return
	}
	fileSize := fileHeader.Size
	var projectUpload service.ProjectService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&projectUpload); err == nil {
		res := projectUpload.Upload(claim.Id, file, fileSize)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func ShowProjectPass(c *gin.Context) {
	var showProject service.ShowProjectService
	if err := c.ShouldBind(&showProject); err == nil {
		res := showProject.ShowPass()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func ShowProjectFail(c *gin.Context) {
	var showProject service.ShowProjectService
	if err := c.ShouldBind(&showProject); err == nil {
		res := showProject.ShowFail()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func ShowProjectUnknown(c *gin.Context) {
	var showProject service.ShowProjectService
	if err := c.ShouldBind(&showProject); err == nil {
		res := showProject.ShowUnknown()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func ShowMyProject(c *gin.Context) {
	var showProject service.ShowProjectService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&showProject); err == nil {
		res := showProject.ShowMy(claim.Id)
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func SearchProject(c *gin.Context) {
	var searchProject service.SearchProjectService
	if err := c.ShouldBind(&searchProject); err == nil {
		res := searchProject.Search()
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func DetailProjectByPid(c *gin.Context) {
	var detailProjectByPid service.DetailService
	if err := c.ShouldBind(&detailProjectByPid); err == nil {
		res := detailProjectByPid.Detail(c.Param("pid"))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
