package models

import "strconv"

type ReportResult struct {
	ServiceName string `json:"service_name"`
	Proceeds    int64  `json:"proceeds"`
}

func (r *ReportResult) ConvertCSV() []string {
	return []string{r.ServiceName, strconv.FormatInt(r.Proceeds, 10)}
}
