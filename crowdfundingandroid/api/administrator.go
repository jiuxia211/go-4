package api

import (
	"jiuxia/crowdfundingandroid/service"
	"jiuxia/crowdfundingandroid/utils"

	"github.com/gin-gonic/gin"
)

func AuditProject(c *gin.Context) {
	var auditProject service.AuditProjectService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&auditProject); err == nil {
		res := auditProject.Audit(claim.Id, c.Param("pid"))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
func DeleteProject(c *gin.Context) {
	var deleteProject service.DeleteProjectService
	claim, _ := utils.ParseToken(c.GetHeader("Authorization"))
	if err := c.ShouldBind(&deleteProject); err == nil {
		res := deleteProject.Delete(claim.Id, c.Param("pid"))
		c.JSON(200, res)
	} else {
		c.JSON(400, err)
	}
}
