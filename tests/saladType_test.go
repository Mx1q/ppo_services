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

func TestSaladTypeService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladTypeRepo := mocks.NewMockISaladTypeRepository(ctrl)
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
	svc := services.NewSaladTypeService(saladTypeRepo, logger)

	saladTypeId := uuid.New()

	tests := []struct {
		name       string
		saladType  *domain.SaladType
		beforeTest func(saladTypeRepo mocks.MockISaladTypeRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное создание",
			saladType: &domain.SaladType{
				ID:          saladTypeId,
				Name:        "meat",
				Description: "",
			},
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					Create(context.Background(), &domain.SaladType{
						ID:          saladTypeId,
						Name:        "meat",
						Description: "",
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное создание
		{
			name: "пустое название",
			saladType: &domain.SaladType{
				ID:          saladTypeId,
				Name:        "",
				Description: "",
			},
			wantErr: true,
			errStr:  errors.New("creating salad type: empty name"),
		}, // пустое название
		{
			name: "ошибка выполнения запроса в репозитории",
			saladType: &domain.SaladType{
				ID:          saladTypeId,
				Name:        "meat",
				Description: "",
			},
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					Create(context.Background(), &domain.SaladType{
						ID:          saladTypeId,
						Name:        "meat",
						Description: "",
					}).
					Return(fmt.Errorf("creating err"))
			},
			wantErr: true,
			errStr:  errors.New("creating salad type: creating err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladTypeRepo)
			}

			err := svc.Create(context.Background(), tt.saladType)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestSaladTypeService_DeleteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladTypeRepo := mocks.NewMockISaladTypeRepository(ctrl)
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
	svc := services.NewSaladTypeService(saladTypeRepo, logger)

	saladTypeId := uuid.New()

	tests := []struct {
		name        string
		saladTypeId uuid.UUID
		beforeTest  func(saladTypeRepo mocks.MockISaladTypeRepository)
		wantErr     bool
		errStr      error
	}{
		{
			name:        "успешное удаление",
			saladTypeId: saladTypeId,
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					DeleteById(context.Background(), saladTypeId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное удаление
		{
			name:        "ошибка выполнения запроса в репозитории",
			saladTypeId: saladTypeId,
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					DeleteById(context.Background(), saladTypeId).
					Return(fmt.Errorf("deleting err"))
			},
			wantErr: true,
			errStr:  errors.New("deleting salad type by id: deleting err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladTypeRepo)
			}

			err := svc.DeleteById(context.Background(), tt.saladTypeId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestSaladTypeService_Link(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladTypeRepo := mocks.NewMockISaladTypeRepository(ctrl)
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
	svc := services.NewSaladTypeService(saladTypeRepo, logger)

	saladId := uuid.New()
	saladTypeId := uuid.New()

	tests := []struct {
		name        string
		saladId     uuid.UUID
		saladTypeId uuid.UUID
		beforeTest  func(saladTypeRepo mocks.MockISaladTypeRepository)
		wantErr     bool
		errStr      error
	}{
		{
			name:        "успешное добавление типа",
			saladId:     saladId,
			saladTypeId: saladTypeId,
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					Link(context.Background(), saladId, saladTypeId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное добавление типа
		{
			name:        "ошибка выполнения запроса в репозитории",
			saladId:     saladId,
			saladTypeId: saladTypeId,
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					Link(context.Background(), saladId, saladTypeId).
					Return(fmt.Errorf("linking err"))
			},
			wantErr: true,
			errStr:  errors.New("linking salad type: linking err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladTypeRepo)
			}

			err := svc.Link(context.Background(), tt.saladId, tt.saladTypeId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestSaladTypeService_Unlink(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladTypeRepo := mocks.NewMockISaladTypeRepository(ctrl)
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
	svc := services.NewSaladTypeService(saladTypeRepo, logger)

	saladId := uuid.New()
	saladTypeId := uuid.New()

	tests := []struct {
		name        string
		saladId     uuid.UUID
		saladTypeId uuid.UUID
		beforeTest  func(saladTypeRepo mocks.MockISaladTypeRepository)
		wantErr     bool
		errStr      error
	}{
		{
			name:        "успешное удаление типа",
			saladId:     saladId,
			saladTypeId: saladTypeId,
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					Unlink(context.Background(), saladId, saladTypeId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное удаление типа
		{
			name:        "ошибка выполнения запроса в репозитории",
			saladId:     saladId,
			saladTypeId: saladTypeId,
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					Unlink(context.Background(), saladId, saladTypeId).
					Return(fmt.Errorf("unlinking err"))
			},
			wantErr: true,
			errStr:  errors.New("unlinking salad type: unlinking err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladTypeRepo)
			}

			err := svc.Unlink(context.Background(), tt.saladId, tt.saladTypeId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestSaladTypeService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladTypeRepo := mocks.NewMockISaladTypeRepository(ctrl)
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
	svc := services.NewSaladTypeService(saladTypeRepo, logger)

	page := 1

	tests := []struct {
		name       string
		page       int
		beforeTest func(saladTypeRepo mocks.MockISaladTypeRepository)
		expected   []*domain.SaladType
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное получение",
			page: page,
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					GetAll(context.Background(), page).
					Return([]*domain.SaladType{
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
					}, 1, nil)
			},
			expected: []*domain.SaladType{
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
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					GetAll(context.Background(), page).
					Return(nil, 0, fmt.Errorf("getting all err"))
			},
			wantErr: true,
			errStr:  errors.New("getting all salad types: getting all err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladTypeRepo)
			}

			types, _, err := svc.GetAll(context.Background(), page)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, types)
			}
		})
	}
}

func TestSaladTypeService_GetAllBySaladId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladTypeRepo := mocks.NewMockISaladTypeRepository(ctrl)
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
	svc := services.NewSaladTypeService(saladTypeRepo, logger)

	saladId := uuid.New()

	tests := []struct {
		name       string
		saladId    uuid.UUID
		beforeTest func(saladTypeRepo mocks.MockISaladTypeRepository)
		expected   []*domain.SaladType
		wantErr    bool
		errStr     error
	}{
		{
			name:    "успешное получение",
			saladId: saladId,
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					GetAllBySaladId(context.Background(), saladId).
					Return([]*domain.SaladType{
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
			expected: []*domain.SaladType{
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
			name:    "ошибка выполнения запроса в репозитории",
			saladId: saladId,
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					GetAllBySaladId(context.Background(), saladId).
					Return(nil, fmt.Errorf("getting all err"))
			},
			wantErr: true,
			errStr:  errors.New("getting all types of salad: getting all err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladTypeRepo)
			}

			types, err := svc.GetAllBySaladId(context.Background(), tt.saladId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, types)
			}
		})
	}
}

func TestSaladTypeService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladTypeRepo := mocks.NewMockISaladTypeRepository(ctrl)
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
	svc := services.NewSaladTypeService(saladTypeRepo, logger)

	typeId := uuid.New()

	tests := []struct {
		name       string
		typeId     uuid.UUID
		beforeTest func(saladTypeRepo mocks.MockISaladTypeRepository)
		expected   *domain.SaladType
		wantErr    bool
		errStr     error
	}{
		{
			name:   "успешное получение",
			typeId: typeId,
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					GetById(context.Background(), typeId).
					Return(&domain.SaladType{
						ID:          typeId,
						Name:        "type",
						Description: "",
					}, nil)
			},
			expected: &domain.SaladType{
				ID:          typeId,
				Name:        "type",
				Description: "",
			},
			wantErr: false,
		}, // успешное получение
		{
			name:   "ошибка выполнения запроса в репозитории",
			typeId: typeId,
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					GetById(context.Background(), typeId).
					Return(nil, fmt.Errorf("getting type err"))
			},
			wantErr: true,
			errStr:  errors.New("getting salad type by id: getting type err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladTypeRepo)
			}

			saladType, err := svc.GetById(context.Background(), tt.typeId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, saladType)
			}
		})
	}
}

func TestSaladTypeService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladTypeRepo := mocks.NewMockISaladTypeRepository(ctrl)
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
	svc := services.NewSaladTypeService(saladTypeRepo, logger)

	saladTypeId := uuid.New()

	tests := []struct {
		name       string
		saladType  *domain.SaladType
		beforeTest func(saladTypeRepo mocks.MockISaladTypeRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное создание",
			saladType: &domain.SaladType{
				ID:          saladTypeId,
				Name:        "meat",
				Description: "",
			},
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					Update(context.Background(), &domain.SaladType{
						ID:          saladTypeId,
						Name:        "meat",
						Description: "",
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное обновление
		{
			name: "пустое название",
			saladType: &domain.SaladType{
				ID:          saladTypeId,
				Name:        "",
				Description: "",
			},
			wantErr: true,
			errStr:  errors.New("updating salad type: empty name"),
		}, // пустое название
		{
			name: "ошибка выполнения запроса в репозитории",
			saladType: &domain.SaladType{
				ID:          saladTypeId,
				Name:        "meat",
				Description: "",
			},
			beforeTest: func(saladTypeRepo mocks.MockISaladTypeRepository) {
				saladTypeRepo.EXPECT().
					Update(context.Background(), &domain.SaladType{
						ID:          saladTypeId,
						Name:        "meat",
						Description: "",
					}).
					Return(fmt.Errorf("updating err"))
			},
			wantErr: true,
			errStr:  errors.New("updating salad type: updating err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladTypeRepo)
			}

			err := svc.Update(context.Background(), tt.saladType)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
