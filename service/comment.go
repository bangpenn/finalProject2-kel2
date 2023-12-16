package service

import (
	"github.com/novita/finalproject2-revisi/dto"
	"github.com/novita/finalproject2-revisi/entity"
	"github.com/novita/finalproject2-revisi/repository"
)

type Comments interface {
	Add(payload *dto.Comment) (*entity.Comment, error)
	GetAll() ([]entity.Comment, error)
	Update(comment *entity.Comment, updatedData *dto.Comment) (*entity.Comment, error)
	Delete(comment *entity.Comment) error
}

type CommentsImpl struct {
	repository repository.Comments
}

func NewComments(repository repository.Comments) *CommentsImpl {
	return &CommentsImpl{repository}
}

func (s *CommentsImpl) Add(payload *dto.Comment) (*entity.Comment, error) {
	comment := entity.Comment{
		Message: payload.Message,
		PhotoID: payload.PhotoID,
		UserID:  payload.UserID,
	}

	return s.repository.Create(&comment)
}

func (s *CommentsImpl) GetAll() ([]entity.Comment, error) {
	return s.repository.GetAll()
}

func (s *CommentsImpl) Update(comment *entity.Comment, updatedData *dto.Comment) (*entity.Comment, error) {
	return s.repository.Update(comment, updatedData)
}

func (s *CommentsImpl) Delete(comment *entity.Comment) error {
	return s.repository.Delete(comment)
}
