package model

type BalanceStatus int

const (
	OPEN BalanceStatus = iota
	PROGRESS
	PENDENT_REOPEN
	CLOSED
)
