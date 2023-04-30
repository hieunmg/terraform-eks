package controller

import (
	"weshare/components/appcontext"
	"weshare/modules/post/model"

	"github.com/gin-gonic/gin"
)

func CreatePost(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.PostCreate
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(err)
		}

	}
}
