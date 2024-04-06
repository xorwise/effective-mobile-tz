package api

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
	"strconv"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/xorwise/effective-mobile-tz/bootstrap"
	"github.com/xorwise/effective-mobile-tz/repository"
)

type ListCarsController struct {
	env     *bootstrap.Env
	timeout time.Duration
	conn    *pgx.Conn
	logger  *slog.Logger
}

func (cc *ListCarsController) Handle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	ctx, cancel := context.WithTimeout(r.Context(), cc.timeout)
	defer cancel()

	page := r.URL.Query().Get("page")
	if page == "" {
		page = "1"
	}
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "page must be integer"})
		cc.logger.Debug("error: page must be integer", "page", page)
		return
	}
	regNum := r.URL.Query().Get("reg-num")
	mark := r.URL.Query().Get("mark")
	model := r.URL.Query().Get("model")
	year := r.URL.Query().Get("year")
	if year == "" {
		year = "0"
	}
	yearInt, err := strconv.Atoi(year)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "year must be integer"})
		cc.logger.Debug("error: year must be integer", "year", year)
		return
	}
	ownerName := r.URL.Query().Get("owner-name")
	ownerSurname := r.URL.Query().Get("owner-surname")
	ownerPatronymic := r.URL.Query().Get("owner-patronymic")

	repo := repository.NewCarRepository(cc.conn, ctx, cc.logger)
	cars, err := repo.List(cc.env.ListLimit, pageInt, regNum, mark, model, yearInt, ownerName, ownerSurname, ownerPatronymic)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(map[string]string{"error": err.Error()})
		cc.logger.Error("error: list cars", "error", err)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(cars)
}
