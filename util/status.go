package util

var (
	PENDING         = "PENDING"
	PAID            = "PAID"
	OVERDUE         = "OVERDUE"
	DRAFT           = "DRAFT"
	PARTIAL_PAYMENT = "PARTIAL_PAYMENT"
)

func GetRandomInvoiceStatus() string {
	status := []string{PENDING, PAID, OVERDUE, DRAFT, PARTIAL_PAYMENT}
	k := int64(len(status) - 1)
	return status[RandomInt(0, k)]
}
