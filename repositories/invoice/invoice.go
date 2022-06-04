package invoice

import (
	"fmt"
	"kost/entities"

	"github.com/paimanbandi/rupiah"
	"github.com/unidoc/unipdf/v3/creator"
	"github.com/unidoc/unipdf/v3/model"
)

type invoiceModel struct {
	c *creator.Creator
}

func NewInvoiceModel(c *creator.Creator) *invoiceModel {
	return &invoiceModel{c: c}
}

func (m *invoiceModel) CreateInvoice(path string, transaction entities.TransactionUpdateResponse) string {
	m.c.NewPage()

	logo, err := m.c.NewImageFromFile(path)
	if err != nil {
		fmt.Println(err)
	}

	invoice := m.c.NewInvoice()

	// Set invoice logo
	invoice.SetLogo(logo)

	// Set invoice information
	invoice.SetNumber(transaction.BookingID)
	invoice.SetDate(transaction.UpdatedAt.Format("2006-01-02"))
	invoice.SetDueDate(transaction.UpdatedAt.Format("2006-01-02"))
	invoice.AddInfo("Payment terms", "Due on receipt")
	invoice.AddInfo("Paid", "No")

	// Set invoice addresses
	invoice.SetSellerAddress(&creator.InvoiceAddress{
		Name:   transaction.Title,
		Street: transaction.Address,
	})

	invoice.SetBuyerAddress(&creator.InvoiceAddress{
		Name:  transaction.Name,
		Phone: transaction.Phone,
	})

	// Add products to invoice

	invoice.AddLine(
		transaction.Title,
		"1",
		rupiah.FormatInt64ToRp(transaction.Price),
		rupiah.FormatInt64ToRp(transaction.Price),
	)

	// Set invoice totals
	invoice.SetSubtotal(rupiah.FormatInt64ToRp(transaction.Price))
	// invoice.AddTotalLine("Tax (10%)", "$10.00")
	// invoice.AddTotalLine("Shipping", "$5.00")
	invoice.SetTotal(rupiah.FormatInt64ToRp(transaction.Price))

	// Set invoice content sections
	invoice.SetNotes("Notes", "Thank you for your rent kost")
	invoice.SetTerms("Terms and conditions", "Full refund for 60 days after purchase.")

	m.CustomizeColor(invoice)
	m.c.Draw(invoice)
	output := m.c.WriteToFile(transaction.BookingID + ".pdf")
	if output != nil {
		fmt.Println(err)
	}

	return transaction.BookingID + ".pdf"
}

func (m *invoiceModel) CustomizeColor(i *creator.Invoice) {
	// Load custom font
	fontHelvetica := model.NewStandard14FontMustCompile(model.HelveticaName)

	// Create colors from RGB
	lightBlue := creator.ColorRGBFrom8bit(217, 240, 250)
	red := creator.ColorRGBFrom8bit(225, 0, 0)

	// Set invoice title text style
	i.SetTitleStyle(creator.TextStyle{
		Color:    red,
		Font:     fontHelvetica,
		FontSize: 32,
	})

	// Set invoice address heading style
	i.SetAddressHeadingStyle(creator.TextStyle{
		Font:     fontHelvetica,
		Color:    red,
		FontSize: 16,
	})

	// Set columns and rows styling
	//  Line formatting can be changed immediately after adding a line
	for cn, col := range i.Columns() {
		col.BackgroundColor = lightBlue
		col.BorderColor = lightBlue
		col.TextStyle.FontSize = 9
		col.Alignment = creator.CellHorizontalAlignmentCenter

		for _, line := range i.Lines() {
			line[cn].BorderColor = lightBlue
			line[cn].TextStyle.FontSize = 9
			line[cn].Alignment = creator.CellHorizontalAlignmentCenter
		}
	}

	// Change Total text syle
	titleCell, contentCell := i.Total()
	titleCell.BackgroundColor = lightBlue
	titleCell.BorderColor = lightBlue
	contentCell.BackgroundColor = lightBlue
	contentCell.BorderColor = lightBlue

	// Set Note text style
	i.SetNoteHeadingStyle(creator.TextStyle{
		Color:    red,
		Font:     fontHelvetica,
		FontSize: 16,
	})
}
