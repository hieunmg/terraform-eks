package controller

import (
	"net/http"
	"weshare/common"
	"weshare/components/appcontext"
	"weshare/components/hasher"
	"weshare/modules/authentication/model"
	"weshare/modules/authentication/service"
	"weshare/modules/authentication/storage"

	"github.com/gin-gonic/gin"
)

func Register(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		var req model.AccountRegisterRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		brcrypt := hasher.NewBcrypt()
		repository := storage.NewSQLRepository(appCtx.GetDBConnection())
		service := service.NewRegisterService(repository, brcrypt)
		err := service.RegisterAccount(c.Request.Context(), &req)
		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, true)
	}
}
