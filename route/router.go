package route

import (
	frpApi "github.com/SpicyChickenFLY/gin-frp/controller/api/frp"
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
		groupTransportService := groupFrp.Group("/transport")
		{
			groupTransportService.GET(":name", frpApi.GetFrpTransportService)
			groupTransportService.POST("", frpApi.CreateFrpTransportService)
			groupTransportService.PUT(":name", frpApi.UpdateFrpTransportService)
			groupTransportService.DELETE(":name", frpApi.DeleteFrpTransportService)
		}
		groupSecureTransportService := groupFrp.Group("/secure-transport")
		{
			groupSecureTransportService.GET(":name", frpApi.GetFrpSecureTransportService)
			groupSecureTransportService.POST("", frpApi.CreateFrpsecureTransportService)
			groupSecureTransportService.PUT(":name", frpApi.UpdateFrpsecureTransportService)
			groupSecureTransportService.DELETE(":name", frpApi.DeleteFrpsecureTransportService)
		}
		groupFileService := groupFrp.Group("/file")
		{
			groupFileService.GET(":name", frpApi.GetFrpFileService)
			groupFileService.POST("", frpApi.CreateFrpFileService)
			groupFileService.PUT(":name", frpApi.UpdateFrpFileService)
			groupFileService.DELETE(":name", frpApi.DeleteFrpFileService)
		}
		// groupHTTPService := groupFrp.Group("/web")
		// {
		// 	groupHTTPService.GET(":name", frpApi.GetFrpHTTPService)
		// 	groupHTTPService.POST("", frpApi.CreateFrpHTTPService)
		// 	groupHTTPService.PUT(":name", frpApi.UpdateFrpHTTPService)
		// 	groupHTTPService.DELETE(":name", frpApi.DeleteFrpHTTPService)
		// }
	}

	return router
}
