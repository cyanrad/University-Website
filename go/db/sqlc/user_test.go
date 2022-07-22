package db

import (
	"context"
	"testing"
	"time"

	"github.com/cyanrad/university/util"
	"github.com/stretchr/testify/require"
)

// >> creates a random user & tests if the creation is done successfully
func CreateRandomUser(t *testing.T) User {
	// >> hashing passowrd
	hash, err := util.HashPassword(util.RandomString(8))
	require.NoError(t, err)

	// >> creating query params
	arg := CreateUserParams{
		ID:             util.RandomID(),
		HashedPassword: hash,
		FullName:       util.RandomString(10),
	}

	// >> doing the query
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)

	// >> checking data correctness
	require.Equal(t, arg.ID, user.ID)
	require.Equal(t, arg.HashedPassword, user.HashedPassword)
	require.Equal(t, arg.FullName, user.FullName)
	require.NotZero(t, user.CreatedAt)

	return user
}

// >> Tests the creation of a random user in the database
func TestCreateUser(t *testing.T) {
	CreateRandomUser(t)
}

// >> Test getting a user from the database
func TestGetUser(t *testing.T) {
	user1 := CreateRandomUser(t) // >> creating a user then fetching them
	user2, err := testQueries.GetUser(context.Background(), user1.ID)
	require.NoError(t, err)

	// >> comparing original created user, and the one we got from the GetUser() func
	require.Equal(t, user1.ID, user2.ID)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.WithinDuration(t, user1.CreatedAt, user2.CreatedAt, time.Second)
}
