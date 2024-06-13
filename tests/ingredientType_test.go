package tests

import (
	"context"
	"errors"
	"fmt"
	"github.com/Mx1q/ppo_services/domain"
	"github.com/Mx1q/ppo_services/services"
	"github.com/Mx1q/ppo_services/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestIngredientTypeService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ingredientTypeRepo := mocks.NewMockIIngredientTypeRepository(ctrl)
	logger := mocks.NewMockILogger(ctrl)
	logger.EXPECT().
		Infof(gomock.Any()).
		AnyTimes()
	logger.EXPECT().
		Infof(gomock.Any(), gomock.Any()).
		AnyTimes()
	logger.EXPECT().
		Warnf(gomock.Any(), gomock.Any()).
		AnyTimes()
	logger.EXPECT().
		Errorf(gomock.Any(), gomock.Any()).
		AnyTimes()
	svc := services.NewIngredientTypeService(ingredientTypeRepo, logger)

	ingredientTypeId := uuid.New()

	tests := []struct {
		name           string
		ingredientType *domain.IngredientType
		beforeTest     func(ingredientTypeRepo mocks.MockIIngredientTypeRepository)
		wantErr        bool
		errStr         error
	}{
		{
			name: "успешное создание",
			ingredientType: &domain.IngredientType{
				ID:          ingredientTypeId,
				Name:        "meat",
				Description: "",
			},
			beforeTest: func(ingredientTypeRepo mocks.MockIIngredientTypeRepository) {
				ingredientTypeRepo.EXPECT().
					Create(context.Background(), &domain.IngredientType{
						ID:          ingredientTypeId,
						Name:        "meat",
						Description: "",
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное создание
		{
			name: "пустое название",
			ingredientType: &domain.IngredientType{
				ID:          ingredientTypeId,
				Name:        "",
				Description: "",
			},
			wantErr: true,
			errStr:  errors.New("creating ingredient type: empty name"),
		}, // пустое название
		{
			name: "ошибка выполнения запроса в репозитории",
			ingredientType: &domain.IngredientType{
				ID:          ingredientTypeId,
				Name:        "meat",
				Description: "",
			},
			beforeTest: func(ingredientTypeRepo mocks.MockIIngredientTypeRepository) {
				ingredientTypeRepo.EXPECT().
					Create(context.Background(), &domain.IngredientType{
						ID:          ingredientTypeId,
						Name:        "meat",
						Description: "",
					}).
					Return(fmt.Errorf("creating err"))
			},
			wantErr: true,
			errStr:  errors.New("creating ingredient type: creating err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*ingredientTypeRepo)
			}

			err := svc.Create(context.Background(), tt.ingredientType)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestIngredientTypeService_DeleteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ingredientTypeRepo := mocks.NewMockIIngredientTypeRepository(ctrl)
	logger := mocks.NewMockILogger(ctrl)
	logger.EXPECT().
		Infof(gomock.Any()).
		AnyTimes()
	logger.EXPECT().
		Infof(gomock.Any(), gomock.Any()).
		AnyTimes()
	logger.EXPECT().
		Warnf(gomock.Any(), gomock.Any()).
		AnyTimes()
	logger.EXPECT().
		Errorf(gomock.Any(), gomock.Any()).
		AnyTimes()
	svc := services.NewIngredientTypeService(ingredientTypeRepo, logger)

	ingredientTypeId := uuid.New()

	tests := []struct {
		name         string
		ingredientId uuid.UUID
		beforeTest   func(ingredientTypeRepo mocks.MockIIngredientTypeRepository)
		wantErr      bool
		errStr       error
	}{
		{
			name:         "успешное удаление",
			ingredientId: ingredientTypeId,
			beforeTest: func(ingredientTypeRepo mocks.MockIIngredientTypeRepository) {
				ingredientTypeRepo.EXPECT().
					DeleteById(context.Background(), ingredientTypeId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное удаление
		{
			name:         "ошибка выполнения запроса в репозитории",
			ingredientId: ingredientTypeId,
			beforeTest: func(ingredientTypeRepo mocks.MockIIngredientTypeRepository) {
				ingredientTypeRepo.EXPECT().
					DeleteById(context.Background(), ingredientTypeId).
					Return(fmt.Errorf("deleting err"))
			},
			wantErr: true,
			errStr:  errors.New("deleting ingredient type by id: deleting err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*ingredientTypeRepo)
			}

			err := svc.DeleteById(context.Background(), tt.ingredientId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestIngredientTypeService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ingredientTypeRepo := mocks.NewMockIIngredientTypeRepository(ctrl)
	logger := mocks.NewMockILogger(ctrl)
	logger.EXPECT().
		Infof(gomock.Any()).
		AnyTimes()
	logger.EXPECT().
		Infof(gomock.Any(), gomock.Any()).
		AnyTimes()
	logger.EXPECT().
		Warnf(gomock.Any(), gomock.Any()).
		AnyTimes()
	logger.EXPECT().
		Errorf(gomock.Any(), gomock.Any()).
		AnyTimes()
	svc := services.NewIngredientTypeService(ingredientTypeRepo, logger)

	tests := []struct {
		name       string
		beforeTest func(ingredientTypeRepo mocks.MockIIngredientTypeRepository)
		expected   []*domain.IngredientType
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное получение",
			beforeTest: func(ingredientTypeRepo mocks.MockIIngredientTypeRepository) {
				ingredientTypeRepo.EXPECT().
					GetAll(context.Background()).
					Return([]*domain.IngredientType{
						{
							ID:          uuid.UUID{1},
							Name:        "first",
							Description: "",
						},
						{
							ID:          uuid.UUID{2},
							Name:        "second",
							Description: "",
						},
						{
							ID:          uuid.UUID{3},
							Name:        "third",
							Description: "",
						},
						{
							ID:          uuid.UUID{4},
							Name:        "fourth",
							Description: "",
						},
					}, nil)
			},
			expected: []*domain.IngredientType{
				{
					ID:          uuid.UUID{1},
					Name:        "first",
					Description: "",
				},
				{
					ID:          uuid.UUID{2},
					Name:        "second",
					Description: "",
				},
				{
					ID:          uuid.UUID{3},
					Name:        "third",
					Description: "",
				},
				{
					ID:          uuid.UUID{4},
					Name:        "fourth",
					Description: "",
				},
			},
			wantErr: false,
		}, // успешное получение
		{
			name: "ошибка выполнения запроса в репозитории",
			beforeTest: func(ingredientTypeRepo mocks.MockIIngredientTypeRepository) {
				ingredientTypeRepo.EXPECT().
					GetAll(context.Background()).
					Return(nil, fmt.Errorf("getting all err"))
			},
			wantErr: true,
			errStr:  errors.New("getting all ingredient types: getting all err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*ingredientTypeRepo)
			}

			types, err := svc.GetAll(context.Background())
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, types)
			}
		})
	}
}

func TestIngredientTypeService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ingredientTypeRepo := mocks.NewMockIIngredientTypeRepository(ctrl)
	logger := mocks.NewMockILogger(ctrl)
	logger.EXPECT().
		Infof(gomock.Any()).
		AnyTimes()
	logger.EXPECT().
		Infof(gomock.Any(), gomock.Any()).
		AnyTimes()
	logger.EXPECT().
		Warnf(gomock.Any(), gomock.Any()).
		AnyTimes()
	logger.EXPECT().
		Errorf(gomock.Any(), gomock.Any()).
		AnyTimes()
	svc := services.NewIngredientTypeService(ingredientTypeRepo, logger)

	ingredientTypeId := uuid.New()

	tests := []struct {
		name       string
		typeId     uuid.UUID
		beforeTest func(ingredientTypeRepo mocks.MockIIngredientTypeRepository)
		expected   *domain.IngredientType
		wantErr    bool
		errStr     error
	}{
		{
			name:   "успешное получение",
			typeId: ingredientTypeId,
			beforeTest: func(ingredientTypeRepo mocks.MockIIngredientTypeRepository) {
				ingredientTypeRepo.EXPECT().
					GetById(context.Background(), ingredientTypeId).
					Return(&domain.IngredientType{
						ID:          ingredientTypeId,
						Name:        "type",
						Description: "",
					}, nil)
			},
			expected: &domain.IngredientType{
				ID:          ingredientTypeId,
				Name:        "type",
				Description: "",
			},
			wantErr: false,
		}, // успешное получение
		{
			name:   "ошибка выполнения запроса в репозитории",
			typeId: ingredientTypeId,
			beforeTest: func(ingredientTypeRepo mocks.MockIIngredientTypeRepository) {
				ingredientTypeRepo.EXPECT().
					GetById(context.Background(), ingredientTypeId).
					Return(nil, fmt.Errorf("getting type err"))
			},
			wantErr: true,
			errStr:  errors.New("getting ingredient type by id: getting type err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*ingredientTypeRepo)
			}

			types, err := svc.GetById(context.Background(), tt.typeId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, types)
			}
		})
	}
}

func TestIngredientTypeService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ingredientTypeRepo := mocks.NewMockIIngredientTypeRepository(ctrl)
	logger := mocks.NewMockILogger(ctrl)
	logger.EXPECT().
		Infof(gomock.Any()).
		AnyTimes()
	logger.EXPECT().
		Infof(gomock.Any(), gomock.Any()).
		AnyTimes()
	logger.EXPECT().
		Warnf(gomock.Any(), gomock.Any()).
		AnyTimes()
	logger.EXPECT().
		Errorf(gomock.Any(), gomock.Any()).
		AnyTimes()
	svc := services.NewIngredientTypeService(ingredientTypeRepo, logger)

	ingredientTypeId := uuid.New()

	tests := []struct {
		name           string
		ingredientType *domain.IngredientType
		beforeTest     func(ingredientTypeRepo mocks.MockIIngredientTypeRepository)
		wantErr        bool
		errStr         error
	}{
		{
			name: "успешное создание",
			ingredientType: &domain.IngredientType{
				ID:          ingredientTypeId,
				Name:        "meat",
				Description: "",
			},
			beforeTest: func(ingredientTypeRepo mocks.MockIIngredientTypeRepository) {
				ingredientTypeRepo.EXPECT().
					Update(context.Background(), &domain.IngredientType{
						ID:          ingredientTypeId,
						Name:        "meat",
						Description: "",
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное обновление
		{
			name: "пустое название",
			ingredientType: &domain.IngredientType{
				ID:          ingredientTypeId,
				Name:        "",
				Description: "",
			},
			wantErr: true,
			errStr:  errors.New("updating ingredient type: empty name"),
		}, // пустое название
		{
			name: "ошибка выполнения запроса в репозитории",
			ingredientType: &domain.IngredientType{
				ID:          ingredientTypeId,
				Name:        "meat",
				Description: "",
			},
			beforeTest: func(ingredientTypeRepo mocks.MockIIngredientTypeRepository) {
				ingredientTypeRepo.EXPECT().
					Update(context.Background(), &domain.IngredientType{
						ID:          ingredientTypeId,
						Name:        "meat",
						Description: "",
					}).
					Return(fmt.Errorf("updating err"))
			},
			wantErr: true,
			errStr:  errors.New("updating ingredient type: updating err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*ingredientTypeRepo)
			}

			err := svc.Update(context.Background(), tt.ingredientType)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
