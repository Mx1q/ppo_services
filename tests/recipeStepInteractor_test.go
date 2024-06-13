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

func TestRecipeStepInteractor_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeStepService := mocks.NewMockIRecipeStepService(ctrl)
	validatorService := mocks.NewMockIValidatorService(ctrl)

	svc := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{validatorService})

	tests := []struct {
		name       string
		step       *domain.RecipeStep
		beforeTest func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное создание",
			step: &domain.RecipeStep{
				ID:          uuid.UUID{1},
				RecipeID:    uuid.UUID{11},
				Name:        "salad",
				Description: "description",
				StepNum:     1,
			},
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(nil).
						Times(2)
				}

				recipeStepService.EXPECT().
					Create(context.Background(), &domain.RecipeStep{
						ID:          uuid.UUID{1},
						RecipeID:    uuid.UUID{11},
						Name:        "salad",
						Description: "description",
						StepNum:     1,
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное создание
		{
			name: "ошибка валидации (название)",
			step: &domain.RecipeStep{
				ID:          uuid.UUID{1},
				RecipeID:    uuid.UUID{11},
				Name:        "salad",
				Description: "description",
				StepNum:     1,
			},
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(fmt.Errorf("invalid"))
				}
			},
			wantErr: true,
			errStr:  errors.New("recipe step interactor (name): invalid"),
		}, // ошибка валидации (название)
		{
			name: "ошибка валидации",
			step: &domain.RecipeStep{
				ID:          uuid.UUID{1},
				RecipeID:    uuid.UUID{11},
				Name:        "salad",
				Description: "description",
				StepNum:     1,
			},
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
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
			errStr:  errors.New("recipe step interactor (description): invalid"),
		}, // ошибка валидации (описание)
		{
			name: "ошибка выполнения запроса в сервисе recipe step",
			step: &domain.RecipeStep{
				ID:          uuid.UUID{1},
				RecipeID:    uuid.UUID{11},
				Name:        "salad",
				Description: "description",
				StepNum:     1,
			},
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(nil).
						Times(2)
				}

				recipeStepService.EXPECT().
					Create(context.Background(), &domain.RecipeStep{
						ID:          uuid.UUID{1},
						RecipeID:    uuid.UUID{11},
						Name:        "salad",
						Description: "description",
						StepNum:     1,
					}).
					Return(fmt.Errorf("creating err"))
			},
			wantErr: true,
			errStr:  errors.New("recipe step interactor: creating err"),
		}, // ошибка выполнения запроса в сервисе recipe step
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeStepService, []mocks.MockIValidatorService{*validatorService})
			}

			err := svc.Create(context.Background(), tt.step)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestRecipeStepInteractor_DeleteAllByRecipeID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeStepService := mocks.NewMockIRecipeStepService(ctrl)
	validatorService := mocks.NewMockIValidatorService(ctrl)

	svc := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{validatorService})

	recipeId := uuid.New()

	tests := []struct {
		name       string
		recipeId   uuid.UUID
		beforeTest func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService)
		wantErr    bool
		errStr     error
	}{
		{
			name:     "успешное удаление",
			recipeId: recipeId,
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
				recipeStepService.EXPECT().
					DeleteAllByRecipeID(context.Background(), recipeId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное удаление
		{
			name:     "ошибка выполнения запроса в сервисе recipe step",
			recipeId: recipeId,
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
				recipeStepService.EXPECT().
					DeleteAllByRecipeID(context.Background(), recipeId).
					Return(fmt.Errorf("deleting err"))
			},
			wantErr: true,
			errStr:  errors.New("deleting err"),
		}, // ошибка выполнения запроса в сервисе recipe step
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeStepService, []mocks.MockIValidatorService{*validatorService})
			}

			err := svc.DeleteAllByRecipeID(context.Background(), tt.recipeId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestRecipeStepInteractor_DeleteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeStepService := mocks.NewMockIRecipeStepService(ctrl)
	validatorService := mocks.NewMockIValidatorService(ctrl)

	svc := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{validatorService})

	stepId := uuid.New()

	tests := []struct {
		name       string
		stepId     uuid.UUID
		beforeTest func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService)
		wantErr    bool
		errStr     error
	}{
		{
			name:   "успешное удаление",
			stepId: stepId,
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
				recipeStepService.EXPECT().
					DeleteById(context.Background(), stepId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное удаление
		{
			name:   "ошибка выполнения запроса в сервисе recipe step",
			stepId: stepId,
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
				recipeStepService.EXPECT().
					DeleteById(context.Background(), stepId).
					Return(fmt.Errorf("deleting err"))
			},
			wantErr: true,
			errStr:  errors.New("deleting err"),
		}, // ошибка выполнения запроса в сервисе recipe step
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeStepService, []mocks.MockIValidatorService{*validatorService})
			}

			err := svc.DeleteById(context.Background(), tt.stepId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestRecipeStepInteractor_GetAllByRecipeID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeStepService := mocks.NewMockIRecipeStepService(ctrl)
	validatorService := mocks.NewMockIValidatorService(ctrl)

	svc := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{validatorService})

	recipeId := uuid.New()

	tests := []struct {
		name       string
		recipeId   uuid.UUID
		beforeTest func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService)
		expected   []*domain.RecipeStep
		wantErr    bool
		errStr     error
	}{
		{
			name:     "успешное удаление",
			recipeId: recipeId,
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
				recipeStepService.EXPECT().
					GetAllByRecipeID(context.Background(), recipeId).
					Return([]*domain.RecipeStep{
						{
							ID:          uuid.UUID{1},
							RecipeID:    recipeId,
							Name:        "first step",
							Description: "description",
							StepNum:     1,
						},
						{
							ID:          uuid.UUID{2},
							RecipeID:    recipeId,
							Name:        "second step",
							Description: "description",
							StepNum:     2,
						},
						{
							ID:          uuid.UUID{3},
							RecipeID:    recipeId,
							Name:        "third step",
							Description: "description",
							StepNum:     3,
						},
					}, nil)
			},
			expected: []*domain.RecipeStep{
				{
					ID:          uuid.UUID{1},
					RecipeID:    recipeId,
					Name:        "first step",
					Description: "description",
					StepNum:     1,
				},
				{
					ID:          uuid.UUID{2},
					RecipeID:    recipeId,
					Name:        "second step",
					Description: "description",
					StepNum:     2,
				},
				{
					ID:          uuid.UUID{3},
					RecipeID:    recipeId,
					Name:        "third step",
					Description: "description",
					StepNum:     3,
				},
			},
			wantErr: false,
		}, // успешное получение
		{
			name:     "ошибка выполнения запроса в сервисе recipe step",
			recipeId: recipeId,
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
				recipeStepService.EXPECT().
					GetAllByRecipeID(context.Background(), recipeId).
					Return(nil, fmt.Errorf("getting err"))
			},
			wantErr: true,
			errStr:  errors.New("getting err"),
		}, // ошибка выполнения запроса в сервисе recipe step
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeStepService, []mocks.MockIValidatorService{*validatorService})
			}

			steps, err := svc.GetAllByRecipeID(context.Background(), tt.recipeId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, steps)
			}
		})
	}
}

func TestRecipeStepInteractor_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeStepService := mocks.NewMockIRecipeStepService(ctrl)
	validatorService := mocks.NewMockIValidatorService(ctrl)

	svc := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{validatorService})

	stepId := uuid.New()

	tests := []struct {
		name       string
		stepId     uuid.UUID
		beforeTest func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService)
		expected   *domain.RecipeStep
		wantErr    bool
		errStr     error
	}{
		{
			name:   "успешное удаление",
			stepId: stepId,
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
				recipeStepService.EXPECT().
					GetById(context.Background(), stepId).
					Return(&domain.RecipeStep{
						ID:          uuid.UUID{1},
						RecipeID:    uuid.UUID{11},
						Name:        "first step",
						Description: "description",
						StepNum:     1,
					}, nil)
			},
			expected: &domain.RecipeStep{
				ID:          uuid.UUID{1},
				RecipeID:    uuid.UUID{11},
				Name:        "first step",
				Description: "description",
				StepNum:     1,
			},
			wantErr: false,
		}, // успешное получение
		{
			name:   "ошибка выполнения запроса в сервисе recipe step",
			stepId: stepId,
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
				recipeStepService.EXPECT().
					GetById(context.Background(), stepId).
					Return(nil, fmt.Errorf("getting err"))
			},
			wantErr: true,
			errStr:  errors.New("getting err"),
		}, // ошибка выполнения запроса в сервисе recipe step
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeStepService, []mocks.MockIValidatorService{*validatorService})
			}

			step, err := svc.GetById(context.Background(), tt.stepId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, step)
			}
		})
	}
}

func TestRecipeStepInteractor_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeStepService := mocks.NewMockIRecipeStepService(ctrl)
	validatorService := mocks.NewMockIValidatorService(ctrl)

	svc := services.NewRecipeStepInteractor(recipeStepService, []domain.IValidatorService{validatorService})

	tests := []struct {
		name       string
		step       *domain.RecipeStep
		beforeTest func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное создание",
			step: &domain.RecipeStep{
				ID:          uuid.UUID{1},
				RecipeID:    uuid.UUID{11},
				Name:        "salad",
				Description: "description",
				StepNum:     1,
			},
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(nil).
						Times(2)
				}

				recipeStepService.EXPECT().
					Update(context.Background(), &domain.RecipeStep{
						ID:          uuid.UUID{1},
						RecipeID:    uuid.UUID{11},
						Name:        "salad",
						Description: "description",
						StepNum:     1,
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное обновление
		{
			name: "ошибка валидации (название)",
			step: &domain.RecipeStep{
				ID:          uuid.UUID{1},
				RecipeID:    uuid.UUID{11},
				Name:        "salad",
				Description: "description",
				StepNum:     1,
			},
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(fmt.Errorf("invalid"))
				}
			},
			wantErr: true,
			errStr:  errors.New("recipe step interactor (name): invalid"),
		}, // ошибка валидации (название)
		{
			name: "ошибка валидации",
			step: &domain.RecipeStep{
				ID:          uuid.UUID{1},
				RecipeID:    uuid.UUID{11},
				Name:        "salad",
				Description: "description",
				StepNum:     1,
			},
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
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
			errStr:  errors.New("recipe step interactor (description): invalid"),
		}, // ошибка валидации (описание)
		{
			name: "ошибка выполнения запроса в сервисе recipe step",
			step: &domain.RecipeStep{
				ID:          uuid.UUID{1},
				RecipeID:    uuid.UUID{11},
				Name:        "salad",
				Description: "description",
				StepNum:     1,
			},
			beforeTest: func(recipeStepService mocks.MockIRecipeStepService, validatorServices []mocks.MockIValidatorService) {
				for _, validator := range validatorServices {
					validator.EXPECT().
						Verify(context.Background(), gomock.Any()).
						Return(nil).
						Times(2)
				}

				recipeStepService.EXPECT().
					Update(context.Background(), &domain.RecipeStep{
						ID:          uuid.UUID{1},
						RecipeID:    uuid.UUID{11},
						Name:        "salad",
						Description: "description",
						StepNum:     1,
					}).
					Return(fmt.Errorf("updating err"))
			},
			wantErr: true,
			errStr:  errors.New("recipe step interactor: updating err"),
		}, // ошибка выполнения запроса в сервисе recipe step
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeStepService, []mocks.MockIValidatorService{*validatorService})
			}

			err := svc.Update(context.Background(), tt.step)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
