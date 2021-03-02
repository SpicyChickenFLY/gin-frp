package route

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
			groupTCPService.GET(":name", controller.GetFrpTCPService)
			groupTCPService.POST("", controller.CreateFrpTCPService)
			groupTCPService.PUT(":name", controller.UpdateFrpTCPService)
			groupTCPService.DELETE(":name", controller.DeleteFrpTCPService)
		}
		groupUDPService := groupFrp.Group("/udp")
		{
			groupUDPService.GET(":name", controller.)
			groupUDPService.POST("", controller.)
			groupUDPService.PUT(":name", controller.)
			groupUDPService.DELETE(":name", controller.)
		}
		groupFileService := groupFrp.Group("/file")
		{
			groupFileService.GET(":name", controller.)
			groupFileService.POST("", controller.)
			groupFileService.PUT(":name", controller.)
			groupFileService.DELETE(":name", controller.)
		}
		groupHTTPService := groupFrp.Group("/http")
		{
			groupHTTPService.GET(":name", controller.)
			groupHTTPService.POST("", controller.)
			groupHTTPService.PUT(":name", controller.)
			groupHTTPService.DELETE(":name", controller.)
		}
		groupHTTPSService := groupFrp.Group("/https")
		{
			groupHTTPSService.GET(":name", controller.)
			groupHTTPSService.POST("", controller.)
			groupHTTPSService.PUT(":name", controller.)
			groupHTTPSService.DELETE(":name", controller.)
		}
	}

	return router
}