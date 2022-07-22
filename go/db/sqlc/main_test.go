package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/cyanrad/university/util"
	_ "github.com/lib/pq"
)

// >> used in *_test file of this package
var testQueries *Queries

// >> database operations unit testing entry point
func TestMain(m *testing.M) {

	// >> loading configuration from env file/variables
	config, err := util.LoadConfig("../../")
	if err != nil {
		log.Fatal("can't read config: ", err)
	}

	// >> Creating connection object
	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("@: Error: cannot connect to db:", err)
	}

	// >> connecting to database
	testQueries = New(conn)

	// >> running the unit tests
	os.Exit(m.Run())
}
