package cmd

import (
	"github.com/kieranajp/langoustine/pkg/database"
	"github.com/kieranajp/langoustine/pkg/handler"
	"github.com/urfave/cli/v2"
)

func Export(c *cli.Context) error {
	db, err := database.Connect(c.String("db-dsn"))
	if err != nil {
		return err
	}
	defer db.Close()

	h := handler.NewExporter(handler.NewBaseHandler(db))
	h.ExportToEpub(c.Args().First())

	return nil
}
