package user

type TopUpBalanceDto struct {
	Amount int `json:"amount" binding:"required"`
}
