package service

import (
	"errors"
	"math/rand"

	"../entity"
	"../repository"
)

type service struct{}

var (
	postRepository repository.PostRepository
)

//NewPostService creates a new service
func NewPostService(repo repository.PostRepository) PostService {
	postRepository = repo
	return &service{}
}

func (*service) Validate(post *entity.Post) error {
	if post == nil {
		err1 := errors.New("The post is empty")
		return err1
	}
	if post.Title == "" {
		err2 := errors.New("The post title is empty")
		return err2
	}
	if post.Text == "" {
		err3 := errors.New("The post text is empty")
		return err3
	}
	return nil
}

func (*service) Create(post *entity.Post) (*entity.Post, error) {
	post.ID = rand.Int63()
	return postRepository.Save(post)
}

func (*service) FindAll() ([]entity.Post, error) {
	return postRepository.FindAll()
}
