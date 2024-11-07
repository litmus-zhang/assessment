package db

import (
	"context"
	"testing"

	"github.com/litmus-zhang/assessment/util"
	"github.com/stretchr/testify/require"
)

func TestCreateItem(t *testing.T) {

	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	cus := CreateTestCustomer(t, c.ID)
	i := CreateTestInvoice(t, c.ID, cus.ID)
	CreateTestInvoiceItem(t, i.ID)

}

func TestDeleteInvoiceItem(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	cus := CreateTestCustomer(t, c.ID)
	i := CreateTestInvoice(t, c.ID, cus.ID)
	ii := CreateTestInvoiceItem(t, i.ID)

	err := testQueries.DeleteItem(context.Background(), ii.ID)
	require.NoError(t, err)

}

func TestGetAllItemsForAnInvoice(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	cus := CreateTestCustomer(t, c.ID)
	i := CreateTestInvoice(t, c.ID, cus.ID)
	for k := 0; k < 10; k++ {
		CreateTestInvoiceItem(t, i.ID)
	}

	arg := GetAlltemsForAnInvoiceParams{
		InvoiceID: i.ID,
		Limit:     5,
		Offset:    0,
	}

	items, err := testQueries.GetAlltemsForAnInvoice(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, items)
	require.Len(t, items, 5)
}

func TestUpdateInvoiceItem(t *testing.T) {
	u := CreateTestUser(t)
	c := CreateTestCompany(t, u.ID)
	cus := CreateTestCustomer(t, c.ID)
	i := CreateTestInvoice(t, c.ID, cus.ID)
	ii := CreateTestInvoiceItem(t, i.ID)

	arg := UpdateItemParams{
		ID:        ii.ID,
		Name:      util.RandomString(6) + " item",
		UnitPrice: "100",
		Quantity:  10,
		InvoiceID: i.ID,
	}
	item, err := testQueries.UpdateItem(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, item)
	require.Equal(t, ii.ID, item.ID)
	require.NotEqual(t, ii, item)

}
