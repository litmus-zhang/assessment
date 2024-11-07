package db

import (
	"context"
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/litmus-zhang/assessment/util"
	"github.com/stretchr/testify/require"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbUrl    = "postgresql://main:main@localhost:4000/main?sslmode=disable"
)

func TestMain(m *testing.M) {
	var err error

	testDB, err = sql.Open(dbDriver, dbUrl)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}
	testQueries = New(testDB)

	code := m.Run()

	testDB.Close()

	os.Exit(code)

}

func CreateTestUser(t *testing.T) User {
	hash, err := util.HashPassword(util.RandomString(8))

	require.NoError(t, err)

	arg := CreateUserParams{
		FirstName: util.RandomUser(),
		LastName:  util.RandomUser(),
		Email:     util.RandomUser() + "@test.com",
		Password:  hash,
	}
	user, err := testQueries.CreateUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, user)
	require.Equal(t, arg.FirstName, user.FirstName)
	require.Equal(t, arg.LastName, user.LastName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)

	return user
}
