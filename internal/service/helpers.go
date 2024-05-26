package service

import (
	"github.com/sirupsen/logrus"
)

func findCombinations(banknotes []int, amount int) ([][]int, error) {
	var result [][]int
	var findCombinationRecursively func(int, int, []int)
	if amount <= 0 {
		logrus.Error("amount must be greater than zero")
		return nil, errZeroOrNegativeAmount
	}

	minVal := minBanknote(banknotes)

	if amount%minVal != 0 {
		logrus.Error("amount must be a multiple of ", minVal)
		return nil, errIncorrectAmount
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
