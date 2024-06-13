package tests

import (
	"context"
	"errors"
	"github.com/Mx1q/ppo_services/services"
	"github.com/Mx1q/ppo_services/tests/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUrlValidatorService_Verify(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
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
	svc := services.NewUrlValidatorService(logger)

	tests := []struct {
		name    string
		text    string
		wantErr bool
		errStr  error
	}{
		{
			name:    "ссылка отсутствует",
			text:    "word",
			wantErr: false,
		}, // ссылка отсутствует
		{
			name:    "передано несколько слов",
			text:    "some words",
			wantErr: true,
			errStr:  errors.New("verifying url: accepts only 1 word"),
		}, // передано несколько слов
		{
			name:    "ссылка без протокола",
			text:    "google.com",
			wantErr: true,
			errStr:  errors.New("verifying url: found google.com"),
		}, // ссылка без протокола
		{
			name:    "ссылка с протоколом",
			text:    "https://bmstu.ru",
			wantErr: true,
			errStr:  errors.New("verifying url: found https://bmstu.ru"),
		}, // ссылка с протоколом
		{
			name:    "протокол + доменное имя + путь",
			text:    "https://bmstu.ru/student",
			wantErr: true,
			errStr:  errors.New("verifying url: found https://bmstu.ru/student"),
		}, // протокол + доменное имя + путь
		{
			name:    "доменное имя + путь",
			text:    "bmstu.ru/student",
			wantErr: true,
			errStr:  errors.New("verifying url: found bmstu.ru/student"),
		}, // доменное имя + путь
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := svc.Verify(context.Background(), tt.text)

			if tt.wantErr {
				require.Equal(t, tt.errStr.Error(), err.Error())
			} else {
				require.Nil(t, err)
			}
		})
	}
}
