package controller

import (
	"net/http"
	"weshare/components/appcontext"
	"weshare/modules/authentication/model"
	"weshare/modules/authentication/service"
	"weshare/modules/authentication/storage"

	"github.com/gin-gonic/gin"
)

func Logout(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req *model.AccountLogoutRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(err)
		}

		repository := storage.NewSQLRepository(appCtx.GetDBConnection())
		service := service.NewLogoutRepository(repository)
		err := service.Logout(c.Request.Context(), req.AccountId)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, true)
	}
}
