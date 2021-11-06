package database

import (
	"context"

	"github.com/jackc/pgtype"
	pgtypeuuid "github.com/jackc/pgtype/ext/gofrs-uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/zerolog/log"
)

func Connect(dsn string) (*pgxpool.Pool, error) {
	dbconfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to parse database config")
		return nil, err
	}
	dbconfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
		conn.ConnInfo().RegisterDataType(pgtype.DataType{
			Value: &pgtypeuuid.UUID{},
			Name:  "uuid",
			OID:   pgtype.UUIDOID,
		})
		return nil
	}
	db, err := pgxpool.ConnectConfig(context.Background(), dbconfig)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to connect to database")
		return nil, err
	}

	return db, nil
}
