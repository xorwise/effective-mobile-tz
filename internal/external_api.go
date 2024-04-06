package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/xorwise/effective-mobile-tz/bootstrap"
	"github.com/xorwise/effective-mobile-tz/domain"
)

func GetCarInfo(env *bootstrap.Env, regNum string, ctx context.Context) (*domain.Car, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/info?regNum=%s", env.APIPath, regNum), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	var car domain.ExternalResponse
	err = json.NewDecoder(resp.Body).Decode(&car)
	if err != nil {
		return nil, err
	}
	return &domain.Car{
		RegNum:          car.RegNum,
		Mark:            car.Mark,
		Model:           car.Model,
		Year:            car.Year,
		OwnerName:       car.Owner.Name,
		OwnerSurname:    car.Owner.Surname,
		OwnerPatronymic: car.Owner.Patronymic,
	}, nil
}
