package db

import (
	"context"
	"database/sql"
	"fmt"
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

func CreateTestCustomer(t *testing.T, companyID int64) Customer {
	arg := CreateCustomerParams{
		FirstName:   util.RandomUser(),
		LastName:    util.RandomUser(),
		Email:       util.RandomUser() + "@customer.com",
		PhoneNumber: util.RandomPhoneNumber(),
		CompanyID:   companyID,
		Address:     sql.NullString{String: "Test Address", Valid: true},
	}
	customer, err := testQueries.CreateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)
	require.Equal(t, arg.FirstName, customer.FirstName)
	require.Equal(t, arg.LastName, customer.LastName)
	require.Equal(t, arg.Email, customer.Email)
	require.Equal(t, arg.PhoneNumber, customer.PhoneNumber)
	require.Equal(t, arg.CompanyID, customer.CompanyID)
	require.NotZero(t, customer.ID)

	return customer
}

func CreateTestInvoice(t *testing.T, companyID int64, customerID int64) Invoice {
	arg := CreateInvoiceParams{
		CustomerID: customerID,
		Name:       util.RandomString(8) + " Invoice",
		DueDate:    util.RandomDateInFuture(7),
		Status:     util.GetRandomInvoiceStatus(),
		CompanyID:  companyID,
		Note:       sql.NullString{String: "Test Note", Valid: true},
		Discount:   fmt.Sprintf("%.2f", 10.0),
	}
	invoice, err := testQueries.CreateInvoice(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, invoice)
	require.Equal(t, arg.CustomerID, invoice.CustomerID)
	require.Equal(t, arg.Name, invoice.Name)
	require.Equal(t, arg.Status, invoice.Status)
	require.Equal(t, arg.CompanyID, invoice.CompanyID)
	require.Equal(t, arg.Note, invoice.Note)
	require.Equal(t, arg.Discount, invoice.Discount)
	require.NotZero(t, invoice.ID)

	return invoice
}

func CreateTestInvoiceItem(t *testing.T, invoiceID int64) Item {
	arg := CreateItemParams{
		Name:        util.RandomString(8),
		UnitPrice:   fmt.Sprintf("%.2f", 10.4),
		Description: "Test Description",
		InvoiceID:   invoiceID,
		Quantity:    int32(util.RandomInt(1, 30)),
	}
	item, err := testQueries.CreateItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)
	require.Equal(t, arg.Name, item.Name)
	require.Equal(t, arg.UnitPrice, item.UnitPrice)
	require.Equal(t, arg.Description, item.Description)
	require.Equal(t, arg.InvoiceID, item.InvoiceID)
	require.NotZero(t, item.ID)

	return item
}
