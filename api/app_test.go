package api

import (
	"bytes"
	"io"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/litmus-zhang/assessment/db"
	"github.com/stretchr/testify/require"
)

var testServer Server

const (
	dbDriver      = "postgres"
	dbUrl         = "postgresql://main:main@localhost:4000/main?sslmode=disable"
	ServerAddress = "localhost:8000"
)

func TestHealthCheck(t *testing.T) {

}

func TestDashboardDataPage(t *testing.T) {

}

func TestInvoicePage(t *testing.T) {

}

func z(t *testing.T) {
	type fields struct {
		db     db.Querier
		router *gin.Engine
	}
	type args struct {
		ctx *gin.Context
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := &Server{
				db:     tt.fields.db,
				router: tt.fields.router,
			}
			server.healthCheck(tt.args.ctx)
		})
	}
}

func requireResponseBodyMatch(t *testing.T, body *bytes.Buffer) {
	_, err := io.ReadAll(body)
	require.NoError(t, err)

}
