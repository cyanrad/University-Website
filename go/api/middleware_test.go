package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/cyanrad/university/token"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

// >> adds token and auth header to request
// used for the setupAuth in the testing
func addAuthorization(
	t *testing.T,
	request *http.Request,
	tokenMaker token.Maker,
	authorizationType string,
	username string,
	duration time.Duration,
) {
	token, err := tokenMaker.CreateToken(username, duration)
	require.NoError(t, err)

	authorizationHeader := fmt.Sprintf("%s %s", authorizationType,
		token)
	request.Header.Set(authorizationHeaderKey, authorizationHeader)
}

// >> the testing function
func TestAuthMiddleware(t *testing.T) {
	testCases := []struct {
		name string
		// >> auth header setup
		setupAuth func(t *testing.T, request *http.Request, tokenMaker token.Maker)
		// >> checks correctness of response
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{ // >> happy case
			name: "OK",
			setupAuth: func(t *testing.T, request *http.Request,
				tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker,
					authorizationTypeBearer, "user",
					time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
			},
		},
		{ // >> Empty header/ no token&auth
			name: "NoAuthorization",
			setupAuth: func(t *testing.T, request *http.Request,
				tokenMaker token.Maker) {
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				// >> should respond 401
				require.Equal(t, http.StatusUnauthorized,
					recorder.Code)
			},
		},
		{ // >> unsupported auth type (supported = bearer)
			name: "UnsupportedAuthorization",
			setupAuth: func(t *testing.T, request *http.Request,
				tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, "unsupported",
					"user", time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized,
					recorder.Code)
			},
		},
		{ // >> invalid header format
			name: "InvalidAuthorizationFormat",
			setupAuth: func(t *testing.T, request *http.Request,
				tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker, "", "user",
					time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized,
					recorder.Code)
			},
		},
		{ // >> expired token (negative time)
			name: "ExpiredToken",
			setupAuth: func(t *testing.T, request *http.Request,
				tokenMaker token.Maker) {
				addAuthorization(t, request, tokenMaker,
					authorizationTypeBearer, "user",
					-time.Minute)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusUnauthorized,
					recorder.Code)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]

		t.Run(tc.name, func(t *testing.T) {
			// >> nil, since we don't need a store
			server := NewTestServer(t, nil)

			// >> adding a temporary path
			authPath := "/auth"
			server.router.GET(
				authPath,
				// >> creating middleware
				authMiddleware(server.tokenMaker),
				func(ctx *gin.Context) {
					// >> @ success send OK
					ctx.JSON(http.StatusOK, gin.H{})
				},
			)

			recorder := httptest.NewRecorder()
			request, err := http.NewRequest(
				http.MethodGet, authPath, nil)
			require.NoError(t, err)

			// >> adding header data
			tc.setupAuth(t, request, server.tokenMaker)

			// >> sending request
			server.router.ServeHTTP(recorder, request)
			tc.checkResponse(t, recorder)
		})
	}
}
