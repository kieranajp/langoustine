package exporter

import (
	"io"

	"github.com/jmoiron/sqlx"
)

type EpubExporter struct {
	db *sqlx.DB
}

func (e *EpubExporter) Export(recipeUUID string) (io.Writer, error) {
	return nil, nil
}
