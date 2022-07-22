package token

import (
	"testing"
	"time"

	"github.com/cyanrad/university/util"
	"github.com/stretchr/testify/require"
)

// >> testing the successful case of paseto maker
func TestPasetoMaker(t *testing.T) {
	// >> creating a new maker
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	ID := util.RandomID()
	duration := time.Minute // expires a minute after

	issuedAt := time.Now() // for testing time and duration
	expireAt := issuedAt.Add(duration)

	// >> creating a new token
	token, err := maker.CreateToken(ID, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// >> reading the token and verifying it's done correctly
	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	// >> checking payload data
	require.NotZero(t, payload.StudentID)
	require.Equal(t, ID, payload.StudentID)
	require.WithinDuration(t, issuedAt, payload.IssuedAt, time.Second)
	require.WithinDuration(t, expireAt, payload.ExpiredAt, time.Second)
}

func TestExpiredToken(t *testing.T) {
	// >> creating new maker
	maker, err := NewPasetoMaker(util.RandomString(32))
	require.NoError(t, err)

	// >> specifying invalid expirationg time
	token, err := maker.CreateToken(util.RandomID(), -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	// >> payload should not contain data, since token is expired.
	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
