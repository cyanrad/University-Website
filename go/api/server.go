package api

import (
	db "github.com/cyanrad/university/db/sqlc"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Srever struct to serve HTTP responses
type Server struct {
	store  db.Store    // for executing database queries
	router *gin.Engine // the gin server
}

// >> purpose: all paths in the same location
// for testing and production
var pathsURI = map[string]string{
	"createUser": "/users",
	"loginUser":  "/login",
}

// >> create new server
func NewServer(store db.Store) *Server {
	server := &Server{store: store}
	router := gin.Default()

	// >> no idea what that does. something about security
	// >> but without it can't communicate with the front-end
	router.Use(cors.Default())

	router.POST(pathsURI["createUser"], server.createUser)
	router.POST(pathsURI["loginUser"], server.loginUser)

	server.router = router
	return server
}

// >> create a response from the generated go error
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
