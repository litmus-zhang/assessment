package db

import (
	"context"
	"testing"

	"github.com/litmus-zhang/assessment/util"
	"github.com/stretchr/testify/require"
)

func TestCreatePaymentDetail(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	CreateTestPaymentDetail(t, c.ID)

}

func TestGetPaymentDetailByID(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	p := CreateTestPaymentDetail(t, c.ID)
	arg := GetACompanyPaymentDetailByIDParams{
		ID:        p.ID,
		CompanyID: c.ID,
	}

	paymentDetail, err := testQueries.GetACompanyPaymentDetailByID(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, paymentDetail)
	require.Equal(t, p, paymentDetail)
}

func TestGetAllPaymentDetailByCompanyID(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	for i := 0; i < 10; i++ {
		CreateTestPaymentDetail(t, c.ID)
	}
	arg := ListAllCompanyPaymentDetailsParams{
		CompanyID: c.ID,
		Limit:     5,
		Offset:    0,
	}
	paymentDetails, err := testQueries.ListAllCompanyPaymentDetails(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, paymentDetails)
	require.Len(t, paymentDetails, 5)
}
func TestDeletePaymentDetail(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	p := CreateTestPaymentDetail(t, c.ID)
	arg := DeletePaymentDetailParams{
		ID:        p.ID,
		CompanyID: c.ID,
	}
	err := testQueries.DeletePaymentDetail(context.Background(), arg)
	require.NoError(t, err)
}

func TestUpdatePaymentDetail(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	p := CreateTestPaymentDetail(t, c.ID)
	arg := UpdatePaymentDetailParams{
		ID:            p.ID,
		AccountName:   p.AccountName + " Updated",
		AccountNumber: util.RandomNumber(10),
		BankName:      p.BankName + " Updated",
		CompanyID:     p.CompanyID,
	}
	paymentDetail, err := testQueries.UpdatePaymentDetail(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, paymentDetail)
	require.NotEqual(t, p, paymentDetail)
}
