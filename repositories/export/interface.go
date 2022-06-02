package export

import "kost/entities"

type RepoExport interface {
	ReportTransaction(time int) ([]entities.House, error)
}
