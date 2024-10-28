package model

type BalanceStatus uint16

const (
	OPEN BalanceStatus = iota
	PROGRESS
	PENDENT_REOPEN
	CLOSED
)
