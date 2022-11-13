package v1

type balanceResponseDTO struct {
	Rubles  string `json:"rubles"`
	Pennies string `json:"pennies"`
}

type replenishmentRequestDTO struct {
	UserID string `json:"userID"`
	Value  string `json:"value"`
	Source string `json:"source"`
}
