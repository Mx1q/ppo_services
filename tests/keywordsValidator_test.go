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

func TestKeywordValidatorService_NewKeywordValidatorService(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	keywordRepo := mocks.NewMockIKeywordValidatorRepository(ctrl)
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

	tests := []struct {
		name       string
		repo       *mocks.MockIKeywordValidatorRepository
		beforeTest func(keywordRepo mocks.MockIKeywordValidatorRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное создание",
			repo: keywordRepo,
			beforeTest: func(keywordRepo mocks.MockIKeywordValidatorRepository) {
				keywordRepo.EXPECT().
					GetAll(context.Background()).
					Return(make(map[string]uuid.UUID), nil)
			},
			wantErr: false,
		}, // успешное создание
		{
			name: "ошибка выполнения запроса в репозитории",
			repo: keywordRepo,
			beforeTest: func(keywordRepo mocks.MockIKeywordValidatorRepository) {
				keywordRepo.EXPECT().
					GetAll(context.Background()).
					Return(nil, fmt.Errorf("getting keywords err"))
			},
			wantErr: true,
			errStr:  errors.New("creating keywords validator: getting keywords err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*keywordRepo)
			}

			_, err := services.NewKeywordValidatorService(context.Background(), tt.repo, logger)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestKeywordValidatorService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	keywordRepo := mocks.NewMockIKeywordValidatorRepository(ctrl)
	keywordRepo.EXPECT().
		GetAll(context.Background()).
		Return(make(map[string]uuid.UUID), nil)
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
	svc, _ := services.NewKeywordValidatorService(context.Background(), keywordRepo, logger)

	tests := []struct {
		name       string
		keyword    *domain.KeyWord
		beforeTest func(keywordRepo mocks.MockIKeywordValidatorRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное создание",
			keyword: &domain.KeyWord{
				ID:   uuid.UUID{1},
				Word: "banned_word",
			},
			beforeTest: func(keywordRepo mocks.MockIKeywordValidatorRepository) {
				keywordRepo.EXPECT().
					Create(context.Background(), &domain.KeyWord{
						ID:   uuid.UUID{1},
						Word: "banned_word",
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное создание
		{
			name: "пустое ключевое слово",
			keyword: &domain.KeyWord{
				ID:   uuid.UUID{1},
				Word: "",
			},
			wantErr: true,
			errStr:  errors.New("creating keyword: empty word"),
		}, // пустое ключевое слово
		{
			name: "несколько слов",
			keyword: &domain.KeyWord{
				ID:   uuid.UUID{1},
				Word: "two words",
			},
			wantErr: true,
			errStr:  errors.New("creating keyword: accepts only 1 word"),
		}, // несколько слов
		{
			name: "ошибка выполнения запроса в репозитории",
			keyword: &domain.KeyWord{
				ID:   uuid.UUID{1},
				Word: "banned_word",
			},
			beforeTest: func(keywordRepo mocks.MockIKeywordValidatorRepository) {
				keywordRepo.EXPECT().
					Create(context.Background(), &domain.KeyWord{
						ID:   uuid.UUID{1},
						Word: "banned_word",
					}).
					Return(fmt.Errorf("creating keyword err"))
			},
			wantErr: true,
			errStr:  errors.New("creating keyword: creating keyword err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*keywordRepo)
			}

			err := svc.Create(context.Background(), tt.keyword)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestKeywordValidatorService_DeleteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	keywordRepo := mocks.NewMockIKeywordValidatorRepository(ctrl)
	keywordRepo.EXPECT().
		GetAll(context.Background()).
		Return(make(map[string]uuid.UUID), nil)
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
	svc, _ := services.NewKeywordValidatorService(context.Background(), keywordRepo, logger)

	keywordId := uuid.New()

	tests := []struct {
		name       string
		keywordId  uuid.UUID
		beforeTest func(keywordRepo mocks.MockIKeywordValidatorRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name:      "успешное удаление",
			keywordId: keywordId,
			beforeTest: func(keywordRepo mocks.MockIKeywordValidatorRepository) {
				keywordRepo.EXPECT().
					DeleteById(context.Background(), keywordId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное удаление
		{
			name:      "ошибка выполнения запроса в репозитории",
			keywordId: keywordId,
			beforeTest: func(keywordRepo mocks.MockIKeywordValidatorRepository) {
				keywordRepo.EXPECT().
					DeleteById(context.Background(), keywordId).
					Return(fmt.Errorf("deleting keyword err"))
			},
			wantErr: true,
			errStr:  errors.New("deleting keyword by id: deleting keyword err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*keywordRepo)
			}

			err := svc.DeleteById(context.Background(), tt.keywordId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestKeywordValidatorService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	keywordRepo := mocks.NewMockIKeywordValidatorRepository(ctrl)
	keywordRepo.EXPECT().
		GetAll(context.Background()).
		Return(make(map[string]uuid.UUID), nil)
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
	svc, _ := services.NewKeywordValidatorService(context.Background(), keywordRepo, logger)

	tests := []struct {
		name       string
		beforeTest func(keywordRepo mocks.MockIKeywordValidatorRepository)
		expected   map[string]uuid.UUID
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное получение",
			beforeTest: func(keywordRepo mocks.MockIKeywordValidatorRepository) {
				keywordRepo.EXPECT().
					GetAll(context.Background()).
					Return(map[string]uuid.UUID{
						"banned1": {1},
						"banned2": {2},
						"banned3": {3},
						"banned4": {4},
					}, nil)
			},
			expected: map[string]uuid.UUID{
				"banned1": {1},
				"banned2": {2},
				"banned3": {3},
				"banned4": {4},
			},
			wantErr: false,
		}, // успешное получение
		{
			name: "ошибка выполнения запроса в репозитории",
			beforeTest: func(keywordRepo mocks.MockIKeywordValidatorRepository) {
				keywordRepo.EXPECT().
					GetAll(context.Background()).
					Return(nil, fmt.Errorf("getting keywords err"))
			},
			wantErr: true,
			errStr:  errors.New("getting all keywords: getting keywords err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*keywordRepo)
			}

			keywords, err := svc.GetAll(context.Background())
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, keywords, tt.expected)
			}
		})
	}
}

func TestKeywordValidatorService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	keywordRepo := mocks.NewMockIKeywordValidatorRepository(ctrl)
	keywordRepo.EXPECT().
		GetAll(context.Background()).
		Return(make(map[string]uuid.UUID), nil)
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
	svc, _ := services.NewKeywordValidatorService(context.Background(), keywordRepo, logger)

	keywordId := uuid.New()

	tests := []struct {
		name       string
		keywordId  uuid.UUID
		beforeTest func(keywordRepo mocks.MockIKeywordValidatorRepository)
		expected   *domain.KeyWord
		wantErr    bool
		errStr     error
	}{
		{
			name:      "успешное получение",
			keywordId: keywordId,
			beforeTest: func(keywordRepo mocks.MockIKeywordValidatorRepository) {
				keywordRepo.EXPECT().
					GetById(context.Background(), keywordId).
					Return(&domain.KeyWord{
						ID:   keywordId,
						Word: "banned_word",
					}, nil)
			},
			expected: &domain.KeyWord{
				ID:   keywordId,
				Word: "banned_word",
			},
			wantErr: false,
		}, // успешное получение
		{
			name:      "ошибка выполнения запроса в репозитории",
			keywordId: keywordId,
			beforeTest: func(keywordRepo mocks.MockIKeywordValidatorRepository) {
				keywordRepo.EXPECT().
					GetById(context.Background(), keywordId).
					Return(nil, fmt.Errorf("getting keyword err"))
			},
			wantErr: true,
			errStr:  errors.New("getting keyword by id: getting keyword err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*keywordRepo)
			}

			keyword, err := svc.GetById(context.Background(), tt.keywordId)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, keyword, tt.expected)
			}
		})
	}
}

func TestKeywordValidatorService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	keywordRepo := mocks.NewMockIKeywordValidatorRepository(ctrl)
	keywordRepo.EXPECT().
		GetAll(context.Background()).
		Return(make(map[string]uuid.UUID), nil)
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
	svc, _ := services.NewKeywordValidatorService(context.Background(), keywordRepo, logger)

	tests := []struct {
		name       string
		keyword    *domain.KeyWord
		beforeTest func(keywordRepo mocks.MockIKeywordValidatorRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное обновление",
			keyword: &domain.KeyWord{
				ID:   uuid.UUID{1},
				Word: "banned_word",
			},
			beforeTest: func(keywordRepo mocks.MockIKeywordValidatorRepository) {
				keywordRepo.EXPECT().
					Update(context.Background(), &domain.KeyWord{
						ID:   uuid.UUID{1},
						Word: "banned_word",
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное обновление
		{
			name: "пустое ключевое слово",
			keyword: &domain.KeyWord{
				ID:   uuid.UUID{1},
				Word: "",
			},
			wantErr: true,
			errStr:  errors.New("updating keyword: empty word"),
		}, // пустое ключевое слово
		{
			name: "несколько слов",
			keyword: &domain.KeyWord{
				ID:   uuid.UUID{1},
				Word: "two words",
			},
			wantErr: true,
			errStr:  errors.New("updating keyword: accepts only 1 word"),
		}, // несколько слов
		{
			name: "ошибка выполнения запроса в репозитории",
			keyword: &domain.KeyWord{
				ID:   uuid.UUID{1},
				Word: "banned_word",
			},
			beforeTest: func(keywordRepo mocks.MockIKeywordValidatorRepository) {
				keywordRepo.EXPECT().
					Update(context.Background(), &domain.KeyWord{
						ID:   uuid.UUID{1},
						Word: "banned_word",
					}).
					Return(fmt.Errorf("updating keyword err"))
			},
			wantErr: true,
			errStr:  errors.New("updating keyword: updating keyword err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*keywordRepo)
			}

			err := svc.Update(context.Background(), tt.keyword)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestKeywordValidatorService_Verify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	keywordRepo := mocks.NewMockIKeywordValidatorRepository(ctrl)
	keywordRepo.EXPECT().
		GetAll(context.Background()).
		Return(map[string]uuid.UUID{
			"banned1": {1},
			"banned2": {2},
			"banned3": {3},
			"banned4": {4},
		}, nil)
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
	svc, _ := services.NewKeywordValidatorService(context.Background(), keywordRepo, logger)

	tests := []struct {
		name       string
		word       string
		beforeTest func(keywordRepo mocks.MockIKeywordValidatorRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name:    "отсутствие ключевых слов",
			word:    "word_for_test",
			wantErr: false,
		}, // отсутствие ключевых слов
		{
			name:    "аргумент - несколько слов",
			word:    "two words",
			wantErr: true,
			errStr:  errors.New("verifying keywords: accepts only 1 word"),
		}, // аргумент - несколько слов
		{
			name:    "найдено ключевое слово",
			word:    "banned4",
			wantErr: true,
			errStr:  errors.New("verifying keywords: found banned4"),
		}, // найдено ключевое слово
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*keywordRepo)
			}

			err := svc.Verify(context.Background(), tt.word)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
