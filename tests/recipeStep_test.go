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

func TestRecipeStepService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeStepRepo := mocks.NewMockIRecipeStepRepository(ctrl)
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
	svc := services.NewRecipeStepService(recipeStepRepo, logger)

	stepId := uuid.New()
	recipeId := uuid.New()

	tests := []struct {
		name       string
		recipeStep *domain.RecipeStep
		beforeTest func(recipeStepRepo mocks.MockIRecipeStepRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное создание первого шага",
			recipeStep: &domain.RecipeStep{
				ID:          stepId,
				RecipeID:    recipeId,
				Name:        "first recipe step",
				Description: "description",
				StepNum:     1,
			},
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
					Create(context.Background(), &domain.RecipeStep{
						ID:          stepId,
						RecipeID:    recipeId,
						Name:        "first recipe step",
						Description: "description",
						StepNum:     1,
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное создание первого шага
		{
			name: "успешное создание второго шага",
			recipeStep: &domain.RecipeStep{
				ID:          stepId,
				RecipeID:    recipeId,
				Name:        "second recipe step",
				Description: "description",
				StepNum:     2,
			},
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
					Create(context.Background(), &domain.RecipeStep{
						ID:          stepId,
						RecipeID:    recipeId,
						Name:        "second recipe step",
						Description: "description",
						StepNum:     2,
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное создание второго шага
		{
			name: "пустое имя шага",
			recipeStep: &domain.RecipeStep{
				ID:          stepId,
				RecipeID:    recipeId,
				Name:        "",
				Description: "empty step name",
				StepNum:     1,
			},
			wantErr: true,
			errStr:  errors.New("creating recipe step: empty name"),
		}, // пустое имя шага
		{
			name: "пустое описание шага",
			recipeStep: &domain.RecipeStep{
				ID:          stepId,
				RecipeID:    recipeId,
				Name:        "empty step description",
				Description: "",
				StepNum:     1,
			},
			wantErr: true,
			errStr:  errors.New("creating recipe step: empty description"),
		}, // пустое описание шага
		{
			name: "невалидный номер шага - 0",
			recipeStep: &domain.RecipeStep{
				ID:          stepId,
				RecipeID:    recipeId,
				Name:        "first recipe step",
				Description: "description",
				StepNum:     0,
			},
			wantErr: true,
			errStr:  errors.New("creating recipe step: negative or zero step num"),
		}, // невалидный номер шага - 0
		{
			name: "невалидный номер шага - отрицательный",
			recipeStep: &domain.RecipeStep{
				ID:          stepId,
				RecipeID:    recipeId,
				Name:        "first recipe step",
				Description: "description",
				StepNum:     -1,
			},
			wantErr: true,
			errStr:  errors.New("creating recipe step: negative or zero step num"),
		}, // невалидный номер шага - отрицательный
		//{
		//	name: "невалидный номер шага - уже существует",
		//	recipeStep: &domain.RecipeStep{
		//		ID:          stepId,
		//		RecipeID:    recipeId,
		//		Name:        "first recipe step",
		//		Description: "description",
		//		StepNum:     1,
		//	},
		//	beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
		//		recipeStepRepo.EXPECT().
		//			GetAllByRecipeID(context.Background(), recipeId).
		//			Return([]*domain.RecipeStep{
		//				{
		//					ID:          uuid.UUID{1},
		//					RecipeID:    recipeId,
		//					Name:        "first step",
		//					Description: "",
		//					StepNum:     1,
		//				},
		//			}, nil)
		//	},
		//	wantErr: true,
		//	errStr:  errors.New("creating recipe step: invalid step num"),
		//}, // невалидный номер шага - уже существует
		{
			name: "ошибка выполнения запроса в репозитории (create)",
			recipeStep: &domain.RecipeStep{
				ID:          stepId,
				RecipeID:    recipeId,
				Name:        "first recipe step",
				Description: "description",
				StepNum:     1,
			},
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
					Create(context.Background(), &domain.RecipeStep{
						ID:          stepId,
						RecipeID:    recipeId,
						Name:        "first recipe step",
						Description: "description",
						StepNum:     1,
					}).
					Return(fmt.Errorf("creating step err"))
			},
			wantErr: true,
			errStr:  errors.New("creating recipe step: creating step err"),
		}, // ошибка выполнения запроса в репозитории (create)
		//{
		//	name: "ошибка выполнения запроса в репозитории (get all by recipe id)",
		//	recipeStep: &domain.RecipeStep{
		//		ID:          stepId,
		//		RecipeID:    recipeId,
		//		Name:        "first recipe step",
		//		Description: "description",
		//		StepNum:     1,
		//	},
		//	beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
		//		recipeStepRepo.EXPECT().
		//			GetAllByRecipeID(context.Background(), recipeId).
		//			Return(nil, fmt.Errorf("getting steps err"))
		//	},
		//	wantErr: true,
		//	errStr:  errors.New("creating recipe step: getting steps err"),
		//}, // ошибка выполнения запроса в репозитории (get all by recipe id)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeStepRepo)
			}

			err := svc.Create(context.Background(), tt.recipeStep)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestRecipeStepService_DeleteAllByRecipeID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeStepRepo := mocks.NewMockIRecipeStepRepository(ctrl)
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
	svc := services.NewRecipeStepService(recipeStepRepo, logger)

	recipeId := uuid.New()

	tests := []struct {
		name       string
		recipeId   uuid.UUID
		beforeTest func(recipeStepRepo mocks.MockIRecipeStepRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name:     "успешное удаление всех шагов рецепта",
			recipeId: recipeId,
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
					DeleteAllByRecipeID(context.Background(), recipeId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное удаление всех шагов рецепта
		{
			name:     "ошибка выполнения запроса в репозитории",
			recipeId: recipeId,
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
					DeleteAllByRecipeID(context.Background(), recipeId).
					Return(fmt.Errorf("deleting all steps of recipe err"))
			},
			wantErr: true,
			errStr:  errors.New("deleting all step of recipe: deleting all steps of recipe err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeStepRepo)
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

func TestRecipeStepService_DeleteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeStepRepo := mocks.NewMockIRecipeStepRepository(ctrl)
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
	svc := services.NewRecipeStepService(recipeStepRepo, logger)

	tests := []struct {
		name       string
		stepId     uuid.UUID
		beforeTest func(recipeStepRepo mocks.MockIRecipeStepRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name:   "успешное удаление единственного шага",
			stepId: uuid.UUID{1},
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
					DeleteById(context.Background(), uuid.UUID{1}).
					Return(nil)
			},
		}, // успешное удаление единственного шага
		{
			name:   "успешное удаление последнего шага",
			stepId: uuid.UUID{3},
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
					DeleteById(context.Background(), uuid.UUID{3}).
					Return(nil)
			},
		}, // успешное удаление последнего шага
		{
			name:   "успешное удаление первого шага",
			stepId: uuid.UUID{1},
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
					DeleteById(context.Background(), uuid.UUID{1}).
					Return(nil)
			},
		}, // успешное удаление первого шага
		{
			name:   "ошибка удаления шага в репозитории",
			stepId: uuid.UUID{1},
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
					DeleteById(context.Background(), uuid.UUID{1}).
					Return(fmt.Errorf("deleting step err"))
			},
			wantErr: true,
			errStr:  errors.New("deleting step by id: deleting step err"),
		}, // ошибка удаления шага в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeStepRepo)
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

func TestRecipeStepService_GetAllByRecipeID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeStepRepo := mocks.NewMockIRecipeStepRepository(ctrl)
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
	svc := services.NewRecipeStepService(recipeStepRepo, logger)

	recipeId := uuid.New()

	tests := []struct {
		name       string
		recipeId   uuid.UUID
		beforeTest func(recipeStepRepo mocks.MockIRecipeStepRepository)
		expected   []*domain.RecipeStep
		wantErr    bool
		errStr     error
	}{
		{
			name:     "успешное получение всех шагов рецепта",
			recipeId: recipeId,
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
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
		}, // успешное получение всех шагов рецепта
		{
			name:     "ошибка выполнения запроса в репозитории",
			recipeId: recipeId,
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
					GetAllByRecipeID(context.Background(), recipeId).
					Return(nil, fmt.Errorf("getting steps err"))
			},
			wantErr: true,
			errStr:  errors.New("getting all steps of recipe: getting steps err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeStepRepo)
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

func TestRecipeStepService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeStepRepo := mocks.NewMockIRecipeStepRepository(ctrl)
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
	svc := services.NewRecipeStepService(recipeStepRepo, logger)

	stepId := uuid.New()

	tests := []struct {
		name       string
		stepId     uuid.UUID
		beforeTest func(recipeStepRepo mocks.MockIRecipeStepRepository)
		expected   *domain.RecipeStep
		wantErr    bool
		errStr     error
	}{
		{
			name:   "успешное получение шага",
			stepId: stepId,
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
					GetById(context.Background(), stepId).
					Return(&domain.RecipeStep{
						ID:          stepId,
						RecipeID:    uuid.UUID{1},
						Name:        "first step",
						Description: "description",
						StepNum:     1,
					}, nil)
			},
			expected: &domain.RecipeStep{
				ID:          stepId,
				RecipeID:    uuid.UUID{1},
				Name:        "first step",
				Description: "description",
				StepNum:     1,
			},
			wantErr: false,
		}, // успешное получение шага
		{
			name:   "ошибка выполнения запроса в репозитории",
			stepId: stepId,
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
					GetById(context.Background(), stepId).
					Return(nil, fmt.Errorf("getting step err"))
			},
			wantErr: true,
			errStr:  errors.New("getting recipe step by id: getting step err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeStepRepo)
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

func TestRecipeStepService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeStepRepo := mocks.NewMockIRecipeStepRepository(ctrl)
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
	svc := services.NewRecipeStepService(recipeStepRepo, logger)

	stepId := uuid.UUID{1}
	recipeId := uuid.UUID{0}

	tests := []struct {
		name       string
		step       *domain.RecipeStep
		beforeTest func(recipeStepRepo mocks.MockIRecipeStepRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное обновление шага (без изменения номера)",
			step: &domain.RecipeStep{
				ID:          stepId,
				RecipeID:    recipeId,
				Name:        "updated first step",
				Description: "description",
				StepNum:     1,
			},
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
					Update(context.Background(), &domain.RecipeStep{
						ID:          stepId,
						RecipeID:    recipeId,
						Name:        "updated first step",
						Description: "description",
						StepNum:     1,
					}).Return(nil)
			},
			wantErr: false,
		}, // успешное обновление шага (без изменения номера)
		{
			name: "успешное обновление шага (с изменением номера)",
			step: &domain.RecipeStep{
				ID:          stepId,
				RecipeID:    recipeId,
				Name:        "updated first step",
				Description: "description",
				StepNum:     2,
			},
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
					Update(context.Background(), gomock.Any()).
					Return(nil)
			},
			wantErr: false,
		}, // успешное обновление шага (с изменением номера)
		{
			name: "пустое имя шага",
			step: &domain.RecipeStep{
				ID:          stepId,
				RecipeID:    recipeId,
				Name:        "",
				Description: "empty step name",
				StepNum:     1,
			},
			wantErr: true,
			errStr:  errors.New("updating recipe step: empty name"),
		}, // пустое имя шага
		{
			name: "пустое описание шага",
			step: &domain.RecipeStep{
				ID:          stepId,
				RecipeID:    recipeId,
				Name:        "empty step description",
				Description: "",
				StepNum:     1,
			},
			wantErr: true,
			errStr:  errors.New("updating recipe step: empty description"),
		}, // пустое описание шага
		{
			name: "невалидный номер шага - 0",
			step: &domain.RecipeStep{
				ID:          stepId,
				RecipeID:    recipeId,
				Name:        "first recipe step",
				Description: "description",
				StepNum:     0,
			},
			wantErr: true,
			errStr:  errors.New("updating recipe step: negative or zero step num"),
		}, // невалидный номер шага - 0
		{
			name: "невалидный номер шага - отрицательный",
			step: &domain.RecipeStep{
				ID:          stepId,
				RecipeID:    recipeId,
				Name:        "first recipe step",
				Description: "description",
				StepNum:     -1,
			},
			wantErr: true,
			errStr:  errors.New("updating recipe step: negative or zero step num"),
		}, // невалидный номер шага - отрицательный
		{
			name: "ошибка выполнения запроса в репозитории (update)",
			step: &domain.RecipeStep{
				ID:          stepId,
				RecipeID:    recipeId,
				Name:        "updated first step",
				Description: "description",
				StepNum:     2,
			},
			beforeTest: func(recipeStepRepo mocks.MockIRecipeStepRepository) {
				recipeStepRepo.EXPECT().
					Update(context.Background(), &domain.RecipeStep{
						ID:          stepId,
						RecipeID:    recipeId,
						Name:        "updated first step",
						Description: "description",
						StepNum:     2,
					}).Return(fmt.Errorf("updating step err"))
			},
			wantErr: true,
			errStr:  errors.New("updating recipe step: updating step err"),
		}, // ошибка выполнения запроса в репозитории (update)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeStepRepo)
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
