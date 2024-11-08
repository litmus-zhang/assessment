package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/litmus-zhang/assessment/db"
)

type Server struct {
	db     db.Querier
	router *gin.Engine
}

func AppSetup(db db.Querier) *Server {
	server := &Server{db: db}

	router := gin.Default()

	api := router.Group("/api/v1")
	api.GET("/health", server.healthCheck)
	api.GET("/dashboard/:company_id", server.dashboard)
	api.GET("/company/:company_id/invoice/:invoice_id/customers/:customer_id", server.getSingleInvoice)
	api.POST("/invoice", server.createInvoice)

	server.router = router
	return server
}

func (server *Server) Start(address string) error {
	return server.router.Run(address)
}
func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) healthCheck(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "All system operational",
	})
}
