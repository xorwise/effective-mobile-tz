package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/xorwise/effective-mobile-tz/bootstrap"
	"github.com/xorwise/effective-mobile-tz/domain"
	externalapi "github.com/xorwise/effective-mobile-tz/internal"
	"github.com/xorwise/effective-mobile-tz/repository"
)

type CreateCarController struct {
	env     *bootstrap.Env
	timeout time.Duration
	conn    *pgx.Conn
	logger  *slog.Logger
}

func (cc *CreateCarController) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(r.Context(), cc.timeout)
	defer cancel()

	var request domain.CreateCarRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "validation error"})
		cc.logger.Debug("error: invalid request", "error", err)
		return
	}

	repo := repository.NewCarRepository(cc.conn, ctx, cc.logger)
	var cars []*domain.Car
	for _, RegNum := range request.RegNums {
		if RegNum == "" {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "validation error"})
			cc.logger.Debug("error: invalid request", "error", err)
			return
		}
		car, err := externalapi.GetCarInfo(cc.env, RegNum, ctx)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "validation error"})
			cc.logger.Debug("error: invalid request", "error", err)
			return
		}
		err = repo.Create(car)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(map[string]string{"error": "could not create car"})
			cc.logger.Debug("error: invalid request", "error", err)
			return
		}
		cars = append(cars, car)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}
