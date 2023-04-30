package middleware

import (
	"weshare/common"
	"weshare/components/appcontext"

	"github.com/gin-gonic/gin"
)

// Recover recover all response when panic called
// https://go.dev/blog/defer-panic-and-recover
func Recover(ctx appcontext.AppContext) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				c.Header("Content-Type", "application/json")

				if appError, ok := err.(*common.AppError); ok {
					c.AbortWithStatusJSON(appError.StatusCode, appError)
					panic(err)
				}

				appError := common.ErrInternal(err.(error))
				c.AbortWithStatusJSON(appError.StatusCode, appError)

				panic(err)
			}
		}()
		c.Next()
	}
}
