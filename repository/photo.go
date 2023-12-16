package repository

import (
	"errors"

	"github.com/novita/finalproject2-revisi/dto"
	"github.com/novita/finalproject2-revisi/entity"
	"gorm.io/gorm"
)

type Photos interface {
	Add(photo *entity.Photo) (*entity.Photo, error)
	Get(id uint) (*entity.Photo, error)
	GetAll() ([]entity.Photo, error)
	Update(photo *entity.Photo, updatedData *dto.Photo) (*entity.Photo, error)
	Delete(photo *entity.Photo) error
}

type PhotosImpl struct {
	db *gorm.DB
}

func NewPhotos(db *gorm.DB) *PhotosImpl {
	return &PhotosImpl{db}
}

func (r *PhotosImpl) Add(photo *entity.Photo) (*entity.Photo, error) {
	if err := r.db.Create(&photo).Error; err != nil {
		return nil, err
	}

	return photo, nil
}

func (r *PhotosImpl) Get(id uint) (*entity.Photo, error) {
	photo := entity.Photo{}

	if err := r.db.First(&photo, id).Error; err != nil {
		return nil, errors.New("photo not found")
	}

	return &photo, nil
}

func (r *PhotosImpl) GetAll() ([]entity.Photo, error) {
	photos := []entity.Photo{}

	if err := r.db.Find(&photos).Error; err != nil {
		return nil, errors.New("failed to get photos")
	}

	return photos, nil
}

func (r *PhotosImpl) Update(photo *entity.Photo, updatedData *dto.Photo) (*entity.Photo, error) {
	if err := r.db.Model(&photo).Updates(&updatedData).Error; err != nil {
		return nil, err
	}

	return photo, nil
}

func (r *PhotosImpl) Delete(photo *entity.Photo) error {
	if err := r.db.Delete(&photo).Error; err != nil {
		return errors.New("failed to delete photo")
	}

	return nil
}
