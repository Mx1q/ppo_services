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

func TestAuthService_Login(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	jwtKey := "abcdefgh123"
	repo := mocks.NewMockIAuthRepository(ctrl)
	crypto := mocks.NewMockIHashCrypto(ctrl)
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
	svc := services.NewAuthService(repo, logger, crypto, jwtKey)

	tests := []struct {
		name       string
		authInfo   *domain.UserAuth
		beforeTest func(authRepo mocks.MockIAuthRepository, crypto mocks.MockIHashCrypto)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешная аутентификация",
			authInfo: &domain.UserAuth{
				Username: "username",
				Password: "pass",
			},
			beforeTest: func(authRepo mocks.MockIAuthRepository, crypto mocks.MockIHashCrypto) {
				authRepo.EXPECT().
					GetByUsername(
						context.Background(),
						"username",
					).
					Return(&domain.UserAuth{
						Username:   "username",
						Password:   "pass",
						HashedPass: "hashedPass",
					}, nil)

				crypto.EXPECT().
					CheckPasswordHash("pass", "hashedPass").
					Return(true)
			},
			wantErr: false,
		}, // успешная аутентификация
		{
			name: "пустое имя пользователя",
			authInfo: &domain.UserAuth{
				Username: "",
				Password: "pass",
			},
			wantErr: true,
			errStr:  errors.New("empty username"),
		}, // пустое имя пользователя
		{
			name: "пустой пароль",
			authInfo: &domain.UserAuth{
				Username: "username",
				Password: "",
			},
			wantErr: true,
			errStr:  errors.New("empty password"),
		}, // пустой пароль
		{
			name: "ошибка получения данных из репозитория",
			authInfo: &domain.UserAuth{
				Username: "username",
				Password: "pass",
			},
			beforeTest: func(authRepo mocks.MockIAuthRepository, crypto mocks.MockIHashCrypto) {
				authRepo.EXPECT().
					GetByUsername(
						context.Background(),
						"username",
					).
					Return(nil, fmt.Errorf("getting user err"))
			},
			wantErr: true,
			errStr:  errors.New("getting user by name: getting user err"),
		}, // ошибка получения данных из репозитория
		{
			name: "неверный пароль",
			authInfo: &domain.UserAuth{
				Username: "username",
				Password: "pass",
			},
			beforeTest: func(authRepo mocks.MockIAuthRepository, crypto mocks.MockIHashCrypto) {
				authRepo.EXPECT().
					GetByUsername(
						context.Background(),
						"username",
					).
					Return(&domain.UserAuth{
						Username:   "username",
						Password:   "pass",
						HashedPass: "hashedPass",
					}, nil)

				crypto.EXPECT().
					CheckPasswordHash("pass", "hashedPass").
					Return(false)
			},
			wantErr: true,
			errStr:  errors.New("invalid password"),
		}, // неверный пароль
		{
			name: "ошибка получения токена",
			authInfo: &domain.UserAuth{
				Username: "username",
				Password: "pass",
			},
			beforeTest: func(authRepo mocks.MockIAuthRepository, crypto mocks.MockIHashCrypto) {
				authRepo.EXPECT().
					GetByUsername(
						context.Background(),
						"username",
					).
					Return(&domain.UserAuth{
						Username:   "username",
						Password:   "pass",
						HashedPass: "hashedPass",
					}, nil)

				crypto.EXPECT().
					CheckPasswordHash("pass", "hashedPass").
					Return(true)

			},
			wantErr: false,
		}, // ошибка получения токена
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*repo, *crypto)
			}

			token, err := svc.Login(context.Background(), tt.authInfo)

			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
				_, errTokenParse := services.VerifyAuthToken(token, jwtKey)
				require.Nil(t, errTokenParse)
			}
		})
	}
}

func TestAuthService_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	repo := mocks.NewMockIAuthRepository(ctrl)
	crypto := mocks.NewMockIHashCrypto(ctrl)
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
	svc := services.NewAuthService(repo, logger, crypto, "abcdefgh123")

	tests := []struct {
		name       string
		authInfo   *domain.User
		beforeTest func(authRepo mocks.MockIAuthRepository, crypto mocks.MockIHashCrypto)
		wantErr    bool
		errStr     error
	}{
		{
			name: "успешная регистрация",
			authInfo: &domain.User{
				Name:     "test",
				Username: "test123",
				Password: "pass123",
				Email: mail.Address{
					Name:    "",
					Address: "test@mail.ru",
				},
			},
			beforeTest: func(authRepo mocks.MockIAuthRepository, crypto mocks.MockIHashCrypto) {
				crypto.EXPECT().
					GenerateHashPass("pass123").
					Return("hashedPass123", nil)

				authRepo.EXPECT().
					Register(
						context.Background(),
						&domain.User{
							Name:     "test",
							Username: "test123",
							Password: "hashedPass123",
							Email: mail.Address{
								Name:    "",
								Address: "test@mail.ru",
							},
						},
					).
					Return(uuid.New(), nil)
			},
			wantErr: false,
		}, // успешная регистрация
		{
			name: "пустое имя пользователя",
			authInfo: &domain.User{
				Name:     "test",
				Username: "",
				Password: "pass123",
				Email: mail.Address{
					Name:    "",
					Address: "test@mail.ru",
				},
			},
			wantErr: true,
			errStr:  errors.New("empty username"),
		}, // пустое имя пользователя
		{
			name: "пустое имя",
			authInfo: &domain.User{
				Name:     "",
				Username: "test",
				Password: "pass123",
				Email: mail.Address{
					Name:    "",
					Address: "test@mail.ru",
				},
			},
			wantErr: true,
			errStr:  errors.New("empty name"),
		}, // пустое имя пользователя
		{
			name: "пустой пароль",
			authInfo: &domain.User{
				Name:     "test",
				Username: "test123",
				Password: "",
				Email: mail.Address{
					Name:    "",
					Address: "test@mail.ru",
				},
			},
			wantErr: true,
			errStr:  errors.New("empty password"),
		}, // пустой пароль
		{
			name: "пустая почта",
			authInfo: &domain.User{
				Username: "emptyMail",
				Name:     "emptyMail",
				Password: "emptyMailPass",
				Email: mail.Address{
					Name:    "emptyMail",
					Address: "",
				},
			},
			wantErr: true,
			errStr:  errors.New("invalid email: mail: no address"),
		}, // пустая почта
		{
			name: "некорректная почта (отсутствует @)",
			authInfo: &domain.User{
				Username: "invalidMail1",
				Name:     "invalidMail1",
				Password: "invalidMail1Pass",
				Email: mail.Address{
					Name:    "invalidMail1",
					Address: "invalidMail1",
				},
			},
			wantErr: true,
			errStr:  errors.New("invalid email: mail: missing '@' or angle-addr"),
		}, // некорректная почта (отсутствует @)
		{
			name: "некорректная почта (отсутствует имя пользователя)",
			authInfo: &domain.User{
				Username: "invalidMail2",
				Name:     "invalidMail2",
				Password: "invalidMail2Pass",
				Email: mail.Address{
					Name:    "invalidMail2",
					Address: "@mail.ru",
				},
			},
			wantErr: true,
			errStr:  errors.New("invalid email: mail: missing word in phrase: mail: invalid string"),
		}, // некорректная почта (отсутствует имя пользователя)
		{
			name: "ошибка выполнения запроса в репозитории",
			authInfo: &domain.User{
				Name:     "test",
				Username: "test123",
				Password: "pass123",
				Email: mail.Address{
					Name:    "",
					Address: "test@mail.ru",
				},
			},
			beforeTest: func(authRepo mocks.MockIAuthRepository, crypto mocks.MockIHashCrypto) {
				crypto.EXPECT().
					GenerateHashPass("pass123").
					Return("hashedPass123", nil)

				authRepo.EXPECT().
					Register(
						context.Background(),
						&domain.User{
							Name:     "test",
							Username: "test123",
							Password: "hashedPass123",
							Email: mail.Address{
								Name:    "",
								Address: "test@mail.ru",
							},
						},
					).
					Return(uuid.Nil, fmt.Errorf("registration err"))
			},
			wantErr: true,
			errStr:  errors.New("registration user: registration err"),
		}, // ошибка выполнения запроса в репозитории
		{
			name: "ошибка получения хэша пароля",
			authInfo: &domain.User{
				Name:     "test",
				Username: "test123",
				Password: "pass123",
				Email: mail.Address{
					Name:    "",
					Address: "test@mail.ru",
				},
			},
			beforeTest: func(authRepo mocks.MockIAuthRepository, crypto mocks.MockIHashCrypto) {
				crypto.EXPECT().
					GenerateHashPass("pass123").
					Return("", errors.New("failed to generate hash"))
			},
			wantErr: true,
			errStr:  errors.New("generating hash: failed to generate hash"),
		}, // ошибка получения хэша пароля
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.beforeTest != nil {
				tt.beforeTest(*repo, *crypto)
			}

			_, err := svc.Register(context.Background(), tt.authInfo)

			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
