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

func TestSaladService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladRepo := mocks.NewMockISaladRepository(ctrl)
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
	svc := services.NewSaladService(saladRepo, logger)

	saladId := uuid.New()
	userId := uuid.New()

	tests := []struct {
		name       string
		salad      *domain.Salad
		beforeTest func(saladRepo mocks.MockISaladRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное создание салата",
			salad: &domain.Salad{
				ID:          saladId,
				AuthorID:    userId,
				Name:        "salad",
				Description: "",
			},
			beforeTest: func(saladRepo mocks.MockISaladRepository) {
				saladRepo.EXPECT().
					Create(context.Background(), &domain.Salad{
						ID:          saladId,
						AuthorID:    userId,
						Name:        "salad",
						Description: "",
					}).
					Return(uuid.Nil, nil)
			},
			wantErr: false,
		}, // успешное создание салата
		{
			name: "пустое название салата",
			salad: &domain.Salad{
				ID:          saladId,
				AuthorID:    userId,
				Name:        "",
				Description: "",
			},
			beforeTest: func(saladRepo mocks.MockISaladRepository) {
				saladRepo.EXPECT().
					Create(context.Background(), &domain.Salad{
						ID:          saladId,
						AuthorID:    userId,
						Name:        "",
						Description: "",
					}).
					Return(uuid.Nil, nil).
					AnyTimes()
			},
			wantErr: true,
			errStr:  errors.New("empty salad name"),
		}, // пустое название салата
		{
			name: "ошибка выполнения запроса в репозитории",
			salad: &domain.Salad{
				ID:          saladId,
				AuthorID:    userId,
				Name:        "salad",
				Description: "",
			},
			beforeTest: func(saladRepo mocks.MockISaladRepository) {
				saladRepo.EXPECT().
					Create(context.Background(), &domain.Salad{
						ID:          saladId,
						AuthorID:    userId,
						Name:        "salad",
						Description: "",
					}).
					Return(uuid.Nil, fmt.Errorf("creating salad err"))
			},
			wantErr: true,
			errStr:  errors.New("creating salad: creating salad err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladRepo)
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

func TestSaladService_DeleteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladRepo := mocks.NewMockISaladRepository(ctrl)
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
	svc := services.NewSaladService(saladRepo, logger)

	saladId := uuid.New()

	tests := []struct {
		name       string
		id         uuid.UUID
		beforeTest func(saladRepo mocks.MockISaladRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное удаление",
			id:   saladId,
			beforeTest: func(saladRepo mocks.MockISaladRepository) {
				saladRepo.EXPECT().
					DeleteById(context.Background(), saladId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное удаление
		{
			name: "ошибка выполнения запроса в репозитории",
			id:   saladId,
			beforeTest: func(saladRepo mocks.MockISaladRepository) {
				saladRepo.EXPECT().
					DeleteById(context.Background(), saladId).
					Return(fmt.Errorf("deleting salad err"))
			},
			wantErr: true,
			errStr:  errors.New("deleting salad by id: deleting salad err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladRepo)
			}

			err := svc.DeleteById(context.Background(), tt.id)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestSaladService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladRepo := mocks.NewMockISaladRepository(ctrl)
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
	svc := services.NewSaladService(saladRepo, logger)

	filter := &domain.RecipeFilter{
		AvailableIngredients: nil,
		MinRate:              0,
		SaladTypes:           nil,
	}
	page := 1

	tests := []struct {
		name       string
		page       int
		beforeTest func(saladRepo mocks.MockISaladRepository)
		expected   []*domain.Salad
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное получение всех салатов",
			page: page,
			beforeTest: func(saladRepo mocks.MockISaladRepository) {
				saladRepo.EXPECT().
					GetAll(context.Background(), filter, page).
					Return([]*domain.Salad{
						{
							ID:          uuid.UUID{1},
							AuthorID:    uuid.UUID{11},
							Name:        "salad1",
							Description: "",
						},
						{
							ID:          uuid.UUID{2},
							AuthorID:    uuid.UUID{22},
							Name:        "salad2",
							Description: "",
						},
						{
							ID:          uuid.UUID{3},
							AuthorID:    uuid.UUID{33},
							Name:        "salad3",
							Description: "",
						},
						{
							ID:          uuid.UUID{4},
							AuthorID:    uuid.UUID{44},
							Name:        "salad1",
							Description: "",
						},
					}, 4, nil)
			},
			expected: []*domain.Salad{
				{
					ID:          uuid.UUID{1},
					AuthorID:    uuid.UUID{11},
					Name:        "salad1",
					Description: "",
				},
				{
					ID:          uuid.UUID{2},
					AuthorID:    uuid.UUID{22},
					Name:        "salad2",
					Description: "",
				},
				{
					ID:          uuid.UUID{3},
					AuthorID:    uuid.UUID{33},
					Name:        "salad3",
					Description: "",
				},
				{
					ID:          uuid.UUID{4},
					AuthorID:    uuid.UUID{44},
					Name:        "salad1",
					Description: "",
				},
			},
			wantErr: false,
		}, // успешное получение всех салатов
		{
			name: "ошибка выполнения запроса в репозитории",
			page: page,
			beforeTest: func(saladRepo mocks.MockISaladRepository) {
				saladRepo.EXPECT().
					GetAll(context.Background(), filter, page).
					Return(nil, 0, fmt.Errorf("getting all salads err"))
			},
			wantErr: true,
			errStr:  errors.New("getting all salads: getting all salads err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladRepo)
			}

			salads, _, err := svc.GetAll(context.Background(), filter, tt.page)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, salads, tt.expected)
			}
		})
	}
}

func TestSaladService_GetAllRatedByUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladRepo := mocks.NewMockISaladRepository(ctrl)
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
	svc := services.NewSaladService(saladRepo, logger)

	userId := uuid.New()
	page := 1

	tests := []struct {
		name       string
		userId     uuid.UUID
		page       int
		beforeTest func(saladRepo mocks.MockISaladRepository)
		expected   []*domain.Salad
		wantErr    bool
		errStr     error
	}{
		{
			name:   "успешное получение салатов, оцененных пользователем",
			userId: userId,
			page:   page,
			beforeTest: func(saladRepo mocks.MockISaladRepository) {
				saladRepo.EXPECT().
					GetAllRatedByUser(context.Background(), userId, page).
					Return([]*domain.Salad{
						{
							ID:          uuid.UUID{1},
							AuthorID:    uuid.UUID{11},
							Name:        "salad1",
							Description: "",
						},
						{
							ID:          uuid.UUID{2},
							AuthorID:    uuid.UUID{22},
							Name:        "salad2",
							Description: "",
						},
						{
							ID:          uuid.UUID{3},
							AuthorID:    uuid.UUID{33},
							Name:        "salad3",
							Description: "",
						},
						{
							ID:          uuid.UUID{4},
							AuthorID:    uuid.UUID{44},
							Name:        "salad1",
							Description: "",
						},
					}, 1, nil)
			},
			expected: []*domain.Salad{
				{
					ID:          uuid.UUID{1},
					AuthorID:    uuid.UUID{11},
					Name:        "salad1",
					Description: "",
				},
				{
					ID:          uuid.UUID{2},
					AuthorID:    uuid.UUID{22},
					Name:        "salad2",
					Description: "",
				},
				{
					ID:          uuid.UUID{3},
					AuthorID:    uuid.UUID{33},
					Name:        "salad3",
					Description: "",
				},
				{
					ID:          uuid.UUID{4},
					AuthorID:    uuid.UUID{44},
					Name:        "salad1",
					Description: "",
				},
			},
			wantErr: false,
		}, // успешное получение салатов, оцененных пользователем
		{
			name:   "ошибка выполнения запроса в репозитории",
			userId: userId,
			page:   page,
			beforeTest: func(saladRepo mocks.MockISaladRepository) {
				saladRepo.EXPECT().
					GetAllRatedByUser(context.Background(), userId, page).
					Return(nil, 0, fmt.Errorf("getting all salads err"))
			},
			wantErr: true,
			errStr:  errors.New("getting all salads rated by user: getting all salads err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladRepo)
			}

			salads, _, err := svc.GetAllRatedByUser(context.Background(), tt.userId, tt.page)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, salads, tt.expected)
			}
		})
	}
}

func TestSaladService_GetAllByUserId(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladRepo := mocks.NewMockISaladRepository(ctrl)
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
	svc := services.NewSaladService(saladRepo, logger)

	userId := uuid.New()

	tests := []struct {
		name       string
		beforeTest func(saladRepo mocks.MockISaladRepository)
		userId     uuid.UUID
		expected   []*domain.Salad
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное получение всех салатов",
			beforeTest: func(saladRepo mocks.MockISaladRepository) {
				saladRepo.EXPECT().
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
			userId: userId,
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
		}, // успешное получение всех салатов
		{
			name: "ошибка выполнения запроса в репозитории",
			beforeTest: func(saladRepo mocks.MockISaladRepository) {
				saladRepo.EXPECT().
					GetAllByUserId(context.Background(), userId).
					Return(nil, fmt.Errorf("getting all salads by uid err"))
			},
			userId:  userId,
			wantErr: true,
			errStr:  errors.New("getting all salads by author id: getting all salads by uid err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladRepo)
			}

			salads, err := svc.GetAllByUserId(context.Background(), tt.userId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, salads, tt.expected)
			}
		})
	}
}

func TestSaladService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladRepo := mocks.NewMockISaladRepository(ctrl)
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
	svc := services.NewSaladService(saladRepo, logger)

	saladId := uuid.New()

	tests := []struct {
		name       string
		beforeTest func(saladRepo mocks.MockISaladRepository)
		saladId    uuid.UUID
		expected   *domain.Salad
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное получение всех салатов",
			beforeTest: func(saladRepo mocks.MockISaladRepository) {
				saladRepo.EXPECT().
					GetById(context.Background(), saladId).
					Return(&domain.Salad{
						ID:          uuid.UUID{1},
						AuthorID:    uuid.UUID{11},
						Name:        "salad1",
						Description: "",
					}, nil)
			},
			saladId: saladId,
			expected: &domain.Salad{
				ID:          uuid.UUID{1},
				AuthorID:    uuid.UUID{11},
				Name:        "salad1",
				Description: "",
			},
			wantErr: false,
		}, // успешное получение всех салатов
		{
			name: "ошибка выполнения запроса в репозитории",
			beforeTest: func(saladRepo mocks.MockISaladRepository) {
				saladRepo.EXPECT().
					GetById(context.Background(), saladId).
					Return(nil, fmt.Errorf("getting salad by id err"))
			},
			saladId: saladId,
			wantErr: true,
			errStr:  errors.New("getting salad by id: getting salad by id err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladRepo)
			}

			salads, err := svc.GetById(context.Background(), tt.saladId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, salads, tt.expected)
			}
		})
	}
}

func TestSaladService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	saladRepo := mocks.NewMockISaladRepository(ctrl)
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
	svc := services.NewSaladService(saladRepo, logger)

	saladId := uuid.New()
	authorId := uuid.New()

	tests := []struct {
		name       string
		salad      *domain.Salad
		beforeTest func(saladRepo mocks.MockISaladRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное обновление",
			salad: &domain.Salad{
				ID:          saladId,
				AuthorID:    authorId,
				Name:        "salad",
				Description: "",
			},
			beforeTest: func(saladRepo mocks.MockISaladRepository) {
				saladRepo.EXPECT().
					Update(context.Background(), &domain.Salad{
						ID:          saladId,
						AuthorID:    authorId,
						Name:        "salad",
						Description: "",
					}).Return(nil)
			},
			wantErr: false,
		}, // успешное обновление
		{
			name: "пустое название салата",
			salad: &domain.Salad{
				ID:          saladId,
				AuthorID:    authorId,
				Name:        "",
				Description: "",
			},
			wantErr: true,
			errStr:  errors.New("updating salad: empty salad name"),
		}, // пустое название салата
		{
			name: "ошибка выполнения запроса в репозитории",
			salad: &domain.Salad{
				ID:          saladId,
				AuthorID:    authorId,
				Name:        "salad",
				Description: "",
			},
			beforeTest: func(saladRepo mocks.MockISaladRepository) {
				saladRepo.EXPECT().
					Update(context.Background(), &domain.Salad{
						ID:          saladId,
						AuthorID:    authorId,
						Name:        "salad",
						Description: "",
					}).Return(fmt.Errorf("updating salad err"))
			},
			wantErr: true,
			errStr:  errors.New("updating salad: updating salad err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*saladRepo)
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
