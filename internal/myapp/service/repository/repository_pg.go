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
	return &ServiceRepo{pg: pg}
}

var _ service.Repository = (*ServiceRepo)(nil)

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