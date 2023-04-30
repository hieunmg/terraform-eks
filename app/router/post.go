package router

import (
	"weshare/components/appcontext"
	postcontroller "weshare/modules/post/controller"

	"github.com/gin-gonic/gin"
)

func PostController(router *gin.RouterGroup, ctx appcontext.AppContext) {
	posts := router.Group("/posts")
	posts.POST("/:id", postcontroller.CreatePost(ctx))
	posts.PUT("/:id", postcontroller.UpdatePost(ctx))
	posts.GET("/:id", postcontroller.GetPost(ctx))
	posts.GET("/", postcontroller.ListPost(ctx))
}
