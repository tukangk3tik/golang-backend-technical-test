package service

import (
	"fmt"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/dto/request"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/entity"
	"gitlab.com/tukangk3tik_/privyid-golang-test/pkg/repository"
)

type UserBalanceService interface {
	GetUserBalance(userId uint) (data map[string]any, err error)
	TopUpBalance(userId uint, amount uint, header request.HeaderDto) (map[string]any, error)
	Transfer(userId uint, to string, amount uint, header request.HeaderDto) (map[string]any, error)
}

type userBalanceService struct {
	balanceRepo repository.UserBalanceRepository
	userRepo    repository.UserRepository
}

func NewUserBalanceService(balanceRepo repository.UserBalanceRepository, userRepo repository.UserRepository) UserBalanceService {
	return &userBalanceService{balanceRepo: balanceRepo, userRepo: userRepo}
}

func (s *userBalanceService) GetUserBalance(userId uint) (map[string]any, error) {
	data := make(map[string]any, 0)

	res, err := s.balanceRepo.GetUserBalance(userId)
	if err == nil {
		data["balance"] = res.Balance
		data["balance_achieve"] = res.BalanceAchieve
	}
	return data, err
}

func (s *userBalanceService) TopUpBalance(userId uint, amount uint, header request.HeaderDto) (map[string]any, error) {
	data := make(map[string]any, 0)
	balance, err := s.balanceRepo.GetUserBalance(userId)
	if err != nil {
		return nil, err
	}

	err = s.balanceRepo.AddBalance(balance.ID, amount)
	if err != nil {
		panic(err)
	}

	balanceNow := balance.Balance + amount
	dataHistory := map[string]any{
		"userBalanceId": balance.ID,
		"balanceBefore": balance.Balance,
		"balanceAfter":  balanceNow,
		"type":          entity.DEBIT,
	}

	user, err := s.userRepo.GetUser(userId)
	header.Author = user.Username
	_ = s.balanceRepo.CreateBalanceHistory(header, dataHistory)

	data["top_up_amount"] = amount
	data["balance_now"] = balanceNow

	return data, nil
}

func (s *userBalanceService) Transfer(userId uint, to string, amount uint, header request.HeaderDto) (map[string]any, error) {
	data := make(map[string]any, 0)
	balance, err := s.balanceRepo.GetUserBalance(userId)
	if err != nil {
		return nil, err
	}

	if balance.Balance < amount {
		return nil, fmt.Errorf("insufficient balance")
	}

	err = s.balanceRepo.DeductBalance(balance.ID, amount)
	if err != nil {
		panic(err)
	}

	balanceNowSender := balance.Balance - amount
	receiver, err1 := s.userRepo.GetUserByUsername(to)
	rBalance, err2 := s.balanceRepo.GetUserBalance(receiver.ID)
	if err1 != nil || err2 != nil {
		panic(err1.Error() + err2.Error())
	}

	balanceNowReceiver := rBalance.Balance + amount
	err = s.balanceRepo.AddBalance(rBalance.ID, amount)
	if err != nil {
		// refund fail update balance
		_ = s.balanceRepo.AddBalance(balance.ID, amount)
		panic(err)
	}

	dataHistorySender := map[string]any{
		"userBalanceId": balance.ID,
		"balanceBefore": balance.Balance,
		"balanceAfter":  balanceNowSender,
		"type":          entity.KREDIT,
	}

	dataHistoryReceiver := map[string]any{
		"userBalanceId": rBalance.ID,
		"balanceBefore": rBalance.Balance,
		"balanceAfter":  balanceNowReceiver,
		"type":          entity.DEBIT,
	}

	user, err := s.userRepo.GetUser(userId)
	header.Author = user.Username
	_ = s.balanceRepo.CreateBalanceHistory(header, dataHistorySender, dataHistoryReceiver)

	data["to"] = receiver.Username
	data["amount_send"] = amount
	data["balance_now"] = balanceNowSender

	return data, nil
}
