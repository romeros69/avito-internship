package repository

import (
	"avito-internship/internal/myapp/service"
	"avito-internship/internal/pkg/postgres"
	"context"
	"fmt"
	"github.com/google/uuid"
)

type ServiceRepo struct {
	pg *postgres.Postgres
}

func NewServiceRepo(pg *postgres.Postgres) *ServiceRepo {
	return &ServiceRepo{
		pg: pg,
	}
}

var _ service.Repository = (*ServiceRepo)(nil)

func (s *ServiceRepo) GetServiceNameByID(ctx context.Context, serviceID uuid.UUID) (string, error) {
	query := `select tittle from service where id=$1`

	rows, err := s.pg.Pool.Query(ctx, query, serviceID)
	if err != nil {
		return "", fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	var nameService string
	err = rows.Scan(&nameService)
	if err != nil {
		return "", fmt.Errorf("error parsing tittle of service: %w", err)
	}
	return nameService, nil
}

func (s *ServiceRepo) ServiceExistsByID(ctx context.Context, serviceID uuid.UUID) (bool, error) {
	query := `select * from service where id = $1`

	rows, err := s.pg.Pool.Query(ctx, query, serviceID)
	if err != nil {
		return false, fmt.Errorf("cannot execute query: %w", err)
	}
	defer rows.Close()
	if rows.Next() {
		return true, nil
	}
	return false, nil
}
