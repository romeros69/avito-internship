package history

import (
	"avito-internship/internal/myapp/models"
	"context"
)

type UseCase interface {
	CreateHistory(context.Context, models.History) error
}
