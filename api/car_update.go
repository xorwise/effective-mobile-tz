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
	"github.com/xorwise/effective-mobile-tz/repository"
)

type UpdateCarController struct {
	env     *bootstrap.Env
	timeout time.Duration
	conn    *pgx.Conn
	logger  *slog.Logger
}

func (cc *UpdateCarController) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(r.Context(), cc.timeout)
	defer cancel()

	regNum := r.PathValue("regNum")
	var updateRequest domain.UpdateCarRequest
	err := json.NewDecoder(r.Body).Decode(&updateRequest)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "validation error"})
		cc.logger.Debug("error: invalid request", "error", err)
		return
	}

	repo := repository.NewCarRepository(cc.conn, ctx, cc.logger)
	car, ok := repo.Update(regNum, updateRequest)
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "car not found"})
		cc.logger.Debug("error: car not found", "reg_num", regNum)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(car)
}
