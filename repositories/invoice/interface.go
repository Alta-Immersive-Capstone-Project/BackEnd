package invoice

import (
	"kost/entities"
)

type InvoiceModel interface {
	CreateInvoice(path string, transaction entities.TransactionUpdateResponse) string
	CreateReport(path string, transactions []entities.TransactionKost) string
}
