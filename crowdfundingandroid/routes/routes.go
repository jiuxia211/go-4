package routes

import (
	"jiuxia/crowdfundingandroid/api"
	"jiuxia/crowdfundingandroid/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("something-very-secret"))
	r.Use(sessions.Sessions("mysession", store))
	v1 := r.Group("api/v1")
	{
		v1.POST("user/register", api.UserRegister)
		v1.POST("user/login", api.UserLogin)
		authorized := v1.Group("/")
		authorized.Use(middleware.JwT())
		{
			//用户端
			authorized.PUT("user/update/email", api.UserUpdateEmail)
			authorized.PUT("user/update/password", api.UserUpdatePassword)
			authorized.POST("user/send-email", api.SendEmail)
			authorized.POST("user/contribution/:pid", api.MakeContribution)
			//项目端
			authorized.POST("project/create", api.ProjectCreate)
			authorized.PUT("project/upload/info/:pid", api.ProjectUploadInfo)
			authorized.PUT("project/upload/file/:pid", api.ProjectUploadFile)
			authorized.POST("project/show/pass", api.ShowProjectPass)
			authorized.POST("project/show/fail", api.ShowProjectFail)
			authorized.POST("project/show/unknown", api.ShowProjectUnknown)
			authorized.POST("project/search", api.SearchProject)
			authorized.POST("project/show/me", api.ShowMyProject)
			authorized.GET("project/:pid", api.DetailProjectByPid)
			//管理员操作
			authorized.POST("project/audit/:pid", api.AuditProject)
			authorized.DELETE("project/delete/:pid", api.DeleteProject)
		}
	}
	return r
}
