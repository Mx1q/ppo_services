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

func TestIngredientService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ingredientRepo := mocks.NewMockIIngredientRepository(ctrl)
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
	svc := services.NewIngredientService(ingredientRepo, logger)

	ingredientId := uuid.New()
	typeId := uuid.New()

	tests := []struct {
		name       string
		ingredient *domain.Ingredient
		beforeTest func(ingredientRepo mocks.MockIIngredientRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное создание",
			ingredient: &domain.Ingredient{
				ID:       ingredientId,
				TypeID:   typeId,
				Name:     "tomato",
				Calories: 10,
			},
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					Create(context.Background(), &domain.Ingredient{
						ID:       ingredientId,
						TypeID:   typeId,
						Name:     "tomato",
						Calories: 10,
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное создание
		{
			name: "пустое название",
			ingredient: &domain.Ingredient{
				ID:       ingredientId,
				TypeID:   typeId,
				Name:     "",
				Calories: 10,
			},
			wantErr: true,
			errStr:  errors.New("creating ingredient: empty name"),
		}, // пустое название
		{
			name: "число калорий  <0",
			ingredient: &domain.Ingredient{
				ID:       ingredientId,
				TypeID:   typeId,
				Name:     "tomato",
				Calories: -1,
			},
			wantErr: true,
			errStr:  errors.New("creating ingredient: negative calories"),
		}, // число калорий <0
		{
			name: "ошибка выполнения запроса в репозитории",
			ingredient: &domain.Ingredient{
				ID:       ingredientId,
				TypeID:   typeId,
				Name:     "tomato",
				Calories: 10,
			},
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					Create(context.Background(), &domain.Ingredient{
						ID:       ingredientId,
						TypeID:   typeId,
						Name:     "tomato",
						Calories: 10,
					}).
					Return(fmt.Errorf("creating ingredient err"))
			},
			wantErr: true,
			errStr:  errors.New("creating ingredient: creating ingredient err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*ingredientRepo)
			}

			err := svc.Create(context.Background(), tt.ingredient)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestIngredientService_DeleteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ingredientRepo := mocks.NewMockIIngredientRepository(ctrl)
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
	svc := services.NewIngredientService(ingredientRepo, logger)

	ingredientId := uuid.New()

	tests := []struct {
		name         string
		ingredientId uuid.UUID
		beforeTest   func(ingredientRepo mocks.MockIIngredientRepository)
		wantErr      bool
		errStr       error
	}{
		{
			name:         "успешное удаление",
			ingredientId: ingredientId,
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					DeleteById(context.Background(), ingredientId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное удаление
		{
			name:         "ошибка выполнения запроса в репозитории",
			ingredientId: ingredientId,
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					DeleteById(context.Background(), ingredientId).
					Return(fmt.Errorf("deleting ingredient err"))
			},
			wantErr: true,
			errStr:  errors.New("deleting ingredient by id: deleting ingredient err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*ingredientRepo)
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

func TestIngredientService_Link(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ingredientRepo := mocks.NewMockIIngredientRepository(ctrl)
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
	svc := services.NewIngredientService(ingredientRepo, logger)

	recipeId := uuid.New()
	ingredientId := uuid.New()

	tests := []struct {
		name         string
		recipeId     uuid.UUID
		ingredientId uuid.UUID
		beforeTest   func(ingredientRepo mocks.MockIIngredientRepository)
		wantErr      bool
		errStr       error
	}{
		{
			name:         "успешное добавление",
			recipeId:     recipeId,
			ingredientId: ingredientId,
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					Link(context.Background(), recipeId, ingredientId).
					Return(uuid.New(), nil)
			},
			wantErr: false,
		}, // успешное добавление
		{
			name:         "ошибка выполнения запроса в репозитории",
			recipeId:     recipeId,
			ingredientId: ingredientId,
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					Link(context.Background(), recipeId, ingredientId).
					Return(uuid.Nil, fmt.Errorf("linking ingredient err"))
			},
			wantErr: true,
			errStr:  errors.New("linking ingredient to recipe: linking ingredient err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*ingredientRepo)
			}

			_, err := svc.Link(context.Background(), tt.recipeId, tt.ingredientId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestIngredientService_Unlink(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ingredientRepo := mocks.NewMockIIngredientRepository(ctrl)
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
	svc := services.NewIngredientService(ingredientRepo, logger)

	recipeId := uuid.New()
	ingredientId := uuid.New()

	tests := []struct {
		name         string
		recipeId     uuid.UUID
		ingredientId uuid.UUID
		beforeTest   func(ingredientRepo mocks.MockIIngredientRepository)
		wantErr      bool
		errStr       error
	}{
		{
			name:         "успешное удаление",
			recipeId:     recipeId,
			ingredientId: ingredientId,
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					Unlink(context.Background(), recipeId, ingredientId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное добавление
		{
			name:         "ошибка выполнения запроса в репозитории",
			recipeId:     recipeId,
			ingredientId: ingredientId,
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					Unlink(context.Background(), recipeId, ingredientId).
					Return(fmt.Errorf("unlinking ingredient err"))
			},
			wantErr: true,
			errStr:  errors.New("unlinking ingredient from recipe: unlinking ingredient err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*ingredientRepo)
			}

			err := svc.Unlink(context.Background(), tt.recipeId, tt.ingredientId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestIngredientService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ingredientRepo := mocks.NewMockIIngredientRepository(ctrl)
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
	svc := services.NewIngredientService(ingredientRepo, logger)

	page := 1

	tests := []struct {
		name       string
		page       int
		beforeTest func(ingredientRepo mocks.MockIIngredientRepository)
		expected   []*domain.Ingredient
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное получение",
			page: page,
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					GetAll(context.Background(), page).
					Return([]*domain.Ingredient{
						{
							ID:       uuid.UUID{1},
							TypeID:   uuid.UUID{11},
							Name:     "tomato",
							Calories: 10,
						},
						{
							ID:       uuid.UUID{2},
							TypeID:   uuid.UUID{2},
							Name:     "salad",
							Calories: 10,
						},
						{
							ID:       uuid.UUID{3},
							TypeID:   uuid.UUID{33},
							Name:     "cabbage",
							Calories: 10,
						},
						{
							ID:       uuid.UUID{4},
							TypeID:   uuid.UUID{44},
							Name:     "cucumber",
							Calories: 10,
						},
					}, 1, nil)
			},
			expected: []*domain.Ingredient{
				{
					ID:       uuid.UUID{1},
					TypeID:   uuid.UUID{11},
					Name:     "tomato",
					Calories: 10,
				},
				{
					ID:       uuid.UUID{2},
					TypeID:   uuid.UUID{2},
					Name:     "salad",
					Calories: 10,
				},
				{
					ID:       uuid.UUID{3},
					TypeID:   uuid.UUID{33},
					Name:     "cabbage",
					Calories: 10,
				},
				{
					ID:       uuid.UUID{4},
					TypeID:   uuid.UUID{44},
					Name:     "cucumber",
					Calories: 10,
				},
			},
			wantErr: false,
		}, // успешное получение
		{
			name: "ошибка выполнения запроса в репозитории",
			page: page,
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					GetAll(context.Background(), page).
					Return(nil, 0, fmt.Errorf("getting ingredients err"))
			},
			wantErr: true,
			errStr:  errors.New("getting all ingredients: getting ingredients err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*ingredientRepo)
			}

			ingredients, _, err := svc.GetAll(context.Background(), tt.page)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, ingredients)
			}
		})
	}
}

func TestIngredientService_GetAllIngredientsByTypeId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ingredientRepo := mocks.NewMockIIngredientRepository(ctrl)
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
	svc := services.NewIngredientService(ingredientRepo, logger)

	recipeId := uuid.New()

	tests := []struct {
		name       string
		recipeId   uuid.UUID
		beforeTest func(ingredientRepo mocks.MockIIngredientRepository)
		expected   []*domain.Ingredient
		wantErr    bool
		errStr     error
	}{
		{
			name:     "успешное получение",
			recipeId: recipeId,
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					GetAllByRecipeId(context.Background(), recipeId).
					Return([]*domain.Ingredient{
						{
							ID:       uuid.UUID{1},
							TypeID:   uuid.UUID{11},
							Name:     "tomato",
							Calories: 10,
						},
						{
							ID:       uuid.UUID{2},
							TypeID:   uuid.UUID{2},
							Name:     "salad",
							Calories: 10,
						},
						{
							ID:       uuid.UUID{3},
							TypeID:   uuid.UUID{33},
							Name:     "cabbage",
							Calories: 10,
						},
						{
							ID:       uuid.UUID{4},
							TypeID:   uuid.UUID{44},
							Name:     "cucumber",
							Calories: 10,
						},
					}, nil)
			},
			expected: []*domain.Ingredient{
				{
					ID:       uuid.UUID{1},
					TypeID:   uuid.UUID{11},
					Name:     "tomato",
					Calories: 10,
				},
				{
					ID:       uuid.UUID{2},
					TypeID:   uuid.UUID{2},
					Name:     "salad",
					Calories: 10,
				},
				{
					ID:       uuid.UUID{3},
					TypeID:   uuid.UUID{33},
					Name:     "cabbage",
					Calories: 10,
				},
				{
					ID:       uuid.UUID{4},
					TypeID:   uuid.UUID{44},
					Name:     "cucumber",
					Calories: 10,
				},
			},
			wantErr: false,
		}, // успешное получение
		{
			name:     "ошибка выполнения запроса в репозитории",
			recipeId: recipeId,
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					GetAllByRecipeId(context.Background(), recipeId).
					Return(nil, fmt.Errorf("getting ingredients err"))
			},
			wantErr: true,
			errStr:  errors.New("getting all ingredients by recipe id: getting ingredients err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*ingredientRepo)
			}

			ingredients, err := svc.GetAllByRecipeId(context.Background(), tt.recipeId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, ingredients)
			}
		})
	}
}

func TestIngredientService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ingredientRepo := mocks.NewMockIIngredientRepository(ctrl)
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
	svc := services.NewIngredientService(ingredientRepo, logger)

	ingredientId := uuid.New()

	tests := []struct {
		name         string
		ingredientId uuid.UUID
		beforeTest   func(ingredientRepo mocks.MockIIngredientRepository)
		expected     *domain.Ingredient
		wantErr      bool
		errStr       error
	}{
		{
			name:         "успешное получение",
			ingredientId: ingredientId,
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					GetById(context.Background(), ingredientId).
					Return(&domain.Ingredient{
						ID:       uuid.UUID{1},
						TypeID:   uuid.UUID{11},
						Name:     "tomato",
						Calories: 10,
					}, nil)
			},
			expected: &domain.Ingredient{
				ID:       uuid.UUID{1},
				TypeID:   uuid.UUID{11},
				Name:     "tomato",
				Calories: 10,
			},
			wantErr: false,
		}, // успешное получение
		{
			name:         "ошибка выполнения запроса в репозитории",
			ingredientId: ingredientId,
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					GetById(context.Background(), ingredientId).
					Return(nil, fmt.Errorf("getting ingredient err"))
			},
			wantErr: true,
			errStr:  errors.New("getting ingredient by id: getting ingredient err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*ingredientRepo)
			}

			ingredient, err := svc.GetById(context.Background(), tt.ingredientId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, ingredient)
			}
		})
	}
}

func TestIngredientService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ingredientRepo := mocks.NewMockIIngredientRepository(ctrl)
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
	svc := services.NewIngredientService(ingredientRepo, logger)

	ingredientId := uuid.New()
	typeId := uuid.New()

	tests := []struct {
		name       string
		ingredient *domain.Ingredient
		beforeTest func(ingredientRepo mocks.MockIIngredientRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное создание",
			ingredient: &domain.Ingredient{
				ID:       ingredientId,
				TypeID:   typeId,
				Name:     "tomato",
				Calories: 10,
			},
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					Update(context.Background(), &domain.Ingredient{
						ID:       ingredientId,
						TypeID:   typeId,
						Name:     "tomato",
						Calories: 10,
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное обновление
		{
			name: "пустое название",
			ingredient: &domain.Ingredient{
				ID:       ingredientId,
				TypeID:   typeId,
				Name:     "",
				Calories: 10,
			},
			wantErr: true,
			errStr:  errors.New("updating ingredient: empty name"),
		}, // пустое название
		{
			name: "число калорий  <0",
			ingredient: &domain.Ingredient{
				ID:       ingredientId,
				TypeID:   typeId,
				Name:     "tomato",
				Calories: -1,
			},
			wantErr: true,
			errStr:  errors.New("updating ingredient: negative calories"),
		}, // число калорий <0
		{
			name: "ошибка выполнения запроса в репозитории",
			ingredient: &domain.Ingredient{
				ID:       ingredientId,
				TypeID:   typeId,
				Name:     "tomato",
				Calories: 10,
			},
			beforeTest: func(ingredientRepo mocks.MockIIngredientRepository) {
				ingredientRepo.EXPECT().
					Update(context.Background(), &domain.Ingredient{
						ID:       ingredientId,
						TypeID:   typeId,
						Name:     "tomato",
						Calories: 10,
					}).
					Return(fmt.Errorf("updating ingredient err"))
			},
			wantErr: true,
			errStr:  errors.New("updating ingredient: updating ingredient err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*ingredientRepo)
			}

			err := svc.Update(context.Background(), tt.ingredient)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
