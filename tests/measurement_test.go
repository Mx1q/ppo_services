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

func TestMeasurementService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	measurementRepo := mocks.NewMockIMeasurementRepository(ctrl)
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
	svc := services.NewMeasurementService(measurementRepo, logger)

	measurementId := uuid.New()

	tests := []struct {
		name        string
		measurement *domain.Measurement
		beforeTest  func(measurementRepo mocks.MockIMeasurementRepository)
		wantErr     bool
		errStr      error
	}{
		{
			name: "успешное создание",
			measurement: &domain.Measurement{
				ID:    measurementId,
				Name:  "spoon",
				Grams: 1,
			},
			beforeTest: func(measurementRepo mocks.MockIMeasurementRepository) {
				measurementRepo.EXPECT().
					Create(context.Background(), &domain.Measurement{
						ID:    measurementId,
						Name:  "spoon",
						Grams: 1,
					}).Return(nil)
			},
			wantErr: false,
		}, // успешное создание
		{
			name: "пустое название",
			measurement: &domain.Measurement{
				ID:    measurementId,
				Name:  "",
				Grams: 1,
			},
			wantErr: true,
			errStr:  errors.New("creating measurement unit: empty name"),
		}, // пустое название
		{
			name: "количество граммов - 0",
			measurement: &domain.Measurement{
				ID:    measurementId,
				Name:  "spoon",
				Grams: 0,
			},
			wantErr: true,
			errStr:  errors.New("creating measurement unit: negative or zero grams count"),
		}, // количество граммов - 0
		{
			name: "количество граммов <0",
			measurement: &domain.Measurement{
				ID:    measurementId,
				Name:  "spoon",
				Grams: 0,
			},
			wantErr: true,
			errStr:  errors.New("creating measurement unit: negative or zero grams count"),
		}, // количество граммов <0
		{
			name: "ошибка выполнения запроса в репозитории",
			measurement: &domain.Measurement{
				ID:    measurementId,
				Name:  "spoon",
				Grams: 1,
			},
			beforeTest: func(measurementRepo mocks.MockIMeasurementRepository) {
				measurementRepo.EXPECT().
					Create(context.Background(), &domain.Measurement{
						ID:    measurementId,
						Name:  "spoon",
						Grams: 1,
					}).Return(fmt.Errorf("creating measurement unit err"))
			},
			wantErr: true,
			errStr:  errors.New("creating measurement unit: creating measurement unit err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*measurementRepo)
			}

			err := svc.Create(context.Background(), tt.measurement)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestMeasurementService_DeleteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	measurementRepo := mocks.NewMockIMeasurementRepository(ctrl)
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
	svc := services.NewMeasurementService(measurementRepo, logger)

	measurementId := uuid.New()

	tests := []struct {
		name          string
		measurementId uuid.UUID
		beforeTest    func(measurementRepo mocks.MockIMeasurementRepository)
		expected      *domain.Measurement
		wantErr       bool
		errStr        error
	}{
		{
			name:          "успешное получение",
			measurementId: measurementId,
			beforeTest: func(measurementRepo mocks.MockIMeasurementRepository) {
				measurementRepo.EXPECT().
					DeleteById(context.Background(), measurementId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное удаление
		{
			name:          "ошибка выполнения запроса в репозитории",
			measurementId: measurementId,
			beforeTest: func(measurementRepo mocks.MockIMeasurementRepository) {
				measurementRepo.EXPECT().
					DeleteById(context.Background(), measurementId).
					Return(fmt.Errorf("getting measurement unit err"))
			},
			wantErr: true,
			errStr:  errors.New("deleting measurement unit by id: getting measurement unit err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*measurementRepo)
			}

			err := svc.DeleteById(context.Background(), tt.measurementId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestMeasurementService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	measurementRepo := mocks.NewMockIMeasurementRepository(ctrl)
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
	svc := services.NewMeasurementService(measurementRepo, logger)

	tests := []struct {
		name       string
		beforeTest func(measurementRepo mocks.MockIMeasurementRepository)
		expected   []*domain.Measurement
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное получение",
			beforeTest: func(measurementRepo mocks.MockIMeasurementRepository) {
				measurementRepo.EXPECT().
					GetAll(context.Background()).
					Return([]*domain.Measurement{
						{
							ID:    uuid.UUID{1},
							Name:  "spoon",
							Grams: 1,
						},
						{
							ID:    uuid.UUID{2},
							Name:  "gram",
							Grams: 1,
						},
						{
							ID:    uuid.UUID{3},
							Name:  "count",
							Grams: 1,
						},
						{
							ID:    uuid.UUID{4},
							Name:  "pinch",
							Grams: 1,
						},
					}, nil)
			},
			expected: []*domain.Measurement{
				{
					ID:    uuid.UUID{1},
					Name:  "spoon",
					Grams: 1,
				},
				{
					ID:    uuid.UUID{2},
					Name:  "gram",
					Grams: 1,
				},
				{
					ID:    uuid.UUID{3},
					Name:  "count",
					Grams: 1,
				},
				{
					ID:    uuid.UUID{4},
					Name:  "pinch",
					Grams: 1,
				},
			},
			wantErr: false,
		}, // успешное получение
		{
			name: "ошибка выполнения запроса в репозитории",
			beforeTest: func(measurementRepo mocks.MockIMeasurementRepository) {
				measurementRepo.EXPECT().
					GetAll(context.Background()).
					Return(nil, fmt.Errorf("getting measurements err"))
			},
			wantErr: true,
			errStr:  errors.New("getting all measurement units: getting measurements err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*measurementRepo)
			}

			measurements, err := svc.GetAll(context.Background())
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, measurements)
			}
		})
	}
}

func TestMeasurementService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	measurementRepo := mocks.NewMockIMeasurementRepository(ctrl)
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
	svc := services.NewMeasurementService(measurementRepo, logger)

	measurementId := uuid.New()

	tests := []struct {
		name          string
		measurementId uuid.UUID
		beforeTest    func(measurementRepo mocks.MockIMeasurementRepository)
		expected      *domain.Measurement
		wantErr       bool
		errStr        error
	}{
		{
			name:          "успешное получение",
			measurementId: measurementId,
			beforeTest: func(measurementRepo mocks.MockIMeasurementRepository) {
				measurementRepo.EXPECT().
					GetById(context.Background(), measurementId).
					Return(&domain.Measurement{
						ID:    measurementId,
						Name:  "spoon",
						Grams: 1,
					}, nil)
			},
			expected: &domain.Measurement{
				ID:    measurementId,
				Name:  "spoon",
				Grams: 1,
			},
			wantErr: false,
		}, // успешное получение
		{
			name:          "ошибка выполнения запроса в репозитории",
			measurementId: measurementId,
			beforeTest: func(measurementRepo mocks.MockIMeasurementRepository) {
				measurementRepo.EXPECT().
					GetById(context.Background(), measurementId).
					Return(nil, fmt.Errorf("getting measurement err"))
			},
			wantErr: true,
			errStr:  errors.New("getting measurement unit by id: getting measurement err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*measurementRepo)
			}

			measurements, err := svc.GetById(context.Background(), tt.measurementId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, measurements)
			}
		})
	}
}

func TestMeasurementService_GetByRecipeId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	measurementRepo := mocks.NewMockIMeasurementRepository(ctrl)
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
	svc := services.NewMeasurementService(measurementRepo, logger)

	recipeId := uuid.New()
	ingredientId := uuid.New()

	tests := []struct {
		name                string
		recipeId            uuid.UUID
		ingredientId        uuid.UUID
		beforeTest          func(measurementRepo mocks.MockIMeasurementRepository)
		expectedMeasurement *domain.Measurement
		expectedCount       int
		wantErr             bool
		errStr              error
	}{
		{
			name:         "успешное получение",
			recipeId:     recipeId,
			ingredientId: ingredientId,
			beforeTest: func(measurementRepo mocks.MockIMeasurementRepository) {
				measurementRepo.EXPECT().
					GetByRecipeId(context.Background(), ingredientId, recipeId).
					Return(&domain.Measurement{
						ID:    uuid.UUID{1},
						Name:  "spoon",
						Grams: 1,
					}, 1, nil)
			},
			expectedMeasurement: &domain.Measurement{
				ID:    uuid.UUID{1},
				Name:  "spoon",
				Grams: 1,
			},
			expectedCount: 1,
			wantErr:       false,
		}, // успешное получение
		{
			name:         "ошибка выполнения запроса в репозитории",
			recipeId:     recipeId,
			ingredientId: ingredientId,
			beforeTest: func(measurementRepo mocks.MockIMeasurementRepository) {
				measurementRepo.EXPECT().
					GetByRecipeId(context.Background(), ingredientId, recipeId).
					Return(nil, 0, fmt.Errorf("getting measurement err"))
			},
			wantErr: true,
			errStr:  errors.New("getting measurement unit by recipe id: getting measurement err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*measurementRepo)
			}

			measurement, count, err := svc.GetByRecipeId(context.Background(), tt.ingredientId, tt.recipeId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expectedMeasurement, measurement)
				require.Equal(t, tt.expectedCount, count)
			}
		})
	}
}

func TestMeasurementService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	measurementRepo := mocks.NewMockIMeasurementRepository(ctrl)
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
	svc := services.NewMeasurementService(measurementRepo, logger)

	measurementId := uuid.New()

	tests := []struct {
		name        string
		measurement *domain.Measurement
		beforeTest  func(measurementRepo mocks.MockIMeasurementRepository)
		wantErr     bool
		errStr      error
	}{
		{
			name: "успешное обновление",
			measurement: &domain.Measurement{
				ID:    measurementId,
				Name:  "spoon",
				Grams: 1,
			},
			beforeTest: func(measurementRepo mocks.MockIMeasurementRepository) {
				measurementRepo.EXPECT().
					Update(context.Background(), &domain.Measurement{
						ID:    measurementId,
						Name:  "spoon",
						Grams: 1,
					}).Return(nil)
			},
			wantErr: false,
		}, // успешное обновление
		{
			name: "пустое название",
			measurement: &domain.Measurement{
				ID:    measurementId,
				Name:  "",
				Grams: 1,
			},
			wantErr: true,
			errStr:  errors.New("updating measurement unit: empty name"),
		}, // пустое название
		{
			name: "количество граммов - 0",
			measurement: &domain.Measurement{
				ID:    measurementId,
				Name:  "spoon",
				Grams: 0,
			},
			wantErr: true,
			errStr:  errors.New("updating measurement unit: negative or zero grams count"),
		}, // количество граммов - 0
		{
			name: "количество граммов <0",
			measurement: &domain.Measurement{
				ID:    measurementId,
				Name:  "spoon",
				Grams: 0,
			},
			wantErr: true,
			errStr:  errors.New("updating measurement unit: negative or zero grams count"),
		}, // количество граммов <0
		{
			name: "ошибка выполнения запроса в репозитории",
			measurement: &domain.Measurement{
				ID:    measurementId,
				Name:  "spoon",
				Grams: 1,
			},
			beforeTest: func(measurementRepo mocks.MockIMeasurementRepository) {
				measurementRepo.EXPECT().
					Update(context.Background(), &domain.Measurement{
						ID:    measurementId,
						Name:  "spoon",
						Grams: 1,
					}).Return(fmt.Errorf("updating measurement unit err"))
			},
			wantErr: true,
			errStr:  errors.New("updating measurement unit: updating measurement unit err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*measurementRepo)
			}

			err := svc.Update(context.Background(), tt.measurement)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestMeasurementService_UpdateLink(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	measurementRepo := mocks.NewMockIMeasurementRepository(ctrl)
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
	svc := services.NewMeasurementService(measurementRepo, logger)

	linkId := uuid.New()
	measurementId := uuid.New()

	tests := []struct {
		name          string
		linkId        uuid.UUID
		measurementId uuid.UUID
		amount        int
		beforeTest    func(measurementRepo mocks.MockIMeasurementRepository)
		wantErr       bool
		errStr        error
	}{
		{
			name:          "успешное обновление",
			linkId:        linkId,
			measurementId: measurementId,
			amount:        1,
			beforeTest: func(measurementRepo mocks.MockIMeasurementRepository) {
				measurementRepo.EXPECT().
					UpdateLink(context.Background(), linkId, measurementId, 1).
					Return(nil)
			},
			wantErr: false,
		}, // успешное обновление
		{
			name:          "количество - 0",
			linkId:        linkId,
			measurementId: measurementId,
			amount:        0,
			wantErr:       true,
			errStr:        errors.New("negative or zero amount"),
		}, // количество - 0
		{
			name:          "количество < 0",
			linkId:        linkId,
			measurementId: measurementId,
			amount:        -1,
			wantErr:       true,
			errStr:        errors.New("negative or zero amount"),
		}, // количество < 0
		{
			name:          "ошибка выполнения запроса в репозитории",
			linkId:        linkId,
			measurementId: measurementId,
			amount:        1,
			beforeTest: func(measurementRepo mocks.MockIMeasurementRepository) {
				measurementRepo.EXPECT().
					UpdateLink(context.Background(), linkId, measurementId, 1).
					Return(fmt.Errorf("updating measurement err"))
			},
			wantErr: true,
			errStr:  errors.New("updating measurement unit by link id: updating measurement err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*measurementRepo)
			}

			err := svc.UpdateLink(context.Background(), tt.linkId, tt.measurementId, tt.amount)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
