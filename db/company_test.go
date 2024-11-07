package db

import (
	"context"
	"testing"

	"github.com/litmus-zhang/assessment/util"
	"github.com/stretchr/testify/require"
)

func TestCreateCompany(t *testing.T) {
	u := CreateTestUser(t)
	CreateTestCompany(t, u.ID)
}

func TestGetCompany(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)

	company, err := testQueries.GetCompany(context.Background(), c.ID)
	require.NoError(t, err)
	require.NotEmpty(t, company)
	require.Equal(t, c, company)
}

func TestUpdateCompany(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)

	arg := UpdateCompanyParams{
		ID:          c.ID,
		Name:        c.Name + " Updated",
		Address:     c.Address + " Updated",
		PhoneNumber: util.RandomPhoneNumber(),
		Email:       util.RandomUser() + "@company.com",
		OwnedBy:     c.OwnedBy,
	}
	company, err := testQueries.UpdateCompany(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, company)
	require.Equal(t, arg.Name, company.Name)
	require.NotEqual(t, c, company)
}
func TestGetUserCompanyByID(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)

	company, err := testQueries.GetCompanyCreatedByUser(context.Background(), GetCompanyCreatedByUserParams{
		OwnedBy: int32(u.ID),
		ID:      c.ID,
	})
	require.NoError(t, err)
	require.NotEmpty(t, company)
	require.Equal(t, c, company)
}

func TestGetAllUserCompany(t *testing.T) {
	u := CreateTestUser(t)
	for i := 0; i < 10; i++ {
		CreateTestCompany(t, u.ID)
	}
	arg := GetCompaniesCreatedByUserParams{
		OwnedBy: int32(u.ID),
		Limit:   5,
		Offset:  0,
	}

	companies, err := testQueries.GetCompaniesCreatedByUser(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, companies)
	require.Len(t, companies, 5)
}

func TestDeleteCompany(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)

	arg := DeleteCompanyParams{
		ID:      c.ID,
		OwnedBy: int32(u.ID),
	}
	err := testQueries.DeleteCompany(context.Background(), arg)
	require.NoError(t, err)

	company, err := testQueries.GetCompany(context.Background(), c.ID)
	require.Error(t, err)
	require.Empty(t, company)
}
