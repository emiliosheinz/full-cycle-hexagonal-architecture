package cli_test

import (
	"fmt"
	"testing"

	"github.com/emiliosheinz/exagonal-architecture/adapters/cli"
	mock_application "github.com/emiliosheinz/exagonal-architecture/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	productName := "Test Product"
	productPrice := 25.0

	productStatus := "enabled"
	productId := "c7bc89c9-3866-4427-9e05-6e92ae9d3ed3"

	productMock := mock_application.NewMockProductInterface(ctrl)
	productMock.EXPECT().GetId().Return(productId).AnyTimes()
	productMock.EXPECT().GetStatus().Return(productStatus).AnyTimes()
	productMock.EXPECT().GetPrice().Return(productPrice).AnyTimes()
	productMock.EXPECT().GetName().Return(productName).AnyTimes()

	service := mock_application.NewMockProductServiceInterface(ctrl)
	service.EXPECT().Create(productName, productPrice).Return(productMock, nil).AnyTimes()
	service.EXPECT().Get(productId).Return(productMock, nil).AnyTimes()
	service.EXPECT().Enable(gomock.Any()).Return(productMock, nil).AnyTimes()
	service.EXPECT().Disable(gomock.Any()).Return(productMock, nil).AnyTimes()

	expectedResult := fmt.Sprintf(
		"Product ID %s with name %s, price %f, and status %s has been created",
		productId,
		productName,
		productPrice,
		productStatus,
	)
	result, err := cli.Run(service, "create", productId, productName, productPrice)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf(
		"Product %s has been enabled",
		productId,
	)
	result, err = cli.Run(service, "enable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf(
		"Product %s has been disabled",
		productId,
	)
	result, err = cli.Run(service, "disable", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)

	expectedResult = fmt.Sprintf(
		"Product ID: %s\nName: %s\nPrice: %f\nStatus: %s",
		productId,
		productName,
		productPrice,
		productStatus,
	)
	result, err = cli.Run(service, "", productId, "", 0)
	require.Nil(t, err)
	require.Equal(t, expectedResult, result)
}
