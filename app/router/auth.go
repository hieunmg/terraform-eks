package router

import (
	"weshare/components/appcontext"
	authcontroller "weshare/modules/authentication/controller"

	"github.com/gin-gonic/gin"
)

func AuthRouter(router *gin.RouterGroup, appCtx appcontext.AppContext) {
	auth := router.Group("/accounts")
	auth.POST("/login", authcontroller.Login(appCtx))
	auth.POST("/register", authcontroller.Register(appCtx))
	auth.POST("/refresh", authcontroller.Refresh(appCtx))
	auth.POST("/logout", authcontroller.Logout(appCtx))
}
