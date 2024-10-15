package routers

import (
	"device_management/api/controller"
	"device_management/api/middleware"
	"device_management/core/configs"
	"net/http"

	"github.com/gin-gonic/gin"
	cors "github.com/rs/cors/wrapper/gin"
)

type ApiRouter struct {
	Engine *gin.Engine
}

func NewApiRouter(
	cf *configs.Configs,
	hanlderFile *controller.ControllerSaveFile,
	user *controller.ControllerUsers,
	jwt *middleware.Middleware,
	device *controller.ControllerDevices,
	file *controller.ControllerFileStore,

) *ApiRouter {
	engine := gin.New()
	gin.DisableConsoleColor()

	engine.Use(gin.Logger())
	engine.Use(cors.AllowAll())
	engine.Use(gin.Recovery())

	r := engine.RouterGroup.Group("/manager")
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	r.POST("/login", user.Login)
	r.POST("/check", user.CheckToken)
	userGroup := r.Group("/user", jwt.Authenticate())
	{
		userGroup.POST("/add", user.AddUser)
		userGroup.GET("/list", user.GetListUser)
		userGroup.DELETE("/delete", user.DeleteUserById)
		userGroup.PATCH("/update/information", user.UpdateUserById)
		userGroup.PATCH("/reset/password", user.ResetPasswordUser)
	}
	fileGroup := r.Group("/files")
	{
		fileGroup.StaticFS("/export", http.Dir("publics"))
		fileGroup.POST("/upload", hanlderFile.SaveFile)
		fileGroup.GET("/list", file.GetListFilesByObjectId)
		fileGroup.DELETE("/delete", jwt.Authenticate(), file.DeleteFileById)
	}
	//view user
	viewForUser := r.Group("/view")
	{
		viewForUser.GET("/list", device.GetListDeviceForUser)
	}
	//device
	deviceGroup := r.Group("/device", jwt.Authenticate())
	{
		deviceGroup.POST("/add", device.AddDevice)
		deviceGroup.GET("/list", device.GetListDevice)
		deviceGroup.DELETE("/delete", device.DeleteDeviceById)
		deviceGroup.PATCH("/update", device.UpdatedDeviceById)
	}
	return &ApiRouter{
		Engine: engine,
	}
}
