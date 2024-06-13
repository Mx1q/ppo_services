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

func TestRecipeService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeRepo := mocks.NewMockIRecipeRepository(ctrl)
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
	svc := services.NewRecipeService(recipeRepo, logger)

	recipeId := uuid.New()
	saladId := uuid.New()

	tests := []struct {
		name       string
		recipe     *domain.Recipe
		beforeTest func(recipeRepo mocks.MockIRecipeRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное создание рецепта",
			recipe: &domain.Recipe{
				ID:               recipeId,
				SaladID:          saladId,
				Status:           0,
				NumberOfServings: 1,
				TimeToCook:       1,
			},
			beforeTest: func(recipeRepo mocks.MockIRecipeRepository) {
				recipeRepo.EXPECT().
					Create(context.Background(), &domain.Recipe{
						ID:               recipeId,
						SaladID:          saladId,
						Status:           0,
						NumberOfServings: 1,
						TimeToCook:       1,
					}).
					Return(uuid.Nil, nil)
			},
			wantErr: false,
		}, // успешное создание рецепта
		{
			name: "ошибка - число порций 0",
			recipe: &domain.Recipe{
				ID:               recipeId,
				SaladID:          saladId,
				Status:           0,
				NumberOfServings: 0,
				TimeToCook:       1,
			},
			wantErr: true,
			errStr:  errors.New("creating recipe: negative or zero number of servings"),
		}, // ошибка - число порций 0
		{
			name: "ошибка - число порций <0",
			recipe: &domain.Recipe{
				ID:               recipeId,
				SaladID:          saladId,
				Status:           0,
				NumberOfServings: -1,
				TimeToCook:       1,
			},
			wantErr: true,
			errStr:  errors.New("creating recipe: negative or zero number of servings"),
		}, // ошибка - число порций <0
		{
			name: "ошибка - время приготовления 0",
			recipe: &domain.Recipe{
				ID:               recipeId,
				SaladID:          saladId,
				Status:           0,
				NumberOfServings: 1,
				TimeToCook:       0,
			},
			wantErr: true,
			errStr:  errors.New("creating recipe: negative or zero time to cook"),
		}, // ошибка - время приготовления 0
		{
			name: "ошибка - время приготовления <0",
			recipe: &domain.Recipe{
				ID:               recipeId,
				SaladID:          saladId,
				Status:           0,
				NumberOfServings: 1,
				TimeToCook:       -1,
			},
			wantErr: true,
			errStr:  errors.New("creating recipe: negative or zero time to cook"),
		}, // ошибка - время приготовления <0
		{
			name: "ошибка выполнения запроса в репозитории",
			recipe: &domain.Recipe{
				ID:               recipeId,
				SaladID:          saladId,
				Status:           0,
				NumberOfServings: 1,
				TimeToCook:       1,
			},
			beforeTest: func(recipeRepo mocks.MockIRecipeRepository) {
				recipeRepo.EXPECT().
					Create(context.Background(), &domain.Recipe{
						ID:               recipeId,
						SaladID:          saladId,
						Status:           0,
						NumberOfServings: 1,
						TimeToCook:       1,
					}).
					Return(uuid.Nil, fmt.Errorf("creating recipe err"))
			},
			wantErr: true,
			errStr:  errors.New("creating recipe: creating recipe err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeRepo)
			}

			_, err := svc.Create(context.Background(), tt.recipe)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestRecipeService_DeleteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeRepo := mocks.NewMockIRecipeRepository(ctrl)
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
	svc := services.NewRecipeService(recipeRepo, logger)

	recipeId := uuid.New()

	tests := []struct {
		name       string
		recipeId   uuid.UUID
		beforeTest func(recipeRepo mocks.MockIRecipeRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name:     "успешное удаление рецепта",
			recipeId: recipeId,
			beforeTest: func(recipeRepo mocks.MockIRecipeRepository) {
				recipeRepo.EXPECT().
					DeleteById(context.Background(), recipeId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное удаление рецепта
		{
			name:     "ошибка выполнения запроса в репозитории",
			recipeId: recipeId,
			beforeTest: func(recipeRepo mocks.MockIRecipeRepository) {
				recipeRepo.EXPECT().
					DeleteById(context.Background(), recipeId).
					Return(fmt.Errorf("deleting recipe err"))
			},
			wantErr: true,
			errStr:  errors.New("deleting recipe by id: deleting recipe err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeRepo)
			}

			err := svc.DeleteById(context.Background(), tt.recipeId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestRecipeService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeRepo := mocks.NewMockIRecipeRepository(ctrl)
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
	svc := services.NewRecipeService(recipeRepo, logger)

	filter := &domain.RecipeFilter{
		AvailableIngredients: nil,
		MinRate:              0,
		SaladTypes:           nil,
	}
	page := 1

	tests := []struct {
		name       string
		filter     *domain.RecipeFilter
		page       int
		beforeTest func(recipeRepo mocks.MockIRecipeRepository)
		expected   []*domain.Recipe
		wantErr    bool
		errStr     error
	}{
		{
			name:   "успешное получение всех рецептов",
			filter: filter,
			page:   page,
			beforeTest: func(recipeRepo mocks.MockIRecipeRepository) {
				recipeRepo.EXPECT().
					GetAll(context.Background(), filter, page).
					Return([]*domain.Recipe{
						{
							ID:               uuid.UUID{1},
							SaladID:          uuid.UUID{11},
							Status:           0,
							NumberOfServings: 1,
							TimeToCook:       1,
						},
						{
							ID:               uuid.UUID{2},
							SaladID:          uuid.UUID{22},
							Status:           0,
							NumberOfServings: 2,
							TimeToCook:       2,
						},
						{
							ID:               uuid.UUID{3},
							SaladID:          uuid.UUID{33},
							Status:           0,
							NumberOfServings: 3,
							TimeToCook:       3,
						},
						{
							ID:               uuid.UUID{4},
							SaladID:          uuid.UUID{44},
							Status:           0,
							NumberOfServings: 4,
							TimeToCook:       4,
						},
					}, nil)
			},
			expected: []*domain.Recipe{
				{
					ID:               uuid.UUID{1},
					SaladID:          uuid.UUID{11},
					Status:           0,
					NumberOfServings: 1,
					TimeToCook:       1,
				},
				{
					ID:               uuid.UUID{2},
					SaladID:          uuid.UUID{22},
					Status:           0,
					NumberOfServings: 2,
					TimeToCook:       2,
				},
				{
					ID:               uuid.UUID{3},
					SaladID:          uuid.UUID{33},
					Status:           0,
					NumberOfServings: 3,
					TimeToCook:       3,
				},
				{
					ID:               uuid.UUID{4},
					SaladID:          uuid.UUID{44},
					Status:           0,
					NumberOfServings: 4,
					TimeToCook:       4,
				},
			},
			wantErr: false,
		}, // успешное получение всех рецептов
		{
			name:   "ошибка выполнения запроса в репозитории",
			filter: filter,
			page:   page,
			beforeTest: func(recipeRepo mocks.MockIRecipeRepository) {
				recipeRepo.EXPECT().
					GetAll(context.Background(), filter, page).
					Return(nil, fmt.Errorf("getting recipes err"))
			},
			wantErr: true,
			errStr:  errors.New("getting all recipes: getting recipes err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeRepo)
			}

			recipes, err := svc.GetAll(context.Background(), tt.filter, tt.page)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, recipes)
			}
		})
	}
}

func TestRecipeService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeRepo := mocks.NewMockIRecipeRepository(ctrl)
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
	svc := services.NewRecipeService(recipeRepo, logger)

	recipeId := uuid.New()

	tests := []struct {
		name       string
		recipeId   uuid.UUID
		beforeTest func(recipeRepo mocks.MockIRecipeRepository)
		expected   *domain.Recipe
		wantErr    bool
		errStr     error
	}{
		{
			name:     "успешное получение рецепта",
			recipeId: recipeId,
			beforeTest: func(recipeRepo mocks.MockIRecipeRepository) {
				recipeRepo.EXPECT().
					GetById(context.Background(), recipeId).
					Return(&domain.Recipe{
						ID:               recipeId,
						SaladID:          uuid.UUID{1},
						Status:           0,
						NumberOfServings: 1,
						TimeToCook:       1,
					}, nil)
			},
			expected: &domain.Recipe{
				ID:               recipeId,
				SaladID:          uuid.UUID{1},
				Status:           0,
				NumberOfServings: 1,
				TimeToCook:       1,
			},
			wantErr: false,
		}, // успешное получение рецепта
		{
			name:     "ошибка выполнения запроса в репозитории",
			recipeId: recipeId,
			beforeTest: func(recipeRepo mocks.MockIRecipeRepository) {
				recipeRepo.EXPECT().
					GetById(context.Background(), recipeId).
					Return(nil, fmt.Errorf("getting recipe err"))
			},
			wantErr: true,
			errStr:  errors.New("getting recipe by id: getting recipe err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeRepo)
			}

			recipe, err := svc.GetById(context.Background(), tt.recipeId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, recipe)
			}
		})
	}
}

func TestRecipeService_GetBySaladId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeRepo := mocks.NewMockIRecipeRepository(ctrl)
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
	svc := services.NewRecipeService(recipeRepo, logger)

	saladId := uuid.New()

	tests := []struct {
		name       string
		saladId    uuid.UUID
		beforeTest func(recipeRepo mocks.MockIRecipeRepository)
		expected   *domain.Recipe
		wantErr    bool
		errStr     error
	}{
		{
			name:    "успешное получение рецепта",
			saladId: saladId,
			beforeTest: func(recipeRepo mocks.MockIRecipeRepository) {
				recipeRepo.EXPECT().
					GetBySaladId(context.Background(), saladId).
					Return(&domain.Recipe{
						ID:               uuid.UUID{1},
						SaladID:          saladId,
						Status:           0,
						NumberOfServings: 1,
						TimeToCook:       1,
					}, nil)
			},
			expected: &domain.Recipe{
				ID:               uuid.UUID{1},
				SaladID:          saladId,
				Status:           0,
				NumberOfServings: 1,
				TimeToCook:       1,
			},
			wantErr: false,
		}, // успешное получение рецепта
		{
			name:    "ошибка выполнения запроса в репозитории",
			saladId: saladId,
			beforeTest: func(recipeRepo mocks.MockIRecipeRepository) {
				recipeRepo.EXPECT().
					GetBySaladId(context.Background(), saladId).
					Return(nil, fmt.Errorf("getting recipe err"))
			},
			wantErr: true,
			errStr:  errors.New("getting recipe by salad id: getting recipe err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeRepo)
			}

			recipe, err := svc.GetBySaladId(context.Background(), tt.saladId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, recipe)
			}
		})
	}
}

func TestRecipeService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	recipeRepo := mocks.NewMockIRecipeRepository(ctrl)
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
	svc := services.NewRecipeService(recipeRepo, logger)

	recipeId := uuid.New()
	saladId := uuid.New()

	tests := []struct {
		name       string
		recipe     *domain.Recipe
		beforeTest func(recipeRepo mocks.MockIRecipeRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное обновление рецепта",
			recipe: &domain.Recipe{
				ID:               recipeId,
				SaladID:          saladId,
				Status:           0,
				NumberOfServings: 1,
				TimeToCook:       1,
			},
			beforeTest: func(recipeRepo mocks.MockIRecipeRepository) {
				recipeRepo.EXPECT().
					Update(context.Background(), &domain.Recipe{
						ID:               recipeId,
						SaladID:          saladId,
						Status:           0,
						NumberOfServings: 1,
						TimeToCook:       1,
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное обновление рецепта
		{
			name: "ошибка - число порций 0",
			recipe: &domain.Recipe{
				ID:               recipeId,
				SaladID:          saladId,
				Status:           0,
				NumberOfServings: 0,
				TimeToCook:       1,
			},
			wantErr: true,
			errStr:  errors.New("updating recipe: negative or zero number of servings"),
		}, // ошибка - число порций 0
		{
			name: "ошибка - число порций <0",
			recipe: &domain.Recipe{
				ID:               recipeId,
				SaladID:          saladId,
				Status:           0,
				NumberOfServings: -1,
				TimeToCook:       1,
			},
			wantErr: true,
			errStr:  errors.New("updating recipe: negative or zero number of servings"),
		}, // ошибка - число порций <0
		{
			name: "ошибка - время приготовления 0",
			recipe: &domain.Recipe{
				ID:               recipeId,
				SaladID:          saladId,
				Status:           0,
				NumberOfServings: 1,
				TimeToCook:       0,
			},
			wantErr: true,
			errStr:  errors.New("updating recipe: negative or zero time to cook"),
		}, // ошибка - время приготовления 0
		{
			name: "ошибка - время приготовления <0",
			recipe: &domain.Recipe{
				ID:               recipeId,
				SaladID:          saladId,
				Status:           0,
				NumberOfServings: 1,
				TimeToCook:       -1,
			},
			wantErr: true,
			errStr:  errors.New("updating recipe: negative or zero time to cook"),
		}, // ошибка - время приготовления <0
		{
			name: "ошибка выполнения запроса в репозитории",
			recipe: &domain.Recipe{
				ID:               recipeId,
				SaladID:          saladId,
				Status:           0,
				NumberOfServings: 1,
				TimeToCook:       1,
			},
			beforeTest: func(recipeRepo mocks.MockIRecipeRepository) {
				recipeRepo.EXPECT().
					Update(context.Background(), &domain.Recipe{
						ID:               recipeId,
						SaladID:          saladId,
						Status:           0,
						NumberOfServings: 1,
						TimeToCook:       1,
					}).
					Return(fmt.Errorf("updating recipe err"))
			},
			wantErr: true,
			errStr:  errors.New("updating recipe: updating recipe err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*recipeRepo)
			}

			err := svc.Update(context.Background(), tt.recipe)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
