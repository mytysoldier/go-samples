package db

import (
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	mock_db "unittests/db/mocks"
	"unittests/db/model"
)

func TestGetUser_success(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expectedUser := model.User{Id: 1, Name: "test_user"}

	mockObj := mock_db.NewMockUserInterface(mockCtrl)
	mockObj.EXPECT().GetUserById(1).Return(expectedUser)

	mockUser := mockObj.GetUserById(1)

	assert.Equal(t, expectedUser, mockUser)
}
