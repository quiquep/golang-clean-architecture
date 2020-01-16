package service

import (
	"fmt"
	"testing"

	"../entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RepoMockImplementation struct {
	mock.Mock
}

func (r *RepoMockImplementation) Save(post *entity.Post) (*entity.Post, error) {
	args := r.Called()
	fmt.Printf("Save Args: %v", args)
	result := args.Get(0)
	return result.(*entity.Post), args.Error(1)
}

func (r *RepoMockImplementation) FindAll() ([]entity.Post, error) {
	args := r.Called()
	fmt.Printf("FindAll Args: %v", args)
	result := args.Get(0)
	return result.([]entity.Post), args.Error(1)
}

func TestValidateEmptyPost(t *testing.T) {
	testService := NewPostService(nil)

	err := testService.Validate(nil)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "The post is empty")
}

func TestValidateEmptyPostTitle(t *testing.T) {
	post := entity.Post{ID: 1, Title: "", Text: "B"}

	testService := NewPostService(nil)

	err := testService.Validate(&post)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "The post title is empty")
}

func TestFindAll(t *testing.T) {
	mockRepo := new(RepoMockImplementation)

	post := entity.Post{ID: 1, Title: "A", Text: "B"}
	// setup expectations
	mockRepo.On("FindAll").Return([]entity.Post{post}, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.FindAll()

	//Mock Assertion (behaviour)
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, result[0].Title, "A")
}

func TestCreate(t *testing.T) {
	mockRepo := new(RepoMockImplementation)

	post := entity.Post{ID: 1, Title: "A", Text: "B"}

	// setup expectations
	mockRepo.On("Save").Return(&post, nil)

	testService := NewPostService(mockRepo)

	result, _ := testService.Create(&post)

	//Mock Assertion (behaviour)
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, result.Text, "B")
}
