package datastore

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
	"log/slog"
	"os"
	"reliab-test/internal/config"
)

func InitDB(dbCfg config.DBConf, log *slog.Logger) *sqlx.DB {
	cfg, _ := pgxpool.ParseConfig("")
	cfg.ConnConfig.Host = dbCfg.Host
	cfg.ConnConfig.Port = dbCfg.Port
	cfg.ConnConfig.User = dbCfg.User
	cfg.ConnConfig.Password = dbCfg.Password
	cfg.ConnConfig.Database = dbCfg.DBName

	dbPool, err := pgxpool.NewWithConfig(context.Background(), cfg)
	if err != nil {
		log.Error("couldn't connect to database", err)
		os.Exit(1)
	}

	conn := stdlib.OpenDBFromPool(dbPool)
	sqlxDb := sqlx.NewDb(conn, "pgx")

	log.Info(fmt.Sprintf("Connected to database. User %s@%s:%d/%s", dbCfg.User, dbCfg.Host, dbCfg.Port, dbCfg.DBName))

	return sqlxDb
}
