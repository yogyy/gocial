package main

import (
	"github.com/gocial/internal/auth"
	"github.com/gocial/internal/db"
	"github.com/gocial/internal/env"
	"github.com/gocial/internal/store"
	"go.uber.org/zap"
)

func main() {
	config := config{
		addr: env.GetString("ADDR", ":8080"),
		db: dbConfig{
			addr:         env.GetString("DB_URL", "postgres://postgres:root@localhost:5432/gocial?sslmode=disable"),
			maxOpenConns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_OPEN_CONNS", "10m"),
		},
	}

	logger := zap.Must(zap.NewProduction()).Sugar()
	defer logger.Sync()

	db, err := db.New(
		config.db.addr,
		config.db.maxOpenConns,
		config.db.maxIdleConns,
		config.db.maxIdleTime,
	)

	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()
	logger.Info("database connection pool established")

	store := store.NewStorage(db)
	authenticator := auth.NewJWTAuthenticator(
		env.GetString("JWT_SECRET", "constantine"),
		"gocial",
		"gocial",
	)

	app := &application{config, store, logger, authenticator}
	mux := app.mount()

	logger.Fatal(app.run(mux))
}
