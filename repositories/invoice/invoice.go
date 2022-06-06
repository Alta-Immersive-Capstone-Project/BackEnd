package invoice

import (
	"fmt"
	"image"
	"kost/entities"
	"math"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/labstack/gommon/log"
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
	now := time.Now()
	loc, _ := time.LoadLocation("Asia/Jakarta")

	// Note: without explicit zone, returns time in given location.
	// Set invoice information
	invoice.SetNumber(transaction.BookingID)
	invoice.SetDate(transaction.UpdatedAt.Format(time.RFC3339))
	invoice.SetDueDate(time.Date(now.Year(), now.Month(), now.Day()+1, now.Hour(), now.Minute(), 0, 0, loc).Format(time.RFC3339))
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
	m.c.Finalize()
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
	m.c.Finalize()
}

var (
	robotoRegular, _ = model.NewPdfFontFromTTFFile("Roboto-Regular.ttf")
	robotoBold, _    = model.NewPdfFontFromTTFFile("Roboto-Bold.ttf")
)

func (m *invoiceModel) DoFeatureOverview(transactions []entities.TransactionKost) {
	// Fonts

	// Ensure that the chapter starts on a new page.
	m.c.NewPage()

	ch := m.c.NewChapter("Feature overview")

	chapterFont := robotoRegular
	chapterFontColor := creator.ColorRGBFrom8bit(72, 86, 95)
	chapterFontSize := 18.0

	normalFont := robotoRegular
	normalFontColor := creator.ColorRGBFrom8bit(72, 86, 95)
	normalFontSize := 10.0

	ch.GetHeading().SetFont(chapterFont)
	ch.GetHeading().SetFontSize(chapterFontSize)
	ch.GetHeading().SetColor(chapterFontColor)

	sc := ch.NewSubchapter("Lists History Transaction")
	sc.GetHeading().SetMargins(0, 0, 20, 0)
	sc.GetHeading().SetFont(chapterFont)
	sc.GetHeading().SetFontSize(chapterFontSize)
	sc.GetHeading().SetColor(chapterFontColor)

	// Create table.
	table := m.c.NewTable(7)
	table.SetMargins(0, 0, 10, 0)

	drawCell := func(text string, font *model.PdfFont, align creator.CellHorizontalAlignment) {
		p := m.c.NewStyledParagraph()
		p.Append(text).Style.Font = font

		cell := table.NewCell()
		cell.SetBorder(creator.CellBorderSideAll, creator.CellBorderStyleSingle, 1)
		cell.SetHorizontalAlignment(align)
		cell.SetContent(p)
	}

	// Draw table header.
	drawCell("No", robotoBold, creator.CellHorizontalAlignmentCenter)
	drawCell("Booking ID", robotoBold, creator.CellHorizontalAlignmentCenter)
	drawCell("Penyewa", robotoBold, creator.CellHorizontalAlignmentCenter)
	drawCell("Duration", robotoBold, creator.CellHorizontalAlignmentCenter)
	drawCell("Price", robotoBold, creator.CellHorizontalAlignmentCenter)
	drawCell("Created At", robotoBold, creator.CellHorizontalAlignmentCenter)
	drawCell("Status", robotoBold, creator.CellHorizontalAlignmentCenter)

	// Draw table content.
	for i := 0; i < len(transactions); i++ {
		num := i + 1
		drawCell(fmt.Sprintf("%d", num), robotoRegular, creator.CellHorizontalAlignmentCenter)
		drawCell(transactions[i].BookingID, robotoRegular, creator.CellHorizontalAlignmentCenter)
		drawCell(transactions[i].Name, robotoRegular, creator.CellHorizontalAlignmentLeft)
		drawCell(fmt.Sprintf("%d Bln", transactions[i].Duration), robotoRegular, creator.CellHorizontalAlignmentLeft)
		drawCell(rupiah.FormatInt64ToRp(transactions[i].Price), robotoRegular, creator.CellHorizontalAlignmentLeft)
		drawCell(transactions[i].CreatedAt.Format("2006-01-02"), robotoRegular, creator.CellHorizontalAlignmentLeft)
		drawCell(transactions[i].TransactionStatus, robotoRegular, creator.CellHorizontalAlignmentLeft)
	}

	sc.Add(table)

	sc = ch.NewSubchapter("QR Codes")
	sc.GetHeading().SetMargins(0, 0, 20, 0)
	sc.GetHeading().SetFont(chapterFont)
	sc.GetHeading().SetFontSize(chapterFontSize)
	sc.GetHeading().SetColor(chapterFontColor)

	p := m.c.NewParagraph("For downloading the generated PDF file, scan the QR code with your barcode reader.")
	p.SetFont(normalFont)
	p.SetFontSize(normalFontSize)
	p.SetColor(normalFontColor)
	p.SetMargins(0, 0, 5, 5)
	sc.Add(p)

	qrCode, _ := makeQrCodeImage("https://localhost:8000/generate", 40, 5)
	img, err := m.c.NewImageFromGoImage(qrCode)
	if err != nil {
		panic(err)
	}
	img.SetWidth(40)
	img.SetHeight(40)
	sc.Add(img)
	m.c.Draw(ch)
}

func (m *invoiceModel) CreateReport(path string, transactions []entities.TransactionKost) string {
	m.c.SetPageMargins(50, 50, 100, 70)

	// Generate the table of contents.

	m.c.AddTOC = true
	toc := m.c.TOC()
	hstyle := m.c.NewTextStyle()
	hstyle.Color = creator.ColorRGBFromArithmetic(0.2, 0.2, 0.2)
	hstyle.FontSize = 28
	toc.SetHeading("Table of Contents", hstyle)
	lstyle := m.c.NewTextStyle()
	lstyle.FontSize = 14
	toc.SetLineStyle(lstyle)

	logoImg, err := m.c.NewImageFromFile(path)
	if err != nil {
		return ""
	}

	logoImg.ScaleToHeight(25)
	logoImg.SetPos(58, 20)

	m.DoFeatureOverview(transactions)

	// Setup a front page (always placed first).
	m.c.CreateFrontPage(func(args creator.FrontpageFunctionArgs) {
		m.DoFirstPage()
	})

	// Draw a header on each page.
	m.c.DrawHeader(func(block *creator.Block, args creator.HeaderFunctionArgs) {
		// Draw the header on a block. The block size is the size of the page's top margins.
		block.Draw(logoImg)
	})

	// Draw footer on each page.
	m.c.DrawFooter(func(block *creator.Block, args creator.FooterFunctionArgs) {
		// Draw the on a block for each page.
		p := m.c.NewParagraph("https://lowkost.com")
		p.SetFont(robotoRegular)
		p.SetFontSize(8)
		p.SetPos(50, 20)
		p.SetColor(creator.ColorRGBFrom8bit(63, 68, 76))
		block.Draw(p)

		strPage := fmt.Sprintf("Page %d of %d", args.PageNum, args.TotalPages)
		p = m.c.NewParagraph(strPage)
		p.SetFont(robotoRegular)
		p.SetFontSize(8)
		p.SetPos(300, 20)
		p.SetColor(creator.ColorRGBFrom8bit(63, 68, 76))
		block.Draw(p)
	})

	time := time.Now().UTC().Format("20060102-150405")
	output := m.c.WriteToFile("LOWKOST-" + time + ".pdf")
	if output != nil {
		log.Warn(output)
	}
	m.c.Finalize()
	return "LOWKOST-" + time
}

// Generates the front page.
func (m *invoiceModel) DoFirstPage() {
	helvetica, _ := model.NewStandard14Font("Helvetica")
	helveticaBold, _ := model.NewStandard14Font("Helvetica-Bold")

	p := m.c.NewParagraph("LOWKOST")
	p.SetFont(helvetica)
	p.SetFontSize(48)
	p.SetMargins(85, 0, 150, 0)
	p.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
	m.c.Draw(p)

	p = m.c.NewParagraph("History Transaction Report")
	p.SetFont(helveticaBold)
	p.SetFontSize(30)
	p.SetMargins(85, 0, 0, 0)
	p.SetColor(creator.ColorRGBFrom8bit(45, 148, 215))
	m.c.Draw(p)

	t := time.Now().UTC()
	dateStr := t.Format("01 Jan 2006 15:04")

	p = m.c.NewParagraph(dateStr)
	p.SetFont(helveticaBold)
	p.SetFontSize(12)
	p.SetMargins(90, 0, 5, 0)
	p.SetColor(creator.ColorRGBFrom8bit(56, 68, 77))
	m.c.Draw(p)
}

func makeQrCodeImage(text string, width float64, oversampling int) (image.Image, error) {
	qrCode, err := qr.Encode(text, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	pixelWidth := oversampling * int(math.Ceil(width))
	qrCode, err = barcode.Scale(qrCode, pixelWidth, pixelWidth)
	if err != nil {
		return nil, err
	}

	return qrCode, nil
}
