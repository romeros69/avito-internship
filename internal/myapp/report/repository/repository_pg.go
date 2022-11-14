package repository

import (
	"avito-internship/internal/myapp/models"
	"avito-internship/internal/myapp/report"
	"avito-internship/internal/pkg/postgres"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type ReportRepo struct {
	pg *postgres.Postgres
}

func NewReportRepo(pg *postgres.Postgres) *ReportRepo {
	return &ReportRepo{pg: pg}
}

var _ report.Repository = (*ReportRepo)(nil)

func (r *ReportRepo) CreateReport(ctx context.Context, report models.Report) (uuid.UUID, error) {
	query := `insert into report (id, service_id, value, date) values ($1, $2, $3, $4) returning id`

	rows, err := r.pg.Pool.Query(
		ctx,
		query,
		report.ID,
		report.ServiceID,
		report.Value,
		report.Date,
	)
	if err != nil {
		return uuid.Nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var reportID uuid.UUID
	if !rows.Next() {
		return uuid.Nil, fmt.Errorf("error in creating report")
	}
	err = rows.Scan(&reportID)
	if err != nil {
		return uuid.Nil, fmt.Errorf("error in parsing id report of create report: %w", err)
	}
	return reportID, nil
}
