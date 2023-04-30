package middleware

import (
	"errors"
	"fmt"
	"strings"

	"weshare/common"
	"weshare/components/token"

	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "Authorization"
	authorizationTypeBearer = "Bearer"
)

type UnAuthUrls []string

var unAuthUrls UnAuthUrls = UnAuthUrls{
	"/api/v1/accounts/login",
	"/api/v1/accounts/register",
	"/api/v1/accounts/refresh",
	"/metrics",
	"/health",
}

func (slices *UnAuthUrls) Contains(url string) bool {
	for _, s := range *slices {
		if s == url {
			return true
		}
	}
	return false
}

// AuthMiddleware creates a gin middleware for authorization
func Authorize(secretKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		if unAuthUrls.Contains(c.FullPath()) {
			c.Next()
			return
		}

		tokenProvider, err := token.NewJWTMaker(secretKey)
		if err != nil {
			panic(err)
		}

		authorizationHeader := c.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			message := "authorization header is not provided"
			panic(common.NewUnauthorized(errors.New(message), message, "Unauthorized"))
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			message := "invalid authorization header format"
			panic(common.NewUnauthorized(errors.New(message), message, "Unauthorized"))
		}

		authorizationType := string(fields[0])
		if authorizationType != authorizationTypeBearer {
			message := fmt.Sprintf("unsupported authorization type %s", authorizationType)
			panic(common.NewUnauthorized(errors.New(message), message, "Unauthorized"))
		}

		accessToken := fields[1]
		payload, err := tokenProvider.VerifyToken(accessToken)
		if err != nil {
			panic(common.NewUnauthorized(err, err.Error(), "Unauthorized"))
		}

		c.Set(token.CurrentUser, payload)
		c.Next()
	}
}
