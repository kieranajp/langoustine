package cmd

import (
	"fmt"
	"os"

	"github.com/kieranajp/langoustine/pkg/database"
	"github.com/kieranajp/langoustine/pkg/exporter"
	"github.com/pkg/errors"
	"github.com/urfave/cli/v2"
)

func Export(c *cli.Context) error {
	db, err := database.Connect(c.String("db-dsn"))
	if err != nil {
		return err
	}
	defer db.Close()

	var f exporter.Format
	var file *os.File
	var e exporter.Exporter
	switch c.String("format") {
	case "kindle":
		f = exporter.Kindle
		file, err = os.OpenFile(fmt.Sprintf("%s.%s", c.String("output"), f.Ext()), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return errors.Wrap(err, "Failed to open output file")
		}
		e = exporter.NewKindleExporter(db, file)
	case "epub":
		f = exporter.Epub
		file, err = os.OpenFile(fmt.Sprintf("%s.%s", c.String("output"), f.Ext()), os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			return errors.Wrap(err, "Failed to open output file")
		}
		//e = exporter.NewEpubExporter(db, file)
	default:
		return errors.New("Invalid format: must be one of 'kindle' or 'epub'")
	}
	defer file.Close()

	e.Export(c.Args().First())

	return nil
}
