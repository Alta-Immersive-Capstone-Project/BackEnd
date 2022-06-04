package unipdf

import (
	"fmt"
	"kost/configs"

	"github.com/unidoc/unipdf/v3/common/license"
	"github.com/unidoc/unipdf/v3/creator"
)

func NewInitPdf(config *configs.AppConfig) *creator.Creator {
	// To get your free API key for metered license, sign up on: https://cloud.unidoc.io
	// Make sure to be using UniPDF v3.19.1 or newer for Metered API key support.
	err := license.SetMeteredKey(config.UniPdfLicense)
	if err != nil {
		fmt.Println(err)
		fmt.Println("Make sure to get a valid key from https://cloud.unidoc.io")
	}

	lk := license.GetLicenseKey()
	if lk == nil {
		fmt.Println("Failed retrieving license key")
	}

	connect := creator.New()
	return connect
}
