package api

import (
	"fmt"

	db "github.com/cyanrad/university/db/sqlc"
	"github.com/cyanrad/university/token"
	"github.com/cyanrad/university/util"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// Srever struct to serve HTTP responses
type Server struct {
	store      db.Store    // for executing database queries
	router     *gin.Engine // the gin server
	config     util.Config // configuration object
	tokenMaker token.Maker // Token maker for user session
}

// >> purpose: all paths in the same location
// for testing and production
var pathsURI = map[string]string{
	"createUser":  "/users",
	"loginUser":   "/users/login",
	"createEvent": "/event",
	"getEvent":    "/event/:id",
}

// >> create new server
func NewServer(config util.Config, store db.Store) (*Server, error) {
	fmt.Println(config.TokenSymmetricKey)
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		tokenMaker: tokenMaker,
		config:     config,
		store:      store,
	}

	server.setupRouter()
	return server, nil
}

// >> create a response from the generated go error
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func (server *Server) setupRouter() {
	router := gin.Default()

	// >> no idea what that does. something about security and stuff
	// >> but without it can't communicate with the front-end
	router.Use(cors.Default())

	// >> user
	router.POST(pathsURI["createUser"], server.createUser)
	router.POST(pathsURI["loginUser"], server.loginUser)

	// >> event
	router.POST(pathsURI["createEvent"], server.createEvent)
	router.GET(pathsURI["getEvent"], server.getEvent)

	// >> creating a new route group
	// for the auth middleware
	// for now, no operation requires it
	/*authRoutes := */
	router.Group("/").Use(authMiddleware(server.tokenMaker))

	server.router = router
}
