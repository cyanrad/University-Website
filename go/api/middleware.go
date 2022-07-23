package api

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/cyanrad/university/token"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// >> creats the gin auth middleware
func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	// >> custom func to avoid abort boilerplate
	abort := func(ctx *gin.Context, err error) {
		ctx.JSON(http.StatusUnauthorized, errorResponse(err))
		ctx.Abort()
	}

	// >> returns middleware handler
	return func(ctx *gin.Context) {
		// >> getting the auth bytes
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		// >> if the authorization tag doesn't exist
		// >> or if it's not provided
		if len(authorizationHeader) == 0 {
			abort(ctx, errors.New("authorization header is not provided"))
			return
		}

		// >> Splitting auth header by space
		// >> into token type and body
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			abort(ctx, errors.New("invalid authorization header format"))
			return
		}

		// >> getting the autherization type
		// converting to lower for comparison ease
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			abort(ctx, fmt.Errorf("unsupported authorization type %s",
				authorizationType))
			return
		}

		// >> getting the token
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			abort(ctx, err)
			return
		}

		// >> storing the payload in the context
		// >> to be accessed by the next header
		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next() // >> forwarding request
	}
}
