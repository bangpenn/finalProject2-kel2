package service

import (
	"github.com/novita/finalproject2-revisi/dto"
	"github.com/novita/finalproject2-revisi/entity"
	"github.com/novita/finalproject2-revisi/repository"
)

type SocialMedias interface {
	Add(payload *dto.SocialMedia) (*entity.SocialMedia, error)
	GetAll() ([]entity.SocialMedia, error)
	Update(socialMedia *entity.SocialMedia) (*entity.SocialMedia, error)
	Delete(socialMedia *entity.SocialMedia) error
}

type SocialMediasImpl struct {
	repository repository.SocialMedias
}

func NewSocialMedias(repository repository.SocialMedias) *SocialMediasImpl {
	return &SocialMediasImpl{repository}
}

func (s *SocialMediasImpl) Add(payload *dto.SocialMedia) (*entity.SocialMedia, error) {
	socialMedia := entity.SocialMedia{
		Name:           payload.Name,
		SocialMediaURL: payload.SocialMediaURL,
		UserID:         payload.UserID,
	}

	return s.repository.Add(&socialMedia)
}

func (s *SocialMediasImpl) GetAll() ([]entity.SocialMedia, error) {
	return s.repository.GetAll()
}

func (s *SocialMediasImpl) Update(socialMedia *entity.SocialMedia) (*entity.SocialMedia, error) {
	return s.repository.Update(socialMedia)
}

func (s *SocialMediasImpl) Delete(socialMedia *entity.SocialMedia) error {
	return s.repository.Delete(socialMedia)
}
