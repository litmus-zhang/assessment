package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/litmus-zhang/assessment/db"
)

type getDashboardData struct {
	CompanyID int64 `uri:"company_id" binding:"required,min=1"`
	Page      int32 `form:"page" default:"1"`
	Size      int32 `form:"size" default:"10"`
}

func (server *Server) dashboard(ctx *gin.Context) {
	var req getDashboardData

	err := ctx.ShouldBindUri(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err = ctx.ShouldBindQuery(&req)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	summary, err := server.db.GetCompanyInvoiceSummary(ctx, req.CompanyID)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	log.Println("req:", req)
	invoices, err := server.db.GetAllInvoices(ctx, db.GetAllInvoicesParams{
		CompanyID: req.CompanyID,
		Limit:     req.Size,
		Offset:    (req.Page - 1) * req.Size,
	})

	if err != nil {
		ctx.JSON(http.StatusNotFound, errorResponse(err))
		return
	}
	type dashboardData struct {
		Summary    []db.GetCompanyInvoiceSummaryRow
		AllInvoice []db.Invoice
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Daashboard Data fetched successfully",
		"data": dashboardData{
			Summary:    summary,
			AllInvoice: invoices,
		},
	})
}
