package controller

import (
	"net/http"
	"weshare/common"
	"weshare/components/appcontext"
	"weshare/components/hasher"
	"weshare/components/token"

	"weshare/modules/authentication/model"
	"weshare/modules/authentication/service"
	"weshare/modules/authentication/storage"

	"github.com/gin-gonic/gin"
)

func Login(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.AccountLoginRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		tokenMaker, err := token.NewJWTMaker(appCtx.GetSecretKey())
		if err != nil {
			panic(err)
		}

		brcrypt := hasher.NewBcrypt()
		repository := storage.NewSQLRepository(appCtx.GetDBConnection())
		service := service.NewLoginService(
			repository,
			tokenMaker,
			brcrypt,
			appCtx.GetAccessTokenDuration(),
			appCtx.GetRefreshTokenDuration())
		rsp, err := service.Login(c.Request.Context(), &req)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, rsp)
	}
}
