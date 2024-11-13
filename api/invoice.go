package api

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/litmus-zhang/assessment/db"
)

type getInvoiceData struct {
	CompanyID  int64 `uri:"company_id" binding:"required,min=1"`
	InvoiceID  int64 `uri:"invoice_id" binding:"required,min=1"`
	CustomerID int64 `uri:"customer_id" binding:"required,min=1"`
}
type fullInvoiceData struct {
	CompanyDetail  db.CompanyDetail
	Payment        []db.PaymentDetail
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
		log.Println("company not found")
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	invoice, err := server.db.GetOneInvoice(ctx, db.GetOneInvoiceParams{
		ID:        req.InvoiceID,
		CompanyID: req.CompanyID,
	})
	if err != nil {
		log.Println("Invoice not found")

		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	customer, err := server.db.GetCustomerByID(ctx, db.GetCustomerByIDParams{
		ID:        req.CustomerID,
		CompanyID: req.CompanyID,
	})
	if err != nil {
		log.Println("customer not found")

		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	payments, err := server.db.ListAllCompanyPaymentDetails(ctx, db.ListAllCompanyPaymentDetailsParams{
		CompanyID: req.CompanyID,
	})
	if err != nil {
		log.Println("payment not found")
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	fullInvoiceData := fullInvoiceData{
		CompanyDetail:  company,
		Payment:        payments,
		CustomerDetail: customer,
		InvoiceDetail:  invoice,
	}
	ctx.JSON(http.StatusOK, fullInvoiceData)

}

type createInvoiceData struct {
	CompanyID  int64     `json:"company_id" binding:"required,min=1"`
	CustomerID int64     `json:"customer_id" binding:"required,min=1"`
	Name       string    `json:"name" binding:"required"`
	DueDate    time.Time `json:"due_date" binding:"required"`
	Status     string    `json:"status" `
	Note       string    `json:"note"`
	Discount   string    `json:"discount"`
}

func (server *Server) createInvoice(ctx *gin.Context) {
	var req createInvoiceData
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}
	if req.Status == "" {
		req.Status = "draft"
	}
	args := db.CreateInvoiceParams{
		CompanyID:  req.CompanyID,
		CustomerID: req.CustomerID,
		Name:       req.Name,
		DueDate:    req.DueDate,
		Status:     req.Status,
		Note:       sql.NullString{String: req.Note, Valid: req.Note != ""},
		Discount:   req.Discount,
	}
	invoice, err := server.db.CreateInvoice(ctx, args)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusCreated, invoice)
}
