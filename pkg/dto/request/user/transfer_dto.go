package user

type TransferDto struct {
	To     string `json:"to" binding:"required"` // to = username
	Amount int    `json:"amount" binding:"required"`
}
