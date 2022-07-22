package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

// >> test is taken from <tech school: backend master class: #16>
func TestPasswordHash(t *testing.T) {
	password := RandomString(6)

	// >> hashing
	hashedPassword1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword1)

	// >> Pass comp func
	err = CheckPassword(password, hashedPassword1)
	require.NoError(t, err)

	// >> purposfully wrong password
	wrongPassword := RandomString(6)
	err = CheckPassword(wrongPassword, hashedPassword1)
	require.EqualError(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	// >> testing hash salt (2 string hashed with different
	// 							salts should not be the same)
	hashedPassword2, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hashedPassword2)
	require.NotEqual(t, hashedPassword1, hashedPassword2)
}
