package route

import (
	"github.com/SpicyChickenFLY/gin-frp/controller/api"
	"github.com/SpicyChickenFLY/gin-frp/pkgs/middleware"
	"github.com/gin-gonic/gin"
)

// InitRouter 初始化路由控制器
func InitRouter() *gin.Engine {
	// 禁用控制台颜色，当你将日志写入到文件的时候，你不需要控制台颜色。
	//gin.DisableConsoleColor()

	// 写入日志的文件
	/*f, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f)*/

	router := gin.New()

	router.Use(middleware.Cors())
	router.Static("/static", "static")
	router.LoadHTMLGlob("templates/*")

	// Group: Todo List
	groupFrp := router.Group("/frp")
	{
		groupTCPService := groupFrp.Group("/tcp")
		{
			groupTCPService.GET(":name", api.GetFrpTCPService)
			groupTCPService.POST("", api.CreateFrpTCPService)
			groupTCPService.PUT(":name", api.UpdateFrpTCPService)
			groupTCPService.DELETE(":name", api.DeleteFrpTCPService)
		}
		groupUDPService := groupFrp.Group("/udp")
		{
			groupUDPService.GET(":name", api.GetFrpUDPService)
			groupUDPService.POST("", api.CreateFrpUDPService)
			groupUDPService.PUT(":name", api.UpdateFrpUDPService)
			groupUDPService.DELETE(":name", api.DeleteFrpUDPService)
		}
		groupFileService := groupFrp.Group("/file")
		{
			groupFileService.GET(":name", api.GetFrpFileService)
			groupFileService.POST("", api.CreateFrpFileService)
			groupFileService.PUT(":name", api.UpdateFrpFileService)
			groupFileService.DELETE(":name", api.DeleteFrpFileService)
		}
		groupHTTPService := groupFrp.Group("/http")
		{
			groupHTTPService.GET(":name", api.GetFrpHTTPService)
			groupHTTPService.POST("", api.CreateFrpHTTPService)
			groupHTTPService.PUT(":name", api.UpdateFrpHTTPService)
			groupHTTPService.DELETE(":name", api.DeleteFrpHTTPService)
		}
		groupHTTPSService := groupFrp.Group("/https")
		{
			groupHTTPSService.GET(":name", api.GetFrpHTTPSService)
			groupHTTPSService.POST("", api.CreateFrpHTTPSService)
			groupHTTPSService.PUT(":name", api.UpdateFrpHTTPSService)
			groupHTTPSService.DELETE(":name", api.DeleteFrpHTTPSService)
		}
	}

	return router
}
