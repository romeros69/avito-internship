package v1

type balanceResponseDTO struct {
	Balance string `json:"balance"`
}

type replenishmentRequestDTO struct {
	UserID string `json:"userID"`
	Value  string `json:"value"`
	Source string `json:"source"`
}
