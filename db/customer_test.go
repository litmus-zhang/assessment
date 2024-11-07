package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/litmus-zhang/assessment/util"
	"github.com/stretchr/testify/require"
)

func TestCreateCustomer(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	customer := CreateTestCustomer(t, c.ID)
	require.NotEmpty(t, customer)

}

func TestGetCustomer(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	customer := CreateTestCustomer(t, c.ID)
	arg := GetCustomerByIDParams{
		ID:        customer.ID,
		CompanyID: c.ID,
	}

	cus, err := testQueries.GetCustomerByID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)
	require.Equal(t, cus, customer)

	arg2 := GetCustomerByEmailParams{
		Email:     customer.Email,
		CompanyID: c.ID,
	}

	cus, err = testQueries.GetCustomerByEmail(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, customer)
	require.Equal(t, cus, customer)
}

func TestListAllCompanyCustomers(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)

	for i := 0; i < 10; i++ {
		CreateTestCustomer(t, c.ID)
	}
	arg := ListCustomersParams{
		CompanyID: c.ID,
		Limit:     5,
		Offset:    0,
	}

	customers, err := testQueries.ListCustomers(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customers)
	require.Len(t, customers, 5)
}
func TestUpdateCustomer(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	customer := CreateTestCustomer(t, c.ID)
	arg := UpdateCustomerParams{
		ID:          customer.ID,
		CompanyID:   c.ID,
		FirstName:   util.RandomUser() + "Updated",
		LastName:    "Updated" + util.RandomUser(),
		Email:       util.RandomUser() + `@example.com`,
		PhoneNumber: util.RandomPhoneNumber(),
		Address:     sql.NullString{String: util.RandomString(5) + "Street", Valid: true},
	}
	cus, err := testQueries.UpdateCustomer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, customer)
	require.Equal(t, cus.ID, customer.ID)
	require.NotEqual(t, cus, customer)
}

func TestDeleteCustomer(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	customer := CreateTestCustomer(t, c.ID)

	arg := DeleteCustomerParams{
		ID:        customer.ID,
		CompanyID: c.ID,
	}
	err := testQueries.DeleteCustomer(context.Background(), arg)
	require.NoError(t, err)
}
