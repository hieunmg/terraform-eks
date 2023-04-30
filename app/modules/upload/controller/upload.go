package controller

import (
	"mime/multipart"
	"net/http"
	"weshare/common"
	"weshare/components/appcontext"
	"weshare/modules/upload/service"

	_ "image/jpeg"
	_ "image/png"

	"github.com/gin-gonic/gin"
)

func Upload(appCtx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		folder := c.DefaultPostForm("folder", "img")
		file, err := fileHeader.Open()

		if err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				panic(common.ErrInternal(err))
			}
		}(file)

		dataBytes := make([]byte, fileHeader.Size)
		if _, err := file.Read(dataBytes); err != nil {
			panic(common.ErrInvalidRequest(err))
		}

		service := service.NewUploadService(appCtx.GetUploadProvider(), nil)
		img, err := service.Upload(c.Request.Context(), dataBytes, folder, fileHeader.Filename)

		if err != nil {
			panic(err)
		}
		c.JSON(http.StatusOK, common.SimpleSuccessResponse(img))
	}
}
