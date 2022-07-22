package api

import (
	"os"
	"testing"

	db "github.com/cyanrad/university/db/sqlc"
	"github.com/gin-gonic/gin"
)

// >> used for testing
func NewTestServer(t *testing.T, store db.Store) *Server {
	// will be added to later
	server := NewServer(store)

	return server
}

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	os.Exit(m.Run())
}
