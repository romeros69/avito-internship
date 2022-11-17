package usecase

import (
	"avito-internship/internal/myapp/models"
	"avito-internship/internal/myapp/report"
	"context"
	"encoding/csv"
	"fmt"
	"github.com/google/uuid"
	"os"
	"time"
)

type ReportUseCase struct {
	repo report.Repository
}

func NewReportUseCase(repo report.Repository) *ReportUseCase {
	return &ReportUseCase{
		repo: repo,
	}
}

var _ report.UseCase = (*ReportUseCase)(nil)

func (r *ReportUseCase) createCSVFile(data [][]string) error {
	_, err := os.Stat("./web/static/report.csv")
	if err == nil {
		err := os.Remove("./web/static/report.csv")
		if err != nil {
			return fmt.Errorf("error in deleting csv file: %w", err)
		}
	}
	f, err := os.Create("./web/static/report.csv")
	defer f.Close()
	if err != nil {
		return fmt.Errorf("error in creating file: %w", err)
	}
	w := csv.NewWriter(f)
	err = w.WriteAll(data)
	if err != nil {
		return fmt.Errorf("error in writting data to csv report: %w", err)
	}
	return nil
}

func (r *ReportUseCase) GetReport(ctx context.Context, year, month int) error {
	var startPeriod time.Time
	startPeriod = startPeriod.AddDate(year, month, 0)
	reportList, err := r.repo.GetReport(ctx, startPeriod)
	if err != nil {
		return err
	}
	var data [][]string
	data = append(data, []string{"service_name", "proceeds"})
	for _, v := range reportList {
		data = append(data, v.ConvertCSV())
	}
	err = r.createCSVFile(data)
	if err != nil {
		return err
	}
	return nil
}

func (r *ReportUseCase) CreateReport(ctx context.Context, report models.Report) (uuid.UUID, error) {
	return r.repo.CreateReport(ctx, report)
}
