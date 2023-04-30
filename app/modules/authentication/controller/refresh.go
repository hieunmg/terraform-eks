package controller

import (
	"net/http"
	"weshare/components/appcontext"
	"weshare/components/token"
	"weshare/modules/authentication/model"
	"weshare/modules/authentication/service"
	"weshare/modules/authentication/storage"

	"github.com/gin-gonic/gin"
)

func Refresh(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.RenewAccessTokenRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(err)
		}

		tokenMaker, err := token.NewJWTMaker(appCtx.GetSecretKey())
		if err != nil {
			panic(err)
		}

		repository := storage.NewSQLRepository(appCtx.GetDBConnection())
		service := service.NewRenewAccessTokenService(repository, tokenMaker, appCtx.GetAccessTokenDuration())

		rsp, err := service.RenewAccessToken(c.Request.Context(), &req)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, rsp)

	}
}
