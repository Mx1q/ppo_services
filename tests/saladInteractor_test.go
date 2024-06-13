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

func TestSaladInteractor_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladService := mocks.NewMockISaladService(ctrl)
	validatorService := mocks.NewMockIValidatorService(ctrl)

	svc := services.NewSaladInteractor(saladService, []domain.IValidatorService{validatorService})

	tests := []struct {
		name       string
		salad      *domain.Salad
		beforeTest func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное создание",
			salad: &domain.Salad{
				ID:          uuid.UUID{1},
				AuthorID:    uuid.UUID{11},
				Name:        "salad",
				Description: "description",
			},
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(nil).
						Times(2)
				}

				saladService.EXPECT().
					Create(context.Background(), &domain.Salad{
						ID:          uuid.UUID{1},
						AuthorID:    uuid.UUID{11},
						Name:        "salad",
						Description: "description",
					}).
					Return(uuid.Nil, nil)
			},
			wantErr: false,
		}, // успешное создание
		{
			name: "ошибка валидации (название)",
			salad: &domain.Salad{
				ID:          uuid.UUID{1},
				AuthorID:    uuid.UUID{11},
				Name:        "salad",
				Description: "",
			},
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(fmt.Errorf("invalid"))
				}
			},
			wantErr: true,
			errStr:  errors.New("salad interactor (name): invalid"),
		}, // ошибка валидации (название)
		{
			name: "ошибка валидации",
			salad: &domain.Salad{
				ID:          uuid.UUID{1},
				AuthorID:    uuid.UUID{11},
				Name:        "salad",
				Description: "description",
			},
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(nil)
				}

				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(fmt.Errorf("invalid"))
				}
			},
			wantErr: true,
			errStr:  errors.New("salad interactor (description): invalid"),
		}, // ошибка валидации (описание)
		{
			name: "ошибка выполнения запроса в сервисе salad",
			salad: &domain.Salad{
				ID:          uuid.UUID{1},
				AuthorID:    uuid.UUID{11},
				Name:        "salad",
				Description: "",
			},
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(nil)
				}

				saladService.EXPECT().
					Create(context.Background(), &domain.Salad{
						ID:          uuid.UUID{1},
						AuthorID:    uuid.UUID{11},
						Name:        "salad",
						Description: "",
					}).
					Return(uuid.Nil, fmt.Errorf("creating err"))
			},
			wantErr: true,
			errStr:  errors.New("salad interactor: creating err"),
		}, // ошибка выполнения запроса в сервисе salad
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladService, []mocks.MockIValidatorService{*validatorService})
			}

			_, err := svc.Create(context.Background(), tt.salad)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestSaladInteractor_DeleteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladService := mocks.NewMockISaladService(ctrl)
	validatorService := mocks.NewMockIValidatorService(ctrl)

	svc := services.NewSaladInteractor(saladService, []domain.IValidatorService{validatorService})

	saladId := uuid.New()

	tests := []struct {
		name       string
		saladId    uuid.UUID
		beforeTest func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService)
		wantErr    bool
		errStr     error
	}{
		{
			name:    "успешное удаление",
			saladId: saladId,
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				saladService.EXPECT().
					DeleteById(context.Background(), saladId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное удаление
		{
			name:    "ошибка выполнения запроса в сервисе salad",
			saladId: saladId,
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				saladService.EXPECT().
					DeleteById(context.Background(), saladId).
					Return(fmt.Errorf("deleting err"))
			},
			wantErr: true,
			errStr:  errors.New("deleting err"),
		}, // ошибка выполнения запроса в сервисе salad
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladService, []mocks.MockIValidatorService{*validatorService})
			}

			err := svc.DeleteById(context.Background(), tt.saladId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestSaladInteractor_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladService := mocks.NewMockISaladService(ctrl)
	validatorService := mocks.NewMockIValidatorService(ctrl)

	svc := services.NewSaladInteractor(saladService, []domain.IValidatorService{validatorService})

	userId := uuid.New()
	filter := &domain.RecipeFilter{
		AvailableIngredients: nil,
		MinRate:              0,
		SaladTypes:           nil,
	}
	page := 1

	tests := []struct {
		name       string
		page       int
		beforeTest func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService)
		expected   []*domain.Salad
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное получение салатов",
			page: page,
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				saladService.EXPECT().
					GetAll(context.Background(), filter, page).
					Return([]*domain.Salad{
						{
							ID:          uuid.UUID{1},
							AuthorID:    userId,
							Name:        "salad1",
							Description: "",
						},
						{
							ID:          uuid.UUID{2},
							AuthorID:    userId,
							Name:        "salad2",
							Description: "",
						},
						{
							ID:          uuid.UUID{3},
							AuthorID:    userId,
							Name:        "salad3",
							Description: "",
						},
						{
							ID:          uuid.UUID{4},
							AuthorID:    userId,
							Name:        "salad1",
							Description: "",
						},
					}, 4, nil)
			},
			expected: []*domain.Salad{
				{
					ID:          uuid.UUID{1},
					AuthorID:    userId,
					Name:        "salad1",
					Description: "",
				},
				{
					ID:          uuid.UUID{2},
					AuthorID:    userId,
					Name:        "salad2",
					Description: "",
				},
				{
					ID:          uuid.UUID{3},
					AuthorID:    userId,
					Name:        "salad3",
					Description: "",
				},
				{
					ID:          uuid.UUID{4},
					AuthorID:    userId,
					Name:        "salad1",
					Description: "",
				},
			},
			wantErr: false,
		}, // успешное получение салатов
		{
			name: "ошибка выполнения запроса в сервисе salad",
			page: page,
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				saladService.EXPECT().
					GetAll(context.Background(), filter, page).
					Return(nil, 0, fmt.Errorf("getting salads err"))
			},
			wantErr: true,
			errStr:  errors.New("getting salads err"),
		}, // ошибка выполнения запроса в сервисе salad
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladService, []mocks.MockIValidatorService{*validatorService})
			}

			salads, _, err := svc.GetAll(context.Background(), filter, tt.page)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, salads)
			}
		})
	}
}

func TestSaladInteractor_GetAllRatedByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladService := mocks.NewMockISaladService(ctrl)
	validatorService := mocks.NewMockIValidatorService(ctrl)

	svc := services.NewSaladInteractor(saladService, []domain.IValidatorService{validatorService})

	userId := uuid.New()
	page := 1

	tests := []struct {
		name       string
		userId     uuid.UUID
		page       int
		beforeTest func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService)
		expected   []*domain.Salad
		wantErr    bool
		errStr     error
	}{
		{
			name:   "успешное получение салатов, оцененных пользователем",
			userId: userId,
			page:   page,
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				saladService.EXPECT().
					GetAllRatedByUser(context.Background(), userId, page).
					Return([]*domain.Salad{
						{
							ID:          uuid.UUID{1},
							AuthorID:    userId,
							Name:        "salad1",
							Description: "",
						},
						{
							ID:          uuid.UUID{2},
							AuthorID:    userId,
							Name:        "salad2",
							Description: "",
						},
						{
							ID:          uuid.UUID{3},
							AuthorID:    userId,
							Name:        "salad3",
							Description: "",
						},
						{
							ID:          uuid.UUID{4},
							AuthorID:    userId,
							Name:        "salad1",
							Description: "",
						},
					}, 1, nil)
			},
			expected: []*domain.Salad{
				{
					ID:          uuid.UUID{1},
					AuthorID:    userId,
					Name:        "salad1",
					Description: "",
				},
				{
					ID:          uuid.UUID{2},
					AuthorID:    userId,
					Name:        "salad2",
					Description: "",
				},
				{
					ID:          uuid.UUID{3},
					AuthorID:    userId,
					Name:        "salad3",
					Description: "",
				},
				{
					ID:          uuid.UUID{4},
					AuthorID:    userId,
					Name:        "salad1",
					Description: "",
				},
			},
			wantErr: false,
		}, // успешное получение салатов, оцененных пользователем
		{
			name:   "ошибка выполнения запроса в сервисе salad",
			userId: userId,
			page:   page,
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				saladService.EXPECT().
					GetAllRatedByUser(context.Background(), userId, page).
					Return(nil, 0, fmt.Errorf("getting salads err"))
			},
			wantErr: true,
			errStr:  errors.New("getting salads err"),
		}, // ошибка выполнения запроса в сервисе salad
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladService, []mocks.MockIValidatorService{*validatorService})
			}

			salads, _, err := svc.GetAllRatedByUser(context.Background(), tt.userId, tt.page)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, salads)
			}
		})
	}
}

func TestSaladInteractor_GetAllByUserId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladService := mocks.NewMockISaladService(ctrl)
	validatorService := mocks.NewMockIValidatorService(ctrl)

	svc := services.NewSaladInteractor(saladService, []domain.IValidatorService{validatorService})

	userId := uuid.New()

	tests := []struct {
		name       string
		userId     uuid.UUID
		beforeTest func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService)
		expected   []*domain.Salad
		wantErr    bool
		errStr     error
	}{
		{
			name:   "успешное получение салатов",
			userId: userId,
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				saladService.EXPECT().
					GetAllByUserId(context.Background(), userId).
					Return([]*domain.Salad{
						{
							ID:          uuid.UUID{1},
							AuthorID:    userId,
							Name:        "salad1",
							Description: "",
						},
						{
							ID:          uuid.UUID{2},
							AuthorID:    userId,
							Name:        "salad2",
							Description: "",
						},
						{
							ID:          uuid.UUID{3},
							AuthorID:    userId,
							Name:        "salad3",
							Description: "",
						},
						{
							ID:          uuid.UUID{4},
							AuthorID:    userId,
							Name:        "salad1",
							Description: "",
						},
					}, nil)
			},
			expected: []*domain.Salad{
				{
					ID:          uuid.UUID{1},
					AuthorID:    userId,
					Name:        "salad1",
					Description: "",
				},
				{
					ID:          uuid.UUID{2},
					AuthorID:    userId,
					Name:        "salad2",
					Description: "",
				},
				{
					ID:          uuid.UUID{3},
					AuthorID:    userId,
					Name:        "salad3",
					Description: "",
				},
				{
					ID:          uuid.UUID{4},
					AuthorID:    userId,
					Name:        "salad1",
					Description: "",
				},
			},
			wantErr: false,
		}, // успешное получение салатов
		{
			name:   "ошибка выполнения запроса в сервисе salad",
			userId: userId,
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				saladService.EXPECT().
					GetAllByUserId(context.Background(), userId).
					Return(nil, fmt.Errorf("getting salads err"))
			},
			wantErr: true,
			errStr:  errors.New("getting salads err"),
		}, // ошибка выполнения запроса в сервисе salad
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladService, []mocks.MockIValidatorService{*validatorService})
			}

			salads, err := svc.GetAllByUserId(context.Background(), tt.userId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, salads)
			}
		})
	}
}

func TestSaladInteractor_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladService := mocks.NewMockISaladService(ctrl)
	validatorService := mocks.NewMockIValidatorService(ctrl)

	svc := services.NewSaladInteractor(saladService, []domain.IValidatorService{validatorService})

	saladId := uuid.New()

	tests := []struct {
		name       string
		saladId    uuid.UUID
		beforeTest func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService)
		expected   *domain.Salad
		wantErr    bool
		errStr     error
	}{
		{
			name:    "успешное получение",
			saladId: saladId,
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				saladService.EXPECT().
					GetById(context.Background(), saladId).
					Return(&domain.Salad{
						ID:          saladId,
						AuthorID:    uuid.UUID{1},
						Name:        "salad1",
						Description: "",
					}, nil)
			},
			expected: &domain.Salad{
				ID:          saladId,
				AuthorID:    uuid.UUID{1},
				Name:        "salad1",
				Description: "",
			},
			wantErr: false,
		}, // успешное получение
		{
			name:    "ошибка выполнения запроса в сервисе salad",
			saladId: saladId,
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				saladService.EXPECT().
					GetById(context.Background(), saladId).
					Return(nil, fmt.Errorf("getting salad err"))
			},
			wantErr: true,
			errStr:  errors.New("getting salad err"),
		}, // ошибка выполнения запроса в сервисе salad
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladService, []mocks.MockIValidatorService{*validatorService})
			}

			salad, err := svc.GetById(context.Background(), tt.saladId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, salad)
			}
		})
	}
}

func TestSaladInteractor_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladService := mocks.NewMockISaladService(ctrl)
	validatorService := mocks.NewMockIValidatorService(ctrl)

	svc := services.NewSaladInteractor(saladService, []domain.IValidatorService{validatorService})

	tests := []struct {
		name       string
		salad      *domain.Salad
		beforeTest func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное обновление",
			salad: &domain.Salad{
				ID:          uuid.UUID{1},
				AuthorID:    uuid.UUID{11},
				Name:        "salad",
				Description: "description",
			},
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(nil).
						Times(2)
				}

				saladService.EXPECT().
					Update(context.Background(), &domain.Salad{
						ID:          uuid.UUID{1},
						AuthorID:    uuid.UUID{11},
						Name:        "salad",
						Description: "description",
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное обновление
		{
			name: "ошибка валидации (название)",
			salad: &domain.Salad{
				ID:          uuid.UUID{1},
				AuthorID:    uuid.UUID{11},
				Name:        "salad",
				Description: "",
			},
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(fmt.Errorf("invalid"))
				}
			},
			wantErr: true,
			errStr:  errors.New("salad interactor (name): invalid"),
		}, // ошибка валидации (название)
		{
			name: "ошибка валидации (описание)",
			salad: &domain.Salad{
				ID:          uuid.UUID{1},
				AuthorID:    uuid.UUID{11},
				Name:        "salad",
				Description: "description",
			},
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(nil)
				}

				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(fmt.Errorf("invalid"))
				}
			},
			wantErr: true,
			errStr:  errors.New("salad interactor (description): invalid"),
		}, // ошибка валидации (описание)
		{
			name: "ошибка выполнения запроса в сервисе salad",
			salad: &domain.Salad{
				ID:          uuid.UUID{1},
				AuthorID:    uuid.UUID{11},
				Name:        "salad",
				Description: "",
			},
			beforeTest: func(saladService mocks.MockISaladService, validatorServices []mocks.MockIValidatorService) {
				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(nil)
				}

				saladService.EXPECT().
					Update(context.Background(), &domain.Salad{
						ID:          uuid.UUID{1},
						AuthorID:    uuid.UUID{11},
						Name:        "salad",
						Description: "",
					}).
					Return(fmt.Errorf("updating err"))
			},
			wantErr: true,
			errStr:  errors.New("salad interactor: updating err"),
		}, // ошибка выполнения запроса в сервисе salad
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladService, []mocks.MockIValidatorService{*validatorService})
			}

			err := svc.Update(context.Background(), tt.salad)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
