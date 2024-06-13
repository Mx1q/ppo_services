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
	"net/mail"
	"testing"
)

func TestUserService_Create(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockIUserRepository(ctrl)
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
	svc := services.NewUserService(userRepo, logger)

	testId := uuid.New()

	tests := []struct {
		name       string
		user       *domain.User
		beforeTest func(userRepo mocks.MockIUserRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное создание пользователя",
			user: &domain.User{
				ID:       testId,
				Username: "successCreate",
				Name:     "successCreate",
				Password: "successCreatePass",
				Email: mail.Address{
					Name:    "successCreate",
					Address: "successCreate@mail.ru",
				},
			},
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					Create(context.Background(), &domain.User{
						ID:       testId,
						Username: "successCreate",
						Name:     "successCreate",
						Password: "successCreatePass",
						Email: mail.Address{
							Name:    "successCreate",
							Address: "successCreate@mail.ru",
						},
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное создание пользователя
		{
			name: "пустое имя пользователя",
			user: &domain.User{
				ID:       testId,
				Username: "",
				Password: "emptyUsernamePass",
				Email: mail.Address{
					Name:    "emptyUsername",
					Address: "emptyUsername@mail.ru",
				},
			},
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					Create(context.Background(), domain.User{
						ID:       testId,
						Username: "",
						Password: "emptyUsernamePass",
						Email: mail.Address{
							Name:    "emptyUsername",
							Address: "emptyUsername@mail.ru",
						},
					}).
					Return(nil).
					AnyTimes()
			},
			wantErr: true,
			errStr:  errors.New("creating user: empty username"),
		}, // пустое имя пользователя
		{
			name: "пустое имя",
			user: &domain.User{
				ID:       testId,
				Username: "emptyName",
				Password: "emptyNamePass",
				Name:     "",
				Email: mail.Address{
					Name:    "emptyUsername",
					Address: "emptyUsername@mail.ru",
				},
			},
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					Create(context.Background(), domain.User{
						ID:       testId,
						Username: "emptyName",
						Password: "emptyNamePass",
						Name:     "",
						Email: mail.Address{
							Name:    "emptyUsername",
							Address: "emptyUsername@mail.ru",
						},
					}).
					Return(nil).
					AnyTimes()
			},
			wantErr: true,
			errStr:  errors.New("creating user: empty name"),
		}, // пустое имя
		{
			name: "пустой пароль",
			user: &domain.User{
				ID:       testId,
				Username: "emptyPassword",
				Password: "",
				Email: mail.Address{
					Name:    "emptyPassword",
					Address: "emptyPassword@mail.ru",
				},
			},
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					Create(context.Background(), &domain.User{
						ID:       testId,
						Username: "emptyPassword",
						Password: "",
						Email: mail.Address{
							Name:    "emptyPassword",
							Address: "emptyPassword@mail.ru",
						},
					}).
					Return(nil).
					AnyTimes()
			},
			wantErr: true,
			errStr:  errors.New("creating user: empty password"),
		}, // пустой пароль
		{
			name: "пустая почта",
			user: &domain.User{
				ID:       testId,
				Username: "emptyMail",
				Name:     "emptyMail",
				Password: "emptyMailPass",
				Email: mail.Address{
					Name:    "emptyMail",
					Address: "",
				},
			},
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					Create(context.Background(), &domain.User{
						ID:       testId,
						Username: "emptyMail",
						Name:     "emptyMail",
						Password: "emptyMailPass",
						Email: mail.Address{
							Name:    "emptyMail",
							Address: "",
						},
					}).
					Return(nil).
					AnyTimes()
			},
			wantErr: true,
			errStr:  errors.New("creating user: mail: no address"),
		}, // пустая почта
		{
			name: "некорректная почта (отсутствует @)",
			user: &domain.User{
				ID:       testId,
				Username: "invalidMail1",
				Name:     "invalidMail1",
				Password: "invalidMail1Pass",
				Email: mail.Address{
					Name:    "invalidMail1",
					Address: "invalidMail1",
				},
			},
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					Create(context.Background(), &domain.User{
						ID:       testId,
						Username: "invalidMail1",
						Name:     "invalidMail1",
						Password: "invalidMail1Pass",
						Email: mail.Address{
							Name:    "invalidMail1",
							Address: "invalidMail1",
						},
					}).
					Return(nil).
					AnyTimes()
			},
			wantErr: true,
			errStr:  errors.New("creating user: mail: missing '@' or angle-addr"),
		}, // некорректная почта (отсутствует @)
		{
			name: "некорректная почта (отсутствует имя пользователя)",
			user: &domain.User{
				ID:       testId,
				Username: "invalidMail2",
				Name:     "invalidMail2",
				Password: "invalidMail2Pass",
				Email: mail.Address{
					Name:    "invalidMail2",
					Address: "@mail.ru",
				},
			},
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					Create(context.Background(), &domain.User{
						ID:       testId,
						Username: "invalidMail2",
						Name:     "invalidMail2",
						Password: "invalidMail2Pass",
						Email: mail.Address{
							Name:    "invalidMail2",
							Address: "invalidMail2",
						},
					}).
					Return(nil).
					AnyTimes()
			},
			wantErr: true,
			errStr:  errors.New("creating user: mail: missing word in phrase: mail: invalid string"),
		}, // некорректная почта (отсутствует имя пользователя)
		{
			name: "ошибка выполнения запроса в репозитории",
			user: &domain.User{
				ID:       testId,
				Username: "create",
				Name:     "create",
				Password: "createPass",
				Email: mail.Address{
					Name:    "create",
					Address: "create@mail.ru",
				},
			},
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					Create(context.Background(), &domain.User{
						ID:       testId,
						Username: "create",
						Name:     "create",
						Password: "createPass",
						Email: mail.Address{
							Name:    "create",
							Address: "create@mail.ru",
						},
					}).
					Return(fmt.Errorf("creating user err")).
					AnyTimes()
			},
			wantErr: true,
			errStr:  errors.New("creating user: creating user err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*userRepo)
			}

			err := svc.Create(context.Background(), tt.user)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}

func TestUserService_DeleteById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockIUserRepository(ctrl)
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
	svc := services.NewUserService(userRepo, logger)

	testId := uuid.New()

	tests := []struct {
		name       string
		id         uuid.UUID
		beforeTest func(userRepo mocks.MockIUserRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное удаление",
			id:   testId,
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					DeleteById(context.Background(), testId).
					Return(nil)
			},
			wantErr: false,
		}, // успешное удаление
		{
			name: "ошибка выполнения запроса в репозитории",
			id:   testId,
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					DeleteById(context.Background(), testId).
					Return(fmt.Errorf("delete by id err"))
			},
			wantErr: true,
			errStr:  errors.New("deleting user by id: delete by id err"),
		}, // ошибка выполнения запроса в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*userRepo)
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

func TestUserService_GetAll(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockIUserRepository(ctrl)
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
	svc := services.NewUserService(userRepo, logger)

	page := 1

	tests := []struct {
		name       string
		page       int
		beforeTest func(userRepo mocks.MockIUserRepository)
		expected   []*domain.User
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное получение списка всех пользователей",
			page: page,
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					GetAll(context.Background(), page).
					Return([]*domain.User{
						{
							ID:       uuid.UUID{1},
							Username: "first",
							Password: "firstPass",
							Email: mail.Address{
								Name:    "first",
								Address: "first@mail.ru",
							},
						},
						{
							ID:       uuid.UUID{2},
							Username: "second",
							Password: "secondPass",
							Email: mail.Address{
								Name:    "second",
								Address: "second@mail.ru",
							},
						},
						{
							ID:       uuid.UUID{3},
							Username: "third",
							Password: "thirdPass",
							Email: mail.Address{
								Name:    "third",
								Address: "third@mail.ru",
							},
						},
						{
							ID:       uuid.UUID{4},
							Username: "fourth",
							Password: "fourthPass",
							Email: mail.Address{
								Name:    "fourth",
								Address: "fourth@mail.ru",
							},
						},
					}, nil)
			},
			expected: []*domain.User{
				{
					ID:       uuid.UUID{1},
					Username: "first",
					Password: "firstPass",
					Email: mail.Address{
						Name:    "first",
						Address: "first@mail.ru",
					},
				},
				{
					ID:       uuid.UUID{2},
					Username: "second",
					Password: "secondPass",
					Email: mail.Address{
						Name:    "second",
						Address: "second@mail.ru",
					},
				},
				{
					ID:       uuid.UUID{3},
					Username: "third",
					Password: "thirdPass",
					Email: mail.Address{
						Name:    "third",
						Address: "third@mail.ru",
					},
				},
				{
					ID:       uuid.UUID{4},
					Username: "fourth",
					Password: "fourthPass",
					Email: mail.Address{
						Name:    "fourth",
						Address: "fourth@mail.ru",
					},
				},
			},
			wantErr: false,
		}, // успешное получение списка всех пользователей
		{
			name: "ошибка получения всех пользователей в репозитории",
			page: page,
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					GetAll(context.Background(), page).
					Return(nil, fmt.Errorf("get all users err"))
			},
			wantErr: true,
			errStr:  errors.New("getting all users: get all users err"),
		}, // ошибка получения всех пользователей в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*userRepo)
			}

			users, err := svc.GetAll(context.Background(), tt.page)
			if tt.wantErr {
				require.Equal(t, err.Error(), tt.errStr.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, users, tt.expected)
			}
		})
	}
}

func TestUserService_GetById(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockIUserRepository(ctrl)
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
	svc := services.NewUserService(userRepo, logger)
	testId := uuid.New()

	tests := []struct {
		name       string
		id         uuid.UUID
		beforeTest func(userRepo mocks.MockIUserRepository)
		expected   *domain.User
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное получение пользователя по id",
			id:   testId,
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					GetById(context.Background(), testId).
					Return(&domain.User{
						ID:       testId,
						Username: "testGetById",
						Password: "testGetById_pass",
						Email: mail.Address{
							Name:    "testGetById",
							Address: "testGetById@mail.ru",
						},
					}, nil)
			},
			expected: &domain.User{
				ID:       testId,
				Username: "testGetById",
				Password: "testGetById_pass",
				Email: mail.Address{
					Name:    "testGetById",
					Address: "testGetById@mail.ru",
				},
			},
			wantErr: false,
		}, // успешное получение пользователя по id
		{
			name: "ошибка получения пользователя по id в репозитории",
			id:   testId,
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					GetById(context.Background(), testId).
					Return(nil, fmt.Errorf("get user by id err"))
			},
			wantErr: true,
			errStr:  errors.New("getting user by id: get user by id err"),
		}, // ошибка получения пользователя по id в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*userRepo)
			}

			user, err := svc.GetById(context.Background(), tt.id)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, user, tt.expected)
			}
		})
	}
}

func TestUserService_GetByUsername(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockIUserRepository(ctrl)
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
	svc := services.NewUserService(userRepo, logger)
	testId := uuid.New()
	testUsername := "username"

	tests := []struct {
		name       string
		username   string
		beforeTest func(userRepo mocks.MockIUserRepository)
		expected   *domain.User
		wantErr    bool
		errStr     error
	}{
		{
			name:     "успешное получение пользователя по username",
			username: testUsername,
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					GetByUsername(context.Background(), testUsername).
					Return(&domain.User{
						ID:       testId,
						Username: testUsername,
						Password: "testGetByUsername_pass",
						Email: mail.Address{
							Name:    "testGetByUsername",
							Address: "testGetByUsername@mail.ru",
						},
					}, nil)
			},
			expected: &domain.User{
				ID:       testId,
				Username: testUsername,
				Password: "testGetByUsername_pass",
				Email: mail.Address{
					Name:    "testGetByUsername",
					Address: "testGetByUsername@mail.ru",
				},
			},
			wantErr: false,
		}, // успешное получение пользователя по username
		{
			name:     "ошибка получения пользователя по username в репозитории",
			username: testUsername,
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					GetByUsername(context.Background(), testUsername).
					Return(nil, fmt.Errorf("get user by username err"))
			},
			wantErr: true,
			errStr:  errors.New("getting user by username: get user by username err"),
		}, // ошибка получения пользователя по username в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*userRepo)
			}

			user, err := svc.GetByUsername(context.Background(), tt.username)
			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				require.Equal(t, user, tt.expected)
			}
		})
	}
}

func TestUserService_Update(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userRepo := mocks.NewMockIUserRepository(ctrl)
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
	svc := services.NewUserService(userRepo, logger)
	testId := uuid.New()

	tests := []struct {
		name       string
		user       *domain.User
		beforeTest func(userRepo mocks.MockIUserRepository)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешное обновление пользователя",
			user: &domain.User{
				ID:       testId,
				Username: "testUpdate",
				Name:     "testUpdate",
				Password: "testUpdate_pass",
				Email: mail.Address{
					Name:    "testUpdate",
					Address: "testUpdate@mail.ru",
				},
			},
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					Update(context.Background(), &domain.User{
						ID:       testId,
						Username: "testUpdate",
						Name:     "testUpdate",
						Password: "testUpdate_pass",
						Email: mail.Address{
							Name:    "testUpdate",
							Address: "testUpdate@mail.ru",
						},
					}).
					Return(nil)
			},
			wantErr: false,
		}, // успешное обновление пользователя
		{
			name: "пустое имя пользователя",
			user: &domain.User{
				ID:       testId,
				Username: "",
				Name:     "testUpdate",
				Password: "testUpdate_pass",
				Email: mail.Address{
					Name:    "testUpdate",
					Address: "testUpdate@mail.ru",
				},
			},
			wantErr: true,
			errStr:  errors.New("updating user: empty username"),
		}, // пустое имя пользователя
		{
			name: "пустое имя",
			user: &domain.User{
				ID:       testId,
				Username: "testUpdate",
				Name:     "",
				Password: "testUpdate_pass",
				Email: mail.Address{
					Name:    "testUpdate",
					Address: "testUpdate@mail.ru",
				},
			},
			wantErr: true,
			errStr:  errors.New("updating user: empty name"),
		}, // пустое имя
		{
			name: "пустой пароль",
			user: &domain.User{
				ID:       testId,
				Username: "emptyPassword",
				Name:     "emptyPassword",
				Password: "",
				Email: mail.Address{
					Name:    "emptyPassword",
					Address: "emptyPassword@mail.ru",
				},
			},
			wantErr: true,
			errStr:  errors.New("updating user: empty password"),
		}, // пустой пароль
		{
			name: "пустая почта",
			user: &domain.User{
				ID:       testId,
				Username: "emptyMail",
				Name:     "emptyMail",
				Password: "emptyMailPass",
				Email: mail.Address{
					Name:    "emptyMail",
					Address: "",
				},
			},
			wantErr: true,
			errStr:  errors.New("updating user: mail: no address"),
		}, // пустая почта
		{
			name: "некорректная почта (отсутствует @)",
			user: &domain.User{
				ID:       testId,
				Username: "invalidMail1",
				Name:     "invalidMail1",
				Password: "invalidMail1Pass",
				Email: mail.Address{
					Name:    "invalidMail1",
					Address: "invalidMail1",
				},
			},
			wantErr: true,
			errStr:  errors.New("updating user: mail: missing '@' or angle-addr"),
		}, // некорректная почта (отсутствует @)
		{
			name: "некорректная почта (отсутствует имя пользователя)",
			user: &domain.User{
				ID:       testId,
				Username: "invalidMail2",
				Name:     "invalidMail2",
				Password: "invalidMail2Pass",
				Email: mail.Address{
					Name:    "invalidMail2",
					Address: "@mail.ru",
				},
			},
			wantErr: true,
			errStr:  errors.New("updating user: mail: missing word in phrase: mail: invalid string"),
		}, // "некорректная почта (отсутствует имя пользователя)
		{
			name: "ошибка обновления пользователя в репозитории",
			user: &domain.User{
				ID:       testId,
				Username: "testUpdate",
				Name:     "testUpdate",
				Password: "testUpdate_pass",
				Email: mail.Address{
					Name:    "testUpdate",
					Address: "testUpdate@mail.ru",
				},
			},
			beforeTest: func(userRepo mocks.MockIUserRepository) {
				userRepo.EXPECT().
					Update(context.Background(), &domain.User{
						ID:       testId,
						Username: "testUpdate",
						Name:     "testUpdate",
						Password: "testUpdate_pass",
						Email: mail.Address{
							Name:    "testUpdate",
							Address: "testUpdate@mail.ru",
						},
					}).
					Return(fmt.Errorf("update user err"))
			},
			wantErr: true,
			errStr:  errors.New("updating user: update user err"),
		}, // ошибка обновления пользователя в репозитории
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*userRepo)
			}

			err := svc.Update(context.Background(), tt.user)

			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
