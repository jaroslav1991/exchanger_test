package service

import (
	"exchanger_test/internal/models"
	"fmt"
)

type ExchangerLogic interface {
	ExchangeAmount(ex models.Exchanger) (models.Exchange, error)
}

type ExchangerService struct{}

func NewExchangerService() *ExchangerService {
	return &ExchangerService{}
}

func (e *ExchangerService) ExchangeAmount(ex models.Exchanger) (models.Exchange, error) {
	var resp models.Exchange

	combinations, err := findCombinations(ex.Banknotes, ex.Amount)
	if err != nil {
		return models.Exchange{Exchanges: [][]int{}}, fmt.Errorf("%v", err)
	}

	for _, combination := range combinations {
		resp.Exchanges = append(resp.Exchanges, combination)
	}

	return resp, nil
}

func findCombinations(banknotes []int, amount int) ([][]int, error) {
	var result [][]int
	var findCombinationRecursively func(int, int, []int)
	if amount <= 0 {
		return nil, fmt.Errorf("amount must be greater than zero: %v", errZeroOrNegativeAmount)
	}

	minVal := minBanknote(banknotes)

	if amount%minVal != 0 {
		return nil, fmt.Errorf("amount must be multiple of %d: %v", minVal, errIncorrectAmount)
	}

	findCombinationRecursively = func(amount int, start int, current []int) {
		if amount == 0 {
			temp := make([]int, len(current))
			copy(temp, current)
			result = append(result, temp)
			return
		}

		for i := start; i < len(banknotes); i++ {
			if banknotes[i] <= amount {
				current = append(current, banknotes[i])
				findCombinationRecursively(amount-banknotes[i], i, current)
				current = current[:len(current)-1]
			}
		}
	}

	findCombinationRecursively(amount, 0, []int{})
	return result, nil
}

func minBanknote(banknotes []int) int {
	minVal := banknotes[0]

	for i := 1; i < len(banknotes); i++ {
		if banknotes[i] < minVal {
			minVal = banknotes[i]
		}
	}

	return minVal
}
