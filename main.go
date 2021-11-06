package main

import (
	"os"

	"github.com/kieranajp/langoustine/cmd"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"
)

type Config struct {
	LogLevel string
}

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "serve",
				Usage: "Start recipe server",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "listen-addr",
						Usage:   "Listen Address",
						Value:   "127.0.0.1:8080",
						EnvVars: []string{"LISTEN_ADDRESS"},
					},
					&cli.StringFlag{
						Name:    "db-dsn",
						Usage:   "Database DSN",
						Value:   "postgres://127.0.0.1:5432/recipes?sslmode=disable",
						EnvVars: []string{"DB_DSN"},
					},
				},
				Action: cmd.Serve,
			},
		},
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal().Err(err).Msg("Exit")
	}
}
