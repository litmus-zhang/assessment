package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/litmus-zhang/assessment/db"
)

type getInvoiceData struct {
	CompanyID  int64 `uri:"company_id" binding:"required,min=1"`
	InvoiceID  int64 `uri:"invoice_id" binding:"required,min=1"`
	CustomerID int64 `uri:"customer_id" binding:"required,min=1"`
}
type FullInvoiceData struct {
	CompanyDetail  db.CompanyDetail
	Payment        db.PaymentDetail
	CustomerDetail db.Customer
	InvoiceDetail  db.Invoice
}

func (server *Server) getSingleInvoice(ctx *gin.Context) {
	var req getInvoiceData

	err := ctx.ShouldBindUri(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	company, err := server.db.GetCompany(ctx, req.CompanyID)
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	invoice, err := server.db.GetOneInvoice(ctx, db.GetOneInvoiceParams{
		ID:        req.InvoiceID,
		CompanyID: req.CompanyID,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	customer, err := server.db.GetCustomerByID(ctx, db.GetCustomerByIDParams{
		ID:        req.CustomerID,
		CompanyID: req.CompanyID,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	payment, err := server.db.ListAllCompanyPaymentDetails(ctx, db.ListAllCompanyPaymentDetailsParams{
		CompanyID: req.CompanyID,
	})
	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	fullInvoiceData := FullInvoiceData{
		CompanyDetail:  company,
		Payment:        payment[0],
		CustomerDetail: customer,
		InvoiceDetail:  invoice,
	}
	ctx.JSON(http.StatusOK, fullInvoiceData)

}
