package service

import (
	"testing"
	"unittest/model"
	mock_db "unittest/repository/mock"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetBookById(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	expectedBook := model.Book{Id: 100, Name: "テストブック", Price: 1500}

	mockObj := mock_db.NewMockBookRepository(mockCtrl)
	mockObj.EXPECT().GetBookById(100).Return(expectedBook)

	service := BookService{Repository: mockObj}

	actualBook := service.GetBookById(200)

	assert.Equal(t, expectedBook, actualBook)
}
