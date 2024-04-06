package repository

import (
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/jackc/pgx/v5"
	"github.com/xorwise/effective-mobile-tz/domain"
)

type carRepository struct {
	conn   *pgx.Conn
	ctx    context.Context
	logger *slog.Logger
}

func NewCarRepository(conn *pgx.Conn, ctx context.Context, logger *slog.Logger) *carRepository {
	return &carRepository{conn, ctx, logger}
}

func buildListQuery(args map[string]any, limit int, offset int) (string, []any) {
	query := "SELECT * FROM cars WHERE 1=1"
	new_args := []any{}
	for k, v := range args {
		query += fmt.Sprintf(" AND %s = $%d", k, len(new_args)+1)
		new_args = append(new_args, v)
	}
	if limit > 0 {
		query += fmt.Sprintf(" LIMIT %d", limit)
	}
	if offset > 0 {
		query += fmt.Sprintf(" OFFSET %d", offset)
	}
	return query, new_args
}

func (cr *carRepository) List(limit int, page int, regNum string, mark string, model string, year int, ownerName string, ownerSurname string, ownerPatronymic string) ([]domain.Car, error) {
	args := make(map[string]any)

	if regNum != "" {
		args["reg_num"] = regNum
	}
	if mark != "" {
		args["mark"] = mark
	}
	if model != "" {
		args["model"] = model
	}
	if year != 0 {
		args["year"] = year
	}
	if ownerName != "" {
		args["owner_name"] = ownerName
	}
	if ownerSurname != "" {
		args["owner_surname"] = ownerSurname
	}
	if ownerPatronymic != "" {
		args["owner_patronymic"] = ownerPatronymic
	}
	query, builtArgs := buildListQuery(args, limit, limit*(page-1))
	rows, err := cr.conn.Query(cr.ctx, query, builtArgs...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cars := []domain.Car{}
	for rows.Next() {
		var car domain.Car
		err = rows.Scan(&car.RegNum, &car.Mark, &car.Model, &car.Year, &car.OwnerName, &car.OwnerSurname, &car.OwnerPatronymic)
		if err != nil {
			return nil, err
		}
		cars = append(cars, car)
	}
	cr.logger.Debug("executed query to database", slog.String("query", query))
	return cars, nil
}

func (cr *carRepository) Delete(regNum string) bool {
	result, err := cr.conn.Exec(cr.ctx, "DELETE FROM cars WHERE reg_num = $1", regNum)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return false
	}

	return true
}

func (cr *carRepository) Update(regNum string, updateRequest domain.UpdateCarRequest) (*domain.Car, bool) {
	query := "UPDATE cars SET mark = $1, model = $2, year = $3, owner_name = $4, owner_surname = $5, owner_patronymic = $6 WHERE reg_num = $7"

	result, err := cr.conn.Exec(
		cr.ctx,
		query,
		updateRequest.Mark,
		updateRequest.Model,
		updateRequest.Year,
		updateRequest.OwnerName,
		updateRequest.OwnerSurname,
		updateRequest.OwnerPatronymic,
		regNum,
	)
	if err != nil {
		log.Fatal(err)
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, false
	}

	return &domain.Car{
		RegNum:          regNum,
		Mark:            updateRequest.Mark,
		Model:           updateRequest.Model,
		Year:            updateRequest.Year,
		OwnerName:       updateRequest.OwnerName,
		OwnerSurname:    updateRequest.OwnerSurname,
		OwnerPatronymic: updateRequest.OwnerPatronymic}, true
}

func (cr *carRepository) Create(car *domain.Car) error {

	_, err := cr.conn.Exec(
		cr.ctx,
		"INSERT INTO cars (reg_num, mark, model, year, owner_name, owner_surname, owner_patronymic) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		car.RegNum,
		car.Mark,
		car.Model,
		car.Year,
		car.OwnerName,
		car.OwnerSurname,
		car.OwnerPatronymic,
	)
	return err
}
