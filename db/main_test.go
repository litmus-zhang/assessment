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

func CreateTestCompany(t *testing.T, userID int64) CompanyDetail {

	arg := CreateCompanyParams{
		Name:        "Test Company",
		OwnedBy:     int32(userID),
		Address:     "Test Address",
		Email:       util.RandomUser() + "@company.com",
		PhoneNumber: util.RandomPhoneNumber(),
	}
	c, err := testQueries.CreateCompany(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, c)
	require.Equal(t, arg.Name, c.Name)
	require.Equal(t, arg.OwnedBy, c.OwnedBy)
	require.Equal(t, arg.Address, c.Address)
	require.Equal(t, arg.Email, c.Email)
	require.Equal(t, arg.PhoneNumber, c.PhoneNumber)
	require.NotZero(t, c.ID)

	return c
}

func CreateTestPaymentDetail(t *testing.T, companyID int64) PaymentDetail {
	arg := CreatePaymentDetailParams{
		AccountName:   util.RandomUser() + " Account",
		AccountNumber: util.RandomNumber(10),
		BankName:      util.RandomUser() + " Bank",
		CompanyID:     companyID,
	}
	p, err := testQueries.CreatePaymentDetail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, p)
	require.Equal(t, arg.AccountName, p.AccountName)
	require.Equal(t, arg.AccountNumber, p.AccountNumber)
	require.Equal(t, arg.BankName, p.BankName)
	require.Equal(t, arg.CompanyID, p.CompanyID)
	require.NotZero(t, p.ID)

	return p
}
