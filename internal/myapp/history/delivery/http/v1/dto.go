package v1

type centralHistoryResponseDTO struct {
	History   []historyResponseDTO `json:"history"`
	HasMore   bool                 `json:"hasMore"`
	TotalPage int                  `json:"totalPage"`
}

type historyResponseDTO struct {
	Date                string `json:"date"`
	TypeOfTransaction   string `json:"typeOfTransaction"`
	SourceReplenishment string `json:"sourceReplenishment,omitempty"`
	Value               string `json:"value"`
	ServiceName         string `json:"serviceName,omitempty"`
	OrderID             string `json:"orderID,omitempty"`
}
