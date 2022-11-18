package usecase

import (
	"avito-internship/internal/myapp/service/mock"
	"context"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestServiceUseCase_ServiceExistsByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockServiceRepo := mock.NewMockServiceRepository(ctrl)
	serviceUC := NewServiceUseCase(mockServiceRepo)

	ctx := context.Background()
	serviceID := uuid.New()

	mockServiceRepo.EXPECT().ServiceExistsByID(gomock.Any(), gomock.Eq(serviceID)).Return(true, nil)

	exists, err := serviceUC.ServiceExistsByID(ctx, serviceID)
	require.NoError(t, err)
	require.Nil(t, err)
	require.Equal(t, true, exists)
}

func TestServiceUseCase_GetServiceNameByID(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockServiceRepo := mock.NewMockServiceRepository(ctrl)
	serviceUC := NewServiceUseCase(mockServiceRepo)

	ctx := context.Background()
	serviceID := uuid.New()

	mockServiceRepo.EXPECT().GetServiceNameByID(gomock.Any(), gomock.Eq(serviceID)).Return("cleaning", nil)

	serviceName, err := serviceUC.GetServiceNameByID(ctx, serviceID)
	require.NoError(t, err)
	require.Nil(t, err)
	require.Equal(t, "cleaning", serviceName)
}
