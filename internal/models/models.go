package models

type Exchanger struct {
	Amount    int   `json:"amount"`
	Banknotes []int `json:"banknotes"`
}

type Exchange struct {
	Exchanges [][]int `json:"exchanges"`
}
