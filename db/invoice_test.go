package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/litmus-zhang/assessment/util"
	"github.com/stretchr/testify/require"
)

func TestCreateInvoice(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	cus := CreateTestCustomer(t, c.ID)
	CreateTestInvoice(t, c.ID, cus.ID)

}

func TestGetOneInvoice(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	cus := CreateTestCustomer(t, c.ID)
	i := CreateTestInvoice(t, c.ID, cus.ID)

	arg := GetOneInvoiceParams{
		ID:        i.ID,
		CompanyID: c.ID,
	}

	invoice, err := testQueries.GetOneInvoice(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, invoice)
	require.Equal(t, i, invoice)
}

func TestGetAllCompanyInvoices(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	cus := CreateTestCustomer(t, c.ID)
	for i := 0; i < 10; i++ {
		CreateTestInvoice(t, c.ID, cus.ID)
	}

	arg := GetAllInvoicesParams{
		CompanyID: c.ID,
		Limit:     5,
		Offset:    0,
	}

	invoices, err := testQueries.GetAllInvoices(context.Background(), arg)
	log.Println("Invoices: ", invoices)
	require.NoError(t, err)
	require.NotEmpty(t, invoices)
	require.Len(t, invoices, 5)
}

func TestUpdateInvoice(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	cus := CreateTestCustomer(t, c.ID)
	i := CreateTestInvoice(t, c.ID, cus.ID)

	arg := UpdateInvoiceParams{
		ID:        i.ID,
		CompanyID: c.ID,
		Name:      "Updated Invoice",
		DueDate:   util.RandomDateInFuture(5),
		Status:    util.GetRandomInvoiceStatus(),
		Note:      sql.NullString{String: "Updated Note", Valid: true},
		Discount:  fmt.Sprintf("%.2f", float64(util.RandomInt(0, 100))),
	}

	invoice, err := testQueries.UpdateInvoice(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, invoice)
	require.Equal(t, arg.Name, invoice.Name)
	require.NotEqual(t, i, invoice)
}

func TestDeleteInvoice(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	cus := CreateTestCustomer(t, c.ID)
	i := CreateTestInvoice(t, c.ID, cus.ID)

	arg := DeleteInvoiceParams{
		ID:        i.ID,
		CompanyID: c.ID,
	}

	err := testQueries.DeleteInvoice(context.Background(), arg)
	require.NoError(t, err)

	invoice, err := testQueries.GetOneInvoice(context.Background(), GetOneInvoiceParams{
		ID:        i.ID,
		CompanyID: c.ID,
	})
	require.Error(t, err)
	require.Empty(t, invoice)
}

func TestGetInvoiceByStatus(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	cus := CreateTestCustomer(t, c.ID)
	for i := 0; i < 10; i++ {
		CreateTestInvoice(t, c.ID, cus.ID)
	}
	arg := GetInvoicesByStatusParams{
		Status: util.PENDING,
		Limit:  5,
		Offset: 0,
	}
	inv, err := testQueries.GetInvoicesByStatus(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, inv)

}

func TestGetCompanyInvoiceSummary(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	cus := CreateTestCustomer(t, c.ID)
	for i := 0; i < 10; i++ {
		inv := CreateTestInvoice(t, c.ID, cus.ID)
		for j := 0; j < 3; j++ {
			CreateTestInvoiceItem(t, inv.ID)
		}
		err := testQueries.GetInvoiceTotalFromItems(context.Background(), inv.ID)
		require.NoError(t, err)
	}
	summary, err := testQueries.GetCompanyInvoiceSummary(context.Background(), c.ID)
	log.Printf("Summary: %+v", summary)
	require.NoError(t, err)
	require.NotEmpty(t, summary)
}
