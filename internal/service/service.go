package service

import (
	"exchanger_test/internal/models"
)

type ExchangerService struct {
}

func NewExchangerService() *ExchangerService {
	return &ExchangerService{}
}

func (e *ExchangerService) ExchangeAmount(ex models.Exchanger) (models.ResponseChanges, error) {
	var resp models.ResponseChanges

	combinations, err := findCombinations(ex.Banknotes, ex.Amount)
	if err != nil {
		return models.ResponseChanges{Exchanges: [][]int{}}, err
	}

	for _, combination := range combinations {
		resp.Exchanges = append(resp.Exchanges, combination)
	}

	return resp, nil
}

type ExchangerLogic interface {
	ExchangeAmount(ex models.Exchanger) (models.ResponseChanges, error)
}
