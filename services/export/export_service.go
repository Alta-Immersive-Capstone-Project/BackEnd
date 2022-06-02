package export

import (
	"encoding/csv"
	"kost/repositories/export"
	"kost/utils/s3"
	"os"

	"github.com/labstack/gommon/log"
)

type ServiceExport struct {
	repo export.RepoExport
	s3   s3.S3Control
}

func NewServiceExport(Repo export.RepoExport, S3 s3.S3Control) *ServiceExport {
	return &ServiceExport{
		repo: Repo,
		s3:   S3,
	}
}

func (s *ServiceExport) CreateFileReport(time int) error {
	data, err := s.repo.ReportTransaction(7)
	if err != nil {
		log.Warn(err)
		return err
	}
	csvFile, err := os.Create("report.csv")
	if err != nil {
		log.Warn("failed creating file: %s", err)
	}
	csvwriter := csv.NewWriter(csvFile)

	for _, empRow := range data {
		_ = csvwriter.Write(empRow)
	}
	csvwriter.Flush()
	csvFile.Close()
}
