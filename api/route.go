package api

import (
	"log/slog"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/xorwise/effective-mobile-tz/bootstrap"
)

func Setup(env *bootstrap.Env, timeout time.Duration, conn *pgx.Conn, router *http.ServeMux, logger *slog.Logger) {
	m := LoggingMiddleware{Logger: logger}

	listCarsController := ListCarsController{env, timeout, conn, logger}
	router.Handle("GET /cars", m.WithLogging(listCarsController.Handle))

	deleteCarController := DeleteCarController{env, timeout, conn, logger}
	router.Handle("DELETE /cars/{regNum}", m.WithLogging(deleteCarController.Handle))

	updateCarController := UpdateCarController{env, timeout, conn, logger}
	router.Handle("PUT /cars/{regNum}", m.WithLogging(updateCarController.Handle))

	createCarController := CreateCarController{env, timeout, conn, logger}
	router.Handle("POST /cars", m.WithLogging(createCarController.Handle))
}
