package v1

type reserveBalanceDTO struct {
	UserID    string `json:"userID"`
	ServiceID string `json:"serviceID"`
	OrderID   string `json:"orderID"`
	Value     string `json:"value"`
}
