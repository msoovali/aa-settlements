package domain

type Usecase interface {
	CreateNextMonthSettlements() error
}
