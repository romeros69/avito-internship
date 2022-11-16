package models

type HistoryTransfer struct {
	Histories    []History
	ServiceNames []string
	Count        int64
}
