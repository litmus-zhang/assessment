package util

import (
	"testing"

	"github.com/stretchr/testify/require"
	"golang.org/x/crypto/bcrypt"
)

func TestPasswordHash(t *testing.T) {
	password := RandomString(8)

	hash1, err := HashPassword(password)
	require.NoError(t, err)
	require.NotEmpty(t, hash1)

	err = CheckPasswordHash(password, hash1)
	require.NoError(t, err)

	wrongPassword := RandomString(8)
	err = CheckPasswordHash(wrongPassword, hash1)
	require.Error(t, err, bcrypt.ErrMismatchedHashAndPassword.Error())

	hash2, err := HashPassword(password)

	require.NoError(t, err)
	require.NotEqual(t, hash1, hash2)

}
