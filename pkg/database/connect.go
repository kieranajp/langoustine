package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func Connect(dsn string) (*sqlx.DB, error) {
	// dbconfig, err := pgxpool.ParseConfig(dsn)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("failed to parse database config")
	// 	return nil, err
	// }
	// dbconfig.AfterConnect = func(ctx context.Context, conn *pgx.Conn) error {
	// 	conn.ConnInfo().RegisterDataType(pgtype.DataType{
	// 		Value: &pgtypeuuid.UUID{},
	// 		Name:  "uuid",
	// 		OID:   pgtype.UUIDOID,
	// 	})
	// 	return nil
	// }
	// db, err := pgxpool.ConnectConfig(context.Background(), dbconfig)
	// if err != nil {
	// 	log.Fatal().Err(err).Msg("failed to connect to database")
	// 	return nil, err
	// }

	var db *sqlx.DB
	db = sqlx.MustConnect("postgres", dsn)
	db.SetMaxOpenConns(1000) // The default is 0 (unlimited)
	db.SetMaxIdleConns(10)   // defaultMaxIdleConns = 2
	db.SetConnMaxLifetime(0) // 0, connections are reused forever.

	return db, nil
}
