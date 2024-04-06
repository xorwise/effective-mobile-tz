package main

import (
	"log"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/xorwise/effective-mobile-tz/api"
	"github.com/xorwise/effective-mobile-tz/bootstrap"
)

func main() {
	env := bootstrap.NewEnv()
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: env.LogLevel,
	}))
	conn, err := bootstrap.NewDatabaseConnection(env)
	if err != nil {
		log.Fatal(err)
	}
	defer bootstrap.CloseDatabaseConnection(conn)

	err = bootstrap.CreateTables(conn)
	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()
	api.Setup(env, env.RequestTimeout*time.Second, conn, router, logger)

	log.Fatal(http.ListenAndServe(":8080", router))
}
