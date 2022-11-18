package repository

import (
	"avito-internship/internal/myapp/models"
	"avito-internship/internal/myapp/report"
	"avito-internship/internal/pkg/postgres"
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"time"
)

type ReportRepo struct {
	pg *postgres.Postgres
}

func NewReportRepo(pg *postgres.Postgres) *ReportRepo {
	return &ReportRepo{
		pg: pg,
	}
}

var _ report.Repository = (*ReportRepo)(nil)

func (r *ReportRepo) GetReport(ctx context.Context, startPeriod time.Time) ([]models.ReportResult, error) {
	var end time.Time
	end = end.AddDate(startPeriod.Year()-1, int(startPeriod.Month())-1, 29)

	query := `select s.tittle, sum(r.value) from report r inner join service s on s.id = r.service_id where r.date > $1 and r.date < $2 group by s.tittle`

	rows, err := r.pg.Pool.Query(ctx, query, startPeriod, end)
	if err != nil {
		log.Printf("cannot execute query: %w", err)
		return nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var reportList []models.ReportResult
	for rows.Next() {
		var reportResult models.ReportResult
		err = rows.Scan(
			&reportResult.ServiceName,
			&reportResult.Proceeds,
		)
		if err != nil {
			log.Printf("error parsing report result: %w ", err)
			return nil, fmt.Errorf("error parsing report result: %w ", err)
		}
		reportList = append(reportList, reportResult)
	}
	return reportList, nil
}

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
		log.Printf("cannot execute query: %w", err)
		return uuid.Nil, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var reportID uuid.UUID
	if !rows.Next() {
		log.Printf("error in creating report")
		return uuid.Nil, fmt.Errorf("error in creating report")
	}
	err = rows.Scan(&reportID)
	if err != nil {
		log.Printf("error in parsing id report of create report: %w", err)
		return uuid.Nil, fmt.Errorf("error in parsing id report of create report: %w", err)
	}
	return reportID, nil
}
