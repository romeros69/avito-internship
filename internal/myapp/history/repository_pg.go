package history

import (
	"avito-internship/internal/myapp/models"
	"context"
)

type Repository interface {
	CreateHistory(context.Context, models.History) error
}
