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

func TestCommentService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepo := mocks.NewMockICommentRepository(ctrl)
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
	svc := services.NewCommentService(commentRepo, logger)

	commentId := uuid.New()
	authorId := uuid.New()
	saladId := uuid.New()

	tests := []struct {
		name       string
		comment    *domain.Comment
		beforeTest func(commentRepo mocks.MockICommentRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное создание",
			comment: &domain.Comment{
				ID:       commentId,
				AuthorID: authorId,
				SaladID:  saladId,
				Text:     "",
				Rating:   5,
			},
			beforeTest: func(commentRepo mocks.MockICommentRepository) {
				commentRepo.EXPECT().
					Create(context.Background(), &domain.Comment{
						ID:       commentId,
						AuthorID: authorId,
						SaladID:  saladId,
						Text:     "",
						Rating:   5,
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное создание
		{
			name: "оценка меньше минимальной",
			comment: &domain.Comment{
				ID:       commentId,
				AuthorID: authorId,
				SaladID:  saladId,
				Text:     "",
				Rating:   0,
			},
			wantErr: true,
			errStr:  errors.New("creating comment: rate out of range"),
		}, // оценка меньше минимальной
		{
			name: "оценка больше максимальной",
			comment: &domain.Comment{
				ID:       commentId,
				AuthorID: authorId,
				SaladID:  saladId,
				Text:     "",
				Rating:   10,
			},
			wantErr: true,
			errStr:  errors.New("creating comment: rate out of range"),
		}, // оценка больше максимальной
		{
			name: "ошибка выполнения запроса в репозитории",
			comment: &domain.Comment{
				ID:       commentId,
				AuthorID: authorId,
				SaladID:  saladId,
				Text:     "",
				Rating:   5,
			},
			beforeTest: func(commentRepo mocks.MockICommentRepository) {
				commentRepo.EXPECT().
					Create(context.Background(), &domain.Comment{
						ID:       commentId,
						AuthorID: authorId,
						SaladID:  saladId,
						Text:     "",
						Rating:   5,
					}).
					Return(fmt.Errorf("creating comment err"))
			},
			wantErr: true,
			errStr:  errors.New("creating comment: creating comment err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*commentRepo)
			}

			err := svc.Create(context.Background(), tt.comment)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestCommentService_DeleteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepo := mocks.NewMockICommentRepository(ctrl)
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
	svc := services.NewCommentService(commentRepo, logger)

	commentId := uuid.New()

	tests := []struct {
		name       string
		commentId  uuid.UUID
		beforeTest func(commentRepo mocks.MockICommentRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name:      "успешное удаление",
			commentId: commentId,
			beforeTest: func(commentRepo mocks.MockICommentRepository) {
				commentRepo.EXPECT().
					DeleteById(context.Background(), commentId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное удаление
		{
			name:      "ошибка выполнения запроса в репозитории",
			commentId: commentId,
			beforeTest: func(commentRepo mocks.MockICommentRepository) {
				commentRepo.EXPECT().
					DeleteById(context.Background(), commentId).
					Return(fmt.Errorf("deleting comment err"))
			},
			wantErr: true,
			errStr:  errors.New("deleting comment by id: deleting comment err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*commentRepo)
			}

			err := svc.DeleteById(context.Background(), tt.commentId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestCommentService_GetAllBySaladID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepo := mocks.NewMockICommentRepository(ctrl)
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
	svc := services.NewCommentService(commentRepo, logger)

	saladId := uuid.New()
	page := 1

	tests := []struct {
		name       string
		saladId    uuid.UUID
		page       int
		beforeTest func(commentRepo mocks.MockICommentRepository)
		expected   []*domain.Comment
		wantErr    bool
		errStr     error
	}{
		{
			name:    "успешное получение",
			saladId: saladId,
			page:    page,
			beforeTest: func(commentRepo mocks.MockICommentRepository) {
				commentRepo.EXPECT().
					GetAllBySaladID(context.Background(), saladId, page).
					Return([]*domain.Comment{
						{
							ID:       uuid.UUID{1},
							AuthorID: uuid.UUID{11},
							SaladID:  saladId,
							Text:     "some text",
							Rating:   5,
						},
						{
							ID:       uuid.UUID{2},
							AuthorID: uuid.UUID{22},
							SaladID:  saladId,
							Text:     "some text",
							Rating:   5,
						},
						{
							ID:       uuid.UUID{3},
							AuthorID: uuid.UUID{33},
							SaladID:  saladId,
							Text:     "some text",
							Rating:   5,
						},
						{
							ID:       uuid.UUID{4},
							AuthorID: uuid.UUID{44},
							SaladID:  saladId,
							Text:     "some text",
							Rating:   5,
						},
					}, 1, nil)
			},
			expected: []*domain.Comment{
				{
					ID:       uuid.UUID{1},
					AuthorID: uuid.UUID{11},
					SaladID:  saladId,
					Text:     "some text",
					Rating:   5,
				},
				{
					ID:       uuid.UUID{2},
					AuthorID: uuid.UUID{22},
					SaladID:  saladId,
					Text:     "some text",
					Rating:   5,
				},
				{
					ID:       uuid.UUID{3},
					AuthorID: uuid.UUID{33},
					SaladID:  saladId,
					Text:     "some text",
					Rating:   5,
				},
				{
					ID:       uuid.UUID{4},
					AuthorID: uuid.UUID{44},
					SaladID:  saladId,
					Text:     "some text",
					Rating:   5,
				},
			},
			wantErr: false,
		}, // успешное получение
		{
			name:    "ошибка выполнения запроса в репозитории",
			saladId: saladId,
			page:    page,
			beforeTest: func(commentRepo mocks.MockICommentRepository) {
				commentRepo.EXPECT().
					GetAllBySaladID(context.Background(), saladId, page).
					Return(nil, 0, fmt.Errorf("getting comments err"))
			},
			wantErr: true,
			errStr:  errors.New("getting all comments by salad id: getting comments err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*commentRepo)
			}

			salads, _, err := svc.GetAllBySaladID(context.Background(), tt.saladId, tt.page)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, salads)
			}
		})
	}
}

func TestCommentService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepo := mocks.NewMockICommentRepository(ctrl)
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
	svc := services.NewCommentService(commentRepo, logger)

	commentId := uuid.New()

	tests := []struct {
		name       string
		commentId  uuid.UUID
		beforeTest func(commentRepo mocks.MockICommentRepository)
		expected   *domain.Comment
		wantErr    bool
		errStr     error
	}{
		{
			name:      "успешное получение",
			commentId: commentId,
			beforeTest: func(commentRepo mocks.MockICommentRepository) {
				commentRepo.EXPECT().
					GetById(context.Background(), commentId).
					Return(&domain.Comment{
						ID:       commentId,
						AuthorID: uuid.UUID{1},
						SaladID:  uuid.UUID{11},
						Text:     "",
						Rating:   5,
					}, nil)
			},
			expected: &domain.Comment{
				ID:       commentId,
				AuthorID: uuid.UUID{1},
				SaladID:  uuid.UUID{11},
				Text:     "",
				Rating:   5,
			},
			wantErr: false,
		}, // успешное получение
		{
			name:      "ошибка выполнения запроса в репозитории",
			commentId: commentId,
			beforeTest: func(commentRepo mocks.MockICommentRepository) {
				commentRepo.EXPECT().
					GetById(context.Background(), commentId).
					Return(nil, fmt.Errorf("getting comment err"))
			},
			wantErr: true,
			errStr:  errors.New("getting comment by id: getting comment err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*commentRepo)
			}

			salad, err := svc.GetById(context.Background(), tt.commentId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, salad)
			}
		})
	}
}

func TestCommentService_GetBySaladAndUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepo := mocks.NewMockICommentRepository(ctrl)
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
	svc := services.NewCommentService(commentRepo, logger)

	commentId := uuid.New()
	userId := uuid.New()
	saladId := uuid.New()

	tests := []struct {
		name       string
		userId     uuid.UUID
		saladId    uuid.UUID
		beforeTest func(commentRepo mocks.MockICommentRepository)
		expected   *domain.Comment
		wantErr    bool
		errStr     error
	}{
		{
			name:    "успешное получение",
			userId:  userId,
			saladId: saladId,
			beforeTest: func(commentRepo mocks.MockICommentRepository) {
				commentRepo.EXPECT().
					GetBySaladAndUser(context.Background(), saladId, userId).
					Return(&domain.Comment{
						ID:       commentId,
						AuthorID: uuid.UUID{1},
						SaladID:  uuid.UUID{11},
						Text:     "",
						Rating:   5,
					}, nil)
			},
			expected: &domain.Comment{
				ID:       commentId,
				AuthorID: uuid.UUID{1},
				SaladID:  uuid.UUID{11},
				Text:     "",
				Rating:   5,
			},
			wantErr: false,
		}, // успешное получение
		{
			name:    "ошибка выполнения запроса в репозитории",
			userId:  userId,
			saladId: saladId,
			beforeTest: func(commentRepo mocks.MockICommentRepository) {
				commentRepo.EXPECT().
					GetBySaladAndUser(context.Background(), saladId, userId).
					Return(nil, fmt.Errorf("getting comment err"))
			},
			wantErr: true,
			errStr:  errors.New("getting comment by salad and user IDs: getting comment err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*commentRepo)
			}

			salad, err := svc.GetBySaladAndUser(context.Background(), tt.saladId, tt.userId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, tt.expected, salad)
			}
		})
	}
}

func TestCommentService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	commentRepo := mocks.NewMockICommentRepository(ctrl)
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
	svc := services.NewCommentService(commentRepo, logger)

	commentId := uuid.New()
	authorId := uuid.New()
	saladId := uuid.New()

	tests := []struct {
		name       string
		comment    *domain.Comment
		beforeTest func(commentRepo mocks.MockICommentRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное обновление",
			comment: &domain.Comment{
				ID:       commentId,
				AuthorID: authorId,
				SaladID:  saladId,
				Text:     "",
				Rating:   5,
			},
			beforeTest: func(commentRepo mocks.MockICommentRepository) {
				commentRepo.EXPECT().
					Update(context.Background(), &domain.Comment{
						ID:       commentId,
						AuthorID: authorId,
						SaladID:  saladId,
						Text:     "",
						Rating:   5,
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное обновление
		{
			name: "оценка меньше минимальной",
			comment: &domain.Comment{
				ID:       commentId,
				AuthorID: authorId,
				SaladID:  saladId,
				Text:     "",
				Rating:   0,
			},
			wantErr: true,
			errStr:  errors.New("updating comment: rate out of range"),
		}, // оценка меньше минимальной
		{
			name: "оценка больше максимальной",
			comment: &domain.Comment{
				ID:       commentId,
				AuthorID: authorId,
				SaladID:  saladId,
				Text:     "",
				Rating:   10,
			},
			wantErr: true,
			errStr:  errors.New("updating comment: rate out of range"),
		}, // оценка больше максимальной
		{
			name: "ошибка выполнения запроса в репозитории",
			comment: &domain.Comment{
				ID:       commentId,
				AuthorID: authorId,
				SaladID:  saladId,
				Text:     "",
				Rating:   5,
			},
			beforeTest: func(commentRepo mocks.MockICommentRepository) {
				commentRepo.EXPECT().
					Update(context.Background(), &domain.Comment{
						ID:       commentId,
						AuthorID: authorId,
						SaladID:  saladId,
						Text:     "",
						Rating:   5,
					}).
					Return(fmt.Errorf("updating comment err"))
			},
			wantErr: true,
			errStr:  errors.New("updating comment: updating comment err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*commentRepo)
			}

			err := svc.Update(context.Background(), tt.comment)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
