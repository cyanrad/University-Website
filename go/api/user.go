package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	db "github.com/cyanrad/university/db/sqlc"
	"github.com/cyanrad/university/util"
	"github.com/gin-gonic/gin"
	"github.com/lib/pq"
)

// >> Create New User parameters
type createUserRequest struct {
	ID       string `json:"id" binding:"required,numeric"`
	Password string `json:"password" binding:"required,min=8"`
	FullName string `json:"full_name" binding:"required"`
}

// >> purpose: so that we don't send the hashed password back when creating a user
type userResponse struct {
	ID        string    `json:"id"`
	FullName  string    `json:"full_name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUserResponse(user db.User) userResponse {
	return userResponse{
		ID:        user.ID,
		FullName:  user.FullName,
		CreatedAt: user.CreatedAt,
	}
}

// >> create user handler
// path: "/users"
// creats a new user and inserts them into the users table
func (server *Server) createUser(ctx *gin.Context) {
	// >> reading request data into req var
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// >> hashing the password
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// >> creating sql query parameters
	arg := db.CreateUserParams{
		ID:             req.ID,
		HashedPassword: hashedPassword,
		FullName:       req.FullName,
	}

	// >> inserting the new account into users table
	user, err := server.store.CreateUser(ctx, arg) // query
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() { // check and respond @ sql violation
			case "unique_violation":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// >> send a response with created account
	resp := NewUserResponse(user)
	ctx.JSON(http.StatusOK, resp)
}

type loginRequest struct {
	ID       string `json:"id" binding:"required,numeric"`
	Password string `json:"password" binding:"required,min=8"`
}

type loginResponse struct {
	AccessToken string `json:"access_token"`
	FullName    string `json:"full_name"`
}

// >> checks the credentials of the user
// response contains either true or false
// path: "/login"
func (server *Server) loginUser(ctx *gin.Context) {
	// >> reading the GET request
	var req loginRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// >> getting the user's data
	credentials, err := server.store.GetUser(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// >> comparing login information
	err = util.CheckPassword(req.Password, credentials.HashedPassword)
	if err != nil {
		err = fmt.Errorf("error: incorrect credentials")
		ctx.JSON(http.StatusForbidden, errorResponse(err))
		return
	}

	// >> creating access token
	accessToken, err := server.tokenMaker.CreateToken(
		credentials.ID,
		server.config.AccessTokenDuration,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// >> if all is good
	resp := loginResponse{
		AccessToken: accessToken,
		FullName:    credentials.FullName,
	}
	ctx.JSON(http.StatusOK, resp)
}
